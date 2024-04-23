package model

type TrainingMetric struct {
	Value float32
}

type InternalMetrics struct {
	Name  string
	Scope string
	Value interface{}
}

type ExternalMetrics struct {
	Tps     float64
	Latency float64
}
