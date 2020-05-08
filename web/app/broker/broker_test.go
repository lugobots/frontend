package broker

import (
	"bitbucket.org/makeitplay/lugo-frontend/web/app"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/lugobots/lugo4go/v2/lugo"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"testing"
	"time"
)

var zapLog *zap.SugaredLogger

func init() {
	var err error
	var logger *zap.Logger
	var configZap zap.Config

	configZap = zap.NewDevelopmentConfig()
	configZap.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err = configZap.Build()
	if err != nil {
		log.Fatalf("could not initiliase looger: %s", err)
	}
	zapLog = logger.Sugar()
}

func TestNewBinder_IgnoreWhenThereIsNoConnections(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOnEventClient := NewMockBroadcast_OnEventClient(ctrl)
	config := app.Config{
		GameDuration:      100,
		ListeningDuration: 1 * time.Second,
		HomeTeam:          app.TeamConfiguration{},
		AwayTeam:          app.TeamConfiguration{},
	}
	binder := NewBinder(mockOnEventClient, config, zapLog)

	listenerCtx, listenerStopper := context.WithTimeout(context.Background(), 1*time.Second)

	expectedGameEvent := &lugo.GameEvent{
		GameSnapshot: &lugo.GameSnapshot{
			HomeTeam: &lugo.Team{Score: 0},
			AwayTeam: &lugo.Team{Score: 0},
			Turn:     12,
		},
		Event: &lugo.GameEvent_NewPlayer{NewPlayer: &lugo.EventNewPlayer{
			Player: &lugo.Player{},
		}},
	}
	mockOnEventClient.EXPECT().Recv().Return(expectedGameEvent, nil)
	mockOnEventClient.EXPECT().Recv().Return(nil, io.EOF)
	go func() {
		err := binder.broadcast()
		assert.NotNil(t, err)
		assert.Equal(t, app.ErrGRPCConnectionClosed, err)
		listenerStopper()
	}()
	//waiting or go routine ends
	<-listenerCtx.Done()
	err := listenerCtx.Err()
	assert.Equal(t, context.Canceled, err)
}

func TestNewBinder_SendsTheEvents(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOnEventClient := NewMockBroadcast_OnEventClient(ctrl)
	config := app.Config{
		GameDuration:      100,
		ListeningDuration: 1 * time.Second,
		HomeTeam:          app.TeamConfiguration{},
		AwayTeam:          app.TeamConfiguration{},
	}
	binder := NewBinder(config, zapLog)

	producerCtx, producerStopper := context.WithTimeout(context.Background(), 1*time.Minute)
	consumerCtx, consumerStopper := context.WithTimeout(context.Background(), 1*time.Minute)

	expectedGameEvent := &lugo.GameEvent{
		GameSnapshot: &lugo.GameSnapshot{
			HomeTeam: &lugo.Team{Score: 0},
			AwayTeam: &lugo.Team{Score: 0},
			Turn:     12,
		},
		Event: &lugo.GameEvent_NewPlayer{NewPlayer: &lugo.EventNewPlayer{
			Player: &lugo.Player{},
		}},
	}
	mockOnEventClient.EXPECT().Recv().Return(expectedGameEvent, nil)
	mockOnEventClient.EXPECT().Recv().DoAndReturn(func() {

	}).Return(nil, io.EOF)
	go func() {
		err := binder.broadcast()
		assert.NotNil(t, err)
		assert.Equal(t, app.ErrGRPCConnectionClosed, err)
		producerStopper()
	}()

	var actualGameEvent app.FrontEndUpdate
	updatesChannel := binder.StreamEventsTo("consumer-a")
	go func() {
		for producerCtx.Err() == nil {
			select {
			case e, ok := <-updatesChannel:
				if !ok {
					return
				}
				actualGameEvent = e
			}
		}
		consumerStopper()
	}()

	//waiting or go routine ends
	<-producerCtx.Done()
	err := producerCtx.Err()
	assert.Equal(t, context.Canceled, err)

	<-consumerCtx.Done()
	err = consumerCtx.Err()
	assert.Equal(t, context.Canceled, err)

	assert.Equal(t, app.EventNewPlayer, actualGameEvent.Type)
	assert.Equal(t, expectedGameEvent, actualGameEvent.Update)
}

func TestNewBinder_DropsIdleConnections(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	oldValue := maxIgnoredMessaged
	maxIgnoredMessaged = 3
	defer func() {
		maxIgnoredMessaged = oldValue
	}()

	mockOnEventClient := NewMockBroadcast_OnEventClient(ctrl)
	config := app.Config{
		GameDuration:      100,
		ListeningDuration: 1 * time.Second,
		HomeTeam:          app.TeamConfiguration{},
		AwayTeam:          app.TeamConfiguration{},
	}
	binder := NewBinder(config, zapLog)

	producerCtx, producerStopper := context.WithTimeout(context.Background(), 1*time.Minute)

	expectedGameEvent := &lugo.GameEvent{
		GameSnapshot: &lugo.GameSnapshot{
			HomeTeam: &lugo.Team{Score: 0},
			AwayTeam: &lugo.Team{Score: 0},
			Turn:     12,
		},
		Event: &lugo.GameEvent_NewPlayer{NewPlayer: &lugo.EventNewPlayer{
			Player: &lugo.Player{},
		}},
	}

	checkBeforeDropping := func() {
		assert.Len(t, binder.consumers, 1)
	}
	mockOnEventClient.EXPECT().Recv().DoAndReturn(checkBeforeDropping).Return(expectedGameEvent, nil)
	mockOnEventClient.EXPECT().Recv().DoAndReturn(checkBeforeDropping).Return(expectedGameEvent, nil)
	mockOnEventClient.EXPECT().Recv().DoAndReturn(checkBeforeDropping).Return(expectedGameEvent, nil)
	mockOnEventClient.EXPECT().Recv().DoAndReturn(checkBeforeDropping).Return(expectedGameEvent, nil)
	mockOnEventClient.EXPECT().Recv().DoAndReturn(func() {
		assert.Len(t, binder.consumers, 0)
	}).Return(expectedGameEvent, nil)
	mockOnEventClient.EXPECT().Recv().DoAndReturn(func() {

	}).Return(nil, io.EOF)
	// let ignore
	_ = binder.StreamEventsTo("consumer-a")

	go func() {
		err := binder.broadcast()
		assert.NotNil(t, err)
		assert.Equal(t, app.ErrGRPCConnectionClosed, err)
		producerStopper()
	}()

	//waiting or go routine ends
	<-producerCtx.Done()
	err := producerCtx.Err()
	assert.Equal(t, context.Canceled, err)
}
