# 包来源
## google/api
项目:https://github.com/googleapis/googleapis  
https://github.com/googleapis/googleapis/tree/master/google/api

## protoc-gen-openapiv2
项目:https://github.com/grpc-ecosystem/grpc-gateway  
https://github.com/grpc-ecosystem/grpc-gateway/tree/master/protoc-gen-openapiv2



# 构建镜像
```
 docker build -t dockerhub.qingcloud.com/qingbilling/nb-proto-builder:latest .
```

# 使用镜像
```
docker run -i -v `pwd`:/proto_builder/api  dockerhub.qingcloud.com/qingbilling/nb-proto-builder:latest
```