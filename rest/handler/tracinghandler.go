package handler

import (
	"net/http"
	"sync"

	"github.com/sanhuanshisanshao/go-zero/core/lang"
	"github.com/sanhuanshisanshao/go-zero/core/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var notTracingSpans sync.Map

// DontTraceSpan disable tracing for the specified span name.
func DontTraceSpan(spanName string) {
	notTracingSpans.Store(spanName, lang.Placeholder)
}

// TracingHandler return a middleware that process the opentelemetry.
func TracingHandler(serviceName, path string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		propagator := otel.GetTextMapPropagator()
		tracer := otel.GetTracerProvider().Tracer(trace.TraceName)

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				next.ServeHTTP(w, r)
			}()

			ctx := propagator.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
			spanName := path
			if len(spanName) == 0 {
				spanName = r.URL.Path
			}

			if _, ok := notTracingSpans.Load(spanName); ok {
				return
			}

			spanCtx, span := tracer.Start(
				ctx,
				spanName,
				oteltrace.WithSpanKind(oteltrace.SpanKindServer),
				oteltrace.WithAttributes(semconv.HTTPServerAttributesFromHTTPRequest(
					serviceName, spanName, r)...),
			)
			defer span.End()

			// convenient for tracking error messages
			propagator.Inject(spanCtx, propagation.HeaderCarrier(w.Header()))
			r = r.WithContext(spanCtx)
		})
	}
}
