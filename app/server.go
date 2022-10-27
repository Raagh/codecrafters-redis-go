package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")

	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	conn := acceptConnection(l)
	for conn != nil {
    go processMessage(conn);
    conn = acceptConnection(l)
	}
}

func acceptConnection(l net.Listener) net.Conn {
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	return conn
}

func processMessage(conn net.Conn) {
	for {
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading: ", err.Error())
		}

		conn.Write([]byte("+PONG\r\n"))
		// fmt.Println(string(buf))
		// fmt.Println(equal("ping", buf))
	}
}
