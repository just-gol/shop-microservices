package utils

import "github.com/shopspring/decimal"

func Sub(a int, b int) int {
	return a - b
}

func Multiplication(price decimal.Decimal, num int) float64 {
	f, _ := price.Mul(decimal.NewFromInt(int64(num))).Float64()
	return f
}

func MulFloat(price float64, num int) float64 {
	return price * float64(num)
}
