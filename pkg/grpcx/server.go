package grpcx

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"git.internal.yunify.com/benchmark/benchmark/pkg/app"
	server_interceptor "git.internal.yunify.com/benchmark/benchmark/pkg/grpcx/server-interceptor"
	"git.internal.yunify.com/benchmark/benchmark/pkg/infra/registry"
	"git.internal.yunify.com/benchmark/benchmark/pkg/util/idutil"
	"git.internal.yunify.com/benchmark/benchmark/pkg/util/proc"
)

type GrpcServer struct {
	appConfig *app.Config
	conf      *Config

	interceptors []grpc.UnaryServerInterceptor //拦截器
	register     registry.Registry             //注册中心

	serviceInfo *registry.Service
}

var grpcConfig = new(Config)

func NewGrpcServer(opts ...Option) *GrpcServer {
	gs := &GrpcServer{
		appConfig: new(app.Config),
		conf:      grpcConfig,
	}

	for _, o := range opts {
		o(gs)
	}

	return gs.initServiceInfo()
}

//ConfigInstance grpc 配置单例
func ConfigInstance() (string, interface{}) {
	return "grpc", grpcConfig
}

func (gs *GrpcServer) initServiceInfo() *GrpcServer {
	gs.serviceInfo = &registry.Service{
		Name: gs.appConfig.Service,
		Nodes: []*registry.Node{
			{
				Id:      idutil.GetUuid("service_"),
				Address: gs.appConfig.Endpoint(),
			},
		},
	}
	return gs
}

func (gs *GrpcServer) Start(registerService func(gs *grpc.Server)) (err error) {
	var (
		lis          net.Listener
		interceptors []grpc.UnaryServerInterceptor
	)
	if lis, err = net.Listen("tcp", fmt.Sprintf(":%s", gs.appConfig.Port)); err != nil {
		return err
	}

	if !gs.conf.ErrorDetail {
		interceptors = append(interceptors, server_interceptor.ClearErrorCause())
	}

	interceptors = append(interceptors, gs.interceptors...)
	builtinOptions := []grpc.ServerOption{
		grpc.KeepaliveEnforcementPolicy(KeepaliveEnforcementPolicy),
		grpc.KeepaliveParams(KeepaliveServerConfig),
		grpc.ChainUnaryInterceptor(interceptors...),
	}

	grpcServer := grpc.NewServer(builtinOptions...)
	reflection.Register(grpcServer)
	//注册服务
	registerService(grpcServer)

	//服务注册
	if err = gs.registry(); err != nil {
		return
	}

	//平滑退出
	waitForCalled := proc.AddShutdownListener(func() {
		//grpc server平滑退出
		grpcServer.GracefulStop()
		logrus.Info("grpc server closed success")
	})
	defer waitForCalled()

	logrus.Infof("grpc server start success[%s:%s]", gs.appConfig.Service, gs.appConfig.Port)
	if err = grpcServer.Serve(lis); err != nil {
		return
	}

	return
}

func (gs *GrpcServer) registry() (err error) {
	if gs.register == nil {
		return
	}

	if err = gs.register.Register(gs.serviceInfo); err != nil {
		return
	}

	logrus.Infof("register success  [%s,%s]", gs.serviceInfo.Name, gs.serviceInfo.Nodes[0].Address)

	//退出时，注销注册中心
	proc.AddWrapUpListener(func() {
		if err := gs.register.Deregister(gs.serviceInfo); err != nil {
			logrus.Infof("service deregister fail err:%v", err)
		} else {
			logrus.Infof("service deregister success [%s,%s]", gs.serviceInfo.Name, gs.serviceInfo.Nodes[0].Address)
		}
	})

	return
}
