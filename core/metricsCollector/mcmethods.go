package metricscollector

import (
	"errors"
	"time"

	"github.com/DotNicolasPenha/Metrics-Tracker/core/drivers"
)

func (mc *MetricsCollector) GetMetricsFromRawStats(rawStats *drivers.RawStats) (*drivers.Stats, error) {
	if rawStats == nil {
		return nil, errors.New("rawstats is null in GetMetricsFromRawStats")
	}
	now := time.Now()
	if mc.lastTime.IsZero() {
		mc.last = *rawStats
		mc.lastTime = now

		return &drivers.Stats{
			QPS:               0,
			ActiveConnections: rawStats.ActiveConnections,
		}, nil
	}
	deltaTime := now.Sub(mc.lastTime).Seconds()
	if deltaTime <= 0 {
		deltaTime = 1
	}
	qps := float64(rawStats.TotalQueries-mc.last.TotalQueries) / deltaTime
	mc.last = *rawStats
	mc.lastTime = now
	return &drivers.Stats{
		QPS:               int64(qps),
		ActiveConnections: rawStats.ActiveConnections,
	}, nil
}
