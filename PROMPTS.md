# AI Prompts Used

This project was built with the help of an AI coding assistant (Claude). Below
are the key prompts and how the AI tooling was used, as requested in the
assignment.

## Primary prompt

> Build a full-stack calculator application with a React frontend and a backend
> microservice. The frontend should consume the backend API to perform basic and
> advanced arithmetic operations. Focus on clean design, maintainable code, and
> testable architecture.
>
> **Functional:** Operations — addition, subtraction, multiplication, division;
> optionally exponentiation, square root, percentage. React frontend with an
> intuitive UI, input validation, error handling, and responsive/mobile support.
> REST backend exposing endpoints for the operations, validating input and
> handling edge cases (division by zero, invalid data), returning JSON.
>
> **Non-functional:** Clean, idiomatic code on both layers; unit tests covering
> key functionality; documentation (setup, API usage, design rationale);
> optionally a Dockerfile for full-stack deployment.
>
> **Constraints:** Frontend in React (TypeScript preferred); backend in Go
> (preferred). Deliverables: git repo with both layers, README with setup / API
> examples / design decisions, unit tests + coverage, optional Docker to run both
> together.

## How the AI was directed

The work was decomposed into tracked steps and executed in order:

1. **Environment check** — verified available tooling (Node present; Go installed
   via Homebrew; chose the stdlib to avoid extra dependencies).
2. **Backend logic first** — "Create a pure `calculator` Go package with
   Add/Subtract/Multiply/Divide/Power/Sqrt/Percentage, returning sentinel errors
   for division by zero, negative square roots, and non-finite results."
3. **HTTP layer** — "Wrap the logic in a `net/http` (Go 1.22) API using an
   operation registry map and `POST /api/v1/{operation}`; validate the JSON body,
   distinguish 400 vs 404 vs 422, and add health/operations/CORS."
4. **Backend tests** — "Write table-driven tests for the calculator (aim for
   100%) and handler tests covering success and every error path; produce a
   coverage report."
5. **Frontend** — "Scaffold a Vite React+TS app. Share operation metadata in one
   module that drives both the UI and validation. Add a typed fetch client that
   surfaces server error messages and network failures. Build an accessible,
   responsive Calculator component with per-field validation and light/dark
   theming."
6. **Frontend tests** — "Use Vitest + Testing Library to cover the validation
   helpers, the API client (success / domain error / network failure), and the
   component behaviour (calculation, validation errors, backend errors, unary vs
   binary rendering)."
7. **Packaging & docs** — "Add multi-stage Dockerfiles (distroless Go binary;
   nginx-served SPA with an API reverse proxy), a docker-compose to run both, and
   a README covering setup, API examples, and design decisions."
8. **Verification** — ran `go vet`, `go test -cover`, `tsc`, and
   `vitest --coverage`, plus a production `vite build`, and fixed anything that
   failed before finishing.

## Guiding principles given to the AI

- Keep business logic transport-agnostic and fully unit-tested.
- Prefer the standard library / minimal dependencies over frameworks.
- Validate on both client (UX) and server (source of truth).
- Use meaningful HTTP status codes and a single consistent error envelope.
- Favour readability and idiomatic style over cleverness or extra features.
