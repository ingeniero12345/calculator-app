package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/linktic/calculator-app/backend/internal/calculator"
	"github.com/linktic/calculator-app/backend/internal/service"
)

type CalculationRequest struct {
	A *float64 `json:"a"`
	B *float64 `json:"b"`
}

type CalculationResponse struct {
	Operation string   `json:"operation"`
	A         float64  `json:"a"`
	B         *float64 `json:"b,omitempty"`
	Result    float64  `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Handler struct {
	calculator *service.Calculator
}

func NewHandler(calculator *service.Calculator) *Handler {
	return &Handler{calculator: calculator}
}

func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request) {
	var request CalculationRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body: "+err.Error())
		return
	}

	result, err := h.calculator.Compute(r.PathValue("operation"), request.A, request.B)
	if err != nil {
		writeError(w, statusForError(err), err.Error())
		return
	}

	writeJSON(w, http.StatusOK, toResponse(result))
}

func toResponse(result service.Result) CalculationResponse {
	return CalculationResponse{
		Operation: result.Operation,
		A:         result.A,
		B:         result.B,
		Result:    result.Value,
	}
}

func statusForError(err error) int {
	switch {
	case errors.Is(err, service.ErrUnknownOperation):
		return http.StatusNotFound
	case errors.Is(err, service.ErrMissingOperand), errors.Is(err, service.ErrInvalidOperand):
		return http.StatusBadRequest
	case errors.Is(err, calculator.ErrDivisionByZero),
		errors.Is(err, calculator.ErrNegativeSquareRoot),
		errors.Is(err, calculator.ErrNotFinite):
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ErrorResponse{Error: message})
}
