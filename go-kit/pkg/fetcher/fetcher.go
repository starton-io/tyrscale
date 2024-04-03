package fetcher

import (
	"github.com/starton-io/tyrscale/go-kit/pkg/responses"
)

type GeneralResponse[T any] struct {
	Status  int               `json:"status"`
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    responses.Body[T] `json:"data"`
}

//go:generate mockery --name IFetcher
type IFetcher[T any] interface {
	FetchAll() ([]T, error)
}
