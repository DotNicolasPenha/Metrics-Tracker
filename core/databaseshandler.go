package core

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type DataBaseHandler struct {
	databaseJsonPath string
	databases        []DataBase
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
	Maxconn int64 `json:"maxconn"`
	Maxram  int64 `json:"maxram"`
	Maxcpu  int64 `json:"maxcpu"`
	Maxdisk int64 `json:"maxdisk"`
}

func NewDataBaseHandler(metrackerdirpath string) (*DataBaseHandler, error) {
	if metrackerdirpath == "" {
		return nil, errors.New("metrackerdirpath param is empty.")
	}

	databaseJsonPath := filepath.Join(metrackerdirpath, "databases.json")

	dbh := &DataBaseHandler{
		databaseJsonPath: databaseJsonPath,
		databases:        []DataBase{},
	}

	if err := dbh.createDataBaseFile(); err != nil {
		return nil, err
	}

	if err := dbh.getDataBaseFile(); err != nil {
		return nil, err
	}

	return dbh, nil
}

func (dbh *DataBaseHandler) AddDatabase(db DataBase) error {
	if db.Name == "" {
		return errors.New("database name is empty")
	}

	for _, existing := range dbh.databases {
		if existing.Name == db.Name {
			return errors.New("database with this name already exists")
		}
	}

	dbh.databases = append(dbh.databases, db)
	return dbh.rewriteDataBaseFile()
}

func (dbh *DataBaseHandler) RmDatabase(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}

	for i, db := range dbh.databases {
		if db.Name == name {
			dbh.databases = append(dbh.databases[:i], dbh.databases[i+1:]...)
			return dbh.rewriteDataBaseFile()
		}
	}

	return errors.New("database not found")
}
func (dbh *DataBaseHandler) FindDataBaseByName(name string) (*DataBase, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}

	for i := range dbh.databases {
		if dbh.databases[i].Name == name {
			return &dbh.databases[i], nil
		}
	}

	return nil, errors.New("database not found")
}

func (dbh *DataBaseHandler) UpdateDataBase(updated DataBase) error {
	if updated.Name == "" {
		return errors.New("name is empty")
	}

	for i, db := range dbh.databases {
		if db.Name == updated.Name {
			dbh.databases[i] = updated
			return dbh.rewriteDataBaseFile()
		}
	}

	return errors.New("database not found")
}

func (dbh *DataBaseHandler) createDataBaseFile() error {
	_, err := os.Stat(dbh.databaseJsonPath)
	if err == nil {
		return nil
	}

	if !os.IsNotExist(err) {
		return err
	}

	file, err := os.Create(dbh.databaseJsonPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte("[]"))
	return err
}
func (dbh *DataBaseHandler) getDataBaseFile() error {
	data, err := os.ReadFile(dbh.databaseJsonPath)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		dbh.databases = []DataBase{}
		return nil
	}

	return json.Unmarshal(data, &dbh.databases)
}
func (dbh *DataBaseHandler) rewriteDataBaseFile() error {
	data, err := json.MarshalIndent(dbh.databases, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dbh.databaseJsonPath, data, 0644)
}
