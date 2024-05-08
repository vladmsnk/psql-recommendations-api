package discovery

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	model "psqlRecommendationsApi/internal/model/discovery"
	desc "psqlRecommendationsApi/pkg/discovery"
)

type Registrator interface {
	RegisterInstance(ctx context.Context, instanceName, dbDsn string) (model.CollectorInstance, error)
}

type Delivery struct {
	desc.DiscoveryServer
	registrator Registrator
}

func New(registrator Registrator) *Delivery {
	return &Delivery{
		registrator: registrator,
	}
}

func (d *Delivery) RegisterInstance(ctx context.Context, req *desc.RegisterInstanceRequest) (*desc.RegisterInstanceResponse, error) {
	instanceName := req.GetInstanceName()
	if instanceName == "" {
		return nil, status.Error(codes.InvalidArgument, "instance_name should not be empty")
	}

	dbDsn := req.GetDbDsn()
	if dbDsn == "" {
		return nil, status.Error(codes.InvalidArgument, "db_dsn should not be empty")
	}

	instance, err := d.registrator.RegisterInstance(ctx, instanceName, dbDsn)
	if err != nil {
		if errors.Is(err, model.ErrInstanceAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, fmt.Errorf("registrator.RegisterInstance: %w", err)
	}

	return &desc.RegisterInstanceResponse{
		Id:           instance.Id,
		InstanceName: instance.Name,
		Host:         instance.Host,
		Port:         instance.Port,
	}, nil
}
