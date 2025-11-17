package ocpi221

import (
	"context"
	"fmt"
	"net/http"
	"sort"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) GetCredential(ctx context.Context) (*ocpi.Response[DateTime, Credentials], error) {
	var res ocpi.Response[DateTime, Credentials]
	if err := c.CallEndpoint(ctx, ModuleIDCredentials, InterfaceRoleSender, http.MethodGet, func(endpoint string) string {
		return endpoint
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PostCredential(ctx context.Context, req Credentials) (*ocpi.Response[DateTime, Credentials], error) {
	var res ocpi.Response[DateTime, Credentials]
	if err := c.CallEndpoint(ctx, ModuleIDCredentials, InterfaceRoleSender, http.MethodPost, func(endpoint string) string {
		return endpoint
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PutCredential(ctx context.Context, req Credentials) (*ocpi.Response[DateTime, Credentials], error) {
	var res ocpi.Response[DateTime, Credentials]
	if err := c.CallEndpoint(ctx, ModuleIDCredentials, InterfaceRoleSender, http.MethodPut, func(endpoint string) string {
		return endpoint
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) RegisterCredential(ctx context.Context, req Credentials) (*ocpi.Response[DateTime, Credentials], error) {
	var versionsResponse ocpi.Response[DateTime, ocpi.Versions]
	if err := c.do(ctx, http.MethodGet, c.versionUrl, nil, &versionsResponse); err != nil {
		return nil, err
	}
	versions, err := versionsResponse.Data()
	if err != nil {
		return nil, err
	}
	mutualVersion, ok := versions.MutualVersion(ocpi.VersionNumber221)
	if !ok {
		return nil, ocpi.ErrNoMutualVersion
	}

	var versionDetailsResponse ocpi.Response[DateTime, ocpi.Versions]
	if err := c.do(ctx, http.MethodGet, mutualVersion.URL, nil, &versionDetailsResponse); err != nil {
		return nil, err
	}

	var res ocpi.Response[DateTime, Credentials]
	if err := c.CallEndpoint(ctx, ModuleIDCredentials, InterfaceRoleSender, http.MethodPost, func(endpoint string) string {
		return endpoint
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// UpdateCredentials can be use for
//   - Updating to a newer version
//   - Changing endpoints for the current version
//   - Updating the credentials and resetting the credentials token
func (c *ClientConn) UpdateCredential(ctx context.Context, credentialWithTokenB Credentials, storeCallback func(VersionDetails) error) (*ocpi.Response[DateTime, Credentials], error) {
	var versionResponse ocpi.Response[DateTime, ocpi.Versions]
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
	version, ok := versions.MutualVersion(ocpi.VersionNumber221)
	if !ok {
		return nil, ocpi.ErrNoMutualVersion
	}

	var versionDetailsResponse ocpi.Response[DateTime, VersionDetails]
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
	return c.PutCredential(ctx, credentialWithTokenB)
}
