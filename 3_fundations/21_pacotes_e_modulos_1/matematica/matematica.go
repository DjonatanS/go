package matematica

type Carro struct {
	Marca string
}

type Decimal interface {
	int | float64
}

func (c Carro) Acelerar() {
	println("carro acelerando")
}

func Soma[T Decimal](a, b T) T {
	return a + b
}
