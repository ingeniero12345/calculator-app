// Package calculator implements the core arithmetic operations exposed by the
// service. It is deliberately free of any HTTP or transport concerns so that
// the business logic can be unit tested in isolation and reused elsewhere.
package calculator

import (
	"errors"
	"math"
)

// Sentinel errors returned by the calculator. Callers (e.g. the HTTP layer) can
// use errors.Is to map these to the appropriate response codes.
var (
	// ErrDivisionByZero is returned by Divide when the divisor is zero.
	ErrDivisionByZero = errors.New("division by zero is undefined")
	// ErrNegativeSquareRoot is returned by Sqrt for negative operands, since we
	// only operate on real numbers.
	ErrNegativeSquareRoot = errors.New("square root of a negative number is undefined")
	// ErrNotFinite is returned when an operation produces a non-finite result
	// (e.g. overflow to infinity or NaN).
	ErrNotFinite = errors.New("result is not a finite number")
)

// Add returns a + b.
func Add(a, b float64) (float64, error) {
	return finite(a + b)
}

// Subtract returns a - b.
func Subtract(a, b float64) (float64, error) {
	return finite(a - b)
}

// Multiply returns a * b.
func Multiply(a, b float64) (float64, error) {
	return finite(a * b)
}

// Divide returns a / b, or ErrDivisionByZero when b is zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return finite(a / b)
}

// Power returns a raised to the power of b.
func Power(a, b float64) (float64, error) {
	return finite(math.Pow(a, b))
}

// Sqrt returns the square root of a, or ErrNegativeSquareRoot when a is
// negative.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, ErrNegativeSquareRoot
	}
	return finite(math.Sqrt(a))
}

// Percentage returns b percent of a (i.e. a * b / 100). For example
// Percentage(200, 10) == 20.
func Percentage(a, b float64) (float64, error) {
	return finite(a * b / 100)
}

// finite guards against results that overflow to +/-Inf or NaN, converting them
// into a well-defined error instead of leaking non-serialisable JSON values.
func finite(v float64) (float64, error) {
	if math.IsInf(v, 0) || math.IsNaN(v) {
		return 0, ErrNotFinite
	}
	return v, nil
}
