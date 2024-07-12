package healthcheck

import (
	"fmt"
	"time"

	"github.com/starton-io/tyrscale/gateway/pkg/proxy"
)

func NewHealthCheck(clientManager proxy.ClientManager, healthCheckConfig *HealthCheckConfig) (HealthCheckInterface, error) {
	switch healthCheckConfig.Type {
	case EthBlockNumberType:
		return NewHealthEthBlockNumber(clientManager,
			WithInterval(time.Duration(healthCheckConfig.Interval)*time.Millisecond),
			WithTimeout(time.Duration(healthCheckConfig.Timeout)*time.Millisecond),
		), nil
	case EthSyncingType:
		return NewHealthEthSyncing(clientManager,
			WithEthSyncingInterval(time.Duration(healthCheckConfig.Interval)*time.Millisecond),
			WithEthSyncingTimeout(time.Duration(healthCheckConfig.Timeout)*time.Millisecond),
		), nil
	case CustomType:
		return NewHealthCustom(clientManager,
			WithCustomHealthCheckInterval(time.Duration(healthCheckConfig.Interval)*time.Millisecond),
			WithCustomHealthCheckTimeout(time.Duration(healthCheckConfig.Timeout)*time.Millisecond),
			WithCustomHealthCheckRequest(healthCheckConfig.Request),
		), nil
	default:
		return nil, fmt.Errorf("invalid healthcheck type: %s", healthCheckConfig.Type)
	}
}
