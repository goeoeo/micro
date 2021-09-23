# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

# TODO 可改成私有仓库地址
DOCKER_REGISTRY:="micro"
DOCKER_NAMESPACE:="micro"
DOCKER_USER:=docker-user
DOCKER_PW:=123456
PROTO_BUILDER_IMAGE_NAME:=$(DOCKER_REGISTRY)/$(DOCKER_NAMESPACE)/proto-builder:latest


default:
	docker build -t $(PROTO_BUILDER_IMAGE_NAME) .

	# TODO 可推送至私有仓库
	# docker login -u $(DOCKER_USER) -p $(DOCKER_PW)  $(DOCKER_REGISTRY)
	# docker push $(BUILDER_IMAGE_LATEST)
	@echo "ok"

run:
	docker run --rm -it -v `pwd`:/root $(PROTO_BUILDER_IMAGE_NAME)

clean:
	@echo "ok"

