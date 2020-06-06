package broker

import (
	"bitbucket.org/makeitplay/lugo-frontend/web/app"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBufferHandler_stageUpdates(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRateCounter := NewMockRateCounter(ctrl)
	b := BufferHandler{
		RateCounter:     mockRateCounter,
		Logger:          zapLog,
		bufferOn:        make(chan bool),
		bufferedUpdates: make(chan app.FrontEndUpdate, MaxUpdateBuffer),
		bufferStage:     make(chan app.FrontEndUpdate, MaxUpdateBuffer),
	}
	defer b.Stop()

	go b.stageUpdates()

	mockRateCounter.EXPECT().Incr(int64(1))
	assert.Len(t, b.bufferedUpdates, 0)
	assert.Len(t, b.bufferStage, 0)
	b.bufferedUpdates <- app.FrontEndUpdate{}
	time.Sleep(50 * time.Millisecond)
	assert.Len(t, b.bufferedUpdates, 0)
	assert.Len(t, b.bufferStage, 1)

	mockRateCounter.EXPECT().Incr(int64(1))
	assert.Len(t, b.bufferedUpdates, 0)
	assert.Len(t, b.bufferStage, 1)
	b.bufferedUpdates <- app.FrontEndUpdate{}
	time.Sleep(50 * time.Millisecond)
	assert.Len(t, b.bufferedUpdates, 0)
	assert.Len(t, b.bufferStage, 2)
}

//func TestBufferHandler_rateCounter(t *testing.T) {
//	c := ratecounter.NewRateCounter(time.Second)
//
//	stop := make(chan bool)
//	go func() {
//		for {
//			select {
//			case <-stop:
//			case <-time.Tick(50 * time.Millisecond):
//				//log.Printf("counted")
//				c.Incr(1)
//			}
//		}
//	}()
//	time.Sleep(3 * time.Second)
//	assert.Equal(t, int64(20), c.Rate())
//	time.Sleep(time.Second)
//	assert.Equal(t, int64(20), c.Rate())
//	time.Sleep(time.Second)
//	assert.Equal(t, int64(20), c.Rate())
//	close(stop)
//}
func TestBufferHandler_streamBuffer(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRateCounter := NewMockRateCounter(ctrl)
	b := BufferHandler{
		RateCounter:      mockRateCounter,
		Logger:           zapLog,
		bufferOn:         make(chan bool),
		bufferedUpdates:  make(chan app.FrontEndUpdate, MaxUpdateBuffer),
		bufferStage:      make(chan app.FrontEndUpdate, MaxUpdateBuffer),
		lastReceivedTurn: 100,
	}
	defer b.Stop()
	mockRateCounter.EXPECT().Rate().Return(int64(20)).AnyTimes()
	called := false
	pulse := make(chan bool)
	go b.streamBuffer(func(data BufferedEvent) {
		called = true

	}, pulse)

	assert.True(t, called)
}

func TestBufferHandler_QueueUp(t *testing.T) {
	ctrl := gomock.NewController(t)

	b := BufferHandler{
		RateCounter: NewMockRateCounter(ctrl),
		Logger:      zapLog,
	}

	called := false
	b.Start(func(data BufferedEvent) {
		called = true
	})

	err := b.QueueUp(app.FrontEndUpdate{}, 0)
	assert.Nil(t, err)
	assert.True(t, called)
}

func TestBufferHandler_Start(t *testing.T) {

}
