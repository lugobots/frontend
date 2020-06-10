package broker

import "bitbucket.org/makeitplay/lugo-frontend/web/app"

type BufferedEvent struct {
	Update app.FrontEndUpdate
}

type BufferHandler interface {
	Stop()
	Start(callback func(data BufferedEvent), expectedTurns uint32) <-chan float32
	QueueUp(update app.FrontEndUpdate) error
}
