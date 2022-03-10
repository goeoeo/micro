package server_interceptor

import (
	"context"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	grpc_meadata "git.internal.yunify.com/benchmark/benchmark/pkg/grpcx/grpc-meadata"
	"git.internal.yunify.com/benchmark/benchmark/pkg/util/pbutil"
)

//日志拦截器
func AccessLogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var err error
		requestId := grpc_meadata.RequestIdKey.Get(ctx)

		method := strings.Split(info.FullMethod, "/")
		action := method[len(method)-1]
		entry := logrus.WithFields(logrus.Fields{
			string(grpc_meadata.RequestIdKey): requestId,
		})
		if p, ok := req.(proto.Message); ok {
			if content, err := pbutil.Marshal(p); err != nil {
				entry.Errorf("Failed to marshal proto message to string [%s] [%+v]", action, err)
			} else {
				entry.Infof("Received request [%s] with params [%s]", action, content)
			}
		}
		start := time.Now()

		resp, err := handler(ctx, req)

		elapsed := time.Since(start)
		entry.Infof("(exec_time: %s) Handled request [%s]", elapsed, action)
		return resp, err
	}
}
