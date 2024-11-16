package main

import "fmt"

func main() {

	defer fmt.Println("First Row")

	fmt.Println("Second Row")
}
