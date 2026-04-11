FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/core/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 3000

CMD ["./app"]