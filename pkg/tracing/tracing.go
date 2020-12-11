package tracing

import (
	"context"
	"net/http"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
)

// ParentSpan is a middleware that creates an OpenTracing parentSpan.
func ParentSpan(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tracer := opentracing.GlobalTracer()

		begin := time.Now()
		var span opentracing.Span

		carrier := opentracing.HTTPHeadersCarrier(r.Header)
		clientContext, err := tracer.Extract(
			opentracing.HTTPHeaders,
			carrier)

		if err == nil {
			span = tracer.StartSpan(
				"get data",
				ext.RPCServerOption(clientContext))
		} else {
			span = tracer.StartSpan("get data")
		}

		ext.Component.Set(span, r.URL.Scheme)
		ext.HTTPMethod.Set(span, r.Method)
		ext.HTTPUrl.Set(span, r.URL.String())

		defer func(begin time.Time) {
			duration := time.Since(begin)
			span.LogFields(
				log.Int64("duration.millis", duration.Milliseconds()),
			)
			span.Finish()
		}(begin)

		tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		r = r.WithContext(opentracing.ContextWithSpan(r.Context(), span))
		next.ServeHTTP(w, r)
	})
}

// StartChildSpan creates an OpenTracing childSpan.
func StartChildSpan(ctx context.Context, name string) (opentracing.Span, time.Time) {
	begin := time.Now()
	childSpan, ctx := opentracing.StartSpanFromContext(ctx, name)
	return childSpan, begin
}

// FinishChildSpan ends an OpenTracing childSpan and logs its duration.
func FinishChildSpan(childSpan opentracing.Span, begin time.Time) {
	duration := time.Since(begin)
	childSpan.LogFields(
		log.Int64("duration.millis", duration.Milliseconds()),
	)
	childSpan.Finish()
}
