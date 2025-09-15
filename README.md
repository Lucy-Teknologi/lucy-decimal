# README — lucy-decimal
## Project Overview
lucy-decimal is a Go library aimed at providing precise and robust decimal arithmetic support—ideal for financial, scientific, or any high-precision calculations.

## Features
- High-precision decimal operations: addition, subtraction, multiplication, division.
- Configurable rounding strategies.
- Seamless integration with Go’s big and math/big standards (if applicable).
- Fluent API for chaining operations (optional description).

## Installation

```bash
go get github.com/Lucy-Teknologi/lucy-decimal
```

## Usage

```go
package main

import (
  "fmt"
  "github.com/Lucy-Teknologi/lucy-decimal"
)

func main() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromInt(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875
}
```
(Adjust constructor names, options, and rounding methods to match the actual API.)


