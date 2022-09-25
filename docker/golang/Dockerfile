FROM golang:1.19-alpine

WORKDIR /go/src/app

COPY src/go.mod ./
COPY src/go.sum ./
RUN go mod download

RUN export PATH=$PATH:$GOPATH
RUN go install github.com/codegangsta/gin

COPY src/ ./

EXPOSE 8080