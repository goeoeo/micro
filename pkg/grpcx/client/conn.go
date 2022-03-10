package grpc_client

import (
	"google.golang.org/grpc"

	"git.internal.yunify.com/benchmark/benchmark/pkg/grpcx"
	"git.internal.yunify.com/benchmark/benchmark/pkg/grpcx/grpclb"
	"git.internal.yunify.com/benchmark/benchmark/pkg/infra/global"
)

func NewClientConn(serviceName string) (*grpc.ClientConn, error) {
	rl := grpclb.NewResolver(global.Registry().Registry, serviceName)
	dailOptions := append(grpcx.ClientOptions, rl.DailOptions()...)

	return grpc.Dial(rl.Endpoint(), dailOptions...)
}
