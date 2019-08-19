package model

import "strings"

const (
	Basic   = "basic"
	Bearer  = "bearer"
	Unknown = "unknown"
)

type Token struct {
	Value string
	Type  string
}

func (t Token) IsBasic() bool {
	return Basic == strings.ToLower(t.Type)
}

func (t Token) IsBearer() bool {
	return Bearer == strings.ToLower(t.Type)
}
