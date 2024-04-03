package sorter

import (
	"fmt"
	"reflect"
	"sort"
)

type ISorter interface {
	Sort(slice interface{}) error
}

// sortByField sorts a slice by a specified field with a custom comparison function.
func sortByField(slice interface{}, jsonTag string, lessFunc func(ifield, jfield reflect.Value) bool) error {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return fmt.Errorf("sortByField error: provided data is not a slice")
	}

	// Cache field index
	var fieldIndex int = -1
	sampleElem := reflect.Indirect(rv.Index(0))
	for k := 0; k < sampleElem.NumField(); k++ {
		field := sampleElem.Type().Field(k)
		tag := field.Tag.Get("json")
		if tag == jsonTag || tag == jsonTag+",omitempty" {
			fieldIndex = k
			break
		}
	}
	if fieldIndex == -1 {
		return fmt.Errorf("field with JSON tag '%s' not found", jsonTag)
	}

	sort.Slice(rv.Interface(), func(i, j int) bool {
		iv := reflect.Indirect(rv.Index(i)).Field(fieldIndex)
		jv := reflect.Indirect(rv.Index(j)).Field(fieldIndex)

		return lessFunc(iv, jv)
	})

	return nil
}

type SortByField struct {
	Field      string
	Descending bool
}

func (s SortByField) Sort(slice interface{}) error {
	if s.Descending {
		return sortByField(slice, s.Field, func(ifield, jfield reflect.Value) bool {
			return ifield.String() > jfield.String()
		})
	}
	return sortByField(slice, s.Field, func(ifield, jfield reflect.Value) bool {
		return ifield.String() < jfield.String()
	})
}
