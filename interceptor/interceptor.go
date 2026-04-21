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
	mu        sync.Mutex
	ProxyAddr string
	DBAddr    string
	Metrics   Metrics
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
		io.Copy(dbConn, conn)
	}()
	go func() {
		defer func() { pipedone <- struct{}{} }()
		io.Copy(conn, dbConn)
	}()

	<-pipedone
}
