FROM golang:1.23.2 AS builder
LABEL authors="Aprian Firlanda Imani"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 3000

CMD ["./main"]
