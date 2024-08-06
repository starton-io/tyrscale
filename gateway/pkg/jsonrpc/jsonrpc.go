package jsonrpc

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"strings"
)

const (
	vsn = "2.0"

	subscribeMethodSuffix   = "_subscribe"
	unsubscribeMethodSuffix = "_unsubscribe"

	defaultErrorCode = -32000

	ParseErrorCode     = -32700
	InvalidRequestCode = -32600
	MethodNotFoundCode = -32601
	invalidParamsCode  = -32602
	InternalErrorCode  = -32603
	methodExistsCode   = -32000
	uRLSchemeErrorCode = -32001
)

//var (
//	null = json.RawMessage("null")
//)

// A value of this type can a JSON-RPC request, notification, successful response or
// error response. Which one it is depends on the fields.
type JsonrpcMessage struct {
	Version string          `json:"jsonrpc,omitempty"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Error   *jsonError      `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

func (msg *JsonrpcMessage) IsNotification() bool {
	return msg.ID == nil && msg.Method != ""
}

func (msg *JsonrpcMessage) IsCall() bool {
	return msg.hasValidID() && msg.Method != ""
}

func (msg *JsonrpcMessage) IsResponse() bool {
	return msg.hasValidID() && msg.Method == "" && msg.Params == nil && (msg.Result != nil || msg.Error != nil)
}

func (msg *JsonrpcMessage) hasValidID() bool {
	return len(msg.ID) > 0 && msg.ID[0] != '{' && msg.ID[0] != '['
}

func (msg *JsonrpcMessage) IsSubscribe() bool {
	return strings.HasSuffix(msg.Method, subscribeMethodSuffix)
}

func (msg *JsonrpcMessage) IsUnsubscribe() bool {
	return strings.HasSuffix(msg.Method, unsubscribeMethodSuffix)
}

func (msg *JsonrpcMessage) IsError() bool {
	return msg.Error != nil
}

func (msg *JsonrpcMessage) MustJSONBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b
}

func (msg *JsonrpcMessage) CopyWithID(id json.RawMessage) (clone *JsonrpcMessage) {
	m := &JsonrpcMessage{}
	*m = *msg // copy
	m.ID = id
	return m
}

func ErrorMessage(err error) *JsonrpcMessage {
	msg := &JsonrpcMessage{Version: vsn, Error: &jsonError{
		Code:    defaultErrorCode,
		Message: err.Error(),
	}}
	return msg
}

func (msg *JsonrpcMessage) ErrorResponse(err error) *JsonrpcMessage {
	resp := ErrorMessage(err)
	resp.ID = msg.ID
	return resp
}

type jsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (err *jsonError) Error() string {
	if err.Message == "" {
		return fmt.Sprintf("json-rpc error %d", err.Code)
	}
	return err.Message
}

func (err *jsonError) ErrorCode() int {
	return err.Code
}

func (err *jsonError) ErrorData() interface{} {
	return err.Data
}

type params []interface{}

func (msg *JsonrpcMessage) CacheKey() (string, error) {
	out := msg.Method
	params := params{}
	err := json.Unmarshal(msg.Params, &params)
	if err != nil {
		return "", err
	}
	for _, p := range params {
		out += fmt.Sprintf("/%v", p)
	}
	return fmt.Sprintf("%x", sha1.Sum([]byte(out))), nil
}

// ParseMessage parses raw bytes as a (batch of) JSON-RPC message(s). There are no error
// checks in this function because the raw message has already been syntax-checked when it
// is called. Any non-JSON-RPC messages in the input return the zero value of
// JsonrpcMessage.
func ParseMessage(raw json.RawMessage) ([]*JsonrpcMessage, bool) {
	if !isBatch(raw) {
		msgs := []*JsonrpcMessage{{}}
		_ = json.Unmarshal(raw, &msgs[0])
		return msgs, false
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	_, _ = dec.Token() // skip '['
	var msgs []*JsonrpcMessage
	for dec.More() {
		msgs = append(msgs, new(JsonrpcMessage))
		_ = dec.Decode(&msgs[len(msgs)-1])
	}
	return msgs, true
}

// isBatch returns true when the first non-whitespace characters is '['
func isBatch(raw json.RawMessage) bool {
	for _, c := range raw {
		// skip insignificant whitespace (http://www.ietf.org/rfc/rfc4627.txt)
		if c == 0x20 || c == 0x09 || c == 0x0a || c == 0x0d {
			continue
		}
		return c == '['
	}
	return false
}
