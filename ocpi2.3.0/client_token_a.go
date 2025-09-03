package ocpi230

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

func (c *unregisteredClient) GetEndpoint(ctx context.Context, module ModuleID, role InterfaceRole) (string, error) {
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
	return "", fmt.Errorf(`ocpi230: missing endpoint for module %q and role %q`, module, role)
}
