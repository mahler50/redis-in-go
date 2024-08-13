package main

import (
	"fmt"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	conn, err := net.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		t.Errorf("Fail to conn to server: %v", err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte("PING"))
	if err != nil {
		t.Errorf("Faile to send PING: %v", err)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		t.Errorf("Fail to read data from buffer: %v", err)
	}
	fmt.Println(string(buf[:n]))
}
