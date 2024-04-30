package selector

import (
	"context"
	"fmt"
	"psqlRecommendationsApi/internal/adapters/collector"
	"psqlRecommendationsApi/internal/model"
)

type Selector interface {
	ListTrainingMetrics(ctx context.Context, instanceName string) ([]model.TrainingMetric, error)
}

type CollectorAdapter interface {
	CollectInternalMetrics(ctx context.Context) ([]collector.InternalMetrics, error)
}

type Implementation struct {
	collector CollectorAdapter
}

func New(collector CollectorAdapter) *Implementation {
	return &Implementation{
		collector: collector,
	}
}

func (i *Implementation) ListTrainingMetrics(ctx context.Context, instanceName string) ([]model.TrainingMetric, error) {
	//todo get collector by instanceName
	var res []model.TrainingMetric

	internalMetrics, err := i.collector.CollectInternalMetrics(ctx)
	if err != nil {
		return nil, fmt.Errorf("collector.CollectInternalMetrics: %w", err)
	}
	for _, internalMetric := range internalMetrics {
		res = append(res, model.TrainingMetric{Value: float32(internalMetric.Value)})
	}

	return res, nil
}
