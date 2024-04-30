package environment

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"psqlRecommendationsApi/internal/model"
	desc "psqlRecommendationsApi/pkg/environment"
)

type Delivery struct {
	desc.EnvironmentServer
	selector Selector
}

func New(selector Selector) *Delivery {
	return &Delivery{
		selector: selector,
	}
}

type Selector interface {
	ListTrainingMetrics(ctx context.Context, instanceName string) ([]model.TrainingMetric, error)
}

func (d *Delivery) GetStates(ctx context.Context, req *desc.GetStatesRequest) (*desc.GetStatesResponse, error) {
	instanceName := req.GetInstanceName()
	if instanceName == "" {
		return nil, status.Error(codes.InvalidArgument, "instance_name should not be empty")
	}

	metrics, err := d.selector.ListTrainingMetrics(ctx, instanceName)
	if err != nil {
		return nil, fmt.Errorf("selector.ListTrainingMetrics: %w", err)
	}

	descMetrics := lo.Map(metrics, func(metric model.TrainingMetric, _ int) float32 {
		return metric.Value
	})

	return &desc.GetStatesResponse{Metrics: descMetrics}, nil
}
