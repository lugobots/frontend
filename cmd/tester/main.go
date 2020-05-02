package main

import (
	"bitbucket.org/makeitplay/lugo-frontend/cmd/tester/samples"
	"bitbucket.org/makeitplay/lugo-frontend/cmd/tester/server"
	"fmt"
	"github.com/lugobots/lugo4go/v2/lugo"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"log"
	"net"
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

func MyCommand() *cobra.Command {
	srv := server.NewServer(zapLog)
	grpcServer := grpc.NewServer()

	lugo.RegisterBroadcastServer(grpcServer, srv)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9090))
	if err != nil {
		panic("did not start to listen at configured port")
	}
	running := make(chan bool)

	go func() {
		zapLog.Infof("running gRPC at port %d", 9090)
		// Service 2 - gRPC server
		if err := grpcServer.Serve(lis); err != nil {
			zapLog.With(zap.Error(err)).Fatalf("gRPC server stopped to serve")
		}
		close(running)
	}()

	return &cobra.Command{
		Use:   "initial",
		Short: "From waiting until first turn",
		Run: func(cmd *cobra.Command, args []string) {
			srv.EventQueue = samples.AllPlayersConnect()
			//max := 0
			//zapLog.Info("starting sender")
			//for {
			//	time.Sleep(1 * time.Second)
			//	zapLog.Info("sending event")
			//
			//	max++
			//	if max > 500 {
			//		break
			//	}
			//}
			<-running
		},
	}
}

func main() {
	rootCmd := cobra.Command{}

	rootCmd.AddCommand(MyCommand())

	if err := rootCmd.Execute(); err != nil {
		zapLog.With(zap.Error(err)).Fatalf("failure executing arguments")
	}
}
