package checks

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	fastreq "github.com/starton-io/tyrscale/go-kit/pkg/fastreq/http"
	healthcheck "github.com/starton-io/tyrscale/go-kit/pkg/health/checks"
	"github.com/valyala/fasthttp"
)

type Ping struct {
	URL            string
	Method         string
	Timeout        int
	client         *fasthttp.Client
	Body           interface{}
	Headers        map[string]string
	ExpectedStatus int
	ExpectBody     string
	httpClient     fastreq.IClient
}

// NewPingChecker : time - millisecond
func NewPingChecker(URL, method string, timeout int, body interface{}, headers map[string]string) *Ping {
	if method == "" {
		method = "GET"
	}

	if timeout == 0 {
		timeout = 500
	}

	pingChecker := Ping{
		httpClient: fastreq.NewBuilder().
			SetHeaders(headers).
			SetResponseTimeout(time.Duration(timeout) * time.Millisecond).Build(),
		URL:     URL,
		Method:  method,
		Timeout: timeout,
		Body:    body,
		Headers: headers,
	}
	pingChecker.client = &fasthttp.Client{}

	return &pingChecker
}

func (p Ping) Check(name string, result *healthcheck.ApplicationHealthDetailed, wg *sync.WaitGroup, checklist chan healthcheck.Integration) {
	defer (*wg).Done()
	var (
		start        = time.Now()
		myStatus     = true
		errorMessage = ""
	)

	var body []byte
	var err error
	if p.Body != nil {
		body, err = json.Marshal(p.Body)
		if err != nil {
			myStatus = false
			result.Status = false
			errorMessage = fmt.Sprintf("request failed: %s -> %s with error: %s", p.Method, p.URL, err)
		}
	}

	response, err := p.httpClient.Do(p.URL, p.Method, p.Headers, body)
	if err != nil {
		myStatus = false
		result.Status = false
		errorMessage = fmt.Sprintf("request failed: %s -> %s. error: %s", p.Method, p.URL, err)
	} else if response.StatusCode >= 400 {
		myStatus = false
		result.Status = false
		errorMessage = fmt.Sprintf("request failed: %s -> %s. code: %d. expected: %d", p.Method, p.URL, response.StatusCode, p.ExpectedStatus)
	}

	checklist <- healthcheck.Integration{
		Name:         name,
		Kind:         "ping",
		Status:       myStatus,
		ResponseTime: time.Since(start).Seconds(),
		Error:        errorMessage,
	}

}
