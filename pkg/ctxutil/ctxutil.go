package ctxutil

import "context"

type contextKey string

const (
	sessionIDKey = "sessionID"
)

func WithSessionID(ctx context.Context, sessionID string) context.Context {
	return context.WithValue(ctx, sessionIDKey, sessionID)
}

func GetSessionID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(sessionIDKey).(string)
	return id, ok
}
