package ocpi221

import (
	"net/http"
	"strings"

	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) GetOcpiVersionDetails(w http.ResponseWriter, r *http.Request) {
	if recv, ok := s.ocpi.(VersionsReceiver); ok {
		ctx := r.Context()
		endpoints, err := recv.OnVersionDetails(ctx, ocpi.GetRequestContext(ctx).Token)
		if err != nil {
			ocpihttp.Response(w, err)
			return
		}

		ocpihttp.Response(w, VersionDetails{
			Version:   ocpi.VersionNumber221,
			Endpoints: endpoints,
		})
		return
	}

	path := strings.TrimSuffix(r.RequestURI, "/details")
	origin := ocpihttp.GetHostname(r) + path
	versionDetails := VersionDetails{}
	versionDetails.Version = ocpi.VersionNumber221
	versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDCredentials, Role: InterfaceRoleSender, URL: origin + "/credentials"})

	if s.cdrsSender != nil {
		versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDCdrs, Role: InterfaceRoleSender, URL: origin + "/cdrs"})
	}
	if s.cdrsReceiver != nil {
		versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDCdrs, Role: InterfaceRoleReceiver, URL: origin + "/cdrs"})
	}
	if s.chargingProfilesReceiver != nil {
		versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDChargingProfiles, Role: InterfaceRoleReceiver, URL: origin + "/chargingprofiles"})
	}
	if s.commandsReceiver != nil {
		versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDCommands, Role: InterfaceRoleReceiver, URL: origin + "/commands"})
	}
	if s.hubClientInfoSender != nil {
		versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDHubClientInfo, Role: InterfaceRoleSender, URL: origin + "/hubclientinfo"})
	}
	if s.hubClientInfoReceiver != nil {
		versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDHubClientInfo, Role: InterfaceRoleReceiver, URL: origin + "/clientinfo"})
	}
	if s.locationsSender != nil {
		versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDLocations, Role: InterfaceRoleSender, URL: origin + "/locations"})
	}
	if s.locationsReceiver != nil {
		versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDLocations, Role: InterfaceRoleReceiver, URL: origin + "/locations"})
	}
	if s.sessionsSender != nil {
		versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDSessions, Role: InterfaceRoleSender, URL: origin + "/sessions"})
	}
	if s.sessionsReceiver != nil {
		versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDSessions, Role: InterfaceRoleReceiver, URL: origin + "/sessions"})
	}
	if s.tariffsSender != nil {
		versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDTariffs, Role: InterfaceRoleSender, URL: origin + "/tariffs"})
	}
	if s.tariffsReceiver != nil {
		versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDTariffs, Role: InterfaceRoleReceiver, URL: origin + "/tariffs"})
	}
	if s.tokensSender != nil {
		versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDTokens, Role: InterfaceRoleSender, URL: origin + "/tokens"})
	}
	if s.tokensReceiver != nil {
		versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDTokens, Role: InterfaceRoleReceiver, URL: origin + "/tokens"})
	}

	ocpihttp.Response(w, versionDetails)
}
