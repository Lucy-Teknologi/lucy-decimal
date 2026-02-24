package decimal_util

import "github.com/shopspring/decimal"

func ParseString(value string) (decimal.Decimal, error) {
	return decimal.NewFromString(value)
}

func ParseTwoString(value1, value2 string) (decimal.Decimal, decimal.Decimal, error) {
	dec1, err := decimal.NewFromString(value1)
	if err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, err
	}
	dec2, err := decimal.NewFromString(value2)
	if err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, err
	}
	return dec1, dec2, nil
}
