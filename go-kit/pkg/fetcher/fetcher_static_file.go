package fetcher

import (
	"io"

	"gopkg.in/yaml.v2"
)

type StaticFileFetcher[T any] struct {
	Entities []T `yaml:"items" json:"items"`
}

func NewStaticFileFetcher[T any](contentFile io.Reader) IFetcher[T] {
	var entities StaticFileFetcher[T]
	decoder := yaml.NewDecoder(contentFile)
	if err := decoder.Decode(&entities); err != nil {
		panic(err)
	}
	return &entities
}

func (f *StaticFileFetcher[T]) FetchAll() ([]T, error) {
	return f.Entities, nil
}
