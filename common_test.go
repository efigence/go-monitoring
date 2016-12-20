package monitoring

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMetrics(t *testing.T) {
	z := NewMetricGauge(3.14)
	Convey("MetricGauge", t, func() {
		So(z, ShouldEqual, 3.14)
		So(z.Type(), ShouldEqual, 'G')
	})
}
