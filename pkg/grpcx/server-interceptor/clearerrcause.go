package server_interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"git.internal.yunify.com/benchmark/benchmark/pkg/util/gerr"
)

//ClearErrorCause 关闭详细错误原因展示
func ClearErrorCause() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			if e, ok := status.FromError(err); ok {
				if e.Code() != codes.OK {
					//清理cause
					err = gerr.ClearErrorCause(err)
				}
			}
		}

		return
	}
}
