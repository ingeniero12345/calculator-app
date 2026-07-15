// Package api exposes the calculator business logic over a small JSON/HTTP API.
package api

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"

	"github.com/linktic/calculator-app/backend/internal/calculator"
)

// operation describes a single arithmetic operation the API exposes.
type operation struct {
	// fn performs the calculation. For unary operations (e.g. sqrt) the second
	// operand is ignored.
	fn func(a, b float64) (float64, error)
	// unary reports whether the operation uses only operand a, which lets us
	// validate requests and document the API accurately.
	unary bool
}

// operations is the registry mapping URL-friendly names to their implementation.
// Adding a new operation is a one-line change here.
var operations = map[string]operation{
	"add":        {fn: calculator.Add},
	"subtract":   {fn: calculator.Subtract},
	"multiply":   {fn: calculator.Multiply},
	"divide":     {fn: calculator.Divide},
	"power":      {fn: calculator.Power},
	"percentage": {fn: calculator.Percentage},
	"sqrt":       {fn: func(a, _ float64) (float64, error) { return calculator.Sqrt(a) }, unary: true},
}

// CalculationRequest is the JSON body accepted by the operation endpoints.
// Pointers are used so we can distinguish an omitted field from an explicit 0.
type CalculationRequest struct {
	A *float64 `json:"a"`
	B *float64 `json:"b"`
}

// CalculationResponse is returned on a successful calculation.
type CalculationResponse struct {
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         *float64 `json:"b,omitempty"`
	Result    float64 `json:"result"`
}

// ErrorResponse is the uniform error envelope returned for all failures.
type ErrorResponse struct {
	Error string `json:"error"`
}

// Calculate handles POST /api/v1/{operation}. It decodes and validates the
// request, dispatches to the matching operation and writes a JSON response.
func Calculate(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("operation")
	op, ok := operations[name]
	if !ok {
		writeError(w, http.StatusNotFound, "unknown operation: "+name)
		return
	}

	var req CalculationRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body: "+err.Error())
		return
	}

	if req.A == nil {
		writeError(w, http.StatusBadRequest, "field 'a' is required")
		return
	}
	if !op.unary && req.B == nil {
		writeError(w, http.StatusBadRequest, "field 'b' is required for operation: "+name)
		return
	}
	if err := validateFinite(req.A, req.B); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	var b float64
	if req.B != nil {
		b = *req.B
	}

	result, err := op.fn(*req.A, b)
	if err != nil {
		// Domain errors (division by zero, etc.) are client-facing 422s: the
		// request was well-formed but mathematically undefined.
		writeError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	resp := CalculationResponse{Operation: name, A: *req.A, Result: result}
	if !op.unary {
		resp.B = req.B
	}
	writeJSON(w, http.StatusOK, resp)
}

// validateFinite rejects NaN/Inf operands up front so we never attempt a
// calculation on non-finite input.
func validateFinite(a, b *float64) error {
	for _, v := range []*float64{a, b} {
		if v == nil {
			continue
		}
		if math.IsNaN(*v) || math.IsInf(*v, 0) {
			return errors.New("operands must be finite numbers")
		}
	}
	return nil
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, ErrorResponse{Error: msg})
}
