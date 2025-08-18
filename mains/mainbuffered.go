package main

import "fmt"

func producerBuffer(ch chan int) {
	fmt.Println("Producer started!")
	ch <- 42
	fmt.Println("Producer sent data!")

	ch <- 43
	fmt.Println("Producer sent data again!")
}

func consumerBuffer(ch chan int) {
	fmt.Println("Consumer started!")
	data1 := <-ch
	fmt.Printf("Consumer received: %d\n", data1)

	data2 := <-ch
	fmt.Printf("Consumer received: %d\n", data2)

	fmt.Println("Consumer finished!")
}

func main_buffered() {
	ch := make(chan int, 2) // Create a buffered channel

	go producerBuffer(ch)
	go consumerBuffer(ch)

	// Wait for goroutines to finish
	fmt.Scanln()
}
