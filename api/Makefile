# Copyright 2019 The Newbilling Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

GOPATH:=$(shell go env GOPATH)
PWD:=$(shell pwd)

GENERATE_PB_DIR=generate/pb
GENERATE_SWAGGER_DIR=generate/swagger

PROTO_FILES=$(sort $(wildcard ./*.proto))
PROTOC_INC_PATH=/usr/local/include
GRPC_GATEWAY_PATH=$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.7
GOOGLEAPIS_PATH=$(GRPC_GATEWAY_PATH)/third_party/googleapis

PROTOC_FLAGS:=-I. -I$(GOOGLEAPIS_PATH) -I$(GRPC_GATEWAY_PATH) -I$(PROTOC_INC_PATH)
# proto_builder 容器镜像
PROTO_BUILDER_IMAGE_NAME:=micro/micro/proto-builder:latest

generate: $(PROTO_FILES) Makefile
	@rm -rf $(GENERATE_PB_DIR)/*

    #To generate newbilling.pb.go file
	protoc $(PROTOC_FLAGS) --go_out=plugins=grpc,paths=source_relative:$(GENERATE_PB_DIR) ${PROTO_FILES}

	#注入tag
	protoc-go-inject-tag -input=$(GENERATE_PB_DIR)/*.pb.go

	#To generate pb.gw.go file
	protoc $(PROTOC_FLAGS) --grpc-gateway_out=logtostderr=true,allow_delete_body=true:$(GENERATE_PB_DIR) ${PROTO_FILES}

	#generate swagger.json file
	protoc $(PROTOC_FLAGS) --swagger_out=logtostderr=true,allow_delete_body=true:$(GENERATE_SWAGGER_DIR)/spec ${PROTO_FILES}

	cd $(GENERATE_SWAGGER_DIR)/spec/ && cat ./*.swagger.json 0.swagger.json | jq --slurp 'reduce .[] as $$item ({}; . * $$item)' | jq -f ./preprocess.jq > ./api.swagger.json

	#fix 生成的./api.swagger.json为空的bug，如果为空，则一直重复执行生成。
	@while [ -s  ./api.swagger.json ];do \
 	cd $(GENERATE_SWAGGER_DIR)/spec/ && cat ./*.swagger.json 0.swagger.json | jq --slurp 'reduce .[] as $$item ({}; . * $$item)' | jq -f ./preprocess.jq > ./api.swagger.json; \
	done
	cd $(GENERATE_SWAGGER_DIR)/spec/ && find . | grep .swagger.json | grep -v "api" | xargs rm
	make -C $(GENERATE_SWAGGER_DIR)/spec
	make -C $(GENERATE_SWAGGER_DIR)/swagger-ui


	goimports -l -w -e  $(GENERATE_PB_DIR)
	gofmt -l -w -e -s $(GENERATE_PB_DIR)


build: Makefile
	docker run -i -v $(GOPATH)/pkg/mod:/go/pkg/mod -v `pwd`:/go/project -w /go/project $(PROTO_BUILDER_IMAGE_NAME)  make generate



