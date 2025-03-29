package main

import (
	"fmt"
	"time"
)

func task(name string) []string {
	data := []string{}

	for i := 0; i < 10; i++ {
		println(name, ":", i)
		time.Sleep(time.Second)
		data = append(data, name)
	}
	return data
}

func main() {
	// Define tasks to run
	tasks := []string{"Task1", "Task2", "Task3"}

	// Create a channel to receive results
	resultChannel := make(chan map[string][]string, len(tasks))

	// Start time measurement
	startTime := time.Now()

	// Launch each task as a goroutine
	for _, taskName := range tasks {
		go func(name string) {
			// Run the task
			result := task(name)
			// Send the result through the channel
			resultChannel <- map[string][]string{name: result}
		}(taskName)
	}

	// Collect results from all tasks
	allResults := make(map[string][]string)
	for i := 0; i < len(tasks); i++ {
		// Receive result from channel
		result := <-resultChannel
		// Merge into final results
		for k, v := range result {
			allResults[k] = v
		}
	}

	// Calculate and display execution time
	duration := time.Since(startTime)
	fmt.Printf("\nAll tasks completed in %v\n", duration)

	// Print results
	fmt.Println("\nResults collected:")
	for name, data := range allResults {
		fmt.Printf("%s: %v\n", name, data)
	}
}
