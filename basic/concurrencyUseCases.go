package basic

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func ConcurrencyUseCasesExample() {
	// gracefulShutdownHTTPServer()
	// timeout()
	// waitGroup()

	go func() {
		mutexHTTP()
	}()

	time.Sleep(2 * time.Second)
	log.Println("starting server calls...")
	mutexServerCall()
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

// // // // // //
// Wait groups //
// // // // // //

func waitGroup() {
	fmt.Println("Program start")
	// initialize waitGroup
	var waitGroup sync.WaitGroup
	// increment waitGroup
	waitGroup.Add(10)
	for i := 0; i < 10; i++ {
		// without waitGroup
		// go concurrentTasksWithout(i)
		go concurrentTasksWith(i, &waitGroup)
	}
	// block current goroutine until all goroutines have finished
	waitGroup.Wait()
	finishTask()
	fmt.Println("Program end")

}

func finishTask() {
	fmt.Println("Executing finish task")
}

func concurrentTasksWithout(taskNumber int) {
	fmt.Printf("BEGIN Execute task number %d\n", taskNumber)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("END Execute task number %d\n", taskNumber)
}

func concurrentTasksWith(taskNumber int, waitGroup *sync.WaitGroup) {
	fmt.Printf("BEGIN Execute task number %d\n", taskNumber)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("END Execute task number %d\n", taskNumber)
	// decrement waitGroup
	waitGroup.Done()
}

// // // // //
// Mutexes //
// // // // //

var mu sync.Mutex
var requestCount int

func mutexHTTP() {
	http.HandleFunc("/status", status)
	log.Println("starting server at :8090")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// status increments requestCount each time a request is sent to the route
func status(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	requestCount++
	mu.Unlock()
	fmt.Fprintf(w, "OK - count: %d\n", requestCount)
}

// mutexServerCall will start 10 goroutines that will each
// send 100 requests to the server concurrently.
func mutexServerCall() {
	var w8 sync.WaitGroup
	w8.Add(10)
	for k := 0; k < 10; k++ {
		go caller(&w8)
	}
	w8.Wait()
}

func caller(w8 *sync.WaitGroup) {
	for k := 0; k < 100; k++ {
		res, err := http.Get("http://localhost:8090/status")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		s, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(s))
	}
	w8.Done()
}

// Note
// There is a pretty common idion called the Mutex hat
// when mutexes are usually used in struct types

type st struct {
	mu sync.Mutex
	// ...
}
