package ocpi211

import (
	"context"
	"fmt"
)

type unregisteredClient struct {
	*ClientConn
	tokenA    string
	endpoints map[string]string
}

func (c *unregisteredClient) GetCredentialsToken(ctx context.Context) (string, error) {
	return c.tokenA, nil
}

func (c *unregisteredClient) GetEndpoint(ctx context.Context, module ModuleID) (string, error) {
	c.rw.RLock()
	if c.endpoints == nil {
		c.rw.RUnlock()
		if c.versionDetails == nil {
			versionDetailsResponse, err := c.GetVersionDetails(ctx)
			if err != nil {
				return "", err
			}
			if _, err := versionDetailsResponse.Data(); err != nil {
				return "", err
			}
		}
		c.rw.Lock()
		c.endpoints = make(map[string]string)
		for _, endpoint := range c.versionDetails.Endpoints {
			c.endpoints[string(endpoint.Identifier)] = endpoint.URL
		}
		c.rw.Unlock()
		c.rw.RLock()
	}
	defer c.rw.RUnlock()
	if endpoint, ok := c.endpoints[string(module)]; ok {
		return endpoint, nil
	}
	return "", fmt.Errorf(`ocpi211: missing endpoint for module %q`, module)
}
