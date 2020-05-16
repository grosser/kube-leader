FROM golang:1.14-alpine AS builder
WORKDIR /app

RUN apk add git

# cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# build code
COPY main.go leader.go ./
RUN go build .

ENTRYPOINT ["/app/kube-leader", "kube-leader-example"]
