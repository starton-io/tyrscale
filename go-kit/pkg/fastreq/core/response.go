package core

import (
	"encoding/json"
)

type Response struct {
	StatusCode int
	Status     string
	Body       []byte
	Header     map[string][]string
}

func (r *Response) Bytes() []byte {
	return r.Body
}

func (r *Response) String() string {
	return string(r.Body)
}

func (r *Response) Unmarshal(v any) error {
	return json.Unmarshal(r.Body, v)
}
