package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
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

    if buf[0] == '*' {
      fmt.Println("is an array")
      message := string(buf)
      spaces := strings.Split(message, "\r\n")

      count, _ := strconv.Atoi(spaces[0][1:])
      for i := 0; i < count; i++ {
        command := spaces[i + 2]
        space := spaces[i]
        fmt.Println(space)
        fmt.Println(command)

        if command == "ping" {
          conn.Write([]byte("+PONG\r\n"))
        } else if command == "echo" {
          fmt.Println(spaces[i + 4])
          i = i + 4
        }
      }
    } else if buf[0] == '+' {
      fmt.Println("is a string")
    } else if buf[0] == '$' {
      fmt.Println("is a bulk string")
    }
	}
}
