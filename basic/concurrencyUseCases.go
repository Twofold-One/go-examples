package basic

import (
	"context"
	"fmt"
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
	waitGroup()
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
