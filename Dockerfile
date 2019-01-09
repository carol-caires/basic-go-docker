FROM golang:1.11-stretch

RUN mkdir /go/src/backend/ && mkdir /go/src/logs/
RUN go get -u github.com/go-sql-driver/mysql
WORKDIR /go/src/backend