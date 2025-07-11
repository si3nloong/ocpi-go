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
	c.rw.Lock()
	versions, err := res.Data()
	if err == nil {
		c.versions = versions
	}
	c.rw.Unlock()
	return &res, nil
}

func (c *ClientConn) GetVersionDetails(ctx context.Context) (*ocpi.Response[VersionDetails], error) {
	c.rw.RLock()
	if c.versions == nil {
		c.rw.RUnlock()
		if _, err := c.GetVersions(ctx); err != nil {
			return nil, err
		}
		c.rw.RLock()
	}
	mutualVersion, ok := c.versions.MutualVersion(ocpi.VersionNumber221)
	if !ok {
		c.rw.RUnlock()
		return nil, ocpi.ErrNoMutualVersion
	}
	c.rw.RUnlock()
	var res ocpi.Response[VersionDetails]
	if err := c.do(ctx, http.MethodGet, mutualVersion.URL, nil, &res); err != nil {
		return nil, err
	}
	c.rw.Lock()
	versionDetails, err := res.Data()
	if err == nil {
		c.versionDetails = &versionDetails
	}
	c.rw.Unlock()
	return &res, nil
}
