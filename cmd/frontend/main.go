package main

import (
	"bitbucket.org/makeitplay/lugo-frontend/web/app"
	"bitbucket.org/makeitplay/lugo-frontend/web/app/broker"
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

var zapLog *zap.SugaredLogger
var internalDebug = false

func init() {
	var err error
	var logger *zap.Logger
	var configZap zap.Config

	if internalDebug {
		configZap = zap.NewDevelopmentConfig()
		configZap.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		configZap = zap.NewProductionConfig()
		configZap.DisableCaller = true
	}

	logger, err = configZap.Build()
	if err != nil {
		log.Fatalf("could not initiliase looger: %s", err)
	}
	zapLog = logger.Sugar()
}

func main() {

	gameConfig := app.Configuration{
		Broadcast: app.BroadcastConfig{
			Address:  "localhost:9090",
			Insecure: true,
		},
		DevMode:           false,
		StartMode:         "",
		TimeRemaining:     "",
		GameDuration:      0,
		ListeningDuration: 0,
		HomeTeam:          app.TeamConfiguration{},
		AwayTeam:          app.TeamConfiguration{},
	}

	eventBroker := broker.NewBinder(gameConfig, zapLog)
	server := app.Newhandler("/home/rubens/go/src/bitbucket.org/makeitplay/lugo-frontend/web", "local", eventBroker)
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: server,
	}
	var once sync.Once

	serviceGroup := sync.WaitGroup{}
	firstStop := ""
	running := make(chan error)

	go func() {
		serviceGroup.Add(1)
		defer serviceGroup.Done()
		err := eventBroker.ListenAndBroadcast()
		zapLog.Errorf("event broker has stopped: %s", err)

		once.Do(func() {
			firstStop = "event-broker"
			close(running)
		})
	}()

	go func() {
		serviceGroup.Add(1)
		defer serviceGroup.Done()
		err := httpServer.ListenAndServeTLS("testdata/server.pem", "testdata/server.key")
		zapLog.Errorf("https has stopped: %s", err)

		once.Do(func() {
			firstStop = "http"
			close(running)
		})
	}()

	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt)
		<-signalChan
		once.Do(func() {
			zapLog.Info("interruption signal sent")
			close(running)
		})
	}()

	<-running

	if firstStop != "event-broker" {
		zapLog.Info("stopping event broker")
		if err := eventBroker.Stop(); err != nil {
			zapLog.Errorf("error stopping event broker: %s", err)
		}
	}

	if firstStop != "http" {
		zapLog.Info("stopping http server")
		ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
		if err := httpServer.Shutdown(ctx); err != nil {
			zapLog.Errorf("error stopping event broker: %s", err)
		}
	}
	serviceGroup.Wait()
}
