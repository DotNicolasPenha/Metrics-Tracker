package core

import (
	"errors"

	"github.com/google/uuid"
)

type DBDriverMetrics interface {
	Connect() error
	GetUsedDisk() (int64, error)
	GetActiveQueries() (int64, error)
	EstimateUsedRAM() (int64, error)
	GetTopQueries(limit int) ([]string, error)
	Close() error
}
type DBDriverHandler struct {
	DBDrivers []DBDriver
}
type DBDriver struct {
	typedb  string
	id      uuid.UUID
	name    string
	Metrics *DBDriverMetrics
}

func NewDBDriverHandler() *DBDriverHandler {
	return &DBDriverHandler{
		DBDrivers: []DBDriver{},
	}
}
func (d *DBDriverHandler) GetDriverDBID(typedb string, name string) (*uuid.UUID, error) {
	if typedb == "" || name == "" {
		return nil, errors.New("typedb or name is empty")
	}

	for _, driver := range d.DBDrivers {
		if driver.typedb == typedb && driver.name == name {
			return &driver.id, nil
		}
	}

	return nil, errors.New("driver not found")
}

func (d *DBDriverHandler) AddDriverDB(typedb string, name string, metrics DBDriverMetrics) (*uuid.UUID, error) {
	if typedb == "" || name == "" {
		return nil, errors.New("typedb or name is empty")
	}

	for _, driver := range d.DBDrivers {
		if driver.typedb == typedb && driver.name == name {
			return nil, errors.New("driver already exists")
		}
	}

	newID := uuid.New()
	driver := DBDriver{
		typedb:  typedb,
		name:    name,
		id:      newID,
		Metrics: &metrics,
	}

	d.DBDrivers = append(d.DBDrivers, driver)
	return &newID, nil
}
