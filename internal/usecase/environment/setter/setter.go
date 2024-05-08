package setter

import (
	"context"
	"fmt"
	"psqlRecommendationsApi/internal/model"
)

type Setter interface {
	SetActions(ctx context.Context, actions []model.Action) error
	InitEnvironment(ctx context.Context, instanceName string) error
}

type CollectorAdapter interface {
	SetKnobs(ctx context.Context, knobs []model.Knob) error
	InitLoad(ctx context.Context) error
}

type Implementation struct {
	collector CollectorAdapter
}

func New(collector CollectorAdapter) *Implementation {
	return &Implementation{
		collector: collector,
	}
}

func (i *Implementation) SetActions(ctx context.Context, actions []model.Action) error {
	var knobs []model.Knob
	for _, action := range actions {
		knobs = append(knobs, model.Knob{Name: action.Name, Value: action.Value})
	}

	err := i.collector.SetKnobs(ctx, knobs)
	if err != nil {
		return fmt.Errorf("setter.SetKnobs: %w", err)
	}

	return nil

}

func (i *Implementation) InitEnvironment(ctx context.Context, instanceName string) error {
	err := i.collector.InitLoad(ctx)
	if err != nil {
		return fmt.Errorf("setter.InitLoad: %w", err)
	}
	return nil
}
