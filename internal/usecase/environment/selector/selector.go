package selector

import (
	"context"
	"fmt"
	"psqlRecommendationsApi/internal/adapters/collector"
	"psqlRecommendationsApi/internal/model"
)

type Selector interface {
	ListTrainingMetrics(ctx context.Context, instanceName string) ([]model.TrainingMetric, error)
	ListRewardMetrics(ctx context.Context, instanceName string) (model.ExternalMetrics, error)
	ListKnobs(ctx context.Context, instanceName string) ([]model.Knob, error)
}

type CollectorAdapter interface {
	CollectInternalMetrics(ctx context.Context) ([]collector.InternalMetrics, error)
	CollectKnobs(ctx context.Context) ([]collector.Knob, error)
	CollectExternalMetrics(ctx context.Context) (collector.ExternalMetrics, error)
}

// connection storage

type Implementation struct {
	collector CollectorAdapter
}

func New(collector CollectorAdapter) *Implementation {
	return &Implementation{
		collector: collector,
	}
}

func (i *Implementation) ListRewardMetrics(ctx context.Context, instanceName string) (model.ExternalMetrics, error) {
	metrics, err := i.collector.CollectExternalMetrics(ctx)
	if err != nil {
		return model.ExternalMetrics{}, fmt.Errorf("collector.CollectExternalMetrics: %w", err)
	}

	return model.ExternalMetrics{Tps: metrics.Tps, Latency: metrics.Latency}, nil
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

func (i *Implementation) ListKnobs(ctx context.Context, instanceName string) ([]model.Knob, error) {
	var res []model.Knob

	knobs, err := i.collector.CollectKnobs(ctx)
	if err != nil {
		return nil, fmt.Errorf("collector.CollectKnobs: %w", err)
	}
	for _, knob := range knobs {
		res = append(res, model.Knob{Name: knob.Name, Value: knob.Value, MinVal: knob.MinVal, MaxVal: knob.MaxVal})
	}

	return res, nil
}
