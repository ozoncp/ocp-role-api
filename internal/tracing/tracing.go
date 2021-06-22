package tracing

import (
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

func InitTracing(serviceName string) {
	cfg := jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	logger := jaegerlog.StdLogger
	metricsFactory := metrics.NullFactory

	tracer, _, err := cfg.NewTracer(
		jaegercfg.Logger(logger),
		jaegercfg.Metrics(metricsFactory),
	)

	if err != nil {
		log.Fatal().Msgf("tracer creation error: %v", err)
	}

	opentracing.SetGlobalTracer(tracer)
}
