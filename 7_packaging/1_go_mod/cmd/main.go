package main

import "github.com/DjonatanS/go/7_packaging/1_go_mod/math"

func main() {
	math_struct := math.Math{A: 1, B: 2}

	println(math_struct.Sum())

}
