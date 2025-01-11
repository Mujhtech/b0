package telemetry

import (
	"context"
	"fmt"

	"github.com/mujhtech/b0/cmd/version"
	"github.com/mujhtech/b0/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc/credentials"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type OTel struct {
	cfg        config.OTEL
	ShutdownFn func(ctx context.Context) error
}

func NewOTel(cfg config.OTEL) *OTel {
	return &OTel{
		cfg: cfg,
		ShutdownFn: func(ctx context.Context) error {
			return nil
		},
	}
}

func (ot *OTel) Init(componentName string) error {
	var opts []otlptracegrpc.Option

	if ot.cfg.CollectorURL == "" {
		return fmt.Errorf("OTEL collector URL is required")
	}

	opts = append(opts, otlptracegrpc.WithEndpoint(ot.cfg.CollectorURL))

	if ot.cfg.OTelAuth != (config.OTelAuthConfiguration{}) {
		opts = append(opts, otlptracegrpc.WithHeaders(
			map[string]string{
				ot.cfg.OTelAuth.HeaderName: ot.cfg.OTelAuth.HeaderValue}))
	}

	if ot.cfg.InsecureSkipVerify {
		secureOption := otlptracegrpc.WithInsecure()
		opts = append(opts, secureOption)
	} else {
		secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
		opts = append(opts, secureOption)
	}

	exporter, err := otlptrace.New(context.Background(), otlptracegrpc.NewClient(opts...))
	if err != nil {
		return err
	}

	// Configure Resources.
	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.KeyValue{
				Key:   semconv.ServiceNameKey,
				Value: attribute.StringValue(componentName),
			},
			attribute.KeyValue{
				Key:   semconv.ServiceVersionKey,
				Value: attribute.StringValue(version.Version),
			},
		),
	)
	if err != nil {
		return err
	}

	// Configure Tracer Provider.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resources),
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(ot.cfg.SampleRate)),
	)

	// Configure OTel SDK
	otel.SetTracerProvider(tp)

	// Configure Propagator
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	ot.ShutdownFn = tp.Shutdown

	return nil
}

func (ot *OTel) Type() config.TelemetryProvider {
	return config.TelemetryProviderOtel
}
func (ot *OTel) Capture() {

}
func (ot *OTel) Shutdown(ctx context.Context) error {
	return ot.ShutdownFn(ctx)
}
