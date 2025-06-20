package ocpi221

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/si3nloong/ocpi-go/ocpi"
)

type Config struct {
	BaseURL string
}

type Server struct {
	baseUrl                  string
	logger                   *slog.Logger
	httpHandler              http.Handler
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

func NewServer(credential Credentials, cfg Config) *Server {
	s := new(Server)
	s.baseUrl = "/ocpi/" + string(ocpi.VersionNumber221)
	if cfg.BaseURL != "" {
		s.baseUrl = cfg.BaseURL
	}
	s.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
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
	router := chi.NewRouter()

	s.baseUrl = ""
	router.Get(s.baseUrl+"/details", s.GetOcpiVersionDetails)
	router.Get(s.baseUrl+"/credentials", s.GetOcpiCredentials)
	router.Post(s.baseUrl+"/credentials", s.PostOcpiCredentials)
	router.Put(s.baseUrl+"/credentials", s.PutOcpiCredentials)
	router.Delete(s.baseUrl+"/credentials", s.DeleteOcpiCredentials)

	if s.locationsSender != nil {
		router.Get(s.baseUrl+"/locations", s.GetOcpiLocations)
		router.Get(s.baseUrl+"/locations/{location_id}(/{evse_uid}(/{connector_id}))", s.GetOcpiLocation)
	}
	if s.locationsReceiver != nil {
		router.Get(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}(/{evse_uid}(/{connector_id}))", s.GetOcpiClientOwnedLocation)
		router.Put(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}(/{evse_uid}(/{connector_id}))", s.PutOcpiLocation)
		router.Patch(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}(/{evse_uid}(/{connector_id}))", s.PatchOcpiLocation)
	}

	if s.sessionsSender != nil {
		router.Get(s.baseUrl+"/sessions", s.GetOcpiSessions)
		router.Put(s.baseUrl+"/sessions/{session_id}/charging_preferences", s.PutOcpiSesionChargingPreferences)
	}
	if s.sessionsReceiver != nil {
		router.Get(s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.GetOcpiSession)
		router.Put(s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.PutOcpiSession)
		router.Patch(s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.PatchOcpiSession)
	}

	if s.cdrsSender != nil {
		router.Get(s.baseUrl+"/tariffs", s.GetOcpiCDRs)
	}
	if s.cdrsReceiver != nil {
		router.Get(s.baseUrl+"/cdrs/{cdr_id}", s.GetOcpiCDR)
		router.Post(s.baseUrl+"/cdrs", s.PostOcpiCDR)
	}

	if s.tariffsSender != nil {
		router.Get(s.baseUrl+"/tariffs", s.GetOcpiTariffs)
	}
	if s.tariffsReceiver != nil {
		router.Get(s.baseUrl+"/tariffs/{country_code}/{party_id}/{tariff_id}", s.GetOcpiTariff)
		router.Put(s.baseUrl+"/tariffs/{country_code}/{party_id}/{tariff_id}", s.PutOcpiTariff)
		router.Delete(s.baseUrl+"/tariffs/{country_code}/{party_id}/{tariff_id}", s.DeleteOcpiTariff)
	}

	if s.tokensSender != nil {
		router.Get(s.baseUrl+"/tokens", s.GetOcpiTokens)
		router.Post(s.baseUrl+"/tokens/{token_uid}/authorize", s.PostOcpiToken)
	}
	if s.tokensReceiver != nil {
		router.Get(s.baseUrl+"/tokens/{country_code}/{party_id}/{token_uid}", s.GetOcpiToken)
		router.Put(s.baseUrl+"/tokens/{country_code}/{party_id}/{token_uid}", s.PutOcpiToken)
		router.Patch(s.baseUrl+"/tokens/{country_code}/{party_id}/{token_uid}", s.PatchOcpiToken)
	}

	if s.commandsReceiver != nil {
		router.Post(s.baseUrl+"/commands/{command_type}", s.PostOcpiCommand)
	}
	if s.commandsSender != nil {
		router.Post(s.baseUrl+"/commands/{command_type}/{uid}", s.PostOcpiCommandResponse)
	}

	s.baseUrl = fmt.Sprintf("/ocpi/%s", ocpi.VersionNumber221)
	return router
}
