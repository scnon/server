FROM golang:1.19 as builder
WORKDIR /go/src/app

COPY . .

ADD https://github.com/protocolbuffers/protobuf/releases/download/v21.11/protoc-21.11-linux-x86_64.zip /tmp/protoc.zip
RUN apt update && apt install zip -y && unzip /tmp/protoc.zip -d /usr/local && make deps && make all && ls bin

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/app/bin/server /root/app
ENTRYPOINT [ "/root/app" ]