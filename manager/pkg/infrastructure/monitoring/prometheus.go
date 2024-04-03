package monitoring

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type ExternalCheckerProm struct {
	client v1.API
}

func NewExternalCheckerProm(url string) *ExternalCheckerProm {
	client, err := api.NewClient(api.Config{
		Address: url,
	})
	api := v1.NewAPI(client)
	if err != nil {
		log.Fatalf("Error creating Prometheus client: %v", err)
	}
	return &ExternalCheckerProm{client: api}
}

func (e *ExternalCheckerProm) Check(promql string) (bool, error) {
	r := v1.Range{
		Start: time.Now().Add(-time.Hour),
		End:   time.Now(),
		Step:  time.Minute,
	}
	_, warnings, err := e.client.QueryRange(context.Background(), promql, r)
	if err != nil {
		return false, err
	}

	//TODO: need to handle warnings depends on the use case
	if len(warnings) > 0 {
		return false, fmt.Errorf("warnings: %v", warnings)
	}
	return true, nil
}
