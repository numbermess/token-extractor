package token

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"strings"
	"token-extractor/pkg/model"
)

func ExtractToken(line string) (*model.Token, error) {

	line = strings.TrimSpace(line)

	fields := strings.Fields(line)
	if len(fields) != 2 {
		return nil, errors.New("the Authorization header line should be splittable into exactly two chunks")
	} else {
		token := &model.Token{}
		switch strings.ToLower(fields[0]) {
		case model.Basic:
			token.Type = model.Basic
		case model.Bearer:
			token.Type = model.Bearer
		default:
			token.Type = model.Unknown
		}
		token.Value = fields[1]

		return token, nil
	}
}

func Extract(ctx context.Context) (*model.Token, error) {
	if md, ok := metadata.FromIncomingContext(ctx); !ok {
		return nil, errors.New("cannot extract metadata from context")
	} else {
		chunks := md.Get("Authorization")
		if len(chunks) == 0 {
			return nil, errors.New("there is no Authorization header present in the metadata")
		}
		if len(chunks) > 0 {
			line := chunks[0]
			return ExtractToken(line)
		}
	}
	return nil, errors.New("unexpected error")
}
