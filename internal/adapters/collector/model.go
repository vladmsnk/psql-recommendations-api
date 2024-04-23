package collector

type ExternalMetrics struct {
	Latency float64
	Tps     float64
}

type InternalMetrics struct {
	Name  string
	Value float64
}
