package main

import "fmt"

// Envia o nome recebido para dentro do channel chan<- preencher o channel
func receiver(name string, ch chan<- string) {
	ch <- name
}

// Consome o channel mostrando o dado <-chan esvaziar o channel
func reader(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	channel := make(chan string)
	go receiver("Teste", channel)
	reader(channel)
}
