// chatclient.go
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func chatCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	connection, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)
	go func() {
		io.Copy(os.Stdout, connection)
		log.Println("done")
		done <- true
	}()
	chatCopy(connection, os.Stdin)
	connection.Close()
	<-done
}
