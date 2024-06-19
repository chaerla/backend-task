package tracing

import (
	"backend-task/bootstrap/config"
	"backend-task/pkg/logger"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"log"
)

func InitTracing(cfg *config.Config) {
	logger.Log.Info("Initializing jaeger")
	ep := fmt.Sprintf("%s:%s/api/traces", cfg.JaegerHost, cfg.JaegerPort)
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(ep)))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(cfg.ServiceName),
				attribute.String("environment", "dev"),
				attribute.Int64("ID", 1),
			),
		),
	)

	otel.SetTracerProvider(tp)
}
