package databases

import (
	"errors"
	"path/filepath"

	"github.com/google/uuid"
)

type DataBaseHandler struct {
	databaseJsonPath string
	Databases        []DataBase
}

type DataBase struct {
	Id       uuid.UUID      `json:"id"`
	Typedb   string         `json:"typedb"`
	Name     string         `json:"name"`
	Url      string         `json:"url"`
	Limits   DataBaseLimits `json:"limits"`
	Driverid uuid.UUID      `json:"driverid"`
}

type DataBaseLimits struct {
	Maxconn        int64 `json:"maxconn"`
	Maxqueriesconn int64 `json:"maxqueriesconn"`
	MaxcpuPerconn  int64 `json:"maxcpuperconn"`
	MaxramPerconn  int64 `json:"maxramperconn"`
	MaxdiskPerconn int64 `json:"maxdiskperconn"`
	Maxram         int64 `json:"maxram"`
	Maxcpu         int64 `json:"maxcpu"`
	Maxdisk        int64 `json:"maxdisk"`
}

func NewDataBaseHandler(metrackerdirpath string) (*DataBaseHandler, error) {
	if metrackerdirpath == "" {
		return nil, errors.New("metrackerdirpath param is empty")
	}

	path := filepath.Join(metrackerdirpath, "databases.json")

	dbh := &DataBaseHandler{
		databaseJsonPath: path,
		Databases:        []DataBase{},
	}

	if err := dbh.load(); err != nil {
		return nil, err
	}

	return dbh, nil
}
