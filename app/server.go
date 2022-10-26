package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.

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

		buf := make([]byte, 1024)

		_, err = conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading: ", err.Error())
		}

    expectedMessage := "ping"
		if strings.Compare(expectedMessage, string(buf)) == 0 {
			// conn.Write([]byte("+PONG\r\n"))
      fmt.Println("YES")
		}

		// fmt.Println(string(buf))
		// fmt.Println(equal("ping", buf))
	}
}

