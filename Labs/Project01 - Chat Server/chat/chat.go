package main

import (
	"log"
	"net"
	//"time"
)

func main() {
	// A listener is created to wait for clients that  are
	// trying to join the chat
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	// The broadcaster is started
	go broadcaster()
	for {
		// The loop will block here until a client requests
		// to join. Then a connection to that client is
		// created
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// A handler for the connection is created, and then
		// the for loop goes back to listening for a new client
		// to request to be connected.
		go handleConn(conn)
	}
}
