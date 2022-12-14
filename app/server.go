package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
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

type MapItem struct {
	value      string
	validUntil int64
}

func handle(conn net.Conn) {
	cache := make(map[string]MapItem)

	for {
		buf := make([]byte, 1024)

		if _, err := conn.Read(buf); err != nil {
			// fmt.Println("Error reading: ", err.Error())
			continue
		}

		if buf[0] == '*' {
			fmt.Println("is an array")
			message := string(buf)
			spaces := strings.Split(message, "\r\n")
			command := spaces[2]
			fmt.Println(message)

			if command == "ping" {
				conn.Write([]byte("+PONG\r\n"))
			} else if command == "echo" {
				parameter := spaces[4]
				conn.Write([]byte(fmt.Sprintf("+%s\r\n", parameter)))
			} else if command == "set" {
				key := spaces[4]
				newValue := spaces[6]
				if len(spaces) > 8 {
					if spaces[8] == "px" {
						milli, _ := strconv.Atoi(spaces[10])
						until := time.Now().Add(time.Duration(milli)).UnixMilli()
						fmt.Println(until)
						cache[key] = MapItem{value: newValue, validUntil: until}
					}
				} else {
					cache[key] = MapItem{value: newValue, validUntil: -1}
				}
				conn.Write([]byte("+OK\r\n"))
			} else if command == "get" {
				key := spaces[4]
				item := cache[key]
				now := time.Now().UnixMilli()
				fmt.Println(item)
				fmt.Println(now)
				if item.validUntil != -1 && item.validUntil < now {
					conn.Write([]byte("$-1\r\n"))
				} else {
					conn.Write([]byte(fmt.Sprintf("+%s\r\n", item.value)))
				}
			}
		} else if buf[0] == '+' {
			fmt.Println("is a string")
		} else if buf[0] == '$' {
			fmt.Println("is a bulk string")
		}
	}
}
