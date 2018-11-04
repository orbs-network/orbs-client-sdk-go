FROM golang:1.10-alpine

WORKDIR /go/src/github.com/orbs-network/orbs-client-sdk-go/

ADD . /go/src/github.com/orbs-network/orbs-client-sdk-go/

RUN env

RUN go env

RUN ./test.sh
