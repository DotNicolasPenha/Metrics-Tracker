package interceptor

import (
	"log"
)

func (i *Interceptor) logConnection(clientAddr string) {
	log.Printf("[%s][NEW CONN] from %s to %s", i.Name, clientAddr, i.ProxyAddr)
	i.incrementConnections()
}
func (i *Interceptor) logDisconnection(clientAddr string) {
	log.Printf("[%s][DISCONN] from %s to %s", i.Name, clientAddr, i.ProxyAddr)
}

func (i *Interceptor) logLimitExceeded(clientAddr string) {
	log.Printf("[%s][LIMIT EXCEEDED] from %s: Active connections exceeded", i.Name, clientAddr)
}

func (i *Interceptor) logErrDBDial(clientAddr string, err error) {
	log.Printf("[%s][DB DIAL ERROR] from %s: %s", i.Name, clientAddr, err.Error())
}
func (i *Interceptor) logErrLnAccept(err error) {
	log.Printf("[%s][LN ACCEPT ERROR] %s", i.Name, err.Error())
}
func (i *Interceptor) logErrListenTCP(err error) {
	log.Printf("[%s][LISTEN TCP ERROR] %s", i.Name, err.Error())
}
