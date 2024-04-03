package utils

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationUUID(t *testing.T) {
	type args struct {
		uuid string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid uuid",
			args: args{uuid: "6ba7b810-9dad-11d1-8e12-00c04fd430c8"},
			want: true,
		},
		{
			name: "invalid uuid",
			args: args{uuid: "6ba7b810-9dad-11d1"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidUUID(tt.args.uuid); got != tt.want {
				t.Errorf("IsValidUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshalMapValuesIntoSlice(t *testing.T) {
	// Test case with valid string JSON representations
	src := map[string]any{
		"one": `{"name":"John", "age":30}`,
		"two": `{"name":"Jane", "age":25}`,
	}
	expected := []struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		{Name: "John", Age: 30},
		{Name: "Jane", Age: 25},
	}

	result, err := UnmarshalMapValuesIntoSlice[struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}](src)

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name > result[j].Name
	})

	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	// Test case with invalid JSON format
	srcInvalid := map[string]any{
		"one": "not a json",
	}
	_, errInvalid := UnmarshalMapValuesIntoSlice[struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}](srcInvalid)

	assert.Error(t, errInvalid)

	// Test case with non-string values
	srcNonString := map[string]any{
		"one": 123,
	}
	_, errNonString := UnmarshalMapValuesIntoSlice[struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}](srcNonString)

	assert.Error(t, errNonString)
}

type TestStruct struct {
	Name    string `json:"name"`
	Age     string `json:"age"`
	Address string `json:"address"`
}

func TestStructToMapStr(t *testing.T) {
	testObj := TestStruct{
		Name:    "John Doe",
		Age:     "30",
		Address: "123 Elm St",
	}

	expected := map[string]string{
		"name":    "John Doe",
		"age":     "30",
		"address": "123 Elm St",
	}

	result, err := StructToMapStr(testObj, "json")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test with non-struct type
	_, err = StructToMapStr("this is not a struct", "json")
	if err == nil {
		t.Errorf("Expected an error for non-struct input, but got nil")
	}
}

func TestCopy(t *testing.T) {
	original := struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  30,
	}

	var copied struct {
		Name string
		Age  int
	}

	Copy(&copied, original)

	if !reflect.DeepEqual(original, copied) {
		t.Errorf("Copy failed: expected %+v, got %+v", original, copied)
	}
}
