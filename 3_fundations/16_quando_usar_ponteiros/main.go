package main

import "fmt"

// receber ponteiro de um inteiro
func soma(a, b *int) int {
	*a = 50
	result := *a + *b
	return result
}

func main() {

	a := 10
	b := 20
	fmt.Println(a)

	//passando endereços de memoria
	fmt.Println(soma(&a, &b))

	//valor da memoria alterado
	fmt.Println(a)

}
