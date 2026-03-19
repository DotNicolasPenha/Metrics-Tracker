package core

import (
	"errors"
	"sync"
)

type MetricsCollector struct {
	dbdriverhandler *DBDriverHandler
}

func NewMetricsCollector(dbdriverhandler *DBDriverHandler) (*MetricsCollector, error) {
	if dbdriverhandler == nil {
		return nil, errors.New("dbdriverhandler is nil")
	}
	return &MetricsCollector{
		dbdriverhandler: dbdriverhandler,
	}, nil
}

func (mc *MetricsCollector) StartConnectAllDriversdb() error {
	var wg sync.WaitGroup
	var errOnce sync.Once
	var connectErr error

	for _, dbdriver := range mc.dbdriverhandler.DBDrivers {
		wg.Add(1)
		go func(d DBDriver) {
			defer wg.Done()
			if err := mc.connectDBDriver(d); err != nil {
				errOnce.Do(func() {
					connectErr = err
				})
			}
		}(dbdriver)
	}

	wg.Wait()
	return connectErr
}

func (mc *MetricsCollector) connectDBDriver(dbdriver DBDriver) error {
	if dbdriver.Actions == nil {
		return errors.New("dbdriver actions is nil for " + dbdriver.name)
	}
	return dbdriver.Actions.Connect()
}

func (mc *MetricsCollector) GetMetricsOfDBName(namedb string) (map[string]int64, error) {
	if namedb == "" {
		return nil, errors.New("namedb is empty")
	}

	for _, driver := range mc.dbdriverhandler.DBDrivers {
		if driver.name == namedb {
			if driver.Actions == nil {
				return nil, errors.New("actions not available for " + namedb)
			}

			metrics := make(map[string]int64)

			disk, _ := driver.Actions.GetUsedDisk()
			ram, _ := driver.Actions.EstimateUsedRAM()
			queries, _ := driver.Actions.GetActiveQueries()

			metrics["used_disk"] = disk
			metrics["used_ram"] = ram
			metrics["active_queries"] = queries

			return metrics, nil
		}
	}

	return nil, errors.New("driver not found for " + namedb)
}
