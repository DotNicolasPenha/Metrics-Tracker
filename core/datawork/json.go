package datawork

import (
	"encoding/json"
	"errors"
	"os"
	"path"
)

func (dwh *DataWorkHandler) Load(v any, jsonpath string) error {
	if jsonpath == "" {
		return errors.New("jsonpath is empty to load")
	}
	if v == nil {
		return errors.New("'v' is empty to load")
	}
	data, err := os.ReadFile(path.Join(dwh.pathToWork, jsonpath))

	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}
func (dwh *DataWorkHandler) Save(v any, jsonpath string) error {
	if jsonpath == "" {
		return errors.New("jsonpath is empty to load")
	}
	if v == nil {
		return errors.New("'v' is empty to load")
	}
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path.Join(dwh.pathToWork, jsonpath), data, 0644)
}
