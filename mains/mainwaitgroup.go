package main

import (
	"fmt"
	"sync"
	"time"
)

func worker_wg(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutines finishes

	// Simulate work
	fmt.Println("Worker", id, "started.")
	time.Sleep(time.Second)

	fmt.Println("Worker", id, "finished.")
}

func main_wait_group() {
	var wg sync.WaitGroup

	// Start multiple goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment the counter for each new goroutine
		go worker_wg(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println("All workers finished.")
}
