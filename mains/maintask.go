package main

import "fmt"

func task(ch chan bool) {
	fmt.Println("Performing task...")
	ch <- true // Signal that the task is complete
}

func main() {
	ch := make(chan bool)

	go task(ch)

	// Wait for the task to complete
	<-ch
	fmt.Println("Task completed successfully.")
}
