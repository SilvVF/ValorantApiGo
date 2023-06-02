FROM golang:1.19.3

WORKDIR /usr/src/app

COPY . .

RUN go install github.com/cosmtrek/air@latest

RUN go mod tidy
