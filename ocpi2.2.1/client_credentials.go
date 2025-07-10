package ocpi221

import (
	"context"
	"fmt"
	"net/http"
	"sort"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) GetCredential(ctx context.Context) (*ocpi.Response[Credential], error) {
	var res ocpi.Response[Credential]
	if err := c.CallEndpoint(ctx, ModuleIDCredentials, InterfaceRoleSender, http.MethodGet, func(endpoint string) string {
		return endpoint
	}, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PostCredential(ctx context.Context, req Credential) (*ocpi.Response[Credential], error) {
	var res ocpi.Response[Credential]
	if err := c.CallEndpoint(ctx, ModuleIDCredentials, InterfaceRoleSender, http.MethodPut, func(endpoint string) string {
		return endpoint
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) PutCredential(ctx context.Context, req Credential) (*ocpi.Response[Credential], error) {
	var res ocpi.Response[Credential]
	if err := c.CallEndpoint(ctx, ModuleIDCredentials, InterfaceRoleSender, http.MethodPut, func(endpoint string) string {
		return endpoint
	}, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ClientConn) RegisterCredential(ctx context.Context, req Credential) (*ocpi.Response[Credential], error) {
	var versionsResponse ocpi.Response[ocpi.Versions]
	if err := c.do(ctx, http.MethodGet, c.versionUrl, nil, &versionsResponse); err != nil {
		return nil, err
	}
	versions, err := versionsResponse.Data()
	if err != nil {
		return nil, err
	}
	versions.LatestMutualVersion(ocpi.VersionNumber221)

	var res ocpi.Response[Credential]
	if err := c.CallEndpoint(ctx, ModuleIDCredentials, InterfaceRoleSender, http.MethodPut, func(endpoint string) string {
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
func (c *ClientConn) UpdateCredential(ctx context.Context, credentialWithTokenB Credential, storeCallback func(VersionDetails) error) (*ocpi.Response[Credential], error) {
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
	version, ok := versions.LatestMutualVersion(ocpi.VersionNumber221)
	if !ok {
		return nil, fmt.Errorf(`ocpi221: cannot find mutual ocpi version`)
	}

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
	return c.PutCredential(ctx, credentialWithTokenB)
}
