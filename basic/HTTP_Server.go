package basic

import (
	"log"
	"net/http"
	"time"
)

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	HTMLToSend := []byte("<html><head></head><body>Hello!</body></html>")
	_, err := w.Write(HTMLToSend)
	if err != nil {
		log.Printf("error occured while writing on the body: %s", err)
	}
}

// basic server
func HTTPServerExample() {
	myServer := &http.Server{
		// set the server address
		Addr: "127.0.0.1:8080",
		// define specific config
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// register handler
		Handler: &myHandler{},
	}
	log.Fatal(myServer.ListenAndServe())
}

// check which program is listening on port 8080
// list open files
// $sudo lsof -i :8080

// A/B testing server
func ABTestingServer() {
	abTestingServer := &http.Server{
		Addr:         "127.0.0.1:9899",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      &AbHandler{},
	}
	log.Fatal(abTestingServer.ListenAndServe())
}

type AbHandler struct{}

func (h *AbHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	designA := []byte("<html><head><title>The Golang Hotel</title></head><body><p>The Golang Hotel is a relaxing place!</p><p>We offer 20%% discount if you call this number: <strong>12345</strong></p></body></html>")
	designB := []byte("<html><head><title>The Golang Hotel</title></head><body><h2>The Golang Hotel is a relaxing place!</h2><h5>We offer 20%% discount if you call this number: <strong>12345</strong></h5></body></html>")

	minutes := time.Now().Minute()

	if minutes%2 == 0 {
		log.Println("serving design B")
		_, err := w.Write(designB)
		if err != nil {
			log.Print("imbossible to serve design A")
		}
	} else {
		log.Println("serving design A")
		_, err := w.Write(designA)
		if err != nil {
			log.Print("impossible to sere design B")
		}

	}
}
