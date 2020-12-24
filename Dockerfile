FROM golang:1.15.2-alpine

RUN apk add -u git protobuf
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app

RUN go get google.golang.org/grpc && \
  go get github.com/golang/protobuf/protoc-gen-go && \
  go get github.com/payjp/payjp-go/v1

EXPOSE 50101
