package decimal

func Equal(a, b string) (bool, error) {
	aDec, bDec, err := ParseTwoString(a, b)
	if err != nil {
		return false, err
	}
	return aDec.Equal(bDec), nil
}

func Gt(a, b string) (bool, error) {
	aDec, bDec, err := ParseTwoString(a, b)
	if err != nil {
		return false, err
	}
	return aDec.GreaterThan(bDec), nil
}

func Lt(a, b string) (bool, error) {
	aDec, bDec, err := ParseTwoString(a, b)
	if err != nil {
		return false, err
	}
	return aDec.LessThan(bDec), nil
}

func Gte(a, b string) (bool, error) {
	aDec, bDec, err := ParseTwoString(a, b)
	if err != nil {
		return false, err
	}
	return aDec.GreaterThanOrEqual(bDec), nil
}

func Lte(a, b string) (bool, error) {
	aDec, bDec, err := ParseTwoString(a, b)
	if err != nil {
		return false, err
	}
	return aDec.LessThanOrEqual(bDec), nil
}
