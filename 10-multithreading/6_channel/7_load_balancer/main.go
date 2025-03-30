package main

import (
	"fmt"
	"time"
)

func worker(workerId int, channel chan int) {
	for value := range channel {
		fmt.Printf("Worker %d received %d\n", workerId, value)
		//time.Sleep(time.Second)
	}
}

func main() {
	channel := make(chan int)
	workersCount := 100

	// iniciar workers
	for i := 0; i < workersCount; i++ {
		go worker(i, channel)
	}

	startTime := time.Now()
	for i := 0; i < 1000000; i++ {
		//channel recebe valor do i
		channel <- i
	}

	close(channel) // Fechamento do channel

	elapsedTime := time.Since(startTime) // Tempo de processamento
	fmt.Printf("Total execution time: %s\n", elapsedTime)

}
