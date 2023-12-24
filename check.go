package main

import (
	"fmt"
	"net"
	"time"
)

func check(destination, port string) string {
	address := destination + ":" + port
	timeout := 5 * time.Second

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return fmt.Sprintf("[DOWN] %s is unreachable: %v\n", address, err)
	}

	return fmt.Sprintf("[UP] %s is is reachable, from %v to %v\n", address, conn.LocalAddr(), conn.RemoteAddr())
}
