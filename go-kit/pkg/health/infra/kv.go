package checks

import (
	"fmt"
	"sync"
	"time"

	healthcheck "github.com/starton-io/tyrscale/go-kit/pkg/health/checks"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
)

type KV struct {
	Client kv.IRedisStore
}

func NewKVChecker(kvdb kv.IRedisStore) *KV {
	return &KV{Client: kvdb}
}

func (r *KV) Check(name string, result *healthcheck.ApplicationHealthDetailed, wg *sync.WaitGroup, checklist chan healthcheck.Integration) {
	defer (*wg).Done()
	var (
		start        = time.Now()
		myStatus     = true
		errorMessage = ""
	)

	if r.Client == nil {
		myStatus = false
		result.Status = false
		errorMessage = "connection is nil"
	}

	if err := r.Client.Ping(); err != nil {
		myStatus = false
		result.Status = false
		errorMessage = fmt.Sprintf("ping failed: %s", err)
	}

	checklist <- healthcheck.Integration{
		Name:         name,
		Kind:         "redis",
		Status:       myStatus,
		ResponseTime: time.Since(start).Seconds(),
		Error:        errorMessage,
	}

}
