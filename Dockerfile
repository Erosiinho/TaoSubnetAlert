# ---- Build Stage ----
FROM golang:1.23.3-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /app
COPY . .
RUN go mod tidy && \
    go build -o server cmd/main.go

# ---- Runtime Stage ----
FROM alpine:latest
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
WORKDIR /app
COPY --from=builder /app/server .
RUN chown -R appuser:appgroup /app
USER appuser
EXPOSE 8080

# Start app
CMD ["./server"]