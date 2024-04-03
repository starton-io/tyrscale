package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func UnmarshalMapValuesIntoSlice[T any](src map[string]any) ([]T, error) {
	dst := make([]T, 0)
	for _, value := range src {
		genericInstance := new(T)
		valueStr, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("cannot convert type %T to string", value)
		}
		if err := json.Unmarshal([]byte(valueStr), genericInstance); err != nil {
			return nil, err
		}
		dst = append(dst, *genericInstance)
	}
	return dst, nil
}

func UnmarshalSliceBytesToProto[T protoreflect.ProtoMessage](src [][]byte) ([]T, error) {
	dst := make([]T, 0, len(src)) // Preallocate slice with capacity equal to the size of the map
	var genericInstance T
	var protoType = reflect.TypeOf(genericInstance).Elem()

	for _, value := range src {
		// Create a new instance of the type
		genericInstance = reflect.New(protoType).Interface().(T)
		if err := proto.Unmarshal(value, genericInstance); err != nil {
			return nil, err
		}
		dst = append(dst, genericInstance)
	}
	return dst, nil
}

func StructToMapStr[T any](s T, tagValue string) (map[string]string, error) {
	result := make(map[string]string)
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Ensure we are dealing with a struct
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input is not a struct")
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		typeField := typ.Field(i)
		tag := typeField.Tag
		key := strings.Split(tag.Get(tagValue), ",")[0]
		if key != "" && field.Kind() == reflect.String && field.String() != "" {
			result[key] = field.String()
		}
	}

	return result, nil
}
