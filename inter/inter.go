package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

var (
	actconns = 0
	maxconns = 10
	paddr    = ":8080"
	exaddr   = "localhost:5432"
)

var mu sync.Mutex

func main() {
	go intercept()
	for {
		fmt.Print("\033[H")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("connections:", actconns)
		time.Sleep(500 * time.Millisecond)
	}
}
func intercept() {
	listener, err := net.Listen("tcp", paddr)
	if err != nil {
		panic(err)
	}

	for {
		cconn, err := listener.Accept()
		if err != nil {
			continue
		}
		go InterHandler(cconn)
	}
}

func InterHandler(cconn net.Conn) {
	mu.Lock()
	if actconns+1 > maxconns {
		mu.Unlock()
		cconn.Close()
		return
	}
	actconns++
	mu.Unlock()

	defer func() {
		mu.Lock()
		actconns--
		mu.Unlock()
	}()

	dbConn, err := net.Dial("tcp", exaddr)
	if err != nil {
		cconn.Close()
		return
	}

	defer cconn.Close()
	defer dbConn.Close()
	go func() {
		buf := make([]byte, 4096)
		for {
			nb, err := cconn.Read(buf)
			if err != nil {
				actconns--
				break
			}
			data := buf[:nb]
			fmt.Println("Client- db:", string(data))
			_, err = dbConn.Write(data)
			if err != nil {
				actconns--
				break
			}
		}
	}()

	io.Copy(cconn, dbConn)
}
