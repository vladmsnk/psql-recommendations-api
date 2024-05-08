package registrator

import (
	"context"
	"errors"
	"fmt"
	"psqlRecommendationsApi/cmd/clients"
	"psqlRecommendationsApi/internal/adapters/connections"
	model "psqlRecommendationsApi/internal/model/discovery"
)

type Registrator interface {
	RegisterInstance(ctx context.Context, instanceName, dbDsn string) (model.CollectorInstance, error)
}

type InstanceGetter interface {
	GetInstance(ctx context.Context, instanceName string) (model.CollectorInstance, error)
}

type Storage interface {
	CheckInstanceExists(ctx context.Context, instanceName string) (bool, error)
	SaveInstance(ctx context.Context, instance model.CollectorInstance) error
	GetInstance(ctx context.Context, instanceName string) (model.CollectorInstance, error)
}

type ConnectionProvider interface {
	GetConnection(ctx context.Context, instanceName string) (*clients.CollectorClient, error)
	SetConnection(_ context.Context, instanceName, host string, port int) error
}

type InstanceCreator interface {
	CreateInstance(ctx context.Context, instanceName, dbDsn string) (model.CollectorInstance, error)
}

type Implementation struct {
	storage            Storage
	instanceCreator    InstanceCreator
	connectionProvider ConnectionProvider
}

func New(storage Storage, instanceCreator InstanceCreator, connectionProvider ConnectionProvider) *Implementation {
	return &Implementation{
		storage:            storage,
		instanceCreator:    instanceCreator,
		connectionProvider: connectionProvider,
	}
}

func (i *Implementation) RegisterInstance(ctx context.Context, instanceName, dbDsn string) (model.CollectorInstance, error) {
	exists, err := i.storage.CheckInstanceExists(ctx, instanceName)
	if err != nil {
		return model.CollectorInstance{}, fmt.Errorf("storage.CheckInstanceExists: %w", err)
	}
	if exists {
		return model.CollectorInstance{}, model.ErrInstanceAlreadyExists
	}

	instance, err := i.instanceCreator.CreateInstance(ctx, instanceName, dbDsn)
	if err != nil {
		return model.CollectorInstance{}, fmt.Errorf("instanceCreator.CreateInstance: %w", err)
	}

	err = i.connectionProvider.SetConnection(ctx, instanceName, instance.Name, instance.Port)
	if err != nil {
		if errors.Is(err, connections.ErrConnectionAlreadySet) {
		}

		return model.CollectorInstance{}, fmt.Errorf("connectionProvider.SetConnection: %w", err)
	}

	err = i.storage.SaveInstance(ctx, instance)
	if err != nil {
		return model.CollectorInstance{}, fmt.Errorf("storage.SaveInstance: %w", err)
	}

	return instance, nil
}

func (i *Implementation) GetInstance(ctx context.Context, instanceName string) (model.CollectorInstance, error) {
	instance, err := i.storage.GetInstance(ctx, instanceName)
	if err != nil {
		return model.CollectorInstance{}, fmt.Errorf("i.storage.GetInstance: %w", err)
	}

	return instance, nil
}
