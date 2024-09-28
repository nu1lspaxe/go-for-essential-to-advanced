package chatroom

import (
	"log"
	"net"
)

/*
	log.Fatal(err) -> print err + os.Exit(1)
	log.Print(err) -> print err

	os.Exit(EXIT_CODE), EXIT_CODE:
	- 0 indicates successful termination.
    - non-zeore generally indicates an error or failure.
*/

func RunChatroom() {
	// open new terminal and type: `telnet localhost 1234`
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}
