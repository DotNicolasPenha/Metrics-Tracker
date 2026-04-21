package configurations

import (
	"encoding/json"
	"os"

	"github.com/DotNicolasPenha/Metrics-Tracker/interceptor"
)

type User struct {
	Configurations Configurations            `json:"configurations"`
	Interceptors   []interceptor.Interceptor `json:"interceptors"`
}
type Configurations struct {
	Limits       Limits        `json:"limits"`
	BlockQueries []BlockQuerie `json:"block_queries"`
}
type BlockQuerie struct {
	Query  string `json:"query"`
	Retrys int    `json:"retrys"`
}
type Limits struct {
	MaxActConnections int `json:"max_active_connections"`
}

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
