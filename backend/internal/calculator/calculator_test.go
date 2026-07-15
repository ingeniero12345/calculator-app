package calculator

import (
	"errors"
	"math"
	"testing"
)

func TestBinaryOperations(t *testing.T) {
	tests := []struct {
		name    string
		fn      func(a, b float64) (float64, error)
		a, b    float64
		want    float64
		wantErr error
	}{
		{"add positives", Add, 2, 3, 5, nil},
		{"add negatives", Add, -2, -3, -5, nil},
		{"add zero", Add, 0, 0, 0, nil},
		{"subtract", Subtract, 10, 4, 6, nil},
		{"subtract negative result", Subtract, 4, 10, -6, nil},
		{"multiply", Multiply, 6, 7, 42, nil},
		{"multiply by zero", Multiply, 6, 0, 0, nil},
		{"divide", Divide, 20, 4, 5, nil},
		{"divide fractional", Divide, 1, 8, 0.125, nil},
		{"divide by zero", Divide, 1, 0, 0, ErrDivisionByZero},
		{"power", Power, 2, 10, 1024, nil},
		{"power zero exponent", Power, 123, 0, 1, nil},
		{"power negative exponent", Power, 2, -1, 0.5, nil},
		{"percentage", Percentage, 200, 10, 20, nil},
		{"percentage of zero", Percentage, 0, 50, 0, nil},
		{"multiply overflow", Multiply, math.MaxFloat64, 2, 0, ErrNotFinite},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.fn(tc.a, tc.b)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("error = %v, want %v", err, tc.wantErr)
			}
			if tc.wantErr == nil && got != tc.want {
				t.Fatalf("result = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		want    float64
		wantErr error
	}{
		{"perfect square", 144, 12, nil},
		{"zero", 0, 0, nil},
		{"non-perfect", 2, math.Sqrt2, nil},
		{"negative", -1, 0, ErrNegativeSquareRoot},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Sqrt(tc.a)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("error = %v, want %v", err, tc.wantErr)
			}
			if tc.wantErr == nil && got != tc.want {
				t.Fatalf("result = %v, want %v", got, tc.want)
			}
		})
	}
}
