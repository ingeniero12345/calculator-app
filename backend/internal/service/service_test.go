package service

import (
	"errors"
	"math"
	"testing"

	"github.com/linktic/calculator-app/backend/internal/calculator"
)

func ptr(v float64) *float64 { return &v }

func TestComputeSuccess(t *testing.T) {
	svc := NewCalculator()
	tests := []struct {
		name      string
		operation string
		a         *float64
		b         *float64
		want      float64
		wantB     bool
	}{
		{"add", "add", ptr(2), ptr(3), 5, true},
		{"subtract", "subtract", ptr(10), ptr(4), 6, true},
		{"multiply", "multiply", ptr(6), ptr(7), 42, true},
		{"divide", "divide", ptr(20), ptr(4), 5, true},
		{"power", "power", ptr(2), ptr(10), 1024, true},
		{"percentage", "percentage", ptr(200), ptr(10), 20, true},
		{"sqrt ignores b", "sqrt", ptr(144), ptr(999), 12, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := svc.Compute(tc.operation, tc.a, tc.b)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.Value != tc.want {
				t.Fatalf("value = %v, want %v", result.Value, tc.want)
			}
			if (result.B != nil) != tc.wantB {
				t.Fatalf("B present = %v, want %v", result.B != nil, tc.wantB)
			}
		})
	}
}

func TestComputeErrors(t *testing.T) {
	svc := NewCalculator()
	tests := []struct {
		name      string
		operation string
		a         *float64
		b         *float64
		wantErr   error
	}{
		{"unknown operation", "modulo", ptr(1), ptr(2), ErrUnknownOperation},
		{"missing a", "add", nil, ptr(2), ErrMissingOperand},
		{"missing b for binary", "add", ptr(2), nil, ErrMissingOperand},
		{"non-finite a", "add", ptr(math.Inf(1)), ptr(2), ErrInvalidOperand},
		{"non-finite b", "add", ptr(2), ptr(math.NaN()), ErrInvalidOperand},
		{"division by zero", "divide", ptr(1), ptr(0), calculator.ErrDivisionByZero},
		{"negative sqrt", "sqrt", ptr(-4), nil, calculator.ErrNegativeSquareRoot},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := svc.Compute(tc.operation, tc.a, tc.b)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("error = %v, want %v", err, tc.wantErr)
			}
		})
	}
}

func TestOperationsAreSortedAndComplete(t *testing.T) {
	svc := NewCalculator()
	operations := svc.Operations()
	want := []string{"add", "divide", "multiply", "percentage", "power", "sqrt", "subtract"}

	if len(operations) != len(want) {
		t.Fatalf("got %d operations, want %d", len(operations), len(want))
	}
	for i, name := range want {
		if operations[i].Name != name {
			t.Fatalf("operations[%d] = %s, want %s", i, operations[i].Name, name)
		}
	}
}
