FROM golang:1.14-alpine AS builder
WORKDIR /app

RUN apk add git

COPY go.mod go.sum main.go leader.go ./
RUN go build .

ENTRYPOINT ["/app/kube-leader", "kube-leader-example"]
