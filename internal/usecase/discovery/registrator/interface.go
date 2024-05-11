package registrator

import (
	"context"
	model "psqlRecommendationsApi/internal/model/discovery"
)

type Registrator interface {
	RegisterInstance(ctx context.Context, instanceName, dbDsn string) (model.CollectorInstance, error)
}

type InstanceInfoGetter interface {
	GetInstanceInfo(ctx context.Context, instanceName string) (model.CollectorInstance, error)
}
