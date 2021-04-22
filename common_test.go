package monitoring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMetrics(t *testing.T) {
	z := NewMetricGauge(3.14)
	assert.Equal(t,3.14,z.Float64())
	assert.Equal(t,'G',int32(z.Type()))

}
