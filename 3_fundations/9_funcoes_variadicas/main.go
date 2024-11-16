package main

import (
	"fmt"
)

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(numeros...))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
