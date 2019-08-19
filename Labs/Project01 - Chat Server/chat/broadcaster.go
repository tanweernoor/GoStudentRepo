// broadcaster.go

package main

// A client is an outgoing channel of strings
type client chan<- string

var (
	// The entering and leaving channels are channels over which we
	// send clients or outgoing channels of strings
	entering = make(chan client)
	leaving  = make(chan client)
	// this is the actual channel over which messages are sent
	messages = make(chan string)
)

func broadcaster() {
	// Needs to keep track of all the chat clients.
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// If there is a message from a client, then
			// loop over all clients and send it out to
			// each one.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			// If a client has entered, then they have sent the
			// their output channel to the broadcaster who adds
			// it to the map
			clients[cli] = true

		case cli := <-leaving:
			// When a client leaves, it sends its outgoing channel
			// over the leacing channel.
			// The broadcaster removes it from the map and closes
			// the channel.
			delete(clients, cli)
			close(cli)
		}
	}
}
