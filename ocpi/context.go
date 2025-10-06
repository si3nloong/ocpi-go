package ocpi

import (
	"context"
	"net/http"
	"net/url"
)

type ctxKey string

const requestCtx ctxKey = "request_context"

type RequestContext struct {
	Token string

	RequestID string

	Host       string
	URL        *url.URL
	RequestURI string

	CorrelationID string

	ToPartyID     string
	ToCountryCode string

	FromPartyID     string
	FromCountryCode string
}

func NewRequestContextWithRequest(r *http.Request, token string) *RequestContext {
	uri, err := url.Parse(r.URL.String())
	if err != nil {
		panic(err)
	}
	return &RequestContext{
		Token:           token,
		URL:             uri,
		RequestURI:      r.RequestURI,
		Host:            r.Host,
		RequestID:       r.Header.Get(HttpHeaderXRequestID),
		CorrelationID:   r.Header.Get(HttpHeaderXCorrelationID),
		ToPartyID:       r.Header.Get(HttpHeaderOCPIToPartyID),
		ToCountryCode:   r.Header.Get(HttpHeaderOCPIToCountryCode),
		FromPartyID:     r.Header.Get(HttpHeaderOCPIFromPartyID),
		FromCountryCode: r.Header.Get(HttpHeaderOCPIFromCountryCode),
	}
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
