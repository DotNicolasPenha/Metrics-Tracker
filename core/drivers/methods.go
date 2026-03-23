package drivers

import (
	"errors"

	"github.com/google/uuid"
)

func (d *DriverHandler) GetDriverDBID(typedb, name string) (*uuid.UUID, error) {
	if typedb == "" || name == "" {
		return nil, errors.New("typedb or name is empty")
	}

	for _, driver := range d.Drivers {
		if driver.typedb == typedb && driver.name == name {
			return &driver.id, nil
		}
	}

	return nil, errors.New("driver not found")
}
func (d *DriverHandler) FindDriverByID(id uuid.UUID) (*Driver, error) {
	for i := range d.Drivers {
		if d.Drivers[i].id == id {
			return &d.Drivers[i], nil
		}
	}
	return nil, errors.New("driver not found")
}

func (d *DriverHandler) AddDriverDB(typedb string, name string, actions DriverActions) (*uuid.UUID, error) {
	if typedb == "" || name == "" {
		return nil, errors.New("typedb or name is empty")
	}

	for _, driver := range d.Drivers {
		if driver.typedb == typedb && driver.name == name {
			return nil, errors.New("driver already exists")
		}
	}

	newID := uuid.New()

	driver := Driver{
		typedb:  typedb,
		name:    name,
		id:      newID,
		Actions: actions,
	}

	d.Drivers = append(d.Drivers, driver)
	return &newID, nil
}
