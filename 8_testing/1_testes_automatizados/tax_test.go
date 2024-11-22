package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	value := 500.0
	expected := 5.0

	result := CalculateTax(value)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type Tax struct {
		value    float64
		expected float64
	}

	dict := []Tax{
		{0.0, 0.0},
		{500.0, 5.0},
		{1000.0, 10.0},
	}

	for _, item := range dict {
		result := CalculateTax(item.value)
		if result != item.expected {
			t.Errorf("Expected %f but got %f", item.expected, result)
		}
	}
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(1000.0)
	}
}
