package ocpi221

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

func (s *Server) GetOcpiVersionDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	endpoints := []Endpoint{
		{Identifier: ModuleIDCredentials, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/credentials"},
		{Identifier: ModuleIDCredentials, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/credentials"},
	}
	if s.cdrsSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDCdrs, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/cdrs"})
	}
	if s.cdrsReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDCdrs, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/cdrs"})
	}
	if s.chargingProfilesSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDChargingProfiles, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/chargingprofiles"})
	}
	if s.chargingProfilesReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDChargingProfiles, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/chargingprofiles"})
	}
	if s.commandsSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDCommands, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/commands"})
	}
	if s.commandsReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDCommands, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/commands"})
	}
	if s.hubClientInfoSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDHubClientInfo, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/clientinfo"})
	}
	if s.hubClientInfoReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDHubClientInfo, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/clientinfo"})
	}
	if s.locationsSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDLocations, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/locations"})
	}
	if s.locationsReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDLocations, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/locations"})
	}
	if s.sessionsSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDSessions, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/sessions"})
	}
	if s.sessionsReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDSessions, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/sessions"})
	}
	if s.tariffsSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDTariffs, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/tariffs"})
	}
	if s.tariffsReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDTariffs, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/tariffs"})
	}
	if s.tokensSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDTokens, Role: InterfaceRoleSender, URL: getHostname(r) + s.baseUrl + "/tokens"})
	}
	if s.tokensReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDTokens, Role: InterfaceRoleReceiver, URL: getHostname(r) + s.baseUrl + "/tokens"})
	}

	b, err := json.Marshal(ocpi.NewResponse(endpoints))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (c *Client) Versions(ctx context.Context) ([]ocpi.Version, error) {
	var o ocpi.Response[[]ocpi.Version]
	if err := c.do(ctx, http.MethodGet, c.versionUrl, nil, &o); err != nil {
		return nil, err
	}
	return o.Data, nil
}
