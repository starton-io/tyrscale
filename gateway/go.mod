module github.com/starton-io/tyrscale/gateway

go 1.21.6

require (
	github.com/ThreeDotsLabs/watermill v1.3.5
	github.com/ThreeDotsLabs/watermill-redisstream v1.3.0
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/carlmjohnson/flowmatic v0.23.4
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/ethereum/go-ethereum v1.14.5
	github.com/golang/protobuf v1.5.4
	github.com/joho/godotenv v1.5.1
	github.com/madflojo/tasks v1.2.0
	github.com/prometheus/client_golang v1.19.1
	github.com/redis/go-redis/v9 v9.5.1
	github.com/sony/gobreaker v0.5.0
	github.com/starton-io/tyrscale/go-kit v0.0.0-00010101000000-000000000000
	github.com/starton-io/tyrscale/manager v0.0.0-00010101000000-000000000000
	github.com/starton-io/tyrscale/sdk/tyrscale-sdk-go v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.9.0
	github.com/valyala/fasthttp v1.55.0
	go.uber.org/zap v1.27.0
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/Rican7/retry v0.3.1 // indirect
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/carlmjohnson/deque v0.23.1 // indirect
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/holiman/uint256 v1.2.4 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/lithammer/shortuuid/v3 v3.0.7 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.53.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rs/xid v1.5.0 // indirect
	github.com/samber/lo v1.39.0 // indirect
	github.com/samber/slog-common v0.16.0 // indirect
	github.com/samber/slog-zap/v2 v2.4.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	go.opentelemetry.io/otel v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.27.0 // indirect
	go.opentelemetry.io/otel/metric v1.27.0 // indirect
	go.opentelemetry.io/otel/sdk v1.27.0 // indirect
	go.opentelemetry.io/otel/trace v1.27.0 // indirect
	go.opentelemetry.io/proto/otlp v1.2.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20240525044651-4c93da0ed11d // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240521202816-d264139d666e // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240521202816-d264139d666e // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// TODO: remove this replace after we publish go-kit in github
replace (
	github.com/starton-io/tyrscale/go-kit => ../go-kit
	github.com/starton-io/tyrscale/manager => ../manager
	github.com/starton-io/tyrscale/sdk/tyrscale-sdk-go => ../sdk/tyrscale-sdk-go
)
