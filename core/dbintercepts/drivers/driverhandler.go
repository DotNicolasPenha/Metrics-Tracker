package drivers

import (
	"errors"

	"github.com/DotNicolasPenha/Metrics-Tracker/core/databases"
	"github.com/google/uuid"
)

type DriverActions interface {
	GetUsedDisk() (int64, error)
	GetActiveQueries() (int64, error)
	EstimateUsedRAM() (int64, error)
	Connect(url string) error
	BlockClient() error
	BlockAll() error
	Close() error
}

type DriverHandler struct {
	Driver    []Driver
	dbHandler *databases.DataBaseHandler
}

type Driver struct {
	typedb  string
	id      uuid.UUID
	name    string
	Actions DriverActions
}

func NewDriverHandler(dbhandler *databases.DataBaseHandler) (*DriverHandler, error) {
	if dbhandler == nil {
		return nil, errors.New("dbhandler is nil")
	}
	return &DriverHandler{
		Driver:    []Driver{},
		dbHandler: dbhandler,
	}, nil
}
