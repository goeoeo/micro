package server_interceptor

import (
	"runtime/debug"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"git.internal.yunify.com/benchmark/benchmark/pkg/util/gerr"
)

//recovery的拦截器
func RecoveryInterceptor() grpc.UnaryServerInterceptor {
	return grpc_recovery.UnaryServerInterceptor(
		grpc_recovery.WithRecoveryHandler(func(p interface{}) error {
			logrus.Errorf("GRPC server recovery with error: %+v", p)
			logrus.Error(string(debug.Stack()))

			if e, ok := p.(error); ok {
				return gerr.NewErr(nil).WithCode(codes.Internal).WithCause(e)
			}
			return gerr.NewErr(nil).WithCode(codes.Internal)
		}),
	)
}
