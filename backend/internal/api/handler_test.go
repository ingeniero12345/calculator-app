package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// doRequest is a helper that runs a request through the full router and returns
// the recorded response.
func doRequest(t *testing.T, method, path, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	NewRouter().ServeHTTP(rec, req)
	return rec
}

func TestCalculateSuccess(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		body       string
		wantResult float64
	}{
		{"add", "/api/v1/add", `{"a":2,"b":3}`, 5},
		{"subtract", "/api/v1/subtract", `{"a":10,"b":4}`, 6},
		{"multiply", "/api/v1/multiply", `{"a":6,"b":7}`, 42},
		{"divide", "/api/v1/divide", `{"a":20,"b":4}`, 5},
		{"power", "/api/v1/power", `{"a":2,"b":10}`, 1024},
		{"percentage", "/api/v1/percentage", `{"a":200,"b":10}`, 20},
		{"sqrt (unary)", "/api/v1/sqrt", `{"a":144}`, 12},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rec := doRequest(t, http.MethodPost, tc.path, tc.body)
			if rec.Code != http.StatusOK {
				t.Fatalf("status = %d, want 200 (body: %s)", rec.Code, rec.Body.String())
			}
			var resp CalculationResponse
			if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
				t.Fatalf("decode: %v", err)
			}
			if resp.Result != tc.wantResult {
				t.Fatalf("result = %v, want %v", resp.Result, tc.wantResult)
			}
		})
	}
}

func TestCalculateErrors(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		body       string
		wantStatus int
	}{
		{"division by zero", "/api/v1/divide", `{"a":1,"b":0}`, http.StatusUnprocessableEntity},
		{"negative sqrt", "/api/v1/sqrt", `{"a":-4}`, http.StatusUnprocessableEntity},
		{"unknown operation", "/api/v1/modulo", `{"a":1,"b":2}`, http.StatusNotFound},
		{"missing a", "/api/v1/add", `{"b":2}`, http.StatusBadRequest},
		{"missing b", "/api/v1/add", `{"a":2}`, http.StatusBadRequest},
		{"invalid json", "/api/v1/add", `{not json}`, http.StatusBadRequest},
		{"unknown field", "/api/v1/add", `{"a":1,"b":2,"c":3}`, http.StatusBadRequest},
		{"empty body", "/api/v1/add", ``, http.StatusBadRequest},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rec := doRequest(t, http.MethodPost, tc.path, tc.body)
			if rec.Code != tc.wantStatus {
				t.Fatalf("status = %d, want %d (body: %s)", rec.Code, tc.wantStatus, rec.Body.String())
			}
			var resp ErrorResponse
			if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
				t.Fatalf("decode error body: %v", err)
			}
			if resp.Error == "" {
				t.Fatal("expected non-empty error message")
			}
		})
	}
}

func TestHealthAndOperations(t *testing.T) {
	rec := doRequest(t, http.MethodGet, "/api/v1/health", "")
	if rec.Code != http.StatusOK {
		t.Fatalf("health status = %d, want 200", rec.Code)
	}

	rec = doRequest(t, http.MethodGet, "/api/v1/operations", "")
	if rec.Code != http.StatusOK {
		t.Fatalf("operations status = %d, want 200", rec.Code)
	}
	var payload struct {
		Operations []struct {
			Name string `json:"name"`
		} `json:"operations"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if len(payload.Operations) != len(operations) {
		t.Fatalf("got %d operations, want %d", len(payload.Operations), len(operations))
	}
}

func TestCORSPreflight(t *testing.T) {
	rec := doRequest(t, http.MethodOptions, "/api/v1/add", "")
	if rec.Code != http.StatusNoContent {
		t.Fatalf("status = %d, want 204", rec.Code)
	}
	if rec.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Fatal("missing CORS header")
	}
}
