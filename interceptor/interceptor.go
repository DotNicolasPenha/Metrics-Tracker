package interceptor

import (
	"bytes"
	"errors"
	"io"
	"net"
	"sync/atomic"
)

var (
	ErrEmptyProxyAddr = errors.New("proxy address is empty")
	ErrEmptyDBAddr    = errors.New("database address is empty")
)

type Interceptor struct {
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
	Retrys int64  `json:"retrys"`
}

type Limits struct {
	MaxActConnections int64 `json:"max_active_connections"`
}

type Metrics struct {
	ActConns int64
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
		metrics: Metrics{
			ActConns: 0,
		},
	}, nil
}

func (i *Interceptor) decrementBlockQueryRetry(query []byte) {
	for id, block := range i.Configurations.BlockQueries {
		if bytes.Equal(block.Query, query) {
			if block.Retrys > 0 {
				i.Configurations.BlockQueries[id].Retrys--
			}
			return
		}
	}
}
func (i *Interceptor) incrementConnections() {
	atomic.AddInt64(&i.metrics.ActConns, 1)
}

func (i *Interceptor) decrementConnections() {
	if i.metrics.ActConns > 0 {
		atomic.AddInt64(&i.metrics.ActConns, -1)
	}
}

func (i *Interceptor) isBlockedIP(ip string) bool {
	for _, authorizedIP := range i.Configurations.AuthorizedIPs {
		if ip == authorizedIP {
			return false
		}
	}
	return true
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
		if i.metrics.ActConns+1 > i.Configurations.Limits.MaxActConnections {
			i.logLimitExceeded(conn.RemoteAddr().String())
			conn.Close()
			continue
		}
		connAddrStr := conn.LocalAddr().String()
		if i.isBlockedIP(connAddrStr) {
			i.logBlockedNotAuthorizedIP(connAddrStr)
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
				// log.Printf("[DEBUG]: [%s]\n", string(rawPack))
				for _, block := range i.Configurations.BlockQueries {
					cleanBlock := bytes.ToLower(bytes.Join(bytes.Fields(block.Query), []byte(" ")))

					checkPack := bytes.ToLower(rawPack)
					for j := 0; j < len(checkPack); j++ {
						if (checkPack[j] < 'a' || checkPack[j] > 'z') && (checkPack[j] < '0' || checkPack[j] > '9') {
							checkPack[j] = ' '
						}
					}
					checkPack = bytes.Join(bytes.Fields(checkPack), []byte(" "))

					if bytes.Contains(checkPack, cleanBlock) {
						i.decrementBlockQueryRetry(block.Query)
						i.logBlockedQuery(conn.RemoteAddr().String(), block.Retrys)
						return
					}
				}
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
