package core

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type DBDriverActions interface {
	Connect(url string) error
	GetUsedDisk() (int64, error)
	GetActiveQueries() (int64, error)
	EstimateUsedRAM() (int64, error)
	GetTopQueries(limit int) ([]string, error)
	Close() error
}

type DBDriverHandler struct {
	DBDrivers []DBDriver
	dbHandler *DataBaseHandler
}

type DBDriver struct {
	typedb  string
	id      uuid.UUID
	name    string
	Actions DBDriverActions
}

func NewDBDriverHandler(dbhandler *DataBaseHandler) (*DBDriverHandler, error) {
	if dbhandler == nil {
		return nil, errors.New("dbhandler is nil")
	}
	return &DBDriverHandler{
		DBDrivers: []DBDriver{},
		dbHandler: dbhandler,
	}, nil
}
func (dh *DBDriverHandler) ConnectAllDriversdb() error {
	var wg sync.WaitGroup
	var errOnce sync.Once
	var connectErr error

	for _, db := range dh.dbHandler.databases {
		driver, err := dh.findDriverByID(db.Driverid)
		if err != nil {
			continue
		}

		wg.Add(1)
		go func(db DataBase, d DBDriver) {
			defer wg.Done()

			if err := dh.ConnectDBDriver(db, d); err != nil {
				errOnce.Do(func() {
					connectErr = err
				})
			}
		}(db, *driver)
	}

	wg.Wait()
	return connectErr
}

func (d *DBDriverHandler) ConnectDBDriver(db DataBase, dbdriver DBDriver) error {
	if dbdriver.Actions == nil {
		return errors.New("dbdriver actions is nil for " + dbdriver.name)
	}
	return dbdriver.Actions.Connect(db.Url)
}

func (d *DBDriverHandler) GetDriverDBID(typedb, name string) (*uuid.UUID, error) {
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
func (d *DBDriverHandler) findDriverByID(id uuid.UUID) (*DBDriver, error) {
	for i := range d.DBDrivers {
		if d.DBDrivers[i].id == id {
			return &d.DBDrivers[i], nil
		}
	}
	return nil, errors.New("driver not found")
}

func (d *DBDriverHandler) AddDriverDB(typedb string, name string, actions DBDriverActions) (*uuid.UUID, error) {
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
		Actions: actions,
	}

	d.DBDrivers = append(d.DBDrivers, driver)
	return &newID, nil
}
