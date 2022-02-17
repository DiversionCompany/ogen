module github.com/ogen-go/ogen/examples

go 1.17

require (
	github.com/go-faster/errors v0.5.0
	github.com/go-faster/jx v0.32.1
	github.com/google/uuid v1.3.0
	github.com/ogen-go/ogen v0.0.0
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/otel v1.4.1
	go.opentelemetry.io/otel/metric v0.27.0
	go.opentelemetry.io/otel/trace v1.4.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-logr/logr v1.2.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/segmentio/asm v1.1.3 // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/ogen-go/ogen v0.0.0 => ./..
