// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

var (
	Type = component.MustNewType("sample")
)

const (
	LogsStability    = component.StabilityLevelDevelopment
	TracesStability  = component.StabilityLevelBeta
	MetricsStability = component.StabilityLevelStable
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("github.com/open-telemetry/opentelemetry-collector-contrib/cmd/mdatagen/internal/samplereceiver")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/open-telemetry/opentelemetry-collector-contrib/cmd/mdatagen/internal/samplereceiver")
}