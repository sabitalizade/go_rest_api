# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /main

EXPOSE 4000

CMD [ "/main" ]
