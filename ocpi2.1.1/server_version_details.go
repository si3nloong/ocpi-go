package ocpi211

import (
	"net/http"
	"strings"

	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) GetOcpiVersionDetails(role Role) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimSuffix(r.RequestURI, "/details")
		origin := ocpihttp.GetHostname(r) + path

		versionDetails := VersionDetails{}
		versionDetails.Version = ocpi.VersionNumber211
		switch role {
		case RoleCPO:
			versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDCredentials, URL: origin + "/credentials"})
			versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDCdrs, URL: origin + "/cdrs"})
			versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDLocations, URL: origin + "/locations"})
			versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDSessions, URL: origin + "/sessions"})
			versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDTariffs, URL: origin + "/tariffs"})
			versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDTokens, URL: origin + "/tokens"})
			versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDCommands, URL: origin + "/commands"})

		case RoleEMSP:
			versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDCredentials, URL: origin + "/credentials"})
			versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDCdrs, URL: origin + "/cdrs"})
			versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDLocations, URL: origin + "/locations"})
			versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDSessions, URL: origin + "/sessions"})
			versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDTariffs, URL: origin + "/tariffs"})
			versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDTokens, URL: origin + "/tokens"})
			versionDetails.Endpoints = append(versionDetails.Endpoints, Endpoint{Identifier: ModuleIDCommands, URL: origin + "/commands"})
		}

		ocpihttp.Response(w, versionDetails)
	}
}
