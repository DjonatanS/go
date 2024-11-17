package main

import "github.com/DjonatanS/go/7_packaging/2_workspace/math"

func main() {
	math_struct := math.Math{A: 1, B: 2}

	println(math_struct.Sum())

}
