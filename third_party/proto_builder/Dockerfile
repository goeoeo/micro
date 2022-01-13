# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

FROM golang:1.16.5

ENV GO111MODULE=on
ENV GOPROXY=http://goproxy.cn

RUN  sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN  apt-get clean

RUN apt-get update && apt-get install -y git curl openssl unzip

# go tag generate
RUN go get github.com/favadi/protoc-go-inject-tag@v1.3.0
# goimports
RUN go get golang.org/x/tools/cmd/goimports

# grpc-gateway
RUN go get \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.6.0 \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.6.0 \
    google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1 \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0


# swagger-0.28.0
# RUN go get github.com/go-swagger/go-swagger/cmd/swagger
RUN mkdir -p /swagger && cd /swagger \
	&& wget https://github.com/go-swagger/go-swagger/releases/download/v0.28.0/swagger_linux_amd64 \
	&& chmod +x swagger_linux_amd64 && mv swagger_linux_amd64 /go/bin/swagger

# the protoc can't run on alpine,
# we only need the protobuf's stdarnd library in the `/protoc/include`.
RUN mkdir -p /protoc && cd /protoc \
	&& wget https://github.com/google/protobuf/releases/download/v3.19.1/protoc-3.19.1-linux-x86_64.zip \
	&& unzip protoc-3.19.1-linux-x86_64.zip

RUN chmod +x /protoc/bin/protoc && mv /protoc/bin/protoc /go/bin && cp -r /protoc/include/* /usr/local/include

RUN apt-get update && apt-get install -y make  jq upx git
COPY . /proto_builder/

ENV CGO_ENABLED=0
WORKDIR /proto_builder
CMD ["make","run"]


