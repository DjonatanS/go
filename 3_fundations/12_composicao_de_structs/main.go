package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
}

func main() {
	djonatan := Cliente{
		Nome:  "Djonatan",
		Idade: 27,
		Ativo: true,
	}

	djonatan.Ativo = false

	djonatan.Address.Cidade = "São Paulo"
	djonatan.Address.Estado = "SP"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", djonatan.Nome, djonatan.Idade, djonatan.Ativo)
}
