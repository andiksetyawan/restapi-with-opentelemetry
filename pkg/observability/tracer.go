package observability

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/credentials"
	"restapi-with-opentelemetry/config"
)

type ShutDownFunc func(ctx context.Context) error

func InitTracerProvider() ShutDownFunc {
	//TODO add tls
	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if len(config.OtelCollectorInsecure) > 0 {
		secureOption = otlptracegrpc.WithInsecure()
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(config.OtelCollectorURL),
		),
	)

	if err != nil {
		log.Fatal().Err(err).Msg("failed to initiate otel exporter")
	}

	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", config.ServiceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Error().Err(err).Msgf("could not set resources: ", err)
	}

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()), //TODO in production: set parent and 30% (.3)
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(resources),
		),
	)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return func(ctx context.Context) error {
		cxt, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := exporter.Shutdown(cxt); err != nil {
			return err
		}
		log.Debug().Msg("tracer exporter has been shutdown")
		return nil
	}
}
