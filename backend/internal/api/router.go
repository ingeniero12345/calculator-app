package api

import (
	"net/http"

	"github.com/linktic/calculator-app/backend/internal/service"
)

func NewRouter(calculator *service.Calculator) http.Handler {
	handler := NewHandler(calculator)
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/{operation}", handler.Calculate)
	mux.HandleFunc("GET /api/v1/health", handleHealth)
	mux.HandleFunc("GET /api/v1/operations", handleOperations(calculator))

	return withMiddleware(mux)
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func handleOperations(calculator *service.Calculator) http.HandlerFunc {
	type operationView struct {
		Name  string `json:"name"`
		Unary bool   `json:"unary"`
	}
	return func(w http.ResponseWriter, _ *http.Request) {
		operations := calculator.Operations()
		views := make([]operationView, 0, len(operations))
		for _, operation := range operations {
			views = append(views, operationView{Name: operation.Name, Unary: operation.Unary})
		}
		writeJSON(w, http.StatusOK, map[string]any{"operations": views})
	}
}

func withMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
