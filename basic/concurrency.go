package basic

import (
	"fmt"
	"log"
	"time"
)

// To see list of processes that run on your computer:
// $ ps xau

// Example of Go race detector, which can help to to detect potential errors:
// $ go build -race -o myProgramName main.go

// In each program there is gorouitine this is the main goroutine

// A goroutine is a function that executes independently from the rest of the program

func ConcurrencyExample() {
	// simpleGoroutineExample()
	// channelsExample()
	// unbufferedChannelExample()
	// bufferedChannelExample()
	// channelUsecase()
	// deadlockBuffered()
	// deadlockUnbuffered()
	// selectWithoutDefault()
	selectWithDefault()
}

func simpleGoroutineExample() {
	fmt.Println("launch first goroutine")
	go printNumber()
	fmt.Println("launch second goroutine")
	go printNumber()
	time.Sleep(1 * time.Minute)
}

// Channels is a way for goroutines to communicates with each other.
// It's a pipeline of data between two goroutins.
// This pipeline can only support a specific type.
func channelsExample() {
	// send channel
	// chan <- int

	// receive channel
	// <- chan int

	// send and receive
	// chan int

	// Initialization of bidirectional unbuffered chan
	ch1 := make(chan int)

	// Initialization of bidirectional buffered chan
	ch2 := make(chan string, 3)
	// 3 is the capacity of the channel (space allocated to store values sent to chan)

	fmt.Printf("%T, %T\n", ch1, ch2)

	// send data to the channel
	ch3 := make(chan int, 2)
	ch3 <- 5
	// closing the channel
	// it indicates that no more values will be sent to the channel
	// you cannot send data to the closed channel
	// you cannot close a receive only channel
	// you cannot close channel already closed
	// you can receive date on a closed channel
	close(ch3)

	// receive data from the channel
	var received int
	ch4 := make(chan int, 2)
	ch4 <- 10
	ch4 <- 11
	received = <-ch4
	fmt.Println(received)

	// multi-valued receive operation
	ch5 := make(chan int, 2)
	ch5 <- 5
	close(ch5)
	received, ok := <-ch5
	fmt.Println(received, ok)
}

func unbufferedChannelExample() {
	ch := make(chan int)
	// this will block unbuffered channel for 3 seconds by dummy func
	go dummy(ch)
	log.Println("waiting for reception...")
	ch <- 45
	log.Println("received")
}

func bufferedChannelExample() {
	ch := make(chan int, 1)
	// this will block buffered channel until the data is copied to the buffer
	go dummy(ch)
	log.Println("waiting for reception...")
	// send operation is not blocking
	ch <- 45
	log.Println("received")
}

// Unbuffered channels are used to synchronize two goroutines
func channelUsecase() {
	syncCh := make(chan bool)
	// launch a second goroutine
	go func() {
		longTask2()
		syncCh <- true
	}()
	longTask1()
	// blocks until the second goroutine has finished
	<-syncCh
}

// Deadlock in bufferd channel
func deadlockBuffered() {
	ch := make(chan int, 1)
	go dDummy(ch)
	log.Println("waiting for reception")
	ch <- 45
	ch <- 58
	ch <- 100
}

// Deadlock unbufferd channel
func deadlockUnbuffered() {
	ch := make(chan int)
	ch <- 5
}

// //
// Select Statement
// //
func selectWithoutDefault() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch1 <- "test"

	select {
	case rec, ok := <-ch1:
		if ok {
			log.Printf("received on ch1: %s", rec)
		}
	case rec, ok := <-ch2:
		if ok {
			log.Printf("received on ch2: %s", rec)
		}
	}
	log.Println("end")
}

func selectWithDefault() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	select {
	case rec, ok := <-ch1:
		if ok {
			log.Printf("received on ch1: %s", rec)
		}
	case rec, ok := <-ch2:
		if ok {
			log.Printf("receved on ch2: %s", rec)
		}
	default:
		log.Println("default case")
	}
	log.Println("end")
}

// blocks the goroutine indefinitively and cause a deadlock
func emptySelect() {
	select {}
}

//
// util functions
//
func printNumber() {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		i++
		fmt.Println(i)
		if i == 10 {
			break
		}
	}
}

func dummy(c chan int) {
	time.Sleep(3 * time.Second)
	<-c
}

func longTask1() {
	time.Sleep(3 * time.Second)
}

func longTask2() {
	time.Sleep(1 * time.Second)
}

func dDummy(c chan int) {
	smth := <-c
	log.Println("has received something", smth)
}
