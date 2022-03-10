package grpc_meadata

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/grpc/metadata"
)

//将kv存入http hearder 中
func SetMetaWithHttp(r *http.Request, key MetaKey, val string) {
	r.Header.Set(fmt.Sprintf("grpc-metadata-%s", key), val)
}

func GetMetaWithHttp(r *http.Request, key MetaKey) string {
	return r.Header.Get(fmt.Sprintf("grpc-metadata-%s", key))
}

func SetMetaContext(ctx context.Context, key MetaKey, val string) context.Context {
	ctx = context.WithValue(ctx, key, val)
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.MD{}
	}
	md[string(key)] = []string{val}
	return metadata.NewOutgoingContext(ctx, md)
}

func GetMetaContextValue(ctx context.Context, key MetaKey) string {
	key1 := strings.Replace(string(key), "grpc-metadata-", "", 1)
	rid := getValueFromContext(ctx, key1)
	if len(rid) == 0 {
		return ""
	}
	return rid[0]
}

type getMetadataFromContext func(ctx context.Context) (md metadata.MD, ok bool)

var getMetadataFromContextFunc = []getMetadataFromContext{
	metadata.FromOutgoingContext,
	metadata.FromIncomingContext,
}

func getValueFromContext(ctx context.Context, key string) []string {
	if ctx == nil {
		return []string{}
	}
	for _, f := range getMetadataFromContextFunc {
		md, ok := f(ctx)
		if !ok {
			continue
		}
		m, ok := md[key]
		if ok && len(m) > 0 {
			return m
		}
	}
	m, ok := ctx.Value(key).([]string)
	if ok && len(m) > 0 {
		return m
	}
	s, ok := ctx.Value(key).(string)
	if ok && len(s) > 0 {
		return []string{s}
	}
	return []string{}
}
