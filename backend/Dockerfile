# --- Build stage -------------------------------------------------------------
FROM golang:1.22-alpine AS build
WORKDIR /src

# Cache dependencies first (go.sum is absent because we use only the stdlib).
COPY go.mod ./
RUN go mod download

COPY . .
# Build a small, static binary.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/calculator ./

# --- Run stage ---------------------------------------------------------------
FROM gcr.io/distroless/static-debian12 AS run
COPY --from=build /bin/calculator /calculator
EXPOSE 8080
ENV PORT=8080
USER nonroot:nonroot
ENTRYPOINT ["/calculator"]
