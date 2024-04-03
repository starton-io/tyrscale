package mocks

import "encoding/json"

// Mock filter for testing
type MockFilter struct{}

func (m *MockFilter) GetFilter() (count int64, preFilterKey string, preFilter bool) {
	return 10, "some-pattern*", true
}

func (m *MockFilter) ShouldInclude(value string) bool {
	// Assuming the filter checks for a specific condition in the JSON
	var data map[string]interface{}
	json.Unmarshal([]byte(value), &data)
	return data["include"] == true
}
