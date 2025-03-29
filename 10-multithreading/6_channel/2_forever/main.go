package main

import "fmt"

// Thread 1
func main() {
	forever := make(chan bool)
	go func() {
		for i := range 10 {
			fmt.Println(i)
		}
		forever <- true
	}()
	<-forever
}
