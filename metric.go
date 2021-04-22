package monitoring

import (
	"encoding/json"
)

type MetricType rune

const (
	MetricTypeGauge = 'G' // float64 gauge
	MetricTypeGaugeInt = 'g' // int64 gauge
	MetricTypeCounter = 'C' // int64 counter
)

var MetricGauge = MetricType('G')


type Metric interface {
	Type() MetricType
	json.Marshaler
	Float64() float64
}

type MetricGaugeFloat float64

func (f MetricGaugeFloat) Type() MetricType {
	return MetricType('G')
}

func (f MetricGaugeFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(f)
}
func (f MetricGaugeFloat) Float64() float64 {
	return float64(f)
}


func NewMetricGauge(f float64) Metric {
	return MetricGaugeFloat(f)
}

