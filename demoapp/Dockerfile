FROM golang:1.9-alpine

ADD ./ /go/src/github.com/skuid/nrdemo/demoapp

WORKDIR /go/src/github.com/skuid/nrdemo/demoapp

RUN go install

ENV CGO_ENABLED 1

ENTRYPOINT ["/go/bin/demoapp"]
