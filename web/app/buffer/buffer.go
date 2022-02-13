package buffer

//
//var (
//	ErrBufferNotInitialized = errors.New("buffer not initialized")
//)
//
//type HitsCounter interface {
//	Incr(int64)
//	Hits() int64
//}
//
//func NewBufferizer(log *zap.SugaredLogger, counter HitsCounter) *Bufferizer {
//	return &Bufferizer{
//		HitsCounter: counter,
//		Logger:      log,
//	}
//}
//
//type Bufferizer struct {
//	HitsCounter      HitsCounter
//	Logger           *zap.SugaredLogger
//	bufferedUpdates  chan app.FrontEndUpdate
//	bufferStage      chan app.FrontEndUpdate
//	bufferOn         chan bool
//	lastReceivedTurn uint32
//	lastSentTurn     uint32
//	expectedTurns    uint32
//}
//
//func (h *Bufferizer) Stop() {
//	if h.bufferOn != nil {
//		close(h.bufferOn)
//	}
//	if h.bufferedUpdates != nil {
//		close(h.bufferedUpdates)
//	}
//	if h.bufferStage != nil {
//		close(h.bufferStage)
//	}
//	h.lastReceivedTurn = 0
//	h.lastSentTurn = 0
//}
//
//func (h *Bufferizer) Start(callback func(data broker.BufferedEvent), expectedTurns uint32) <-chan float32 {
//	h.bufferOn = make(chan bool)
//	h.bufferedUpdates = make(chan app.FrontEndUpdate, broker.MaxUpdateBuffer)
//	h.bufferStage = make(chan app.FrontEndUpdate, broker.MaxUpdateBuffer)
//	h.expectedTurns = expectedTurns
//	go h.stageUpdates()
//	pulse := make(chan float32)
//	go func() {
//		h.streamBuffer(callback, pulse)
//		close(pulse)
//	}()
//	return pulse
//}
//
//func (h *Bufferizer) QueueUp(update app.FrontEndUpdate) error {
//	select {
//	case h.bufferedUpdates <- update:
//	default:
//		return ErrBufferNotInitialized
//	}
//	return nil
//}
//
//func (h *Bufferizer) stageUpdates() {
//	for {
//		select {
//		case <-h.bufferOn:
//			return
//		case update, ok := <-h.bufferedUpdates:
//			if !ok {
//				return
//			}
//			if update.Snapshot.Turn > h.lastReceivedTurn {
//				h.lastReceivedTurn = update.Snapshot.Turn
//				h.HitsCounter.Incr(1)
//			}
//			h.bufferStage <- update
//		}
//	}
//}
//
//func (h *Bufferizer) streamBuffer(callback func(data broker.BufferedEvent), pulse chan<- float32) {
//	//var minBufferSize int
//	stream := true
//	streamer := func() {
//		for stream {
//			select {
//			case <-h.bufferOn:
//				return
//			case update, ok := <-h.bufferStage:
//				if !ok {
//					h.Logger.Infof("buffer closed")
//					return
//				}
//				h.lastSentTurn = update.Snapshot.Turn
//				//h.Logger.Infof("buffer size: %d", len(h.bufferStage))
//				callback(broker.BufferedEvent{Update: update})
//			}
//		}
//	}
//	// ideally we want 20 FPS, but a little slower won't hurt and avoid buffering too much
//	minAcceptableRate := int64(17) // FPS
//	minAcceptableBufferSize := 5 * float64(minAcceptableRate)
//	for {
//		select {
//		case <-h.bufferOn:
//			return
//		case <-time.Tick(1 * time.Second):
//			histLast5Sec := h.HitsCounter.Hits()
//			rate := histLast5Sec / broker.MessagesRateMeasureTimeWindow
//			missingTurns := h.expectedTurns - h.lastReceivedTurn
//			// timeToBeBuffered is the missing frames translated to the TIME dimension
//			timeToBeBuffered := float64(missingTurns) * (1 / float64(minAcceptableRate))
//
//			// s = s1 + vt --->
//			desiredBufferSize := math.Floor(timeToBeBuffered * float64(minAcceptableRate-rate))
//			if missingTurns <= 0 {
//				desiredBufferSize = 0
//			} else if desiredBufferSize <= 0 {
//				//even if the server is faster than necessary, let's buffer 5 secons
//				desiredBufferSize = minAcceptableBufferSize
//			} else if desiredBufferSize > broker.MaxUpdateBuffer {
//				desiredBufferSize = broker.MaxUpdateBuffer * 0.8
//			}
//			currentSize := int(h.lastReceivedTurn - h.lastSentTurn)
//			h.Logger.Infof("rate: %d (%d last sec) (missing turns: %d): buffering %.0f sec (desired: %f): current: %d (%d -> %d)",
//				rate, histLast5Sec, missingTurns, timeToBeBuffered, desiredBufferSize, currentSize,
//				h.lastReceivedTurn, h.lastSentTurn,
//			)
//
//			if stream {
//				if currentSize < int(math.Floor(desiredBufferSize*0.8)) { //80% of the expected buffer {
//					stream = false
//					helperNonBlockingPulse(float32(currentSize)/float32(desiredBufferSize), pulse)
//				}
//			} else if currentSize >= int(desiredBufferSize) {
//				stream = true
//				go streamer()
//				helperNonBlockingPulse(1.0, pulse)
//			} else {
//				helperNonBlockingPulse(float32(currentSize)/float32(desiredBufferSize), pulse)
//			}
//		}
//	}
//}
//
//func helperNonBlockingPulse(p float32, pulse chan<- float32) {
//	select {
//	case pulse <- p:
//	default:
//	}
//}
