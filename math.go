package decimal_util

import (
	"fmt"
)

func Add(a, b string) (string, error) {
	aDec, bDec, err := ParseTwoString(a, b)
	if err != nil {
		return "", err
	}
	result := aDec.Add(bDec)
	return result.String(), nil
}

func Subtract(a, b string) (string, error) {
	aDec, bDec, err := ParseTwoString(a, b)
	if err != nil {
		return "", err
	}
	result := aDec.Sub(bDec)
	return result.String(), nil
}

func Negate(value string) (string, error) {
	dec, err := ParseString(value)
	if err != nil {
		return "", fmt.Errorf("failed to parse value: %w", err)
	}
	result := dec.Neg()
	return result.String(), nil
}

func Multiply(a, b string) (string, error) {
	aDec, bDec, err := ParseTwoString(a, b)
	if err != nil {
		return "", err
	}
	result := aDec.Mul(bDec)
	return result.String(), nil
}

func Divide(a, b string) (string, error) {
	aDec, bDec, err := ParseTwoString(a, b)
	if err != nil {
		return "", err
	}
	if bDec.IsZero() {
		return "", fmt.Errorf("division by zero")
	}
	result := aDec.Div(bDec)
	return result.String(), nil
}
