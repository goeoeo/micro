package boot

import (
	"google.golang.org/grpc"

	server_interceptor "git.internal.yunify.com/benchmark/benchmark/pkg/grpcx/server-interceptor"

	"git.internal.yunify.com/benchmark/benchmark/pkg/app"
	"git.internal.yunify.com/benchmark/benchmark/pkg/global"
	"git.internal.yunify.com/benchmark/benchmark/pkg/grpcx"
)

type GrpcServer struct {
	configPath string
	name       string
	port       string

	components  []global.Init
	interceptor []grpc.UnaryServerInterceptor

	//配置挂载后需要执行的函数
	afterConfigFun func() error

	//注册rpc服务
	register func(gs *grpc.Server)

	//命令行解析工具
	Cmd *app.Cmd
}

//NewGrpcServer grpcserver启动实例
func NewGrpcServer(name, port, configPath string, register func(gs *grpc.Server)) *GrpcServer {
	gs := &GrpcServer{
		configPath: configPath,
		name:       name,
		port:       port,
		//内置组件
		components: []global.Init{
			global.Log(),
			global.Registry(),
		},
		interceptor: []grpc.UnaryServerInterceptor{
			server_interceptor.RecoveryInterceptor(),
			server_interceptor.Validation(),
			server_interceptor.AccessLogInterceptor(),
			server_interceptor.ErrLocation(),
			server_interceptor.UnaryTracingInterceptor,
		},
		register: register,
	}

	gs.Cmd = app.NewCmd(gs.run)

	return gs
}

//WithComponents 自定义组件
func (gs *GrpcServer) WithComponents(components ...global.Init) *GrpcServer {
	gs.components = append(gs.components, components...)

	return gs
}

//WithInterceptor 自定义拦截器
func (gs *GrpcServer) WithInterceptor(interceptor ...grpc.UnaryServerInterceptor) *GrpcServer {
	gs.interceptor = append(gs.interceptor, interceptor...)

	return gs
}

//AfterConfigInit 挂载配置后需要执行的函数
func (gs *GrpcServer) AfterConfig(f func() error) *GrpcServer {
	gs.afterConfigFun = f
	return gs
}

//Run 通过cmd启动
func (gs *GrpcServer) Run() error {
	return gs.Cmd.Run()
}

//run 启动grpc server
func (gs *GrpcServer) run(cmd *app.Cmd) (err error) {
	gs.configPath = cmd.ConfigPath(gs.configPath)

	//global配置选项
	options := []global.Option{
		global.WithConfigPath(gs.configPath),
		global.WithAppConfig(&app.Config{Service: gs.name, Port: gs.port}),
		global.WithConfig(grpcx.ConfigInstance()),
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

	//grpc 配置
	grpcOptions := []grpcx.Option{
		grpcx.WithAppConfig(global.AppConfig()),
		grpcx.WithRegister(global.Registry().Registry),
		grpcx.WithInterceptors(gs.interceptor...),
	}

	//启动
	return grpcx.NewGrpcServer(grpcOptions...).Start(gs.register)
}
