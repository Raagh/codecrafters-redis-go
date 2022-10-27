package main

import (
	"fmt"
	"net"
	"os"
	"strings"
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
			// fmt.Println("Error reading: ", err.Error())
      continue;
		}

    message := string(buf)

    if message == "" {
      continue
    }

    // splitMessage := strings.Split(message, "\r\n")
    fmt.Println(message)

    if strings.ToLower(message) == "ping" {
      fmt.Println("PONG")
      conn.Write([]byte("+PONG\r\n"))
    } else {
      // numberOfParameters := splitMessage[0][1:]
      // fmt.Println(numberOfParameters)
    }
	}
}
