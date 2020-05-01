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
		GameDuration:      0,
		ListeningDuration: 0,
		HomeTeam: app.TeamConfiguration{
			Name:   "Canada",
			Avatar: "external/profile-team-home.jpg",
			Colors: app.TeamColors{
				PrimaryColor: app.Color{
					R: 100,
					G: 146,
					B: 250,
				},
				SecondaryColor: app.Color{
					R: 240,
					G: 50,
					B: 150,
				},
			},
		},
		AwayTeam: app.TeamConfiguration{
			Name:   "Canada",
			Avatar: "external/profile-team-away.jpg",
			Colors: app.TeamColors{
				PrimaryColor: app.Color{
					R: 100,
					G: 255,
					B: 150,
				},
				SecondaryColor: app.Color{
					R: 100,
					G: 200,
					B: 50,
				},
			},
		},
	}

	eventBroker := broker.NewBinder(gameConfig, zapLog)
	server := app.Newhandler("/home/rubens/go/src/bitbucket.org/makeitplay/lugo-frontend/web", "local", eventBroker)
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: server,
	}
	var evenBrokerStopped sync.Once
	var httpStopped sync.Once
	var somethingStopped sync.Once

	serviceGroup := sync.WaitGroup{}
	running := make(chan error)

	stoppingEventBroker := func() {
		evenBrokerStopped.Do(func() {
			zapLog.Info("stopping event broker")
			if err := eventBroker.Stop(); err != nil {
				zapLog.Errorf("error stopping event broker: %s", err)
			}
		})
	}
	startingEventBroker := func() {
		serviceGroup.Add(1)
		defer serviceGroup.Done()
		zapLog.Errorf("starting http server at %s", httpServer.Addr)
		err := eventBroker.ListenAndBroadcast()
		zapLog.Errorf("event broker has stopped: %s", err)

		somethingStopped.Do(func() {
			close(running)
		})
		stoppingEventBroker()
	}

	stoppingHttpServer := func() {
		httpStopped.Do(func() {
			zapLog.Info("stopping http server")
			ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
			if err := httpServer.Shutdown(ctx); err != nil {
				zapLog.Errorf("error stopping event broker: %s", err)
			}
		})
	}

	startingHttpServer := func() {
		serviceGroup.Add(1)
		defer serviceGroup.Done()
		err := httpServer.ListenAndServe()
		zapLog.Errorf("https has stopped: %s", err)

		somethingStopped.Do(func() {
			close(running)
		})
		stoppingHttpServer()
	}

	monitorInterruptionSignal := func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt)
		<-signalChan
		somethingStopped.Do(func() {
			zapLog.Info("interruption signal sent")
			close(running)
		})
	}

	go startingEventBroker()
	go startingHttpServer()
	go monitorInterruptionSignal()

	<-running

	stoppingHttpServer()
	stoppingEventBroker()
	serviceGroup.Wait()
}
