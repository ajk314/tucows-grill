FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o tucows-grill-service ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/tucows-grill-service .
EXPOSE 8080
CMD ["./tucows-grill-service"]