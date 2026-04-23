package interceptor

import (
	"log"
)

func (i *Interceptor) logConnection(clientAddr string) {
	log.Printf("[%s] New connection from %s to %s", i.Name, clientAddr, i.ProxyAddr)
	i.incrementConnections()
}
func (i *Interceptor) logDisconnection(clientAddr string) {
	log.Printf("[%s] Connection from %s to %s closed", i.Name, clientAddr, i.ProxyAddr)
}
