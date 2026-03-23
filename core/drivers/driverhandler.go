package drivers

import (
	"errors"

	"github.com/DotNicolasPenha/Metrics-Tracker/core/databases"
	"github.com/google/uuid"
)

type DriverActions interface {
	Connect(url string) error
	Close() error
	GetStats() (Stats, error)
	Query(query string, args ...any) (any, error)
	Exec(query string, args ...any) error
}
type Stats struct {
	QPS               int64
	ActiveConnections int64
}
type RawStats struct {
	TotalQueries      int64
	ActiveConnections int64
}

type DriverHandler struct {
	Drivers   []Driver
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
		Drivers:   []Driver{},
		dbHandler: dbhandler,
	}, nil
}
