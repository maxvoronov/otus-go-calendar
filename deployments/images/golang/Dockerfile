FROM golang:1.13-alpine

ENV GOBIN "/usr/local/bin/"
RUN apk --no-cache add g++ gcc git

# Hot reload tool
RUN go get github.com/cosmtrek/air
COPY deployments/images/golang/configs /etc/air

WORKDIR /app

ENTRYPOINT ["/usr/local/bin/air"]