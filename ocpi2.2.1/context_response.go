package ocpi221

import "context"

const responseCtx ctxKey = "response_context"

type ResponseContext struct {
	token string

	requestID string

	correlationID string
}

func GetResponseContext(ctx context.Context) *ResponseContext {
	if val, ok := ctx.Value(responseCtx).(*ResponseContext); ok {
		return val
	}
	return nil
}

func WithResponseContext(ctx context.Context, rc *ResponseContext) context.Context {
	return context.WithValue(ctx, responseCtx, rc)
}
