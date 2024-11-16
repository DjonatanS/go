package main

func main() {

	for i := 0; i < 10; i++ {
		println(i)
	}

	lista := []int{1, 2, 3}
	for _, v := range lista {
		println(v)
	}

	dict := map[string]int{
		"Djonatan": 1000,
		"Thais":    1000,
	}

	for key, value := range dict {
		println(key, value)
	}
}
