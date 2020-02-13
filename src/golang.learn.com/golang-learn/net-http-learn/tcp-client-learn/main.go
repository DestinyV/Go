package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("client dial error", err)
		return
	}

	conn.Write([]byte("hello tcp server"))
	conn.Close()

}
