package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco Endereco
}

func (c *Cliente) DesativarCliente() {
	c.Ativo = false
	fmt.Printf("O Cliente %s foi desativado", c.Nome)
}

func main() {
	djonatan := Cliente{
		Nome:  "Djonatan",
		Idade: 27,
		Ativo: true,
	}

	djonatan.DesativarCliente()

	print(djonatan.Ativo)

	fmt.Printf("")
}
