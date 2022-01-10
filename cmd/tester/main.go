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

func initTheServer(srv *server.Broadcaster) chan bool {

	grpcServer := grpc.NewServer()

	lugo.RegisterBroadcastServer(grpcServer, srv)
	lugo.RegisterRemoteServer(grpcServer, srv)

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
	return running
}

func main() {
	rootCmd := cobra.Command{}
	srv := server.NewServer(zapLog)
	cms := []struct {
		command string
		sample  samples.Sample
	}{
		{command: "initial", sample: samples.SampleServerIsUp()},
		{command: "players_connect", sample: samples.SamplePlayersConnect()},
		{command: "move_ball", sample: samples.SampleMoveBall()},
		{command: "move_player", sample: samples.SampleMovePlayers()},
		{command: "rotate_player", sample: samples.SampleRotatePlayers()},
		{command: "score_time", sample: samples.SampleScoreTime()},
		{command: "game_over", sample: samples.SampleGameOver()},
		{command: "buffer", sample: samples.SampleBuffering()},
	}

	for _, opt := range cms {
		ddd := opt
		cmd := &cobra.Command{
			Use: ddd.command,
			Run: func(cmd *cobra.Command, args []string) {
				srv.EventQueue = ddd.sample.Events
				srv.Setup = ddd.sample.Setup
				srv.Setup.DevMode = true
				<-initTheServer(srv)
			},
		}
		cmd.Flags().BoolP("dev-mode", "d", false, "Start on dev mode")
		rootCmd.AddCommand(cmd)
	}

	if err := rootCmd.Execute(); err != nil {
		zapLog.With(zap.Error(err)).Fatalf("failure executing arguments")
	}
}
