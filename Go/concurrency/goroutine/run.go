package goroutine

import (
	"context"
	"log"
	"net/http"
)

func Run() {
	done := make(chan error, 3)
	stop := make(chan struct{})

	// debug
	go func() {
		done <- pprof(stop)
	}()

	// main service
	go func() {
		done <- app(stop)
	}()

	go func() {
		reporter.Run(stop)
		done <- nil
	}()

	var stoped bool
	for i := 0; i < cap(done); i++ {
		if err := <-done; err != nil {
			log.Printf("server exit err: %+v", err)
		}

		if !stoped {
			stoped = true
			close(stop)
		}
	}
}

var reporter = NewReporter(3, 10)

func app(stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		reporter.Report("ping pong")
		w.Write([]byte("pong"))
	})
	return server(mux, ":8080", stop)
}

func pprof(stop <-chan struct{}) error {
	return server(http.DefaultServeMux, ":8081", stop)
}

func server(handler http.Handler, addr string, stop <-chan struct{}) error {
	s := http.Server{
		Handler: handler,
		Addr:    addr,
	}

	go func() {
		<-stop
		log.Printf("server will exiting, addr: %s", addr)
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}
