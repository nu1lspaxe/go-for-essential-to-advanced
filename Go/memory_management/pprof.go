package memorymanagement

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func RunPprof() {
	go func() {
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}()

	select {}
}

// go tool pprof http://localhost:6060/debug/pprof/heap
// command: top, list [function name], web...
