package service

import (
	"fmt"
	"math"
	"sort"

	"github.com/linktic/calculator-app/backend/internal/calculator"
)

var (
	ErrUnknownOperation = fmt.Errorf("unknown operation")
	ErrMissingOperand   = fmt.Errorf("missing required operand")
	ErrInvalidOperand   = fmt.Errorf("operands must be finite numbers")
)

type strategy struct {
	compute func(a, b float64) (float64, error)
	unary   bool
}

type Result struct {
	Operation string
	A         float64
	B         *float64
	Value     float64
}

type OperationInfo struct {
	Name  string
	Unary bool
}

type Calculator struct {
	strategies map[string]strategy
}

func NewCalculator() *Calculator {
	return &Calculator{
		strategies: map[string]strategy{
			"add":        {compute: calculator.Add},
			"subtract":   {compute: calculator.Subtract},
			"multiply":   {compute: calculator.Multiply},
			"divide":     {compute: calculator.Divide},
			"power":      {compute: calculator.Power},
			"percentage": {compute: calculator.Percentage},
			"sqrt":       {compute: sqrtStrategy, unary: true},
		},
	}
}

func sqrtStrategy(a, _ float64) (float64, error) {
	return calculator.Sqrt(a)
}

func (c *Calculator) Compute(operation string, a, b *float64) (Result, error) {
	selected, known := c.strategies[operation]
	if !known {
		return Result{}, fmt.Errorf("%w: %s", ErrUnknownOperation, operation)
	}
	if a == nil {
		return Result{}, fmt.Errorf("%w: 'a'", ErrMissingOperand)
	}
	if !selected.unary && b == nil {
		return Result{}, fmt.Errorf("%w: '%s' requires 'b'", ErrMissingOperand, operation)
	}
	if !isFinite(a) || !isFinite(b) {
		return Result{}, ErrInvalidOperand
	}

	value, err := selected.compute(*a, secondOperand(b))
	if err != nil {
		return Result{}, err
	}

	result := Result{Operation: operation, A: *a, Value: value}
	if !selected.unary {
		result.B = b
	}
	return result, nil
}

func (c *Calculator) Operations() []OperationInfo {
	operations := make([]OperationInfo, 0, len(c.strategies))
	for name, s := range c.strategies {
		operations = append(operations, OperationInfo{Name: name, Unary: s.unary})
	}
	sort.Slice(operations, func(i, j int) bool {
		return operations[i].Name < operations[j].Name
	})
	return operations
}

func isFinite(operand *float64) bool {
	return operand == nil || (!math.IsNaN(*operand) && !math.IsInf(*operand, 0))
}

func secondOperand(b *float64) float64 {
	if b == nil {
		return 0
	}
	return *b
}
