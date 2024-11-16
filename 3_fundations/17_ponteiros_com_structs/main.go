package main

import "fmt"

type Cliente struct {
	Nome string
}

type Conta struct {
	Saldo int
}

func (conta Conta) SimularEmprestimo(Valor int) int {
	conta.Saldo += Valor
	return conta.Saldo
}

func NewConta() *Conta {
	return &Conta{Saldo: 0}
}

func (cliente *Cliente) Andou() {
	cliente.Nome = "Djonatan Schvambach"
	fmt.Printf("O Cliente %v andou \n", cliente.Nome)
}

func main() {

	thais := NewConta()

	println(thais.Saldo)

	djonatan := Cliente{
		Nome: "Djonatan",
	}
	djonatan.Andou()

	Conta := Conta{
		Saldo: 100,
	}

	SaldoTotal := Conta.SimularEmprestimo(500)

	fmt.Println("valor simulado", SaldoTotal)

	fmt.Println("valor atual", Conta.Saldo)

	fmt.Println(djonatan.Nome)
}
