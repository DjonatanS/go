package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	djonatan := Cliente{
		Nome:  "Djonatan",
		Idade: 27,
		Ativo: true,
	}

	djonatan.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", djonatan.Nome, djonatan.Idade, djonatan.Ativo)
}
