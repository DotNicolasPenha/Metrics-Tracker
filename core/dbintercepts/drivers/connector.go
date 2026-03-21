package drivers

import (
	"errors"
	"sync"

	"github.com/DotNicolasPenha/Metrics-Tracker/core/databases"
)

func (dh *DriverHandler) ConnectAllDriverdb() error {
	var wg sync.WaitGroup
	var errOnce sync.Once
	var connectErr error

	for _, db := range dh.dbHandler.Databases {
		driver, err := dh.findDriverByID(db.Driverid)
		if err != nil {
			continue
		}

		wg.Add(1)
		go func(db databases.DataBase, d Driver) {
			defer wg.Done()

			if err := dh.ConnectDriver(db, d); err != nil {
				errOnce.Do(func() {
					connectErr = err
				})
			}
		}(db, *driver)
	}

	wg.Wait()
	return connectErr
}

func (d *DriverHandler) ConnectDriver(db databases.DataBase, dbdriver Driver) error {
	if dbdriver.Actions == nil {
		return errors.New("dbdriver actions is nil for " + dbdriver.name)
	}
	return dbdriver.Actions.Connect(db.Url)
}
