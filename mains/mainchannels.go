package main

import (
	"fmt"
	"time"
)

func sendData(ch chan int) {
	fmt.Println("Sending data...")
	ch <- 42
}

// Commented to the mainChannel because inside de package can`t have main
func mainChannel() {
	ch := make(chan int)

	// Launch a goroutine to send data
	go sendData(ch)

	// Receive data from the channel
	value := <-ch
	fmt.Println("Received data:", value)
	time.Sleep(1 * time.Second) // Wait for the goroutine to finish
}
