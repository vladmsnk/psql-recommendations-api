package instance_storage

import (
	"context"

	"psqlRecommendationsApi/cmd/clients"
	model "psqlRecommendationsApi/internal/model/discovery"
)

type Adapter interface {
	CheckInstanceExists(ctx context.Context, instanceName string) (bool, error)
	SaveInstance(ctx context.Context, instance model.CollectorInstance) error
	GetInstance(ctx context.Context, instanceName string) (model.CollectorInstance, error)
}

type Implementation struct {
	redisClient *clients.RedisClient
}

func New(redisClient *clients.RedisClient) *Implementation {
	return &Implementation{
		redisClient: redisClient,
	}
}

func (i *Implementation) CheckInstanceExists(ctx context.Context, instanceName string) (bool, error) {
	return true, nil
}

func (i Implementation) SaveInstance(ctx context.Context, instance model.CollectorInstance) error {
	return nil
}

func (i *Implementation) GetInstance(ctx context.Context, instanceName string) (model.CollectorInstance, error) {
	return model.CollectorInstance{}, nil
}
