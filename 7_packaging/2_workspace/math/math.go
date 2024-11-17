package math

type Math struct {
	A int
	B int
}

func GetMath(a, b int) Math {
	return Math{A: a, B: b}
}

func (m *Math) Sum() int {
	return m.A + m.B
}
