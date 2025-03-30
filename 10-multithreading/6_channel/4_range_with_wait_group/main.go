package main

import (
	"fmt"
	"sync"
)

func main() {
	//Criação de um channel vazio
	ch := make(chan int)
	//Criação do waitGroup
	wg := sync.WaitGroup{}
	//Adicionando 10 espaços no waitGroup
	wg.Add(10)
	go publisher(ch)
	go reader(ch, &wg)
	wg.Wait()
}

func publisher(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	//Fechar o Canal para evitar deadlock não é mais necessário com waitGroups
	//close(ch)
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Printf("Received %d\n", i)
		wg.Done()
	}
}
