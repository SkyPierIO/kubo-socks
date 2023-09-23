package utils

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func GetFirstAvailablePort() string {
	for port := 3333; port <= 3999; port++ {
		p := strconv.Itoa(port)
		host := "127.0.0.1"
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, p), timeout)
		if err != nil {
			fmt.Println("Port is closed, ", err)
			return p
		}
		if conn != nil {
			defer conn.Close()
			fmt.Println("Opened", net.JoinHostPort(host, p))
		}
	}
	return ""
}
