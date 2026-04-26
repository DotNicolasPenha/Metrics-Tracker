package interceptor

import (
	"log"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	Bold   = "\033[1m"
)

func (i *Interceptor) logConnection(clientAddr string) {
	// clientAddr: Green + Bold | ProxyAddr: Cyan + Bold
	log.Printf("%s[%s]%s %s[NEW CONN]%s from %s%s%s%s to %s%s%s%s",
		Gray, i.Name, Reset, Cyan, Reset,
		Bold, Green, clientAddr, Reset,
		Bold, Cyan, i.ProxyAddr, Reset)
}

func (i *Interceptor) logDisconnection(clientAddr string) {
	log.Printf("%s[%s]%s %s[DISCONN]%s from %s%s%s%s to %s%s%s%s",
		Gray, i.Name, Reset, Gray, Reset,
		Bold, Green, clientAddr, Reset,
		Bold, Cyan, i.ProxyAddr, Reset)
}

func (i *Interceptor) logLimitExceeded(clientAddr string) {
	log.Printf("%s[%s]%s %s[LIMIT EXCEEDED]%s from %s%s%s%s: Active connections exceeded",
		Gray, i.Name, Reset, Yellow, Reset,
		Bold, Green, clientAddr, Reset)
}

func (i *Interceptor) logBlockedQuery(clientAddr string, remainingRetries int64) {
	log.Printf("%s[%s]%s %s%s[BLOCKED QUERY]%s from %s%s%s%s: Query blocked. Retries: %d",
		Gray, i.Name, Reset, Red, Bold, Reset,
		Bold, Green, clientAddr, Reset, remainingRetries)
}

func (i *Interceptor) logErrDBDial(clientAddr string, err error) {
	log.Printf("%s[%s]%s %s[DB DIAL ERROR]%s from %s%s%s%s: %v",
		Gray, i.Name, Reset, Red, Reset,
		Bold, Green, clientAddr, Reset, err)
}
func (i *Interceptor) logBlockedNotAuthorizedIP(clientAddr string) {
	log.Printf("%s[%s]%s %s[BLOCKED IP]%s from %s%s%s%s: IP not authorized",
		Gray, i.Name, Reset, Red, Reset,
		Bold, Green, clientAddr, Reset)
}
func (i *Interceptor) logErrLnAccept(err error) {
	log.Printf("%s[%s]%s %s[LN ACCEPT ERROR]%s %v",
		Gray, i.Name, Reset, Red, Reset, err)
}

func (i *Interceptor) logErrListenTCP(err error) {
	log.Printf("%s[%s]%s %s[LISTEN TCP ERROR]%s %v",
		Gray, i.Name, Reset, Red, Reset, err)
}
