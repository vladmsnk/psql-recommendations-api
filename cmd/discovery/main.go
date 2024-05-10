package main

import (
	"log"

	"psqlRecommendationsApi/cmd"
	"psqlRecommendationsApi/cmd/clients"
	"psqlRecommendationsApi/internal/adapters/connections"
	"psqlRecommendationsApi/internal/adapters/docker"
	"psqlRecommendationsApi/internal/adapters/instance_storage"
	"psqlRecommendationsApi/internal/app/discovery"
	"psqlRecommendationsApi/internal/config"
	"psqlRecommendationsApi/internal/usecase/discovery/registrator"
)

func main() {
	redisClient, err := clients.NewRedisClient(config.ConfigStruct.Redis)
	if err != nil {
		log.Fatal(err)
	}
	defer redisClient.Close()

	dockerClient, err := clients.NewDockerClient()
	if err != nil {
		log.Fatal(err)
	}

	var (
		storage            = instance_storage.New(redisClient)
		instanceCreator    = docker.New(dockerClient)
		connectionProvider = connections.New()
		reg                = registrator.New(storage, instanceCreator, connectionProvider)
		app                = discovery.New(reg)
	)

	grpcServer, err := cmd.RunDiscoveryRPCServer(app, config.ConfigStruct.GRPC)
	if err != nil {
		log.Fatal(err)
	}

	defer grpcServer.Close()
}
