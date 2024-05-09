package docker

import (
	"context"
	"psqlRecommendationsApi/cmd/clients"
	model "psqlRecommendationsApi/internal/model/discovery"
)

const imageName = ""

type Adapter interface {
	CreateInstance(ctx context.Context, instanceName, dbDsn string) (model.CollectorInstance, error)
}

type Implementation struct {
	dockerClient *clients.DockerClient
}

func New(dockerClient *clients.DockerClient) *Implementation {
	return &Implementation{
		dockerClient: dockerClient,
	}
}

func (i *Implementation) CreateInstance(ctx context.Context, instanceName, dbDsn string) (model.CollectorInstance, error) {
	return model.CollectorInstance{}, nil
}
