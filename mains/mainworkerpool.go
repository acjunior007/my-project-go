package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work
		results <- job * 2      // Send result back

		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}

func main_workpool() {
	const numWorkers = 3
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, results)
	}

	// Send tasks to the workers
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // Close the jobs channel to signal no more jobs

	// Collection results
	for k := 1; k <= 5; k++ {
		fmt.Println("Result:", <-results)
	}
}
