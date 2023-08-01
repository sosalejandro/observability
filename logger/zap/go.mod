module github.com/sosalejandro/observability/logger/zap

go 1.20

replace github.com/sosalejandro/observability => ../../

require (
	github.com/sosalejandro/observability v0.0.0-20230731162132-8f574250c779
	github.com/stretchr/testify v1.8.4
	go.uber.org/zap v1.24.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/otel v1.16.0 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
