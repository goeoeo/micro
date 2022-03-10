package client_interceptor

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	gcodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	ztrace "git.internal.yunify.com/benchmark/benchmark/pkg/util/trace"
)

// UnaryTracingInterceptor returns a grpc.UnaryClientInterceptor for opentelemetry.
func UnaryTracingInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx, span := startSpan(ctx, method, cc.Target())
	defer span.End()

	ztrace.MessageSent.Event(ctx, 1, req)
	ztrace.MessageReceived.Event(ctx, 1, reply)

	if err := invoker(ctx, method, req, reply, cc, opts...); err != nil {
		s, ok := status.FromError(err)
		if ok {
			span.SetStatus(codes.Error, s.Message())
			span.SetAttributes(ztrace.StatusCodeAttr(s.Code()))
		} else {
			span.SetStatus(codes.Error, err.Error())
		}
		return err
	}

	span.SetAttributes(ztrace.StatusCodeAttr(gcodes.OK))
	return nil
}

func startSpan(ctx context.Context, method, target string) (context.Context, trace.Span) {
	var md metadata.MD
	requestMetadata, ok := metadata.FromOutgoingContext(ctx)
	if ok {
		md = requestMetadata.Copy()
	} else {
		md = metadata.MD{}
	}
	tr := otel.Tracer(ztrace.TraceName)
	name, attr := ztrace.SpanInfo(method, target)
	ctx, span := tr.Start(ctx, name, trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(attr...))
	ztrace.Inject(ctx, otel.GetTextMapPropagator(), &md)
	ctx = metadata.NewOutgoingContext(ctx, md)

	return ctx, span
}
