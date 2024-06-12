FROM golang:1.22.1

WORKDIR /app

COPY . .

RUN go install github.com/air-verse/air@latest

RUN alias air='$(go env GOPATH)/bin/air'