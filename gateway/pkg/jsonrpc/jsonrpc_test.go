package jsonrpc

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestIsNotification(t *testing.T) {
	msg := &JsonrpcMessage{Method: "test_method"}
	if !msg.IsNotification() {
		t.Errorf("Expected IsNotification to be true")
	}

	msg.ID = json.RawMessage(`1`)
	if msg.IsNotification() {
		t.Errorf("Expected IsNotification to be false")
	}
}

func TestIsCall(t *testing.T) {
	msg := &JsonrpcMessage{ID: json.RawMessage(`1`), Method: "test_method"}
	if !msg.IsCall() {
		t.Errorf("Expected IsCall to be true")
	}

	msg.ID = nil
	if msg.IsCall() {
		t.Errorf("Expected IsCall to be false")
	}
}

func TestIsResponse(t *testing.T) {
	msg := &JsonrpcMessage{ID: json.RawMessage(`1`), Result: json.RawMessage(`"result"`)}

	if !msg.IsResponse() {
		t.Errorf("Expected IsResponse to be true")
	}

	msg.Method = "test_method"
	if msg.IsResponse() {
		t.Errorf("Expected IsResponse to be false")
	}
}

func TestIsSubscribe(t *testing.T) {
	msg := &JsonrpcMessage{Method: "test_method_subscribe"}
	if !msg.IsSubscribe() {
		t.Errorf("Expected IsSubscribe to be true")
	}

	msg.Method = "test_method"
	if msg.IsSubscribe() {
		t.Errorf("Expected IsSubscribe to be false")
	}
}

func TestIsUnsubscribe(t *testing.T) {
	msg := &JsonrpcMessage{Method: "test_method_unsubscribe"}
	if !msg.IsUnsubscribe() {
		t.Errorf("Expected IsUnsubscribe to be true")
	}

	msg.Method = "test_method"
	if msg.IsUnsubscribe() {
		t.Errorf("Expected IsUnsubscribe to be false")
	}
}

func TestIsError(t *testing.T) {
	msg := &JsonrpcMessage{Error: &jsonError{Code: -32000, Message: "error"}}
	if !msg.IsError() {
		t.Errorf("Expected IsError to be true")
	}

	msg.Error = nil
	if msg.IsError() {
		t.Errorf("Expected IsError to be false")
	}
}

func TestMustJSONBytes(t *testing.T) {
	msg := &JsonrpcMessage{Version: "2.0", Method: "test_method"}
	bytes := msg.MustJSONBytes()
	expected := `{"jsonrpc":"2.0","method":"test_method"}`
	if string(bytes) != expected {
		t.Errorf("Expected %s, got %s", expected, string(bytes))
	}
}

func TestCopyWithID(t *testing.T) {
	msg := &JsonrpcMessage{Method: "test_method"}
	newID := json.RawMessage(`2`)
	clone := msg.CopyWithID(newID)
	if string(clone.ID) != "2" {
		t.Errorf("Expected ID to be 2, got %s", string(clone.ID))
	}
	if clone.Method != "test_method" {
		t.Errorf("Expected Method to be test_method, got %s", clone.Method)
	}
}

func TestErrorMessage(t *testing.T) {
	err := ErrorMessage(fmt.Errorf("test error"))
	if err.Error.Message != "test error" {
		t.Errorf("Expected error message to be 'test error', got %s", err.Error.Message)
	}
	if err.Error.Code != defaultErrorCode {
		t.Errorf("Expected error code to be %d, got %d", defaultErrorCode, err.Error.Code)
	}
}

func TestErrorResponse(t *testing.T) {
	msg := &JsonrpcMessage{ID: json.RawMessage(`1`)}
	err := msg.ErrorResponse(fmt.Errorf("test error"))
	if string(err.ID) != "1" {
		t.Errorf("Expected ID to be 1, got %s", string(err.ID))
	}
	if err.Error.Message != "test error" {
		t.Errorf("Expected error message to be 'test error', got %s", err.Error.Message)
	}
}

func TestCacheKey(t *testing.T) {
	msg := &JsonrpcMessage{Method: "test_method", Params: json.RawMessage(`[1, "param"]`)}
	key, err := msg.CacheKey()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := "570840821cb3e12af62547e9241924ec357cf4cd"
	if key != expected {
		t.Errorf("Expected %s, got %s", expected, key)
	}
}

func TestParseMessage(t *testing.T) {
	raw := json.RawMessage(`{"jsonrpc":"2.0","method":"test_method"}`)
	msgs, isBatch := ParseMessage(raw)
	if isBatch {
		t.Errorf("Expected isBatch to be false")
	}
	if len(msgs) != 1 {
		t.Errorf("Expected 1 message, got %d", len(msgs))
	}
	if msgs[0].Method != "test_method" {
		t.Errorf("Expected method to be 'test_method', got %s", msgs[0].Method)
	}
}

func TestIsBatch(t *testing.T) {
	raw := json.RawMessage(`[{"jsonrpc":"2.0","method":"test_method"}]`)
	if !isBatch(raw) {
		t.Errorf("Expected isBatch to be true")
	}

	raw = json.RawMessage(`{"jsonrpc":"2.0","method":"test_method"}`)
	if isBatch(raw) {
		t.Errorf("Expected isBatch to be false")
	}
}
