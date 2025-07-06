package ocpi221

import "context"

const requestCtx ctxKey = "request_context"

type RequestContext struct {
	token string

	requestID string

	requestURI string

	correlationID string

	ToPartyID     string
	ToCountryCode string

	FromPartyID     string
	FromCountryCode string
}

func (r *RequestContext) Token() string {
	return r.token
}

func (r *RequestContext) RequestID() string {
	return r.requestID
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
