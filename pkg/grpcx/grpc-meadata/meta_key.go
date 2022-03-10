package grpc_meadata

import (
	"context"
	"net/http"
)

const (
	RequestIdKey MetaKey = "request_id"
	OplogKey     MetaKey = "oplog"
	MessageIdKey MetaKey = "message_id"
)

type MetaKey string

func (s MetaKey) SetHttp(r *http.Request, val string) {
	SetMetaWithHttp(r, s, val)
}

func (s MetaKey) GetHttp(r *http.Request) string {
	return GetMetaWithHttp(r, s)
}

func (s MetaKey) Get(ctx context.Context) string {
	return GetMetaContextValue(ctx, s)
}

func (s MetaKey) Set(ctx context.Context, val string) context.Context {
	return SetMetaContext(ctx, s, val)
}
