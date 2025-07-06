package ocpi221

import (
	"context"
	"fmt"
	"net/http"
	"sort"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) RegisterCredentials(ctx context.Context, req Credential) (*ocpi.Response[Credential], error) {
	var versionResponse ocpi.Response[ocpi.Versions]
	if err := c.do(ctx, http.MethodGet, c.versionUrl, nil, &versionResponse); err != nil {
		return nil, err
	}
	var res ocpi.Response[Credential]
	if err := c.CallEndpoint(ctx, ModuleIDCredentials, InterfaceRoleSender, http.MethodPut, func(endpoint string) string {
		return endpoint
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) UpdateCredentials(ctx context.Context, credential Credential, storeCallback func(VersionDetails) error) (*ocpi.Response[Credential], error) {
	var versionResponse ocpi.Response[ocpi.Versions]
	if err := c.do(ctx, http.MethodGet, c.versionUrl, nil, &versionResponse); err != nil {
		return nil, err
	}
	versions, err := versionResponse.Data()
	if err != nil {
		return nil, err
	}
	if len(versions) == 0 {
		return nil, fmt.Errorf(`ocpi221: empty versions`)
	}
	sort.Sort(versions)
	version := versions[0]

	var versionDetailsResponse ocpi.Response[VersionDetails]
	if err := c.do(ctx, http.MethodGet, version.URL, nil, &versionDetailsResponse); err != nil {
		return nil, err
	}
	versionDetails, err := versionDetailsResponse.Data()
	if err != nil {
		return nil, err
	}
	if err := storeCallback(versionDetails); err != nil {
		return nil, err
	}
	return c.PutCredentials(ctx, credential)
}

func (c *ClientConn) GetCredentials(ctx context.Context) (*ocpi.Response[Credential], error) {
	var res ocpi.Response[Credential]
	if err := c.CallEndpoint(ctx, ModuleIDCredentials, InterfaceRoleSender, http.MethodGet, func(endpoint string) string {
		return endpoint
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PutCredentials(ctx context.Context, req Credential) (*ocpi.Response[Credential], error) {
	var res ocpi.Response[Credential]
	if err := c.CallEndpoint(ctx, ModuleIDCredentials, InterfaceRoleSender, http.MethodPut, func(endpoint string) string {
		return endpoint
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
