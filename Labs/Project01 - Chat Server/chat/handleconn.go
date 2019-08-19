// broadcaster.go
package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConn(conn net.Conn) {
	// A new channel for this client is created
	// and a clientWriter is started to write all
	// incoming messages from the broadcaster to
	// the client on the other end of the connection.
	ch := make(chan string)
	go clientWriter(conn, ch)

	// The address (port) of the client is found and saved
	// as "who" to be used as prefixes to the messages
	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	// Announce arrival and then send the broadcaster the
	// channel so it can send messages
	messages <- who + " has arrived"
	entering <- ch

	// We create a scanner to read lines off of the connection.
	input := bufio.NewScanner(conn)
	for input.Scan() {
		// For each line, the prefix is added to the text of the
		// line and then send over the messages channel to the
		// broadcaster to be relayed out to all the other clients
		messages <- who + ": " + input.Text()
	}
	// Once the client stop, then the Scan() function exits.
	// The broadcaster is sent the channel so it can be removed
	// from the map, and an exit message is sent to everyone
	leaving <- ch
	messages <- who + " has left"
	// The final step is to close the connection which causes the
	// clientWriter() to terminate.
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		// The client writer takes each messages sent over the channel from
		// the broadcaster and writes it to the connection
		fmt.Fprintln(conn, msg)
	}
}
