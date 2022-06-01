package contextEx

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func ContextExServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[Handler] request received")
		// retreive the context of the request
		rCtx := r.Context()
		// create the result chan
		resChan := make(chan int)
		// launch the function doWork in a goroutine
		go doWork(rCtx, resChan)
		// Wait for
		// 1. The client drops the connection.
		// 2. the function doWork to finish it works
		select {
		case <-rCtx.Done():
			log.Println("[Handler] context canceled in main handler, client has diconnected")
			return
		case result := <-resChan:
			log.Println("[Handler] Received 1000")
			log.Println("[Handler] Send response")
			fmt.Fprintf(w, "Response %d", result)
			return
		}

	})

	err := http.ListenAndServe("127.0.0.1:8989", nil)
	if err != nil {
		panic(err)
	}
}

func doWork(ctx context.Context, resChan chan int) {
	log.Println("[doWork] launch the doWork")
	sum := 0
	for {
		log.Println("[doWork] one iteration")
		time.Sleep(time.Millisecond)
		select {
		case <-ctx.Done():
			log.Println("[doWork] ctx Done is received inside doWork")
			return
		default:
			sum++
			if sum > 1000 {
				log.Println("[doWork] sum has reached 1000")
				resChan <- sum
				return
			}
		}
	}
}
