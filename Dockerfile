# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . ./

ENTRYPOINT ["go", "test", "-v", "./...", "-coverprofile", "cover.out", "-bench=.", "-benchmem"]