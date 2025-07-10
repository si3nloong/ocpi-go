package ocpi221

import (
	"context"
	"fmt"
	"sync"

	"github.com/si3nloong/ocpi-go/ocpi"
)

type ocpiClient struct {
	conn      *ClientConn
	rw        sync.RWMutex
	tokenA    string
	endpoints map[string]string
}

func (c *ocpiClient) GetCredentialsToken(ctx context.Context) (string, error) {
	return c.tokenA, nil
}

func (c *ocpiClient) GetEndpoint(ctx context.Context, module ModuleID, role InterfaceRole) (string, error) {
	c.rw.RLock()
	if c.endpoints == nil {
		c.rw.RUnlock()
		versionsResponse, err := c.conn.GetVersions(ctx)
		if err != nil {
			return "", err
		}
		versions, err := versionsResponse.Data()
		if err != nil {
			return "", err
		}
		mutualVersion, ok := versions.LatestMutualVersion(ocpi.VersionNumber221)
		if !ok {
			return "", fmt.Errorf(`ocpi221: cannot find mutual version for %s client`, ocpi.VersionNumber221)
		}
		versionDetailsResponse, err := c.conn.GetVersionDetails(ctx, mutualVersion)
		if err != nil {
			return "", err
		}
		versionDetails, err := versionDetailsResponse.Data()
		if err != nil {
			return "", err
		}
		c.rw.Lock()
		c.endpoints = make(map[string]string)
		for _, endpoint := range versionDetails.Endpoints {
			if endpoint.Identifier == ModuleIDCredentials {
				c.endpoints[string(endpoint.Identifier)] = endpoint.URL
			} else {
				c.endpoints[string(endpoint.Identifier)+":"+string(endpoint.Role)] = endpoint.URL
			}
		}
		c.rw.Unlock()
		c.rw.RLock()
	}
	defer c.rw.RUnlock()
	key := string(module) + ":" + string(role)
	if module == ModuleIDCredentials {
		key = string(module)
	}
	if endpoint, ok := c.endpoints[key]; ok {
		return endpoint, nil
	}
	return "", fmt.Errorf(`ocpi221: missing endpoint for module %q and role %q`, module, role)
}
