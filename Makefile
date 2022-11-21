
# TODO 可改成私有仓库地址
DOCKER_REGISTRY:="micro"
DOCKER_NAMESPACE:="micro"

# micro/micro/proto-builder:latest
PROTO_BUILDER_IMAGE_NAME:=$(DOCKER_REGISTRY)/$(DOCKER_NAMESPACE)/proto-builder:latest

RUN_IN_DOCKER:=docker run -i  -v `pwd`:/go/project -w /go/project $(PROTO_BUILDER_IMAGE_NAME)

default:
	git pull
	cd pkg && ./generate
	git add .
	git commit -m "code update"
	git push
