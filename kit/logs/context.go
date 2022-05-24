package logs

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type contextKey struct{}

func ToContext(ctx context.Context, log *zerolog.Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, log)
}

func For(ctx context.Context) *zerolog.Logger {
	if l, ok := ctx.Value(contextKey{}).(*zerolog.Logger); ok {
		return l
	}

	return &log.Logger
}
