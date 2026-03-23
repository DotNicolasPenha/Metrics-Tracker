package metricscollector

import (
	"errors"
	"time"

	"github.com/DotNicolasPenha/Metrics-Tracker/core/drivers"
)

type MetricsCollector struct {
	dbdriverhandler *drivers.DriverHandler
	last            drivers.RawStats
	lastTime        time.Time
}

func NewMetricsCollector(dbdriverhandler *drivers.DriverHandler) (*MetricsCollector, error) {
	if dbdriverhandler == nil {
		return nil, errors.New("dbdriverhandler is nil")
	}
	return &MetricsCollector{
		dbdriverhandler: dbdriverhandler,
	}, nil
}
