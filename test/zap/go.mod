module zap

go 1.22

replace github.com/alibaba/opentelemetry-go-auto-instrumentation => ../../../opentelemetry-go-auto-instrumentation

require (
	github.com/alibaba/opentelemetry-go-auto-instrumentation v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/otel/sdk v1.30.0
	go.uber.org/zap v1.20.0
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	go.opentelemetry.io/otel v1.30.0 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	go.opentelemetry.io/otel/trace v1.30.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
)
