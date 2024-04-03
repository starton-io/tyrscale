package core

import (
	"reflect"
	"testing"
)

func TestResponse_Bytes(t *testing.T) {
	resp := Response{
		Body: []byte("test body"),
	}

	expected := []byte("test body")
	got := resp.Bytes()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Bytes() = %v, want %v", got, expected)
	}
}

func TestResponse_String(t *testing.T) {
	resp := Response{
		Body: []byte("test body"),
	}

	expected := "test body"
	got := resp.String()

	if got != expected {
		t.Errorf("String() = %v, want %v", got, expected)
	}
}

func TestResponse_Unmarshal(t *testing.T) {
	resp := Response{
		Body: []byte(`{"name":"test","value":123}`),
	}

	type Data struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}

	expected := Data{Name: "test", Value: 123}
	var got Data

	err := resp.Unmarshal(&got)
	if err != nil {
		t.Errorf("Unmarshal() error = %v", err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Unmarshal() = %v, want %v", got, expected)
	}
}
