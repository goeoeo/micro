# micro

集工程实践搭建的golang微服务框架  
目录结构参考[https://github.com/golang-standards/project-layout ](https://github.com/golang-standards/project-layout )
每个组件都由较流行的开源组件组成，站在巨人的肩膀上开发。  

## api 
OpenAPI/Swagger 规范，proto协议定义文件。  

## configs 
存放配置文件  
http://github.com/spf13/viper  
### pkg/config
配置对象包含以下功能: 
* 从配置文件解析配置
* 环境变量可以覆盖配置 （类容器方式）
* 配置文件路径，可由flag参数解析
* 热更新
