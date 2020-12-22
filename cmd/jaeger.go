package cmd

import (
	"io"
	"os"
	"strconv"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

// InitJaeger configures a Jaeger tracer and sets it as OpenTracing's
// global tracer.
func InitJaeger() (io.Closer, error) {
	name := os.Getenv("SERVICE_NAME")

	var cfg *jaegercfg.Configuration
	if disabled, _ := strconv.ParseBool(os.Getenv("JAEGER_DISABLED")); disabled {
		cfg, _ = jaegercfg.FromEnv()
	} else if dev, _ := strconv.ParseBool(os.Getenv("DEV_ENV")); dev {
		cfg = &jaegercfg.Configuration{
			ServiceName: name,
			Sampler: &jaegercfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &jaegercfg.ReporterConfig{
				LogSpans: true,
			},
		}
	} else {
		cfg = &jaegercfg.Configuration{}
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	closer, err := cfg.InitGlobalTracer(
		name,
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		return nil, err
	}

	return closer, nil
}
