package ocpi221

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gofrs/uuid/v5"
)

type ServerOption interface {
	apply(*serverOptions)
}

type serverOptions struct {
	baseUrl string
	logger  *slog.Logger
}

type Server struct {
	serverOptions
	roles                    map[Role]struct{}
	cdrsSender               CDRsSender
	cdrsReceiver             CDRsReceiver
	chargingProfilesSender   ChargingProfilesSender
	chargingProfilesReceiver ChargingProfilesReceiver
	credentials              Credentials
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
	versions                 Versions
}

func NewServer(credential Credentials, options ...ServerOption) *Server {
	s := new(Server)
	s.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	// s.baseUrl = fmt.Sprintf("/ocpi/%s", ocpi.VersionNumber221)
	for _, opt := range options {
		opt.apply(&s.serverOptions)
	}
	s.roles = make(map[Role]struct{})
	s.credentials = credential
	return s
}

func (s *Server) SetCPO(cpo CPO) {
	s.roles[RoleCPO] = struct{}{}
	s.cdrsSender = cpo
	s.chargingProfilesReceiver = cpo
	s.commandsReceiver = cpo
	s.credentials = cpo
	s.hubClientInfoReceiver = cpo
	s.locationsSender = cpo
	s.sessionsSender = cpo
	s.tariffsSender = cpo
	s.tokensReceiver = cpo
	s.versions = cpo
}

func (s *Server) SetEMSP(emsp EMSP) {
	s.roles[RoleEMSP] = struct{}{}
	s.cdrsReceiver = emsp
	s.commandsSender = emsp
	s.credentials = emsp
	s.hubClientInfoReceiver = emsp
	s.locationsReceiver = emsp
	s.sessionsReceiver = emsp
	s.tariffsReceiver = emsp
	s.tokensSender = emsp
	s.versions = emsp
}

func (s *Server) SetHub(hub Hub) {
	s.roles[RoleHUB] = struct{}{}
	s.cdrsSender = hub
	s.cdrsReceiver = hub
	s.chargingProfilesSender = hub
	s.chargingProfilesReceiver = hub
	s.commandsSender = hub
	s.commandsReceiver = hub
	s.credentials = hub
	s.hubClientInfoSender = hub
	s.locationsSender = hub
	s.locationsReceiver = hub
	s.sessionsSender = hub
	s.sessionsReceiver = hub
	s.tariffsSender = hub
	s.tariffsReceiver = hub
	s.tokensSender = hub
	s.tokensReceiver = hub
	s.versions = hub
}

func (s *Server) SetNSP(nsp NSP) {
	s.roles[RoleNSP] = struct{}{}
	s.credentials = nsp
	s.hubClientInfoReceiver = nsp
	s.locationsReceiver = nsp
	s.tariffsReceiver = nsp
	s.versions = nsp
}

func (s *Server) SetNAP(nap NAP) {
	s.roles[RoleNAP] = struct{}{}
	s.credentials = nap
	s.hubClientInfoReceiver = nap
	s.locationsSender = nap
	s.locationsReceiver = nap
	s.tariffsSender = nap
	s.tariffsReceiver = nap
	s.versions = nap
}

func (s *Server) SetSCSP(scsp SCSP) {
	s.roles[RoleSCSP] = struct{}{}
	s.chargingProfilesSender = scsp
	s.credentials = scsp
	s.hubClientInfoReceiver = scsp
	s.sessionsReceiver = scsp
	s.versions = scsp
}

func (s *Server) Handler() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc(s.baseUrl+"/details", s.GetOcpiVersionDetails)
	router.HandleFunc(s.baseUrl+"/credentials", s.GetOcpiCredentials)
	router.HandleFunc("POST "+s.baseUrl+"/credentials", s.PostOcpiCredentials)
	router.HandleFunc("PUT "+s.baseUrl+"/credentials", s.PutOcpiCredentials)
	router.HandleFunc("DELETE "+s.baseUrl+"/credentials", s.DeleteOcpiCredentials)

	if s.locationsSender != nil {
		router.HandleFunc(s.baseUrl+"/locations", s.GetOcpiLocations)
		router.HandleFunc(s.baseUrl+"/locations/{location_id}", s.GetOcpiLocation)
		router.HandleFunc(s.baseUrl+"/locations/{location_id}/{evse_uid}", s.GetOcpiLocation)
		router.HandleFunc(s.baseUrl+"/locations/{location_id}/{evse_uid}/{connector_id}", s.GetOcpiLocation)
	}
	if s.locationsReceiver != nil {
		router.HandleFunc(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}", s.GetOcpiClientOwnedLocation)
		router.HandleFunc(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}/{evse_uid}", s.GetOcpiClientOwnedLocation)
		router.HandleFunc(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id}", s.GetOcpiClientOwnedLocation)
		router.HandleFunc("PUT "+s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}", s.PutOcpiLocation)
		router.HandleFunc("PUT "+s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}/{evse_uid}", s.PutOcpiLocation)
		router.HandleFunc("PUT "+s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id}", s.PutOcpiLocation)
		router.HandleFunc("PATCH "+s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}", s.PatchOcpiLocation)
		router.HandleFunc("PATCH "+s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}/{evse_uid}", s.PatchOcpiLocation)
		router.HandleFunc("PATCH "+s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id}", s.PatchOcpiLocation)
	}

	if s.sessionsSender != nil {
		router.HandleFunc(s.baseUrl+"/sessions", s.GetOcpiSessions)
		router.HandleFunc("PUT "+s.baseUrl+"/sessions/{session_id}/charging_preferences", s.PutOcpiSesionChargingPreferences)
	}
	if s.sessionsReceiver != nil {
		router.HandleFunc(s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.GetOcpiSession)
		router.HandleFunc("PUT "+s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.PutOcpiSession)
		router.HandleFunc("PATCH "+s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.PatchOcpiSession)
	}

	if s.cdrsSender != nil {
		router.HandleFunc(s.baseUrl+"/cdrs", s.GetOcpiCDRs)
	}
	if s.cdrsReceiver != nil {
		router.HandleFunc(s.baseUrl+"/cdrs/{cdr_id}", s.GetOcpiCDR)
		router.HandleFunc("POST "+s.baseUrl+"/cdrs", s.PostOcpiCDR)
	}

	if s.tariffsSender != nil {
		router.HandleFunc(s.baseUrl+"/tariffs", s.GetOcpiTariffs)
	}
	if s.tariffsReceiver != nil {
		router.HandleFunc(s.baseUrl+"/tariffs/{country_code}/{party_id}/{tariff_id}", s.GetOcpiTariff)
		router.HandleFunc("PUT "+s.baseUrl+"/tariffs/{country_code}/{party_id}/{tariff_id}", s.PutOcpiTariff)
		router.HandleFunc("DELETE "+s.baseUrl+"/tariffs/{country_code}/{party_id}/{tariff_id}", s.DeleteOcpiTariff)
	}

	if s.tokensSender != nil {
		router.HandleFunc(s.baseUrl+"/tokens", s.GetOcpiTokens)
		router.HandleFunc("POST "+s.baseUrl+"/tokens/{token_uid}/authorize", s.PostOcpiToken)
	}
	if s.tokensReceiver != nil {
		router.HandleFunc(s.baseUrl+"/tokens/{country_code}/{party_id}/{token_uid}", s.GetOcpiToken)
		router.HandleFunc("PUT "+s.baseUrl+"/tokens/{country_code}/{party_id}/{token_uid}", s.PutOcpiToken)
		router.HandleFunc("PATCH "+s.baseUrl+"/tokens/{country_code}/{party_id}/{token_uid}", s.PatchOcpiToken)
	}

	if s.commandsReceiver != nil {
		router.HandleFunc("POST "+s.baseUrl+"/commands/{command_type}", s.PostOcpiCommand)
	}
	if s.commandsSender != nil {
		router.HandleFunc("POST "+s.baseUrl+"/commands/{command_type}/{uid}", s.PostOcpiCommandResponse)
	}
	return router
}

func writeOkResponse(w http.ResponseWriter, r *http.Request, b []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set(HttpHeaderXRequestID, uuid.Must(uuid.NewV7()).String())
	w.Header().Set(HttpHeaderXCorrelationID, r.Header.Get(HttpHeaderXCorrelationID))
	w.Write(b)
}
