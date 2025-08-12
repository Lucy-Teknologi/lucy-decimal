package decimal

import (
	lucy_decimal "github.com/Lucy-Teknologi/lucy-decimal"
)

func ParseString(value string) (lucy_decimal.Decimal, error) {
	return lucy_decimal.NewFromString(value)
}

func ParseTwoString(value1, value2 string) (lucy_decimal.Decimal, lucy_decimal.Decimal, error) {
	dec1, err := lucy_decimal.NewFromString(value1)
	if err != nil {
		return lucy_decimal.Decimal{}, lucy_decimal.Decimal{}, err
	}
	dec2, err := lucy_decimal.NewFromString(value2)
	if err != nil {
		return lucy_decimal.Decimal{}, lucy_decimal.Decimal{}, err
	}
	return dec1, dec2, nil
}
