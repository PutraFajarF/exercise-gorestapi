# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/go-docker main.go

EXPOSE 1234

ENTRYPOINT [ "/app/go-docker" ]