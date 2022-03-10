package grpcx

import (
	"google.golang.org/grpc"

	client_interceptor "git.internal.yunify.com/benchmark/benchmark/pkg/grpcx/client-interceptor"
)

var ClientOptions = []grpc.DialOption{
	grpc.WithInsecure(),
	grpc.WithConnectParams(grpc.ConnectParams{Backoff: BackOffConfig}),
	grpc.WithKeepaliveParams(KeepaliveClientConfig),
	grpc.WithChainUnaryInterceptor(client_interceptor.UnaryTracingInterceptor), //客户端拦截器
}

//type adminCredentials struct {
//	useTls bool
//}
//
//func (a adminCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
//	res := map[string]string{
//		"nb": "1",
//	}
//	md, ok := metadata.FromIncomingContext(ctx)
//	if !ok {
//		return res, nil
//	}
//	auth := md[iamcommon.AUTHORIZATION]
//	if len(auth) == 0 || len(auth[0]) <= len(iamcommon.AUTHORIZATIONPREX) {
//		return res, nil
//	}
//	token := auth[0][len(iamcommon.AUTHORIZATIONPREX):]
//	res[iamcommon.AUTHORIZATION] = iamcommon.AUTHORIZATIONPREX + token
//	return res, nil
//}
//
//func (a adminCredentials) RequireTransportSecurity() bool {
//	return a.useTls
//}
