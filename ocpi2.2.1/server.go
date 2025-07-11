package ocpi221

import (
	"context"
	"encoding/base64"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

type TokenResolver func(token string) string

type OCPIServer interface {
	IsClientRegistered(ctx context.Context, tokenA string) bool
	VerifyCredentialsToken(ctx context.Context, token string) error
	StoreCredentialsTokenB(ctx context.Context, credentialsTokenB Credentials) error
	StoreVersionDetails(ctx context.Context, versionDetails VersionDetails) error
	GenerateCredentialsTokenC(ctx context.Context, tokenA string) (*Credentials, error)
	CredentialsReceiver
	// Versions
}

type ServerConfig struct {
	OmitRole      bool
	TokenResolver TokenResolver
}

type Server struct {
	logger                   *slog.Logger
	enabledRoleRoute         bool
	errs                     chan error
	ocpi                     OCPIServer
	tokenResolver            TokenResolver
	roles                    map[Role]struct{}
	cdrsSender               CDRsSender
	cdrsReceiver             CDRsReceiver
	chargingProfilesSender   ChargingProfilesSender
	chargingProfilesReceiver ChargingProfilesReceiver
	commandsSender           CommandsSender
	commandsReceiver         CommandsReceiver
	hubClientInfoSender      HubClientInfoSender
	hubClientInfoReceiver    HubClientInfoReceiver
	locationsSender          LocationsSender
	locationsReceiver        LocationsReceiver
	sessionsSender           SessionsSender
	sessionsReceiver         SessionsReceiver
	tariffsSender            TariffsSender
	tariffsReceiver          TariffsReceiver
	tokensSender             TokensSender
	tokensReceiver           TokensReceiver
}

func NewServer(ocpi OCPIServer, cfg ...ServerConfig) *Server {
	s := new(Server)
	s.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	s.tokenResolver = func(token string) string {
		token = strings.TrimPrefix(token, "Token ")
		b, err := base64.StdEncoding.DecodeString(token)
		if err == nil {
			return string(b)
		}
		return token
	}
	s.roles = make(map[Role]struct{})
	s.ocpi = ocpi
	s.errs = make(chan error, 1)
	return s
}

func (s *Server) SetCPO(cpo CPO) {
	s.roles[RoleCPO] = struct{}{}
	s.cdrsSender = cpo
	s.chargingProfilesReceiver = cpo
	s.commandsReceiver = cpo
	s.hubClientInfoReceiver = cpo
	s.locationsSender = cpo
	s.sessionsSender = cpo
	s.tariffsSender = cpo
	s.tokensReceiver = cpo
}

func (s *Server) SetEMSP(emsp EMSP) {
	s.roles[RoleEMSP] = struct{}{}
	s.cdrsReceiver = emsp
	s.commandsSender = emsp
	s.hubClientInfoReceiver = emsp
	s.locationsReceiver = emsp
	s.sessionsReceiver = emsp
	s.tariffsReceiver = emsp
	s.tokensSender = emsp
}

func (s *Server) SetHub(hub Hub) {
	s.roles[RoleHUB] = struct{}{}
	s.cdrsSender = hub
	s.cdrsReceiver = hub
	s.chargingProfilesSender = hub
	s.chargingProfilesReceiver = hub
	s.commandsSender = hub
	s.commandsReceiver = hub
	s.hubClientInfoSender = hub
	s.locationsSender = hub
	s.locationsReceiver = hub
	s.sessionsSender = hub
	s.sessionsReceiver = hub
	s.tariffsSender = hub
	s.tariffsReceiver = hub
	s.tokensSender = hub
	s.tokensReceiver = hub
}

func (s *Server) SetNSP(nsp NSP) {
	s.roles[RoleNSP] = struct{}{}
	s.hubClientInfoReceiver = nsp
	s.locationsReceiver = nsp
	s.tariffsReceiver = nsp
}

func (s *Server) SetNAP(nap NAP) {
	s.roles[RoleNAP] = struct{}{}
	s.hubClientInfoReceiver = nap
	s.locationsSender = nap
	s.locationsReceiver = nap
	s.tariffsSender = nap
	s.tariffsReceiver = nap
}

func (s *Server) SetSCSP(scsp SCSP) {
	s.roles[RoleSCSP] = struct{}{}
	s.chargingProfilesSender = scsp
	s.hubClientInfoReceiver = scsp
	s.sessionsReceiver = scsp
}

func (s *Server) Handler() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/2.2.1/details", s.GetOcpiVersionDetails)
	router.HandleFunc("/2.2.1/credentials", s.GetOcpiCredentials)
	router.HandleFunc("POST /2.2.1/credentials", s.PostOcpiCredentials)
	router.HandleFunc("PUT /2.2.1/credentials", s.PutOcpiCredentials)
	router.HandleFunc("DELETE /2.2.1/credentials", s.DeleteOcpiCredentials)

	if s.cdrsSender != nil {
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.2.1/cdrs"), s.GetOcpiCDRs)
	}
	if s.cdrsReceiver != nil {
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.2.1/cdrs/{cdr_id}"), s.GetOcpiCDR)
		router.HandleFunc("POST "+s.withRole(InterfaceRoleReceiver, "/2.2.1/cdrs"), s.PostOcpiCDR)
	}

	if s.chargingProfilesReceiver != nil {
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.2.1/chargingprofiles/{session_id}"), s.GetOcpiActiveChargingProfile)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleReceiver, "/2.2.1/chargingprofiles/{session_id}"), s.PutOcpiChargingProfile)
		router.HandleFunc("DELETE "+s.withRole(InterfaceRoleReceiver, "/2.2.1/chargingprofiles/{session_id}"), s.DeleteOcpiChargingProfile)
	}
	if s.chargingProfilesSender != nil {
		router.HandleFunc("POST "+s.withRole(InterfaceRoleSender, "/2.2.1/activechargingprofile/{session_id}"), s.PostOcpiActiveChargingProfile)
		router.HandleFunc("POST "+s.withRole(InterfaceRoleSender, "/2.2.1/chargingprofiles/chargingprofile/{session_id}"), s.PostOcpiChargingProfile)
		router.HandleFunc("POST "+s.withRole(InterfaceRoleSender, "/2.2.1/clearprofile/{session_id}"), s.PostOcpiClearProfile)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleSender, "/2.2.1/activechargingprofile/{session_id}"), s.PutOcpiActiveChargingProfile)
	}

	if s.commandsReceiver != nil {
		router.HandleFunc("POST "+s.withRole(InterfaceRoleReceiver, "/2.2.1/commands/{command_type}"), s.PostOcpiCommand)
	}
	if s.commandsSender != nil {
		router.HandleFunc("POST "+s.withRole(InterfaceRoleSender, "/2.2.1/commands/{command_type}/{uid}"), s.PostOcpiCommandResponse)
	}

	if s.hubClientInfoSender != nil {
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.2.1/hubclientinfo"), s.GetOcpiClientInfos)
	}
	if s.hubClientInfoReceiver != nil {
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.2.1/clientinfo/{country_code}/{party_id}"), s.GetOcpiClientInfo)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleReceiver, "/2.2.1/clientinfo/{country_code}/{party_id}"), s.PutOcpiClientInfo)
	}

	if s.locationsSender != nil {
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.2.1/locations"), s.GetOcpiLocations)
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.2.1/locations/{location_id}"), s.GetOcpiLocation)
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.2.1/locations/{location_id}/{evse_uid}"), s.GetOcpiLocation)
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.2.1/locations/{location_id}/{evse_uid}/{connector_id}"), s.GetOcpiLocation)
	}
	if s.locationsReceiver != nil {
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.2.1/locations/{country_code}/{party_id}/{location_id}"), s.GetOcpiClientOwnedLocation)
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}"), s.GetOcpiClientOwnedLocation)
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id}"), s.GetOcpiClientOwnedLocation)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleReceiver, "/2.2.1/locations/{country_code}/{party_id}/{location_id}"), s.PutOcpiLocation)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleReceiver, "/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}"), s.PutOcpiLocation)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleReceiver, "/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id}"), s.PutOcpiLocation)
		router.HandleFunc("PATCH "+s.withRole(InterfaceRoleReceiver, "/2.2.1/locations/{country_code}/{party_id}/{location_id}"), s.PatchOcpiLocation)
		router.HandleFunc("PATCH "+s.withRole(InterfaceRoleReceiver, "/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}"), s.PatchOcpiLocation)
		router.HandleFunc("PATCH "+s.withRole(InterfaceRoleReceiver, "/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id}"), s.PatchOcpiLocation)
	}

	if s.sessionsSender != nil {
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.2.1/sessions"), s.GetOcpiSessions)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleSender, "/2.2.1/sessions/{session_id}/charging_preferences"), s.PutOcpiSesionChargingPreferences)
	}
	if s.sessionsReceiver != nil {
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.2.1/sessions/{country_code}/{party_id}/{session_id}"), s.GetOcpiSession)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleReceiver, "/2.2.1/sessions/{country_code}/{party_id}/{session_id}"), s.PutOcpiSession)
		router.HandleFunc("PATCH "+s.withRole(InterfaceRoleReceiver, "/2.2.1/sessions/{country_code}/{party_id}/{session_id}"), s.PatchOcpiSession)
	}

	if s.tariffsSender != nil {
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.2.1/tariffs"), s.GetOcpiTariffs)
	}
	if s.tariffsReceiver != nil {
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.2.1/tariffs/{country_code}/{party_id}/{tariff_id}"), s.GetOcpiTariff)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleReceiver, "/2.2.1/tariffs/{country_code}/{party_id}/{tariff_id}"), s.PutOcpiTariff)
		router.HandleFunc("DELETE "+s.withRole(InterfaceRoleReceiver, "/2.2.1/tariffs/{country_code}/{party_id}/{tariff_id}"), s.DeleteOcpiTariff)
	}

	if s.tokensSender != nil {
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.2.1/tokens"), s.GetOcpiTokens)
		router.HandleFunc("POST "+s.withRole(InterfaceRoleSender, "/2.2.1/tokens/{token_uid}/authorize"), s.PostOcpiToken)
	}
	if s.tokensReceiver != nil {
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.2.1/tokens/{country_code}/{party_id}/{token_uid}"), s.GetOcpiToken)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleReceiver, "/2.2.1/tokens/{country_code}/{party_id}/{token_uid}"), s.PutOcpiToken)
		router.HandleFunc("PATCH "+s.withRole(InterfaceRoleReceiver, "/2.2.1/tokens/{country_code}/{party_id}/{token_uid}"), s.PatchOcpiToken)
	}

	return s.authorizeMiddleware(router)
}

func (s *Server) LogError(err error) {
	s.errs <- err
}

func (s *Server) Errors() <-chan error {
	return s.errs
}

func (s *Server) withRole(role InterfaceRole, path string) string {
	if !s.enabledRoleRoute {
		return path
	}
	switch role {
	case InterfaceRoleSender:
		return "/emsp" + path
	case InterfaceRoleReceiver:
		return "/cpo" + path
	default:
		panic(fmt.Sprintf("ocpi221: invalid role %q", role))
	}
}
