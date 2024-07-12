module github.com/starton-io/tyrscale/plugins/middleware/exemple

go 1.21.6

replace github.com/starton-io/tyrscale/gateway => ../../../gateway

require (
	github.com/starton-io/tyrscale/gateway v0.0.0-00010101000000-000000000000
	github.com/valyala/fasthttp v1.55.0
)

require (
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
)
