package telemetry

import (
	"context"

	"github.com/getsentry/sentry-go"
	sentryotel "github.com/getsentry/sentry-go/otel"
	"github.com/mujhtech/b0/config"
	"go.opentelemetry.io/otel"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type Sentry struct {
	cfg        config.Telemetry
	ShutdownFn func(ctx context.Context) error
}

func NewSentry(cfg config.Telemetry) *Sentry {
	return &Sentry{
		cfg: cfg,
		ShutdownFn: func(ctx context.Context) error {
			return nil
		},
	}
}

func (st *Sentry) Init(componentName string) error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              st.cfg.Sentry.DSN,
		ServerName:       componentName,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
		Debug:            true,
	})
	if err != nil {
		return err
	}

	// Configure Tracer Provider.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(sentryotel.NewSentrySpanProcessor()),
	)

	// Configure OTel SDK.
	otel.SetTracerProvider(tp)

	// Configure Propagator.
	otel.SetTextMapPropagator(sentryotel.NewSentryPropagator())

	st.ShutdownFn = tp.Shutdown

	return nil
}

func (st *Sentry) Type() config.TelemetryProvider {
	return st.cfg.Provider
}
func (st *Sentry) Capture() {

}
func (st *Sentry) Shutdown(ctx context.Context) error {
	return st.ShutdownFn(ctx)
}
