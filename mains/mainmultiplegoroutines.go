package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan string) {
	// Simulate some work
	fmt.Printf("Worker %d is starting...\n", id)
	time.Sleep(2 * time.Second)

	// Send a message back to the channel
	ch <- fmt.Sprintf("Worker %d completed", id)
}

func main_multiplegourotines() {
	ch := make(chan string)

	// Start multiple goroutines
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	// Receive messages from the channel
	for i := 1; i <= 3; i++ {
		msg := <-ch
		fmt.Println(msg)
	}

}
