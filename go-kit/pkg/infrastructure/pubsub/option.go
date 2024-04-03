package pubsub

import "strings"

const (
	OptionTypeGlobalPrefix = 0x600
)

type Option interface {
	Type() int
}

type globalPrefix struct {
	prefix string
}

func (to globalPrefix) Type() int {
	return OptionTypeGlobalPrefix
}

func NewGlobalPrefix(prefix string) Option {
	if !strings.HasSuffix(prefix, ":") {
		prefix += ":"
	}
	return globalPrefix{prefix: prefix}
}
