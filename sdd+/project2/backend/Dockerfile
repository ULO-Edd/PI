# Dockerfile for Go API

# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/api

# Final stage
FROM alpine:latest

WORKDIR /

COPY --from=builder /api /api
COPY --from=builder /app/db /db

EXPOSE 8080

ENTRYPOINT ["/api"]
