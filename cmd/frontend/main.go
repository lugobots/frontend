package main

import (
	"context"
	"fmt"
	"github.com/lugobots/frontend/web/app"
	"github.com/lugobots/frontend/web/app/broker"
	"github.com/lugobots/frontend/web/app/propagator"
	"github.com/lugobots/lugo4go/v2/proto"
	"github.com/pkg/errors"
	"github.com/rubens21/srvmgr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
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

	eventBroker := broker.NewBinder(app.Config{
		GRPCAddress:         "localhost:9090",
		GRPCInsecure:        true,
		StaysIfDisconnected: true,
	}, zapLog.Named("broker"))

	broadcasterSyncer := propagator.NewPropagator(eventBroker)

	grpcServer := grpc.NewServer()
	proto.RegisterBroadcastServer(grpcServer, broadcasterSyncer)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5001))
	if err != nil {
		zapLog.Errorf("failed on listen grpc port: %s", err)
		return
	}

	server := app.NewHandler("", "local", eventBroker)
	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: server,
	}

	serviceManager := srvmgr.NewManager(zapLog, 10*time.Second)

	serviceManager.AddTask(eventBroker)
	serviceManager.AddTask(srvmgr.GrpcServerAsTask("grpc-server", grpcServer, lis))
	serviceManager.AddTask(srvmgr.MakeTask(
		"http-server",
		func() error {
			return httpServer.ListenAndServe()
		},
		func(ctx context.Context) error {
			if err := httpServer.Shutdown(ctx); err != nil {
				return errors.Wrap(err, "error stopping http server")
			}
			return nil
		},
	))

	if err := serviceManager.Run(); err != nil {
		zapLog.Errorf(err.Error())
	}
}
