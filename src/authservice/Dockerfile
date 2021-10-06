# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /service .

##
## Deploy
##
FROM alpine

WORKDIR /

COPY --from=builder /service ./server

ENV PORT=50051
EXPOSE 50051

ENTRYPOINT ["/server"]
