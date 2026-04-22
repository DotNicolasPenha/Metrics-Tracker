package interceptor

import (
	"errors"
	"io"
	"net"
	"sync"
)

var (
	ErrEmptyProxyAddr = errors.New("proxy address is empty")
	ErrEmptyDBAddr    = errors.New("database address is empty")
)

type Interceptor struct {
	mu             sync.Mutex
	Name           string         `json:"name"`
	ProxyAddr      string         `json:"proxy_addr"`
	DBAddr         string         `json:"db_addr"`
	Configurations Configurations `json:"configurations"`
	Metrics        Metrics
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

type Metrics struct {
	ActConns int
}

func NewInterceptor(proxyAddr, dbAddr string) (*Interceptor, error) {
	if proxyAddr == "" {
		return nil, ErrEmptyProxyAddr
	}
	if dbAddr == "" {
		return nil, ErrEmptyDBAddr
	}
	return &Interceptor{
		ProxyAddr: proxyAddr,
		DBAddr:    dbAddr,
		Metrics:   Metrics{},
	}, nil
}

func (i *Interceptor) incrementConnections() {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.Metrics.ActConns++
}

func (i *Interceptor) decrementConnections() {
	i.mu.Lock()
	defer i.mu.Unlock()
	if i.Metrics.ActConns > 0 {
		i.Metrics.ActConns--
	}
}

func (i *Interceptor) Run() error {
	ln, err := net.Listen("tcp", i.ProxyAddr)
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		i.incrementConnections()
		go i.ConnHandler(conn)
	}
}
func (i *Interceptor) ConnHandler(conn net.Conn) {
	defer i.decrementConnections()
	defer conn.Close()

	dbConn, err := net.Dial("tcp", i.DBAddr)
	if err != nil {
		return
	}
	defer dbConn.Close()

	pipedone := make(chan struct{}, 2)

	go func() {
		defer func() { pipedone <- struct{}{} }()
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				return
			}
			if n > 0 {
				_, err := dbConn.Write(buf[:n])
				if err != nil {
					return
				}
			}
		}
	}()
	go func() {
		defer func() { pipedone <- struct{}{} }()
		io.Copy(conn, dbConn)
	}()

	<-pipedone
}
