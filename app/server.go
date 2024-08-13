package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	log.Printf("Listening on %v", l.Addr())

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// make a buffer to read the data
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("Error reading data: %v\n", err.Error())
	}

	// print the cmd data by log
	log.Printf("Read command: %s\n", string(buf[:n]))

	// return PONG
	_, err = conn.Write([]byte("+PONG\r\n"))
	if err != nil {
		fmt.Printf("Error returning data: %v\n", err.Error())
	}
}
