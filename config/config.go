package config

import "restapi-with-opentelemetry/pkg/helper"

// first load from os env
var (
	ServiceName           = helper.GetEnv("SERVICE_NAME", "restapi-with-opentelemetry")
	ServiceAddress        = helper.GetEnv("SERVICE_ADDRESS", "localhost:8080")
	OtelCollectorURL      = helper.GetEnv("OTEL_EXPORTER_OTLP_ENDPOINT", "localhost:4317")
	OtelCollectorInsecure = helper.GetEnv("OTEL_EXPORTER_OTLP_INSECURE_MODE", "true")
	SecretToken           = helper.GetEnv("SECRET_TOKEN", "SDas12cAS21312312")
)

//Load maybe be used if want to overwrite env from .env file
func Load() {
	ServiceName = helper.GetEnv("SERVICE_NAME", "restapi-with-opentelemetry")
	ServiceAddress = helper.GetEnv("SERVICE_ADDRESS", "localhost:8080")
	OtelCollectorURL = helper.GetEnv("OTEL_EXPORTER_OTLP_ENDPOINT", "localhost:4317")
	OtelCollectorInsecure = helper.GetEnv("OTEL_EXPORTER_OTLP_INSECURE_MODE", "true")
	SecretToken = helper.GetEnv("SECRET_TOKEN", "SDas12cAS21312312")
}
