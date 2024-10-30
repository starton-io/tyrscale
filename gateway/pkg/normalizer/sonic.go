package normalizer

import (
	"reflect"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/option"
)

var SonicCfg sonic.API

func init() {
	err := sonic.Pretouch(
		reflect.TypeOf(RPCErrorResponse{}),
		option.WithCompileMaxInlineDepth(1),
	)
	if err != nil {
		panic(err)
	}
	err = sonic.Pretouch(
		reflect.TypeOf(RPCRequest{}),
		option.WithCompileMaxInlineDepth(1),
	)
	if err != nil {
		panic(err)
	}
	err = sonic.Pretouch(
		reflect.TypeOf(NormalizedRequest{}),
		option.WithCompileMaxInlineDepth(1),
	)
	if err != nil {
		panic(err)
	}
	SonicCfg = sonic.Config{
		CopyString:              false,
		NoQuoteTextMarshaler:    true,
		NoValidateJSONMarshaler: true,
		NoValidateJSONSkip:      true,
		EscapeHTML:              false,
		SortMapKeys:             false,
		CompactMarshaler:        true,
		ValidateString:          false,
	}.Froze()
}
