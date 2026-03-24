package databases

import (
	"errors"
	"path"

	"github.com/DotNicolasPenha/Metrics-Tracker/core/datawork"
	"github.com/google/uuid"
)

type DataBaseHandler struct {
	Databases        []DataBase
	databasedirpath  string
	databasejsonname string
	dwjson           datawork.DataWorkHandler
}

type DataBase struct {
	Id       uuid.UUID      `json:"id"`
	Typedb   string         `json:"typedb"`
	Name     string         `json:"name"`
	Url      string         `json:"url"`
	Limits   DataBaseLimits `json:"limits"`
	Blocks   Blocks         `json:"blocks"`
	Driverid uuid.UUID      `json:"driverid"`
	Proxyid  uuid.UUID      `json:"proxyid"`
}
type Blocks struct {
	Queries []string           `json:"queries"`
	Clients ClientsBlockLimits `json:"clients"`
}
type ClientsBlockLimits struct {
	MaxQueriesPerDay int64 `json:"mqpDay"`
	MaxQueriesPerMin int64 `json:"mpqMin"`
}

// block de database when break the limit
type DataBaseLimits struct {
	Maxconn int64 `json:"maxconn"`
	MaxQPS  int64 `json:"maxqps"`
}

func NewDataBaseHandler(mtdbpath string, namejsonfile string) (*DataBaseHandler, error) {
	if mtdbpath == "" {
		return nil, errors.New("mtdbpath param is empty")
	}
	if namejsonfile == "" {
		return nil, errors.New("namejsonfile param is empty")
	}
	dbh := &DataBaseHandler{
		databasedirpath:  mtdbpath,
		Databases:        []DataBase{},
		databasejsonname: namejsonfile,
	}
	if err := dbh.dataWorkJsonSetup(); err != nil {
		return nil, err
	}
	return dbh, nil
}
func (dbh *DataBaseHandler) dataWorkJsonSetup() error {
	dwh, err := datawork.NewDataWorkHandler(dbh.databasedirpath)
	if err != nil {
		return errors.New(err.Error())
	}
	dbh.dwjson = *dwh
	return dwh.Load([]DataBase{}, path.Join(dbh.databasedirpath, dbh.databasejsonname))
}
