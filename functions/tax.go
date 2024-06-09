package functions

func CalculateTax(amount float64) float64 {
	if amount >= 1000000 {
		return amount + (amount * 0.20)
	} else if amount >= 1000 {
		return amount + (amount * 0.10)
	} else if amount > 0 {
		return amount + (amount * 0.05)
	}
	return 0
}
