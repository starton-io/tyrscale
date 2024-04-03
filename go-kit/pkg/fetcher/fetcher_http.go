package fetcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	fastreq "github.com/starton-io/tyrscale/go-kit/pkg/fastreq/http"
)

type httpFetcher[T any] struct {
	ManagerURL string
	httpClient fastreq.IClient
	filter     map[string]string
}

type Option[T any] func(*httpFetcher[T])

func WithFilter[T any](filter map[string]string) Option[T] {
	return func(fetcher *httpFetcher[T]) {
		fetcher.filter = filter
	}
}

func NewHttpFetcher[T any](managerURL string, httpClient fastreq.IClient, opts ...Option[T]) IFetcher[T] {
	httpFetcher := &httpFetcher[T]{ManagerURL: managerURL, httpClient: httpClient}
	for _, opt := range opts {
		opt(httpFetcher)
	}
	return httpFetcher
}

func (r *httpFetcher[T]) FetchAll() ([]T, error) {
	params := url.Values{}
	for k, v := range r.filter {
		params.Add(k, v)
	}
	if len(params) > 0 {
		r.ManagerURL = fmt.Sprintf("%s?%s", r.ManagerURL, params.Encode())
	}
	resp, err := r.httpClient.Get(r.ManagerURL, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch rpc endpoints: %s", resp.Status)
	}

	responses := new(GeneralResponse[T])
	err = json.Unmarshal(resp.Body, responses)
	if err != nil {
		return nil, err
	}

	entities := make([]T, 0)
	entities = append(entities, responses.Data.Items...)
	return entities, nil
}
