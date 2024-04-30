package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"psqlRecommendationsApi/lib/grpc_server"
	desc "psqlRecommendationsApi/pkg/environment"
)

func RunGRPCServer(implementation desc.EnvironmentServer, cfg *grpc_server.GRPCConfig) (*grpc_server.GRPCServer, error) {
	grpcServer, err := grpc_server.NewGRPCServer(cfg)
	if err != nil {
		return nil, fmt.Errorf("grpc_server.NewGRPCServer: %w", err)
	}

	desc.RegisterEnvironmentServer(grpcServer.Ser, implementation)
	grpcServer.Run()

	log.Printf("started grpc server at %s:%s", cfg.Host, strconv.Itoa(cfg.Port))
	return grpcServer, nil
}

func Lock(ch chan os.Signal) {
	defer func() {
		ch <- os.Interrupt
	}()
	signal.Notify(ch,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-ch
}
