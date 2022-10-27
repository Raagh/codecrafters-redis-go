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

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
    
		if _, err := conn.Read(buf); err != nil {
			fmt.Println("Error reading: ", err.Error())
      continue;
		}

		conn.Write([]byte("+PONG\r\n"))
		// fmt.Println(string(buf))
		// fmt.Println(equal("ping", buf))
	}
}
