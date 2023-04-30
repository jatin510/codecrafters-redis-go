package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func acceptLoop(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		log.Println("connection is created")

		go readLoop(conn)
	}
}

func readLoop(conn net.Conn) {
	defer conn.Close()

	buff := make([]byte, 1024)

	for {
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println("Error reading: ", err.Error())
			return
		}

		log.Println("read from connection", string(buff[:n]))
		if strings.Contains(string(buff[:n]), "ping") {
			log.Println("ping is received")
			conn.Write([]byte("+PONG\r\n"))
		}
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	acceptLoop(l)
}
