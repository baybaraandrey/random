package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func handleC(conn net.Conn) {
	defer func() {
		fmt.Println("Connection closed")
		conn.Close()
	}()

	var buf [512]byte
	
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		
		msg := strings.TrimSpace(string(buf[:n]))
		if msg == "close" {
			return
		}
		fmt.Println("Read message:", msg)
		_, err = conn.Write([]byte(fmt.Sprintf(">> %s\n>> ", msg)))
		if err != nil {
			return
		}
	}
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on :8080")
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("Accepted new connection")
		go handleC(conn)
	}
}
