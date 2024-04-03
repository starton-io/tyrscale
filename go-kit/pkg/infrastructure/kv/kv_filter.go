package kv

import (
	"encoding/json"
	"fmt"
)

// FilterStrategy defines the interface for filtering key-value pairs.
//
//go:generate mockery --name=IFilterStrategy --output=./mocks
type IFilterStrategy interface {
	ShouldInclude([]byte) bool
	GetFilter() (int64, string, bool)
}

type ParamsFilter struct {
	MatchCriteria    map[string]string
	EnablePrefilter  bool
	PrefilterPattern string // Optional: Use if you want to specify a pattern dynamically
	Count            int64
}

// value is stringified JSON, it should be parsed to get the key-value pairs
func (pf *ParamsFilter) ShouldInclude(value []byte) bool {
	//check if prefilter is enable
	if pf.EnablePrefilter {
		return true
	}

	// Parse the JSON string into a map
	var data map[string]interface{}
	if err := json.Unmarshal(value, &data); err != nil {
		fmt.Println("error parsing JSON:", err)
		return false // Assuming we want to exclude items if there's an error parsing
	}

	// Check if all match criteria are met
	for key, expectedValue := range pf.MatchCriteria {
		// convert actualValue to string
		actualValue, ok := data[key]
		actualValueStr := fmt.Sprintf("%v", actualValue)
		if !ok || actualValueStr != expectedValue {
			return false // Criteria not met, exclude this item
		}
	}

	// All criteria met, include this item
	return true
}

func (pf *ParamsFilter) GetFilter() (int64, string, bool) {
	return pf.Count, pf.PrefilterPattern, pf.EnablePrefilter
}
