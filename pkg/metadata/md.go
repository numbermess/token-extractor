package metadata

import (
	"context"
	"strings"
)

type MD map[string][]string

func (md MD) Get(k string) []string {
	if val, ok := md[k]; ok {
		return val
	}
	k = strings.ToLower(k)
	return md[k]
}

type mdIncomingKey struct{}

func FromIncomingContext(ctx context.Context) (md MD, ok bool) {
	md, ok = ctx.Value(mdIncomingKey{}).(MD)
	return
}
