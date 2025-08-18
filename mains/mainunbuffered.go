package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	fmt.Println("Producer started")

	time.Sleep(1 * time.Second) // Simulate work

	ch <- 42

	fmt.Println("Producer finished")
}

func consumer(ch chan int) {
	fmt.Println("Consumer started")

	data := <-ch // Receive data from channel

	fmt.Printf("Consumer received: %d\n", data)

	fmt.Println("Consumer finished")
}

func main() {
	ch := make(chan int)

	// Start goroutines
	go producer(ch)
	go consumer(ch)

	// Wait for goroutines to complete
	time.Sleep(2 * time.Second)
}
