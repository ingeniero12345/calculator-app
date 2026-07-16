package calculator

import (
	"errors"
	"math"
)

var (
	ErrDivisionByZero     = errors.New("division by zero is undefined")
	ErrNegativeSquareRoot = errors.New("square root of a negative number is undefined")
	ErrNotFinite          = errors.New("result is not a finite number")
)

func Add(a, b float64) (float64, error) {
	return finite(a + b)
}

func Subtract(a, b float64) (float64, error) {
	return finite(a - b)
}

func Multiply(a, b float64) (float64, error) {
	return finite(a * b)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return finite(a / b)
}

func Power(a, b float64) (float64, error) {
	return finite(math.Pow(a, b))
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, ErrNegativeSquareRoot
	}
	return finite(math.Sqrt(a))
}

const percentDivisor = 100

func Percentage(a, b float64) (float64, error) {
	return finite(a * b / percentDivisor)
}

func finite(value float64) (float64, error) {
	if math.IsInf(value, 0) || math.IsNaN(value) {
		return 0, ErrNotFinite
	}
	return value, nil
}
