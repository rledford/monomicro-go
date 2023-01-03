VERSION 0.6
FROM golang:1.19-alpine3.17
WORKDIR /go-workdir

RUN apk add build-base

deps:
  COPY go.mod go.sum ./
  RUN go mod download
  SAVE ARTIFACT go.mod AS LOCAL go.mod
  SAVE ARTIFACT go.sum AS LOCAL go.sum

proto-deps:
  FROM golang:buster
  RUN apt-get update && apt-get install -y wget unzip
  RUN wget -O protoc.zip \
      "https://github.com/protocolbuffers/protobuf/releases/download/v3.13.0/protoc-3.13.0-$(uname -s)-$(uname -m).zip"
  RUN unzip protoc.zip -d /usr/local/
  RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
  RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
  RUN VERSION="1.11.0" && \
      curl -sSL "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" -o "/usr/local/bin/buf" 
  RUN chmod +x "/usr/local/bin/buf"
