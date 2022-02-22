package boot

import (
	"github.com/gin-gonic/gin"

	"git.internal.yunify.com/benchmark/benchmark/pkg/gateway/grpc_proxy"

	"git.internal.yunify.com/benchmark/benchmark/pkg/app"
	"git.internal.yunify.com/benchmark/benchmark/pkg/gateway"
	"git.internal.yunify.com/benchmark/benchmark/pkg/global"
)

type WebServer struct {
	configPath string
	name       string
	port       string

	components []global.Init
	options    []gateway.GatewayOption
	router     func(r *gin.Engine)
	//配置挂载后需要执行的函数
	afterConfigFun func() error

	//命令行解析工具
	Cmd *app.Cmd
}

func NewWebServer(name, port, configPath string) *WebServer {
	ws := &WebServer{
		configPath: configPath,
		name:       name,
		port:       port,
		//内置组件
		components: []global.Init{
			global.Log(),
			global.Registry(),
		},
	}
	//cmd执行
	ws.Cmd = app.NewCmd(ws.run)
	return ws
}

//WithComponents 自定义组件
func (gs *WebServer) WithComponents(components ...global.Init) *WebServer {
	gs.components = append(gs.components, components...)

	return gs
}

func (gs *WebServer) WithOption(options ...gateway.GatewayOption) *WebServer {
	gs.options = append(gs.options, options...)
	return gs
}

//WithRegister 注册grpc反向代理
func (gs *WebServer) WithRegister(endpoints map[string][]grpc_proxy.EndpointFun) *WebServer {
	for k, v := range endpoints {
		gs.WithOption(gateway.WithEndpoint(k, v...))
	}

	return gs
}

//WithRouter 路由
func (gs *WebServer) WithRouter(f func(r *gin.Engine)) *WebServer {
	gs.router = f
	return gs
}

//AfterConfigInit 挂载配置后需要执行的函数
func (gs *WebServer) AfterConfig(f func() error) *WebServer {
	gs.afterConfigFun = f
	return gs
}

//Run 通过cmd启动
func (gs *WebServer) Run() error {
	return gs.Cmd.Run()
}

//run 启动web server
func (gs *WebServer) run(cmd *app.Cmd) (err error) {
	gs.configPath = cmd.ConfigPath(gs.configPath)

	//global配置选项
	options := []global.Option{
		global.WithConfigPath(gs.configPath),
		global.WithAppConfig(&app.Config{Service: gs.name, Port: gs.port}),
		global.WithConfig(gateway.ConfigInstance()),
	}

	//初始化global
	if err = global.Configure(options...).Use(gs.components...); err != nil {
		return
	}

	//执行回调函数
	if gs.afterConfigFun != nil {
		if err = gs.afterConfigFun(); err != nil {
			return
		}
	}

	gs.options = append(gs.options, gateway.WithAppConfig(global.AppConfig()))

	return gateway.NewGateway(gs.options...).Router(gs.router).Run()
}
