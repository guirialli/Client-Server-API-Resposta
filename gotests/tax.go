package gotests

import "errors"

type Repository interface {
	Save(amount float64) error
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax(amount)
	if err := repository.Save(tax); err != nil {
		return err
	}
	return nil
}

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

func CalculateTaxError(amount float64) (float64, error) {
	if amount >= 1000000 {
		return amount + (amount * 0.20), nil
	} else if amount >= 1000 {
		return amount + (amount * 0.10), nil
	} else if amount > 0 {
		return amount + (amount * 0.05), nil
	}
	return 0, errors.New("invalid amount value has to be than 0")
}
