package utils

import (
	"encoding/json"

	"google.golang.org/protobuf/proto"
)

func Copy(dest interface{}, src interface{}) {
	data, _ := json.Marshal(src)
	_ = json.Unmarshal(data, dest)
}

func ConvertToPB(dest proto.Message, src interface{}) {
	data, _ := json.Marshal(src)
	_ = proto.Unmarshal(data, dest)
}
