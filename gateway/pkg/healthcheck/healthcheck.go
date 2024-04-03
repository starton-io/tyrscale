package healthcheck

import (
	"fmt"
	"time"

	"github.com/starton-io/tyrscale/gateway/pkg/proxy"
)

func NewHealthCheck(healthCheckType HealthCheckType, clientManager proxy.ClientManager, interval uint32, timeout uint32) (HealthCheckInterface, error) {
	switch healthCheckType {
	case EthBlockNumberType:
		return NewHealthEthBlockNumber(clientManager, WithInterval(time.Duration(interval)*time.Millisecond), WithTimeout(time.Duration(timeout)*time.Millisecond)), nil
	default:
		return nil, fmt.Errorf("invalid healthcheck type: %s", healthCheckType)
	}
}
