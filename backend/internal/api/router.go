package api

import (
	"encoding/json"
	"net/http"
)

// NewRouter wires up all routes and cross-cutting middleware and returns the
// http.Handler ready to be served.
func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Operation endpoints, e.g. POST /api/v1/add.
	mux.HandleFunc("POST /api/v1/{operation}", Calculate)

	// Lightweight liveness probe for container orchestration.
	mux.HandleFunc("GET /api/v1/health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	// Self-describing endpoint listing the supported operations.
	mux.HandleFunc("GET /api/v1/operations", listOperations)

	return withMiddleware(mux)
}

// listOperations returns the registry as JSON so clients can discover what the
// API supports without reading the docs.
func listOperations(w http.ResponseWriter, _ *http.Request) {
	type opInfo struct {
		Name  string `json:"name"`
		Unary bool   `json:"unary"`
	}
	list := make([]opInfo, 0, len(operations))
	for name, op := range operations {
		list = append(list, opInfo{Name: name, Unary: op.unary})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]any{"operations": list})
}

// withMiddleware applies logging-free, dependency-free cross-cutting concerns:
// permissive CORS (so the SPA can call the API from another origin) and JSON
// content negotiation.
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
