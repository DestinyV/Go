package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("start tcp server error:", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept tcp server error:", err)
			return
		}
		go resolveConn(conn)
	}
}
func resolveConn(conn net.Conn) {
	var rev [128]byte
	count, err := conn.Read(rev[:])
	if err != nil {
		fmt.Println("Read tcp server error:", err)
		return
	}
	fmt.Println(string(rev[:count]))
}
