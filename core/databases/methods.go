package databases

import "errors"

func (dbh *DataBaseHandler) AddDatabase(db DataBase) error {
	if db.Name == "" {
		return errors.New("database name is empty")
	}

	for _, existing := range dbh.Databases {
		if existing.Name == db.Name {
			return errors.New("database with this name already exists")
		}
	}

	dbh.Databases = append(dbh.Databases, db)
	return dbh.save()
}
func (dbh *DataBaseHandler) RmDatabase(name string) error {
	if name == "" {
		return errors.New("name is empty")
	}

	for i, db := range dbh.Databases {
		if db.Name == name {
			dbh.Databases = append(dbh.Databases[:i], dbh.Databases[i+1:]...)
			return dbh.save()
		}
	}

	return errors.New("database not found")
}
func (dbh *DataBaseHandler) UpdateDataBase(updated DataBase) error {
	if updated.Name == "" {
		return errors.New("name is empty")
	}

	for i, db := range dbh.Databases {
		if db.Name == updated.Name {
			dbh.Databases[i] = updated
			return dbh.save()
		}
	}

	return errors.New("database not found")
}
func (dbh *DataBaseHandler) FindDataBaseByName(name string) (*DataBase, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}

	for i := range dbh.Databases {
		if dbh.Databases[i].Name == name {
			return &dbh.Databases[i], nil
		}
	}

	return nil, errors.New("database not found")
}
