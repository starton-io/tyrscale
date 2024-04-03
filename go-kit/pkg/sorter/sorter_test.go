package sorter

import (
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestSortByField(t *testing.T) {
	people := []Person{
		{Name: "Clara", Age: 25},
		{Name: "Bob", Age: 22},
		{Name: "Alice", Age: 30},
	}

	// Test ascending sort by name
	sorterAsc := SortByField{Field: "name", Descending: false}
	err := sorterAsc.Sort(people)
	if err != nil {
		t.Errorf("Failed to sort: %v", err)
	}
	if !isSortedByName(people, false) {
		t.Errorf("Failed to sort ascending: got %v", people)
	}

	// Test descending sort by name
	sorterDesc := SortByField{Field: "name", Descending: true}
	err = sorterDesc.Sort(people)
	if err != nil {
		t.Errorf("Failed to sort: %v", err)
	}
	if !isSortedByName(people, true) {
		t.Errorf("Failed to sort descending: got %v", people)
	}
}

func isSortedByName(people []Person, desc bool) bool {
	for i := 0; i < len(people)-1; i++ {
		if desc {
			if people[i].Name < people[i+1].Name {
				return false
			}
		} else {
			if people[i].Name > people[i+1].Name {
				return false
			}
		}
	}
	return true
}
