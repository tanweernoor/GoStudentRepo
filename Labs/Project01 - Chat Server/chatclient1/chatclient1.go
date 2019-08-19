// chatclient.go
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func chatCopy(destination io.Writer, source io.Reader) {
	/// The chatCopy just copies from one stream to another
	if _, err := io.Copy(destination, source); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// A chat connection is requested, if the chat server is
	// not listenting then we just exit
	connection, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	// Makes sure we close the connection if we exit for any reason
	defer connection.Close()
	// All we do now is just copy what comes over the connection
	// to the console
	chatCopy(os.Stdout, connection)
}
