package main

import (
	"fmt"
	"time"
)

func sendData(ch chan string, delay time.Duration) {
	time.Sleep(delay)
	ch <- "Data received"
}

func main_multiplechannels() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendData(ch1, 5*time.Second)
	go sendData(ch2, 3*time.Second)

	// Using select to handle multiple channels

	select {
	case msg1 := <-ch1:
		fmt.Println("From ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("From ch2:", msg2)

	case <-time.After(2 * time.Second):
		fmt.Println("Timeout: No message received in time.")
	}
}
