FROM golang:1.14-alpine AS builder
WORKDIR /app

RUN apk add git dumb-init

# cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# build code
COPY main.go leader.go ./
RUN go build .

# using dumb-init to not run as pid 1 like
ENTRYPOINT ["dumb-init", "/app/kube-leader", "kube-leader-example"]
