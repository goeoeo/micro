# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

FROM golang:1.15-alpine3.12 as builder

ENV GO111MODULE=on
ENV GOPROXY=http://goproxy.cn

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache git curl openssl

RUN go get -u -v github.com/golang/protobuf/protoc-gen-go@v1.3.2
# go tag generate
RUN go get -u -v github.com/favadi/protoc-go-inject-tag@v1.3.0
RUN go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.14.7
RUN go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v1.14.7
RUN go get golang.org/x/tools/cmd/goimports

# swagger-0.13.0
# RUN go get github.com/go-swagger/go-swagger/cmd/swagger
RUN mkdir -p /swagger && cd /swagger \
	&& wget https://github.com/go-swagger/go-swagger/releases/download/v0.25.0/swagger_linux_amd64 \
	&& chmod +x swagger_linux_amd64 && mv swagger_linux_amd64 /go/bin/swagger

# the protoc can't run on alpine,
# we only need the protobuf's stdarnd library in the `/protoc/include`.
RUN mkdir -p /protoc && cd /protoc \
	&& wget https://github.com/google/protobuf/releases/download/v3.5.0/protoc-3.5.0-linux-x86_64.zip \
	&& unzip protoc-3.5.0-linux-x86_64.zip


FROM golang:1.15-alpine3.12

ENV GO111MODULE=on

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache git protobuf make curl openssl jq rsync upx

COPY --from=builder /protoc/include /usr/local/include
COPY --from=builder /go/bin /go/bin
