package tracing

import (
	"fmt"
	"io"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

var (
	// Tracer - tracer instance
	Tracer opentracing.Tracer
	// Closer - Closer instance
	Closer io.Closer
)

// Init returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func Init(service string) (opentracing.Tracer, io.Closer) {
	cfg, err := config.FromEnv()
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger Conf from ENV: %v\n", err))
	}
	cfg.Sampler = &config.SamplerConfig{
		Type:  "const",
		Param: 1,
	}
	cfg.Reporter.LogSpans = true
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	Tracer, Closer = tracer, closer
	// in order to use StartSpanFromContext, we need SetGlobalTracer
	opentracing.SetGlobalTracer(Tracer)
	return tracer, closer
}
