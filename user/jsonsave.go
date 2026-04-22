package user

import (
	"encoding/json"
	"os"
)

func SaveUser(user User, path string) error {
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
func LoadUser(path string) (User, error) {
	var user User
	jsonData, err := os.ReadFile(path)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}
