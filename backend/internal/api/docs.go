package api

import (
	_ "embed"
	"net/http"
)

//go:embed docs/openapi.yaml
var openAPISpec []byte

// swaggerUIPage is a self-contained Swagger UI shell that renders the embedded
// OpenAPI document. The UI assets are loaded from a CDN at view time; the API
// itself keeps zero third-party dependencies (stdlib only).
const swaggerUIPage = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Calculator API — Docs</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.ui = SwaggerUIBundle({
      url: "/api/v1/openapi.yaml",
      dom_id: "#swagger-ui",
    });
  </script>
</body>
</html>`

func handleOpenAPISpec(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/yaml")
	_, _ = w.Write(openAPISpec)
}

func handleDocs(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(swaggerUIPage))
}
