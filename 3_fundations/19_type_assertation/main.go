package main

import (
	"fmt"
)

func main() {
	var minhaVar interface{} = "Djonatan"

	//especificando que é string
	fmt.Println(minhaVar.(string))

	res, ok := minhaVar.(int)
	println(res)
	println(ok)

}
