package matematica

type Carro struct {
	Marca string
}

func (c Carro) Acelerar() {
	println("carro acelerando")
}

func Soma[T int | float64](a, b T) T {
	return a + b
}
