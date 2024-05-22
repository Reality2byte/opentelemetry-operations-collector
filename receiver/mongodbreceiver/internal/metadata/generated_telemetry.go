// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter("github.com/GoogleCloudPlatform/opentelemetry-operations-collector/receiver/mongodbreceiver")
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer("github.com/GoogleCloudPlatform/opentelemetry-operations-collector/receiver/mongodbreceiver")
}
