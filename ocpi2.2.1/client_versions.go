package ocpi221

import (
	"context"
	"net/http"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) GetVersions(ctx context.Context) (*ocpi.Response[ocpi.Versions], error) {
	token, err := c.ocpi.GetCredentialsTokenC(ctx)
	if err != nil {
		return nil, err
	}
	var res ocpi.Response[ocpi.Versions]
	if err := c.do(ctx, token, http.MethodGet, c.versionUrl, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) GetVersionDetails(ctx context.Context, version ocpi.Version) (*ocpi.Response[VersionDetails], error) {
	token, err := c.ocpi.GetCredentialsTokenC(ctx)
	if err != nil {
		return nil, err
	}
	var res ocpi.Response[VersionDetails]
	if err := c.do(ctx, token, http.MethodGet, version.URL, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
