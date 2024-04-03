package sorter

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type ISorterPB[T protoreflect.ProtoMessage] interface {
	Sort(protoMessage []T) ([]T, error)
}

type SortByFieldPB[T protoreflect.ProtoMessage] struct {
	Field      string
	Descending bool
}

func NewSortByFieldPB[T protoreflect.ProtoMessage](field string, descending bool) ISorterPB[T] {
	return SortByFieldPB[T]{
		Field:      field,
		Descending: descending,
	}
}

func (s SortByFieldPB[T]) Sort(protoMessage []T) ([]T, error) {
	if len(protoMessage) == 0 {
		return protoMessage, nil
	}

	// Retrieve the field descriptor once
	fieldDescriptor := protoMessage[0].ProtoReflect().Descriptor().Fields().ByName(protoreflect.Name(s.Field))
	if fieldDescriptor == nil {
		return nil, fmt.Errorf("field %s does not exist", s.Field)
	}
	kind := fieldDescriptor.Kind().String()

	// Determine the appropriate comparison function
	var compare func(a, b protoreflect.Value) bool
	if s.Descending {
		compare = func(a, b protoreflect.Value) bool {
			return greater(a, b, kind)
		}
	} else {
		compare = func(a, b protoreflect.Value) bool {
			return less(a, b, kind)
		}
	}

	// Sort the slice of proto messages using the determined comparison function
	sort.Slice(protoMessage, func(i, j int) bool {
		valI := protoMessage[i].ProtoReflect().Get(fieldDescriptor)
		valJ := protoMessage[j].ProtoReflect().Get(fieldDescriptor)
		return compare(valI, valJ)
	})

	return protoMessage, nil
}

// less is a helper function to compare values of different types
func less(a protoreflect.Value, b protoreflect.Value, kind string) bool {
	switch kind {
	case "int64":
		return a.Int() < b.Int()
	case "uint64":
		return a.Uint() < b.Uint()
	case "float64":
		return a.Float() < b.Float()
	case "string":
		return a.String() < b.String()
	default:
		fmt.Println("unsupported type")
		// Handle other types or throw an error
	}
	return false
}

func greater(a protoreflect.Value, b protoreflect.Value, kind string) bool {
	switch kind {
	case "int64":
		return a.Int() > b.Int()
	case "uint64":
		return a.Uint() > b.Uint()
	case "float64":
		return a.Float() > b.Float()
	case "string":
		return a.String() > b.String()
	default:
		fmt.Println("unsupported type")
		// Handle other types or throw an error
	}
	return false
}
