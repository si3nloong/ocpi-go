package ocpi221

import (
	"context"
	"net/http"
)

func (c *Client) Versions(ctx context.Context) ([]Version, error) {
	var o VersionsResponse
	if err := c.do(ctx, http.MethodGet, c.versionUrl, nil, &o); err != nil {
		return nil, err
	}
	return o.Data, nil
}
