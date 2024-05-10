package discovery

import (
	"context"
	"psqlRecommendationsApi/cmd/clients"
	model "psqlRecommendationsApi/internal/model/discovery"
	desc "psqlRecommendationsApi/pkg/discovery"
)

type Adapter interface {
	RegisterInstance(ctx context.Context, instanceName string, config []byte) (model.CollectorInstance, error)
	GetCollector(ctx context.Context, instanceName string) (model.CollectorInstance, error)
}

type Implementation struct {
	client *clients.DiscoveryClient
}

func New(client *clients.DiscoveryClient) *Implementation {
	return &Implementation{
		client: client,
	}
}

func (i *Implementation) RegisterInstance(ctx context.Context, instanceName string, config []byte) (model.CollectorInstance, error) {
	instance, err := i.client.Client.RegisterInstance(ctx, &desc.RegisterInstanceRequest{
		InstanceName: instanceName,
		Config:       config,
	})
	if err != nil {
		return model.CollectorInstance{}, err
	}
	return model.CollectorInstance{
		Id:   instance.Id,
		Name: instance.InstanceName,
		Host: instance.Host,
		Port: int(instance.Port),
	}, nil
}

func (i *Implementation) GetCollector(ctx context.Context, instanceName string) (model.CollectorInstance, error) {
	return model.CollectorInstance{}, nil
}
