package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
}

func (cliente Cliente) UpdateName(nome string) string {
	cliente.Nome = nome
	return cliente.Nome
}

func main() {
	valor := 10

	//crio uma variavel que recebe valor do inteiro e o endereço de valor
	var ponteiro *int = &valor
	fmt.Println(*ponteiro)

	valor = 20
	fmt.Println(*ponteiro)
}
