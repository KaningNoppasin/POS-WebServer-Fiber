FROM golang:1.23-alpine3.21 AS builder

WORKDIR /web-server-fiber

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /web-server-fiber .

EXPOSE 8080

CMD ["./main"]