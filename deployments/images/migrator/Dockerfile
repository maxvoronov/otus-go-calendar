FROM golang:1.13-alpine

ENV GOBIN "/usr/local/bin/"
RUN apk --no-cache add g++ gcc git

RUN go get -u github.com/pressly/goose/cmd/goose

WORKDIR /migrations

ENTRYPOINT ["/usr/local/bin/goose"]