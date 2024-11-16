package main

type MyNumber int

// Constraint
type Number interface {
	~int | ~float64
}

type Decimal interface {
	float32 | float64
}

func Soma(m map[string]int) int {
	var soma int
	for _, value := range m {
		soma += value
	}
	return soma
}

func SomaGenerics[T Number](m map[string]T) T {
	var soma T
	for _, value := range m {
		soma += value
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Djonatan": 1000, "Thais": 1000}

	print(SomaGenerics(m))

	m2 := map[string]float64{"Djonatan": 1000.23, "Thais": 1000.23}

	print(SomaGenerics(m2))

	m3 := map[string]MyNumber{"Djonatan": 1000, "Thais": 1000}
	print(SomaGenerics(m3))

	println("")
	println(Compara(10, 10.1))

}
