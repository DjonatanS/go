package main

import "fmt"

func main() {
	type x interface{}
	type y interface{}

	var a x = 10
	var b y = "Hello World"

	showType(a)
	showType(b)
}

func showType(T interface{}) {
	fmt.Printf("the var type is %T \n", T)
}
