package kv

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ParamsFilterPB[T protoreflect.ProtoMessage] struct {
	MatchCriteria    map[string]string
	EnablePrefilter  bool
	PrefilterPattern string // Optional: Use if you want to specify a pattern dynamically
	Count            int64
	ProtoType        reflect.Type
}

func NewParamsFilterPB[T protoreflect.ProtoMessage](matchCriteria map[string]string, enablePrefilter bool, prefilterPattern string, count int64) *ParamsFilterPB[T] {
	var msg T
	return &ParamsFilterPB[T]{
		MatchCriteria:    matchCriteria,
		EnablePrefilter:  enablePrefilter,
		PrefilterPattern: prefilterPattern,
		Count:            count,
		ProtoType:        reflect.TypeOf(msg).Elem(),
	}
}

func (pf *ParamsFilterPB[T]) ShouldInclude(b []byte) bool {
	if pf.EnablePrefilter {
		return true
	}

	msg := reflect.New(pf.ProtoType).Interface().(T)
	if err := proto.Unmarshal(b, msg); err != nil {
		fmt.Println("err", err)
		return false
	}

	for key, expectedValue := range pf.MatchCriteria {
		fieldDescriptor := msg.ProtoReflect().Descriptor().Fields().ByName(protoreflect.Name(key))
		if fieldDescriptor == nil || msg.ProtoReflect().Get(fieldDescriptor).String() != expectedValue {
			return false
		}
	}
	return true
}

func (pf *ParamsFilterPB[T]) GetFilter() (int64, string, bool) {
	return pf.Count, pf.PrefilterPattern, pf.EnablePrefilter
}
