package setter

import (
	"context"
	"fmt"
	"psqlRecommendationsApi/internal/model"
)

type Setter interface {
	SetActions(ctx context.Context, actions []model.Action) error
}

type CollectorAdapter interface {
	SetKnobs(ctx context.Context, knobs []model.Knob) error
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
