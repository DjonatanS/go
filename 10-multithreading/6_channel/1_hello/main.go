package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string) //Vazio

	//Thread 2
	go func() {
		canal <- "teste" //Cheio
	}()

	//Thread 1
	msg := <-canal //vazio
	fmt.Println(msg)
}
