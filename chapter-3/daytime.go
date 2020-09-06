package main

import (
	"fmt"
	"net"
	"log"
	"time"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":37")
	if err != nil {
		log.Fatal(err)
	}
	
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on :37")
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("Accept new client")
		
		conn.Write([]byte(time.Now().String()))
		conn.Close()
		fmt.Println("Connection closed")
	}
}

