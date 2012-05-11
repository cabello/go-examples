package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	listenHost = "127.0.0.1"
	listenPort = "1978"
)

func printErr(e error) {
	fmt.Fprintf(os.Stderr, "An error occurred: %s\n", e.Error())
}

func handler(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 24)
	for {
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			printErr(err)
			return
		}

		yell := strings.ToUpper(string(buffer[0:bytesRead]))
		conn.Write([]byte(yell))
	}
}

func main() {
	listener, err := net.Listen("tcp", listenHost+":"+listenPort)
	if err != nil {
		printErr(err)
		return
	}

	fmt.Printf("Listening %s:%s\n", listenHost, listenPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			printErr(err)
			return
		}
		go handler(conn)
	}
}
