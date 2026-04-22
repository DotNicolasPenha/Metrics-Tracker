package user

import "github.com/DotNicolasPenha/Metrics-Tracker/interceptor"

type User struct {
	Configurations Configurations             `json:"configurations"`
	Interceptors   []*interceptor.Interceptor `json:"interceptors"`
}
type Configurations struct {
	Limits       Limits        `json:"limits"`
	BlockQueries []BlockQuerie `json:"block_queries"`
}
type BlockQuerie struct {
	Query  []byte `json:"query"`
	Retrys int    `json:"retrys"`
}
type Limits struct {
	MaxActConnections int `json:"max_active_connections"`
}
