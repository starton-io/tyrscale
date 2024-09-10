package probes

import (
	"github.com/starton-io/tyrscale/gateway/pkg/config"
	healthcheck "github.com/starton-io/tyrscale/go-kit/pkg/health/checks"
	infraHealthcheck "github.com/starton-io/tyrscale/go-kit/pkg/health/infra"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
)

type HealthCheckApplication interface {
	LiveEndpoint() healthcheck.ApplicationHealthDetailed
	ReadyEndpoint() healthcheck.ApplicationHealthDetailed
}

type healthCheckApplication struct {
	healthcheckApp healthcheck.IHealthCheckApplication
}

func NewHealthChecker(configuration *config.Schema, redisClient kv.IRedisStore) HealthCheckApplication {
	// init healcheck app
	healthcheckApp := healthcheck.NewApplication(configuration.ServerName, configuration.AppVersion)

	// kv checker
	redisChecker := infraHealthcheck.NewKVChecker(redisClient)
	healthcheckApp.AddReadinessCheck("redis checker", redisChecker)

	return &healthCheckApplication{
		healthcheckApp: healthcheckApp,
	}
}

func (app healthCheckApplication) LiveEndpoint() healthcheck.ApplicationHealthDetailed {
	return app.healthcheckApp.LiveChecker()
}

func (app healthCheckApplication) ReadyEndpoint() healthcheck.ApplicationHealthDetailed {
	return app.healthcheckApp.ReadyChecker()
}
