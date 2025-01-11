package telemetry

import (
	"context"

	"github.com/mujhtech/b0/config"
)

type Telemetry interface {
	Init(component string) error
	Capture()
	Type() config.TelemetryProvider
	Shutdown(ctx context.Context) error
}

type telemetry struct{}

func (t *telemetry) Init(component string) error {
	return nil
}

func (t *telemetry) Capture() {}

func (t *telemetry) Type() config.TelemetryProvider {
	return config.TelemetryProviderNone
}

func (t *telemetry) Shutdown(ctx context.Context) error {
	return nil
}

func New(cfg config.Telemetry, component string) (Telemetry, error) {

	switch cfg.Provider {

	case config.TelemetryProviderSentry:
		sentry := NewSentry(cfg)

		err := sentry.Init(component)

		if err != nil {
			return nil, err
		}

		return sentry, nil

	case config.TelemetryProviderOtel:
		otel := NewOTel(cfg.OTEL)

		err := otel.Init(component)

		if err != nil {
			return nil, err
		}

		return otel, nil
	default:
		return &telemetry{}, nil
	}
}
