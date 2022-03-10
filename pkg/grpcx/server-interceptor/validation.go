package server_interceptor

import (
	"context"

	"google.golang.org/grpc"

	valid "git.internal.yunify.com/benchmark/benchmark/pkg/util/validation"
)

//格式校验拦截器
func Validation() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if err = valid.DefaultValidation.Valid(req); err != nil {
			return
		}

		return handler(ctx, req)
	}
}
