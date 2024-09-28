package chatroom

import (
	"bufio"
	"fmt"
	"net"
)

func handleConn(conn net.Conn) {
	fmt.Fprint(conn, "Enter your name: ")
	clientScanner := bufio.NewScanner(conn)
	if !clientScanner.Scan() {
		conn.Close()
		fmt.Fprintln(conn, "Have to enter client name.")
		return
	}
	clientName := clientScanner.Text()

	clientCh := make(chan string) // outgoing client messages
	go clientWriter(conn, clientCh)

	// who := conn.RemoteAddr().String()
	clientCh <- "You are " + clientName
	messages <- clientName + " has arrived"
	entering <- Client{record: clientCh, name: clientName}

	fmt.Fprint(conn, "Enter something interesting: ")
	for clientScanner.Scan() {
		message := clientScanner.Text()
		messages <- clientName + ": " + message
	}

	leaving <- Client{record: clientCh, name: clientName}
	messages <- clientName + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, clientCh <-chan string) {
	for msg := range clientCh {
		fmt.Fprintln(conn, msg)
		// note:  ignoring network errors
	}
}
