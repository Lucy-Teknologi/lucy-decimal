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
  a := lucy_decimal.New("123.45")
  b := lucy_decimal.New("6.789")

  sum := a.Add(b)
  diff := a.Sub(b)
  prod := a.Mul(b)
  quot, err := a.Div(b, lucy_decimal.WithPrecision(10), lucy_decimal.WithRounding(lucy_decimal.RoundHalfUp))
  if err != nil {
    // handle error
  }

  fmt.Println("Sum:     ", sum)
  fmt.Println("Diff:    ", diff)
  fmt.Println("Product: ", prod)
  fmt.Println("Quotient:", quot)
}
```
(Adjust constructor names, options, and rounding methods to match the actual API.)

