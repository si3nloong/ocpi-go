package ocpi221

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

const (
	HttpHeaderXRequestID     = "X-Request-ID"
	HttpHeaderXCorrelationID = "X-Correlation-ID"
)

func (s *Server) GetOcpiVersionDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := strings.TrimSuffix(r.RequestURI, "/details")
	origin := httputil.GetHostname(r) + s.baseUrl + path
	endpoints := []Endpoint{
		{Identifier: ModuleIDCredentials, Role: InterfaceRoleSender, URL: origin + "/credentials"},
		{Identifier: ModuleIDCredentials, Role: InterfaceRoleReceiver, URL: origin + "/credentials"},
	}
	if s.cdrsSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDCdrs, Role: InterfaceRoleSender, URL: origin + "/cdrs"})
	}
	if s.cdrsReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDCdrs, Role: InterfaceRoleReceiver, URL: origin + "/cdrs"})
	}
	if s.chargingProfilesSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDChargingProfiles, Role: InterfaceRoleSender, URL: origin + "/chargingprofiles"})
	}
	if s.chargingProfilesReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDChargingProfiles, Role: InterfaceRoleReceiver, URL: origin + "/chargingprofiles"})
	}
	if s.commandsSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDCommands, Role: InterfaceRoleSender, URL: origin + "/commands"})
	}
	if s.commandsReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDCommands, Role: InterfaceRoleReceiver, URL: origin + "/commands"})
	}
	if s.hubClientInfoSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDHubClientInfo, Role: InterfaceRoleSender, URL: origin + "/clientinfo"})
	}
	if s.hubClientInfoReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDHubClientInfo, Role: InterfaceRoleReceiver, URL: origin + "/clientinfo"})
	}
	if s.locationsSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDLocations, Role: InterfaceRoleSender, URL: origin + "/locations"})
	}
	if s.locationsReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDLocations, Role: InterfaceRoleReceiver, URL: origin + "/locations"})
	}
	if s.sessionsSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDSessions, Role: InterfaceRoleSender, URL: origin + "/sessions"})
	}
	if s.sessionsReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDSessions, Role: InterfaceRoleReceiver, URL: origin + "/sessions"})
	}
	if s.tariffsSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDTariffs, Role: InterfaceRoleSender, URL: origin + "/tariffs"})
	}
	if s.tariffsReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDTariffs, Role: InterfaceRoleReceiver, URL: origin + "/tariffs"})
	}
	if s.tokensSender != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDTokens, Role: InterfaceRoleSender, URL: origin + "/tokens"})
	}
	if s.tokensReceiver != nil {
		endpoints = append(endpoints, Endpoint{Identifier: ModuleIDTokens, Role: InterfaceRoleReceiver, URL: origin + "/tokens"})
	}

	b, err := json.Marshal(ocpi.NewResponse(endpoints))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	writeOkResponse(w, r, b)
}

func (c *ClientConn) Versions(ctx context.Context) ([]ocpi.Version, error) {
	var o ocpi.Response[[]ocpi.Version]
	if err := c.do(ctx, http.MethodGet, c.versionUrl, nil, &o); err != nil {
		return nil, err
	}
	return o.Data, nil
}
