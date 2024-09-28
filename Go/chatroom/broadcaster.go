package chatroom

type Client struct {
	record chan<- string // an outgoing message channel
	name   string
}

var (
	entering = make(chan Client)
	leaving  = make(chan Client)
	messages = make(chan string) // all incoming client message
)

func broadcaster() {
	clients := make(map[Client]chan<- string)

	for {
		select {

		case msg := <-messages:
			// Broadcast incoming message to all clients' outgoing message channels
			for _, clientCh := range clients {
				select {
				case clientCh <- msg:
				default:
					// If the client's channel is full, skip this message
				}
			}

		case client := <-entering:
			clients[client] = client.record

		case client := <-leaving:
			delete(clients, client)
			close(client.record)
		}
	}
}
