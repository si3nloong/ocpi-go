package ocpi221

import (
	"context"
	"fmt"
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

func (c *ClientConn) SetVersion(ctx context.Context, version ocpi.Version) error {
	c.rw.RLock()
	if c.versions == nil {
		c.rw.RUnlock()
		var res ocpi.Response[ocpi.Versions]
		if err := c.do(ctx, http.MethodGet, c.versionUrl, nil, &res); err != nil {
			return err
		}
		versions, err := res.Data()
		if err != nil {
			return err
		}

		c.rw.Lock()
		c.versions = versions
		c.rw.Unlock()
		c.rw.RLock()
	}
	selectedVersion, ok := c.versions.LatestMutualVersion(version.Version)
	c.rw.RUnlock()
	if !ok {
		return fmt.Errorf(`ocpi221: missing mutual version for version %q`, version.Version)
	}
	c.rw.Lock()
	c.selectedVersion = selectedVersion
	c.rw.Unlock()
	return nil
}

func (c *ClientConn) GetVersionDetails(ctx context.Context, version ocpi.Version) (*ocpi.Response[VersionDetails], error) {
	var res ocpi.Response[VersionDetails]
	if err := c.do(ctx, http.MethodGet, version.URL, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
