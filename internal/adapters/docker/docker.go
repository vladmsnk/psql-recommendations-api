package docker

import (
	"context"
	model "psqlRecommendationsApi/internal/model/discovery"
)

type Adapter interface {
	CreateInstance(ctx context.Context, instanceName, dbDsn string) (model.CollectorInstance, error)
}

type Implementation struct {
}

func New() *Implementation {
	return &Implementation{}
}

func (i *Implementation) CreateInstance(ctx context.Context, instanceName, dbDsn string) (model.CollectorInstance, error) {
	return model.CollectorInstance{}, nil
}
