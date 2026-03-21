package databases

import (
	"encoding/json"
	"os"
)

func (dbh *DataBaseHandler) load() error {
	data, err := os.ReadFile(dbh.databaseJsonPath)

	if os.IsNotExist(err) {
		dbh.Databases = []DataBase{}
		return dbh.save()
	}

	if err != nil {
		return err
	}

	if len(data) == 0 {
		dbh.Databases = []DataBase{}
		return nil
	}

	return json.Unmarshal(data, &dbh.Databases)
}
func (dbh *DataBaseHandler) save() error {
	data, err := json.MarshalIndent(dbh.Databases, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dbh.databaseJsonPath, data, 0644)
}
