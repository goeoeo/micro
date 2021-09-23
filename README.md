# micro

集工程实践搭建的golang微服务框架  
目录结构参考[https://github.com/golang-standards/project-layout ](https://github.com/golang-standards/project-layout )
每个组件都由较流行的开源组件组成，站在巨人的肩膀上开发。  



## 工程目录介绍

### api 
OpenAPI/Swagger 规范，JSON 模式文件，协议定义文件。    
* generate 由编译插件生成的文件
* google 相关proto文件
* proto 存放proto源文件
* protoc-gen-swagger swagger文档相关文件

### build
打包和持续集成  
将你的云( AMI )、容器( Docker )、操作系统( deb、rpm、pkg )包配置和脚本放在 /build/package 目录下。  

#### build/package/proto-builder
proto文件编译，docker容器  
基于golang:1.15-alpine3.12基镜像构建   




### configs 
存放配置文件  





## grpc 相关编译插件 
### 组件
* protoc #grpc proto文件编译工具
* protoc-gen-go #生成go代码
* protoc-gen-grpc-gateway # 
* protoc-gen-swagger # 生成swagger文档

### 生成3个文件
* pb.go 由protoc-gen-go插件生成
* pb.gw.go 由protoc-gen-grpc-gateway 插件生成
* swagger.json 由protoc-gen-swagger 插件生成


## pkg 介绍

### pkg/config
http://github.com/spf13/viper  

配置对象包含以下功能:
* 从配置文件解析配置
* 环境变量可以覆盖配置 （类容器方式）
* 配置文件路径，可由flag参数解析
* 热更新


