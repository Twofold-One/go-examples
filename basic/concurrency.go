package basic

import (
	"fmt"
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
	channelsExample()
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


// util functions
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
