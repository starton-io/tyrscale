package interceptor

import (
	"errors"

	"github.com/valyala/fasthttp"
)

// var (
//
//	requestCount = prometheus.NewCounterVec(
//		prometheus.CounterOpts{
//			Name: "http_requests_total",
//			Help: "Number of HTTP requests",
//		},
//		[]string{"route_uuid", "method", "status", "host", "scheme"},
//	)
//	requestDuration = prometheus.NewHistogramVec(
//		prometheus.HistogramOpts{
//			Name:    "http_request_duration_seconds",
//			Help:    "Duration of HTTP requests",
//			Buckets: prometheus.DefBuckets,
//		},
//		[]string{"route_uuid", "method", "status", "host", "scheme"},
//	)
//
// )
//
// //func init() {
// //	prometheus.MustRegister(requestCount)
// //	prometheus.MustRegister(requestDuration)
// //}

type DefaultResponseInterceptor struct {
}

func (r *DefaultResponseInterceptor) Intercept(res *fasthttp.Response) error {

	//body, err := res.BodyGunzip()
	//if err != nil {
	//	return err
	//}
	//res.SetBody(body)

	if res.StatusCode() == fasthttp.StatusTooManyRequests {
		return errors.New("too many requests")
	}
	if res.StatusCode() >= fasthttp.StatusInternalServerError {
		return errors.New("invalid response")
	}

	return nil
}

// TODO: implement default last response interceptor GZIP
//type DefaultLastResponseInterceptor struct {
//}
//
//func gZipData(data []byte) (compressedData []byte, err error) {
//	var b bytes.Buffer
//	gz := gzip.NewWriter(&b)
//	_, err = gz.Write(data)
//	if err != nil {
//		return
//	}
//	if err = gz.Flush(); err != nil {
//		return
//	}
//	if err = gz.Close(); err != nil {
//		return
//	}
//	compressedData = b.Bytes()
//	return
//}
//
//func (r *DefaultLastResponseInterceptor) Intercept(res *fasthttp.Response) error {
//	//gzip the body
//	body, err := gZipData(res.Body())
//	if err != nil {
//		return err
//	}
//	res.SetBody(body)
//	return nil
//}
