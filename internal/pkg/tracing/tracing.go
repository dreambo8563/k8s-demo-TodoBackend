package tracing

import (
	"context"
	"fmt"
	"io"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

//Client - traceing client
var Client *TraceClient

//TraceClient -
type TraceClient struct {
	Tracer opentracing.Tracer
	Closer io.Closer
}

// NewTraceClient returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func NewTraceClient() *TraceClient {
	if Client != nil {
		return Client
	}
	cfg, err := config.FromEnv()
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger Conf from ENV: %v\n", err))
	}
	cfg.Sampler = &config.SamplerConfig{
		Type:  "const",
		Param: 1,
	}
	cfg.Reporter.LogSpans = true
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))

	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	Client = &TraceClient{
		Tracer: tracer,
		Closer: closer,
	}
	// in order to use StartSpanFromContext, we need SetGlobalTracer
	opentracing.SetGlobalTracer(tracer)
	return Client
}

//StartSpan -
func (t *TraceClient) StartSpan(spanName string) opentracing.Span {
	return t.Tracer.StartSpan(spanName)
}

//GetTracer - get Tracer
func (t *TraceClient) GetTracer() opentracing.Tracer {
	return t.Tracer
}

//ContextWithSpan -
func ContextWithSpan(ctx context.Context, span opentracing.Span) context.Context {
	return opentracing.ContextWithSpan(ctx, span)
}

//StartSpanFromContext -
func StartSpanFromContext(ctx context.Context, span string) (opentracing.Span, context.Context) {
	return opentracing.StartSpanFromContext(ctx, span)
}
