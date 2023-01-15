package logs

import (
	"context"
	"github.com/labstack/gommon/random"
)

const (
	TraceHeaderKey = "X-REQUEST-ID"
)

const (
	traceContextKey = "trace_id"
)

func WithContext(ctx context.Context, traceID string) context.Context {
	if traceID == "" {
		traceID = random.String(32)
	}

	return context.WithValue(ctx, traceContextKey, traceID)
}

func TraceID(ctx context.Context) string {
	return ctx.Value(traceContextKey).(string)
}
