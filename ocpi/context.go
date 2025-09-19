package ocpi

import "context"

type ctxKey string

const requestCtx ctxKey = "request_context"

type RequestContext struct {
	Token string

	RequestID string

	RequestURI string

	CorrelationID string

	ToPartyID     string
	ToCountryCode string

	FromPartyID     string
	FromCountryCode string
}

func GetRequestContext(ctx context.Context) *RequestContext {
	if val, ok := ctx.Value(requestCtx).(*RequestContext); ok {
		return val
	}
	return &RequestContext{}
}

func WithRequestContext(ctx context.Context, rc *RequestContext) context.Context {
	return context.WithValue(ctx, requestCtx, rc)
}

const responseCtx ctxKey = "response_context"

type ResponseContext struct {
	Token string

	RequestID string

	CorrelationID string

	ToPartyID     string
	ToCountryCode string

	FromPartyID     string
	FromCountryCode string
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
