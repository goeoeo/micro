package grpcx

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/keepalive"

	"git.internal.yunify.com/benchmark/benchmark/pkg/app"
	"git.internal.yunify.com/benchmark/benchmark/pkg/infra/registry"
)

var (
	BackOffConfig = backoff.Config{
		BaseDelay:  1 * time.Second,
		Multiplier: 1,
		Jitter:     0.2,
		MaxDelay:   60 * time.Second,
	}

	KeepaliveClientConfig = keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             time.Second,
		PermitWithoutStream: true,
	}

	KeepaliveServerConfig = keepalive.ServerParameters{
		MaxConnectionIdle: 10 * time.Minute, // If a client is idle for 10 seconds, send a GOAWAY
		Time:              5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout:           time.Second,      // Wait 1 second for the ping ack before assuming the connection is dead
	}

	KeepaliveEnforcementPolicy = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second,
		PermitWithoutStream: true,
	}
)

type Config struct {
	ErrorDetail bool `default:"true"` //展示详细错误信息
}

type Option func(*GrpcServer)

//WithAppConfig 服务配置
func WithAppConfig(app *app.Config) Option {
	return func(g *GrpcServer) {
		g.appConfig = app
	}
}

//WithConfig 服务配置
func WithConfig(c *Config) Option {
	return func(g *GrpcServer) {
		g.conf = c
	}
}

//WithInterceptors 拦截器
func WithInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option {
	return func(g *GrpcServer) {
		g.interceptors = interceptors
	}
}

//WithRegister 注册中心
func WithRegister(r registry.Registry) Option {
	return func(g *GrpcServer) {
		g.register = r
	}
}
