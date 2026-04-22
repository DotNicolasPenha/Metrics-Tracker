package main

import "os"

func main() {
	if len(os.Args) < 5 {
		println("Usage: mt new inter <proxy addr> <db addr>")
		return
	}
	if os.Args[1] != "new" {
		println("Usage: mt new inter <proxy addr> <db addr>")
		return
	}
	if os.Args[2] != "inter" {
		println("Usage: mt new inter <proxy addr> <db addr>")
		return
	}
	if os.Args[3] == "" {
		println("Usage: mt new inter <proxy addr> <db addr>")
		return
	}
	if os.Args[4] == "" {
		println("Usage: mt new inter <proxy addr> <db addr>")
		return
	}

	proxyAddr := os.Args[3]
	dbAddr := os.Args[4]
	println("Proxy Address:", proxyAddr)
	println("Database Address:", dbAddr)
}
