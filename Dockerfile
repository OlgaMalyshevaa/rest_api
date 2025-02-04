FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o file_rest_api ./cmd/main/app.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/file_rest_api .

ENV DATABASE_URL=postgres://postgres@db:5432/financial_db?sslmode=disable

CMD ["./file_rest_api"]
