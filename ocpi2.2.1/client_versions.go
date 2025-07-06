package ocpi221

import (
	"context"
	"net/http"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) GetVersions(ctx context.Context) (*ocpi.Response[ocpi.Versions], error) {
	var res ocpi.Response[ocpi.Versions]
	if err := c.do(ctx, http.MethodGet, c.versionUrl, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
