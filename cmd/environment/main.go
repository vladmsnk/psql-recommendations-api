package main

import (
	"log"
	"os"

	"psqlRecommendationsApi/cmd"
	"psqlRecommendationsApi/cmd/clients"
	"psqlRecommendationsApi/internal/adapters/collector"
	"psqlRecommendationsApi/internal/app/environment"
	"psqlRecommendationsApi/internal/config"
	"psqlRecommendationsApi/internal/usecase/environment/selector"
	"psqlRecommendationsApi/internal/usecase/environment/setter"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatal(err)
	}
	collectorClient, err := clients.NewCollectorClient(config.ConfigStruct.Collector)
	if err != nil {
		log.Fatal(err)
	}
	defer collectorClient.Close()

	redisClient, err := clients.NewRedisClient(config.ConfigStruct.Redis)
	if err != nil {
		log.Fatal(err)
	}
	defer redisClient.Close()

	var (
		collectorAdapter = collector.New(collectorClient)
		metricsSelector  = selector.New(collectorAdapter)
		metricsSetter    = setter.New(collectorAdapter)
		app              = environment.New(metricsSelector, metricsSetter)
	)

	grpcServer, err := cmd.RunEnvironmentGRPCServer(app, config.ConfigStruct.GRPC)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcServer.Close()

	cmd.Lock(make(chan os.Signal, 1))
}
