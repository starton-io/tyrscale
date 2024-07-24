package circuitbreaker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCircuitBreaker(t *testing.T) {
	settings := Settings{
		Enabled:                true,
		Name:                   "test",
		MaxRequests:            5,
		MaxConsecutiveFailures: 2,
		Interval:               1000,
		Timeout:                5000,
	}

	cb := NewCircuitBreaker(settings)
	assert.NotNil(t, cb)
}

func TestAddAndGetCircuitBreaker(t *testing.T) {
	settings := Settings{
		Enabled:                true,
		Name:                   "test",
		MaxRequests:            5,
		MaxConsecutiveFailures: 2,
		Interval:               1000,
		Timeout:                5000,
	}

	cb := NewCircuitBreaker(settings)
	cb.Add("testKey")

	circuitBreaker := cb.Get("testKey")
	assert.NotNil(t, circuitBreaker)
}

func TestRemoveCircuitBreaker(t *testing.T) {
	settings := Settings{
		Enabled:                true,
		Name:                   "test",
		MaxRequests:            5,
		MaxConsecutiveFailures: 2,
		Interval:               1000,
		Timeout:                5000,
	}

	cb := NewCircuitBreaker(settings)
	cb.Add("testKey")
	cb.Remove("testKey")

	circuitBreaker := cb.Get("testKey")
	assert.Nil(t, circuitBreaker)
}

func TestCleanCircuitBreakers(t *testing.T) {
	settings := Settings{
		Enabled:                true,
		Name:                   "test",
		MaxRequests:            5,
		MaxConsecutiveFailures: 2,
		Interval:               1000,
		Timeout:                5000,
	}

	cb := NewCircuitBreaker(settings)
	cb.Add("testKey1")
	cb.Add("testKey2")
	cb.Clean()

	assert.Nil(t, cb.Get("testKey1"))
	assert.Nil(t, cb.Get("testKey2"))
}
