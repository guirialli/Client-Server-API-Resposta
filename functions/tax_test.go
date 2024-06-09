package functions_test

import (
	"testing"

	"github.com/guirialli/pos/test/functions"
)

func TestTaxMoreEq1000(t *testing.T) {
	amount := 1000.00
	expected := 1100.00
	result := functions.CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestTaxLess1000(t *testing.T) {
	amount := 500.00
	expected := 525.00
	result := functions.CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}
	tests := []calcTax{
		{1000.00, 1100.00},
		{500.00, 525.00},
		{1000000.00, 1200000.00},
	}
	for _, test := range tests {
		result := functions.CalculateTax(test.amount)
		if result != test.expected {
			t.Errorf("Expected %f, got %f", test.expected, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		functions.CalculateTax(1000.00)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, 500.00, 1000.00, 1501.00}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := functions.CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("amount = %f, result = %f", amount, result)
		}
	})

}
