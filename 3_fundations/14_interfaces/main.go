package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

//Interface pode conter apenas funções

type Pessoa interface {
	DesativarCliente()
}

type Empresa struct {
	Nome string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco Endereco
}

func (e Empresa) Desativar() {

}

func (cliente Cliente) DesativarCliente() {
	cliente.Ativo = false
	fmt.Printf("O cliente %s foi desativado", cliente.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.DesativarCliente()
}

func main() {
	djonatan := Cliente{
		Nome:  "Djonatan",
		Idade: 27,
		Ativo: true,
	}

	Desativacao(djonatan)

	fmt.Printf("")
}
