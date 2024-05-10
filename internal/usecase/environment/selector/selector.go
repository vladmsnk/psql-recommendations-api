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

type Discovery interface {
	GetCollector(ctx context.Context, instanceName string) (collector.Adapter, error)
}

type Implementation struct {
	discovery Discovery
}

func New(discovery Discovery) *Implementation {
	return &Implementation{
		discovery: discovery,
	}
}

func (i *Implementation) ListRewardMetrics(ctx context.Context, instanceName string) (model.ExternalMetrics, error) {
	collectorAdapter, err := i.discovery.GetCollector(ctx, instanceName)
	if err != nil {
		return model.ExternalMetrics{}, fmt.Errorf("discovery.GetCollector instance_name=%s: %w", instanceName, err)
	}

	metrics, err := collectorAdapter.CollectExternalMetrics(ctx)
	if err != nil {
		return model.ExternalMetrics{}, fmt.Errorf("collector.CollectExternalMetrics: %w", err)
	}

	return model.ExternalMetrics{Tps: metrics.Tps, Latency: metrics.Latency}, nil
}

func (i *Implementation) ListTrainingMetrics(ctx context.Context, instanceName string) ([]model.TrainingMetric, error) {
	var res []model.TrainingMetric

	collectorAdapter, err := i.discovery.GetCollector(ctx, instanceName)
	if err != nil {
		return nil, fmt.Errorf("discovery.GetCollector instance_name=%s: %w", instanceName, err)
	}

	internalMetrics, err := collectorAdapter.CollectInternalMetrics(ctx)
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

	collectorAdapter, err := i.discovery.GetCollector(ctx, instanceName)
	if err != nil {
		return nil, fmt.Errorf("discovery.GetCollector instance_name=%s: %w", instanceName, err)
	}

	knobs, err := collectorAdapter.CollectKnobs(ctx)
	if err != nil {
		return nil, fmt.Errorf("collector.CollectKnobs: %w", err)
	}
	for _, knob := range knobs {
		res = append(res, model.Knob{Name: knob.Name, Value: knob.Value, MinVal: knob.MinVal, MaxVal: knob.MaxVal})
	}

	return res, nil
}
