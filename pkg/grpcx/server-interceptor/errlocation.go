package server_interceptor

import (
	"context"

	"google.golang.org/grpc"

	"git.internal.yunify.com/benchmark/benchmark/pkg/util/gerr"
)

//填充错误发生的位置
func ErrLocation() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if resp, err = handler(ctx, req); err != nil {
			err = gerr.NewErr(ctx).WithCause(err, info.FullMethod)
		}

		return
	}
}
