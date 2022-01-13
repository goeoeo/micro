# Copyright 2019 The Newbilling Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

PROTO_BUILDER=/proto_builder
GENERATE_PB_DIR=./generate/pb
GENERATE_Swagger_DIR=./generate/swagger
GENERATE_SWAGGER_DIR=$(PROTO_BUILDER)/swagger

PROTO_FILES=$(sort $(wildcard ./*.proto))
PROTOC_INC_PATH=/usr/local/include
PROTOC_FLAGS:=-I. -I$(PROTO_BUILDER)/proto  -I$(PROTOC_INC_PATH)

generate: $(PROTO_FILES) Makefile
	@rm -rf $(GENERATE_PB_DIR)/*
	@mkdir -p  $(GENERATE_PB_DIR)

    #To generate pb.go file
	protoc $(PROTOC_FLAGS) --go-grpc_out=paths=source_relative:$(GENERATE_PB_DIR) ${PROTO_FILES}
	protoc $(PROTOC_FLAGS) --go_out=paths=source_relative:$(GENERATE_PB_DIR) ${PROTO_FILES}


	#To generate pb.gw.go file
	protoc $(PROTOC_FLAGS) --grpc-gateway_out=logtostderr=true,paths=source_relative,generate_unbound_methods=true,allow_delete_body=true:$(GENERATE_PB_DIR) ${PROTO_FILES}

	#generate api.swagger.json file
	protoc $(PROTOC_FLAGS) --openapiv2_out=logtostderr=true,json_names_for_fields=false,allow_delete_body=true:$(GENERATE_SWAGGER_DIR)/spec ${PROTO_FILES}


	#注入tag
	protoc-go-inject-tag -input=$(GENERATE_PB_DIR)/*.pb.go

	cd $(GENERATE_SWAGGER_DIR)/spec/ && cat ./*.swagger.json  | jq --slurp 'reduce .[] as $$item ({}; . * $$item)' | jq -f ./preprocess.jq > ./api.swagger.json

	#fix 生成的./api.swagger.json为空的bug，如果为空，则一直重复执行生成。
	@while [ -s  ./api.swagger.json ];do \
 	cd $(GENERATE_SWAGGER_DIR)/spec/ && cat ./*.swagger.json | jq --slurp 'reduce .[] as $$item ({}; . * $$item)' | jq -f ./preprocess.jq > ./api.swagger.json; \
	done
	cd $(GENERATE_SWAGGER_DIR)/spec/ && find . | grep .swagger.json | grep -v "api" | xargs rm

	make -C $(GENERATE_SWAGGER_DIR)/spec
	make -C $(GENERATE_SWAGGER_DIR)/swagger-ui

	# copy
	mkdir -p $(GENERATE_Swagger_DIR)/spec
	mkdir -p $(GENERATE_Swagger_DIR)/swagger-ui

	mv $(GENERATE_SWAGGER_DIR)/spec/api.swagger.json  $(GENERATE_Swagger_DIR)/spec
	mv $(GENERATE_SWAGGER_DIR)/spec/static.go  $(GENERATE_Swagger_DIR)/spec
	mv $(GENERATE_SWAGGER_DIR)/swagger-ui/static.go $(GENERATE_Swagger_DIR)/swagger-ui


	goimports -l -w -e  $(GENERATE_PB_DIR)
	gofmt -l -w -e -s $(GENERATE_PB_DIR)

run:
	cp Makefile api/
	cd api && make generate
	rm api/Makefile