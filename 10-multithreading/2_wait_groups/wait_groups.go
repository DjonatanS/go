package main

import (
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {

	for i := 0; i < 10; i++ {
		println(name, ":", i)
		time.Sleep(time.Second)
		wg.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(20)
	go task("A", &waitGroup)
	go task("B", &waitGroup)

	time.Sleep(10 * time.Second)
	waitGroup.Wait()
}
