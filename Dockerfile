FROM golang:1.17-alpine

RUN apk add build-base

WORKDIR /app

COPY go.mod .
COPY go.sum .

ENV GO111MODULE=on

RUN go mod download

COPY . .

RUN go build -o /app/personal-tracker

EXPOSE 8001

CMD ["/app/personal-tracker"]