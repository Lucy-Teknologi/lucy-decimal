package decimal

func ParseString(value string) (Decimal, error) {
	return NewFromString(value)
}

func ParseTwoString(value1, value2 string) (Decimal, Decimal, error) {
	dec1, err := NewFromString(value1)
	if err != nil {
		return Decimal{}, Decimal{}, err
	}
	dec2, err := NewFromString(value2)
	if err != nil {
		return Decimal{}, Decimal{}, err
	}
	return dec1, dec2, nil
}
