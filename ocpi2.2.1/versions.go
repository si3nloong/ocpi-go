package ocpi221

import (
	"context"
	"net/http"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *Client) Versions(ctx context.Context) ([]ocpi.Version, error) {
	var o ocpi.Response[[]ocpi.Version]
	if err := c.do(ctx, http.MethodGet, c.versionUrl, nil, &o); err != nil {
		return nil, err
	}
	return o.Data, nil
}
