FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o main ./cmd/api/main.go

FROM debian:bullseye-slim

WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
