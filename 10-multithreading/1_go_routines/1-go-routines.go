package main

import (
	"time"
)

func task(name string) {

	for i := 0; i < 10; i++ {
		println(name, ":", i)
		time.Sleep(time.Second)
	}
}

func main() {

	go task("A")
	go task("B")

	time.Sleep(10 * time.Second)
}
