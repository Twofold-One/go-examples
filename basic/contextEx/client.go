package contextEx

import (
	"context"
	"log"
	"net/http"
	"time"
)

func ContextExClient() {
	rootCtx := context.Background()
	// create context
	ctx, cancel := context.WithTimeout(rootCtx, 1500*time.Millisecond)
	defer cancel()
	req, err := http.NewRequest("GET", "http://127.0.0.1:8989", nil)
	if err != nil {
		panic(err)
	}
	// attach context to the request
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	log.Println("resp received", resp)
}
