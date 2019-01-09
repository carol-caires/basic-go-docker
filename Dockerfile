FROM golang:1.11-stretch

ENV GOBIN /go/bin
RUN mkdir /go/src/backend/ && mkdir /go/src/logs/
WORKDIR /go/src/backend