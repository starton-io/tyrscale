module github.com/starton-io/tyrscale/plugins/interceptor/response

go 1.21.6

require github.com/starton-io/tyrscale/gateway v0.0.0

require github.com/valyala/fasthttp v1.55.0

require (
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/klauspost/compress v1.17.8 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
)

replace github.com/starton-io/tyrscale/gateway => ../../../gateway
