package broker

import (
	"bitbucket.org/makeitplay/lugo-frontend/web/app"
	"errors"
	"go.uber.org/zap"
	"math"
	"time"
)

var (
	ErrBufferNotInitialized = errors.New("buffer not initialized")
)

type BufferedEvent struct {
	Update app.FrontEndUpdate
}

type HitsCounter interface {
	Incr(int64)
	Hits() int64
}

type BufferHandler struct {
	HitsCounter     HitsCounter
	Logger          *zap.SugaredLogger
	bufferedUpdates chan app.FrontEndUpdate
	bufferStage     chan app.FrontEndUpdate
	bufferOn        chan bool
	lastSentTurn    uint32
	expectedTurns   uint32
}

func (h *BufferHandler) Stop() {
	if h.bufferOn != nil {
		close(h.bufferOn)
	}
	if h.bufferedUpdates != nil {
		close(h.bufferedUpdates)
	}
	if h.bufferStage != nil {
		close(h.bufferStage)
	}
}

func (h *BufferHandler) Start(callback func(data BufferedEvent), expectedTurns uint32) <- chan float32 {
	h.bufferOn = make(chan bool)
	h.bufferedUpdates = make(chan app.FrontEndUpdate, MaxUpdateBuffer)
	h.bufferStage = make(chan app.FrontEndUpdate, MaxUpdateBuffer)
	h.expectedTurns = expectedTurns
	go h.stageUpdates()
	pulse := make(chan float32)
	go func() {
		h.streamBuffer(callback, pulse)
		close(pulse)
	}()

	go func() {
		for {
			select {
			case on := <-pulse:
				h.Logger.Infof("buffer: %0.2f%%", on)
			}
		}
	}()
	return pulse
	//h.Logger.Infof("TURNING BUFFER ON")
	//time.Sleep(5 * time.Second)
}

func (h *BufferHandler) QueueUp(update app.FrontEndUpdate) error {
	//h.Logger.Warnf("added update: %v", update.Type)
	select {
	case h.bufferedUpdates <- update:
	default:
		return ErrBufferNotInitialized
	}
	return nil
}

func (h *BufferHandler) stageUpdates() {
	lastTurn := uint32(0)
	for {
		select {
		case <-h.bufferOn:
			return
		case update, ok := <-h.bufferedUpdates:
			if !ok {
				return
			}
			if update.Snapshot.Turn > lastTurn {
				lastTurn = update.Snapshot.Turn
				h.HitsCounter.Incr(1)
			}
			h.bufferStage <- update
		}
	}
}

func (h *BufferHandler) streamBuffer(callback func(data BufferedEvent), pulse chan<- float32) {
	//var minBufferSize int
	stream := true
	streamer := func() {
		for stream {
			select {
			case <-h.bufferOn:
				return
			case update, ok := <-h.bufferStage:
				if !ok {
					h.Logger.Infof("buffer closed")
					return
				}
				h.lastSentTurn = update.Snapshot.Turn
				h.Logger.Infof("buffer size: %d", len(h.bufferStage))
				callback(BufferedEvent{Update: update})
				//if len(h.bufferStage) < minBufferSize {
				//	return
				//}
			}
		}
	}
	// ideally we want 20 FPS, but a little slower won't hurt and avoid buffering too much
	minAcceptableRate := int64(17) // FPS
	minAcceptableBufferSize := 5 * float64(minAcceptableRate)
	for {
		select {
		case <-h.bufferOn:
			return







			O calculo do buffer ta sendo feito em cima do numero de itens no channel stage, porem
			o calculo do andamento do jogo eh feito baseado no numero do turn, e o numero to turn eh usado
			como base pra estimar quantos frames faltam/
				Entao o calculo da errado pq vai ter muito mais items no channel do que turnos :

					solucao: esquecer a contagem dos channels e usar um inteiro simples ou talvez atomico
			pra contar quantos turnos faltam, e estimar o tempo baseado em turnos, nao em eventos.















		case <-time.Tick(1 * time.Second):
			histLast5Sec := h.HitsCounter.Hits()
			rate := histLast5Sec / MessagesRateMeasureTimeWindow
			missingFrames := float64(int(h.expectedTurns) - int(h.lastSentTurn))
			// timeToBeBuffered is the missing frames translated to the TIME dimension
			timeToBeBuffered := missingFrames * (1 / float64(minAcceptableRate))

			// s = s1 + vt --->
			bufferSize := math.Floor(timeToBeBuffered * float64(minAcceptableRate-rate))
			if missingFrames <= 0 {
				bufferSize = 0
			} else if bufferSize <= 0 {
				//even if the server is faster than necessary, let's buffer 5 secons
				bufferSize = minAcceptableBufferSize
			} else if bufferSize > MaxUpdateBuffer {
				bufferSize = MaxUpdateBuffer * 0.8
			}
			h.Logger.Infof("rate: %d (%d last sec) (missing frames: %f): buffering %f sec (size: %f): current: %d",
				rate, histLast5Sec, missingFrames, timeToBeBuffered, bufferSize, len(h.bufferStage),
			)
			currentSize := len(h.bufferStage)
			if stream {
				if currentSize < int(math.Floor(bufferSize * 0.8)) { //80% of the expected buffer {
					stream = false
					helperNonBlockingPulse(float32(currentSize)/float32(bufferSize), pulse)
				}
			} else if currentSize >= int(bufferSize) {
					stream = true
					go streamer()
					helperNonBlockingPulse(1.0, pulse)
			} else {
				helperNonBlockingPulse(float32(currentSize)/float32(bufferSize), pulse)
			}
		}
	}
}

func helperNonBlockingPulse(p float32, pulse chan<- float32) {
	select {
	case pulse <- p:
	default:
	}
}
