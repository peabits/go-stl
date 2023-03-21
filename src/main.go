package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Printf("Hello, GO!\n")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	fmt.Printf("%#v", addrs)
}
