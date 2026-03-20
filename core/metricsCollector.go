package core

import (
	"errors"
)

type MetricsCollector struct {
	dbdriverhandler *DBDriverHandler
	dbhandler       *DataBaseHandler
}

func NewMetricsCollector(dbdriverhandler *DBDriverHandler, dbhandler *DataBaseHandler) (*MetricsCollector, error) {
	if dbdriverhandler == nil {
		return nil, errors.New("dbdriverhandler is nil")
	}
	if dbhandler == nil {
		return nil, errors.New("dbhandler is nil")
	}
	return &MetricsCollector{
		dbdriverhandler: dbdriverhandler,
		dbhandler:       dbhandler,
	}, nil
}

func (mc *MetricsCollector) GetMetricsOfdbByName(namedb string) (map[string]int64, error) {
	if namedb == "" {
		return nil, errors.New("namedb is empty")
	}

	db, err := mc.dbhandler.FindDataBaseByName(namedb)
	if err != nil {
		return nil, err
	}

	driver, err := mc.dbdriverhandler.findDriverByID(db.Driverid)
	if err != nil {
		return nil, err
	}

	if driver.Actions == nil {
		return nil, errors.New("actions not available for " + namedb)
	}

	if err := driver.Actions.Connect(db.Url); err != nil {
		return nil, err
	}

	metrics := make(map[string]int64)

	disk, err := driver.Actions.GetUsedDisk()
	if err != nil {
		return nil, err
	}

	ram, err := driver.Actions.EstimateUsedRAM()
	if err != nil {
		return nil, err
	}

	queries, err := driver.Actions.GetActiveQueries()
	if err != nil {
		return nil, err
	}

	metrics["used_disk"] = disk
	metrics["used_ram"] = ram
	metrics["active_queries"] = queries

	return metrics, nil
}
