package main

import (
	"fmt"
	"goproj/3_fundations/21_pacotes_e_modulos_1/matematica"

	"github.com/google/uuid"
)

func main() {

	carro := matematica.Carro{Marca: "Teste"}

	carro.Acelerar()

	fmt.Println(carro.Marca)

	fmt.Print(matematica.Soma(10, 20))

	fmt.Println(uuid.New())
}
