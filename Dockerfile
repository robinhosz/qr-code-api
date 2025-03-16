FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o qr-code-api .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/qr-code-api .

EXPOSE 8080

CMD ["./qr-code-api"]