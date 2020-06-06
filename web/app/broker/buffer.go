package broker

import (
	"bitbucket.org/makeitplay/lugo-frontend/web/app"
	"bitbucket.org/makeitplay/lugo-server/machine"
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

type RateCounter interface {
	Incr(int64)
	Rate() int64
}

type BufferHandler struct {
	RateCounter      RateCounter
	Logger           *zap.SugaredLogger
	bufferedUpdates  chan app.FrontEndUpdate
	bufferStage      chan app.FrontEndUpdate
	bufferOn         chan bool
	lastReceivedTurn uint32
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

func (h *BufferHandler) Start(callback func(data BufferedEvent)) {
	if h.bufferedUpdates != nil {
		close(h.bufferedUpdates)
	}
	h.bufferOn = make(chan bool)
	h.bufferedUpdates = make(chan app.FrontEndUpdate, MaxUpdateBuffer)
	h.bufferStage = make(chan app.FrontEndUpdate, MaxUpdateBuffer)
	go h.stageUpdates()
	pulse := make(chan bool)
	go func() {
		h.streamBuffer(callback, pulse)
		close(pulse)
	}()

	go func() {
		for {
			select {
			case on := <-pulse:
				h.Logger.Infof("Streaming? %v", on)
			}
		}
	}()

	//h.Logger.Infof("TURNING BUFFER ON")
	//time.Sleep(5 * time.Second)
}

func (h *BufferHandler) QueueUp(update app.FrontEndUpdate, currentTurn uint32) error {
	h.Logger.Warnf("added update: %v", update.Type)
	select {
	case h.bufferedUpdates <- update:
		h.lastReceivedTurn = currentTurn
	default:
		return ErrBufferNotInitialized
	}
	return nil
}

func (h *BufferHandler) stageUpdates() {
	for {
		select {
		case <-h.bufferOn:
			return
		case update := <-h.bufferedUpdates:
			h.RateCounter.Incr(1)
			h.bufferStage <- update
		}
	}
}

func (h *BufferHandler) streamBuffer(callback func(data BufferedEvent), pulse chan<- bool) {
	var minBufferSize int
	streamer := func() {
		for {
			select {
			case <-h.bufferOn:
				return
			case update := <-h.bufferStage:
				h.Logger.Infof("buffer size: %d", len(h.bufferStage))
				callback(BufferedEvent{Update: update})
				if len(h.bufferStage) < minBufferSize {
					return
				}
			}
		}
	}
	// ideally we want 20 FPS, but a little slower won't hurt and avoid buffering too much
	minAcceptableRate := int64(17) // FPS
	for {
		rate := h.RateCounter.Rate()
		missingFrames := float64(machine.GameDuration - int(h.lastReceivedTurn))
		// timeToBeBuffered is the missing frames translated to the TIME dimension
		timeToBeBuffered := missingFrames * (1 / float64(minAcceptableRate))

		// s = s1 + vt --->
		bufferSize := math.Floor(timeToBeBuffered * float64(rate-minAcceptableRate))
		if bufferSize <= 0 {
			//even if the server is faster than necessary, let's buffer 5 secons
			bufferSize = 5 * float64(minAcceptableRate)
		}
		minBufferSize = int(math.Floor(bufferSize * 0.8)) //80% of the expected buffer

		if len(h.bufferStage) >= int(bufferSize) {
			helperNonBlockingPulse(true, pulse)
			streamer()
		} else {
			helperNonBlockingPulse(false, pulse)
			time.Sleep(1 * time.Second)
		}
	}
}

func helperNonBlockingPulse(p bool, pulse chan<- bool) {
	select {
	case pulse <- p:
	default:
	}
}
