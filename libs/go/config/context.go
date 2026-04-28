package config

import (
	"context"
	"errors"
)

const (
	IdentityId = "identityId"
	SessionId  = "sessionId"
)

func MustIdentityId(ctx context.Context) string {
	id, ok := ctx.Value(IdentityId).(string)
	if !ok {
		panic(errors.New("invalid identity id"))
	}
	return id
}
