package main

import "fmt"

func main() {
	ch := make(chan int)
	go publisher(ch)
	reader(ch)

}

func publisher(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	//Fechar o Canal para evitar deadlock
	close(ch)
}

func reader(ch chan int) {
	for i := range ch {
		fmt.Printf("Received %d\n", i)
	}
}
