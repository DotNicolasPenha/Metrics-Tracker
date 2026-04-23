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
	metrics        Metrics
}
type Configurations struct {
	Limits        Limits        `json:"limits"`
	BlockQueries  []BlockQuerie `json:"block_queries"`
	AuthorizedIPs []string      `json:"authorized_ips"`
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
		metrics:   Metrics{},
	}, nil
}

func (i *Interceptor) incrementConnections() {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.metrics.ActConns++
}

func (i *Interceptor) decrementConnections() {
	i.mu.Lock()
	defer i.mu.Unlock()
	if i.metrics.ActConns > 0 {
		i.metrics.ActConns--
	}
}

func (i *Interceptor) Run() error {
	ln, err := net.Listen("tcp", i.ProxyAddr)
	if err != nil {
		i.logErrListenTCP(err)
		return err
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			i.logErrLnAccept(err)
			continue
		}
		if i.metrics.ActConns > i.Configurations.Limits.MaxActConnections {
			i.logLimitExceeded(conn.RemoteAddr().String())
			conn.Close()
			continue
		}
		i.incrementConnections()
		go i.ConnHandler(conn)
	}
}
func (i *Interceptor) ConnHandler(conn net.Conn) {
	defer i.decrementConnections()
	defer conn.Close()
	defer i.logDisconnection(conn.RemoteAddr().String())

	dbConn, err := net.Dial("tcp", i.DBAddr)
	if err != nil {
		i.logErrDBDial(conn.RemoteAddr().String(), err)
		return
	}
	defer dbConn.Close()
	i.logConnection(conn.RemoteAddr().String())
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
				rawPack := buf[:n]
				_, err := dbConn.Write(rawPack)
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
