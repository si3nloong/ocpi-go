package ocpi230

import (
	"log/slog"
	"net/http"
	"os"

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
	bookingsSender           BookingsSender
	bookingsReceiver         BookingsReceiver
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
	paymentsSender           PaymentsSender
	paymentsReceiver         PaymentsReceiver
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
	s.bookingsSender = cpo
	s.cdrsSender = cpo
	s.chargingProfilesReceiver = cpo
	s.commandsReceiver = cpo
	s.credentials = cpo
	s.hubClientInfoReceiver = cpo
	s.locationsSender = cpo
	s.sessionsSender = cpo
	s.tariffsSender = cpo
	s.tokensReceiver = cpo
	s.paymentsReceiver = cpo
	s.versions = cpo
}

func (s *Server) SetEMSP(emsp EMSP) {
	s.roles[RoleEMSP] = struct{}{}
	s.bookingsReceiver = emsp
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

func (s *Server) SetHub(hub RoamingHub) {
	s.roles[RoleOther] = struct{}{}
	s.bookingsSender = hub
	s.bookingsReceiver = hub
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
