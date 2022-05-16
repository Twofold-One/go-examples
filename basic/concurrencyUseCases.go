package basic

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ConcurrencyUseCasesExample() {
	// gracefulShutdownHTTPServer()
	timeout()
}

func gracefulShutdownHTTPServer() {
	// create the notification channel
	bye := make(chan os.Signal)
	signal.Notify(bye, os.Interrupt, syscall.SIGTERM)

	mux := http.NewServeMux()
	mux.Handle("/status", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "OK")
		},
	))
	srv := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
	// lauch the server in another goroutine
	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("server: %q\n", err)
		}
	}()
	// wait for os signal
	sig := <-bye
	// code below is executed when we receive an os.Signal
	log.Printf("detected os signal %s\n", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	err := srv.Shutdown(ctx)
	cancel()
	if err != nil {
		log.Fatal(err)
	}
}

func timeout() {
	ch := make(chan int, 1)
	select {
	case rec, ok := <-ch:
		if ok {
			log.Printf("received %d", rec)
		}
	case rec, ok := <-time.After(time.Second * 3):
		if ok {
			log.Printf("operation timed out at %s", rec)
		}
	}

}

// //
// Wait groups
// //

// ToDo
