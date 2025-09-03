package ocpi230

import (
	"context"
	"encoding/base64"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/si3nloong/ocpi-go/ocpi"
)

type OCPIServer interface {
	IsClientRegistered(ctx context.Context, tokenA string) bool
	VerifyCredentialsToken(ctx context.Context, token string) error
	StoreCredentialsTokenB(ctx context.Context, credentialsTokenB Credentials) error
	StoreVersionDetails(ctx context.Context, versionDetails VersionDetails) error
	GenerateCredentialsTokenC(ctx context.Context, tokenA string) (*Credentials, error)
	CredentialsReceiver
	// Versions
}

var defaultServerOptions = ServerOptions{}

type ServerOptions struct {
	EnabledRole bool
}

type Server struct {
	logger                   *slog.Logger
	enabledRole              bool
	errs                     chan error
	ocpi                     OCPIServer
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
	paymentsSender           PaymentsSender
	paymentsReceiver         PaymentsReceiver
}

func NewServer(ocpi OCPIServer, opts *ServerOptions) *Server {
	if opts == nil {
		opts = &defaultServerOptions
	}
	s := new(Server)
	s.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	s.enabledRole = opts.EnabledRole
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

	router.HandleFunc("/2.3.0/details", s.GetOcpiVersionDetails)
	router.HandleFunc("/2.3.0/credentials", s.GetOcpiCredentials)
	router.HandleFunc("POST /2.3.0/credentials", s.PostOcpiCredentials)
	router.HandleFunc("PUT /2.3.0/credentials", s.PutOcpiCredentials)
	router.HandleFunc("DELETE /2.3.0/credentials", s.DeleteOcpiCredentials)

	if s.cdrsSender != nil {
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.3.0/cdrs"), s.GetOcpiCDRs)
	}
	if s.cdrsReceiver != nil {
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.3.0/cdrs/{cdr_id}"), s.GetOcpiCDR)
		router.HandleFunc("POST "+s.withRole(InterfaceRoleReceiver, "/2.3.0/cdrs"), s.PostOcpiCDR)
	}

	if s.chargingProfilesReceiver != nil {
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.3.0/chargingprofiles/{session_id}"), s.GetOcpiActiveChargingProfile)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleReceiver, "/2.3.0/chargingprofiles/{session_id}"), s.PutOcpiChargingProfile)
		router.HandleFunc("DELETE "+s.withRole(InterfaceRoleReceiver, "/2.3.0/chargingprofiles/{session_id}"), s.DeleteOcpiChargingProfile)
	}
	if s.chargingProfilesSender != nil {
		router.HandleFunc("POST "+s.withRole(InterfaceRoleSender, "/2.3.0/activechargingprofile/{session_id}"), s.PostOcpiActiveChargingProfile)
		router.HandleFunc("POST "+s.withRole(InterfaceRoleSender, "/2.3.0/chargingprofiles/chargingprofile/{session_id}"), s.PostOcpiChargingProfile)
		router.HandleFunc("POST "+s.withRole(InterfaceRoleSender, "/2.3.0/clearprofile/{session_id}"), s.PostOcpiClearProfile)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleSender, "/2.3.0/activechargingprofile/{session_id}"), s.PutOcpiActiveChargingProfile)
	}

	if s.commandsReceiver != nil {
		router.HandleFunc("POST "+s.withRole(InterfaceRoleReceiver, "/2.3.0/commands/{command_type}"), s.PostOcpiCommand)
	}
	if s.commandsSender != nil {
		router.HandleFunc("POST "+s.withRole(InterfaceRoleSender, "/2.3.0/commands/{command_type}/{uid}"), s.PostOcpiCommandResponse)
	}

	if s.hubClientInfoSender != nil {
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.3.0/hubclientinfo"), s.GetOcpiClientInfos)
	}
	if s.hubClientInfoReceiver != nil {
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.3.0/clientinfo/{country_code}/{party_id}"), s.GetOcpiClientInfo)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleReceiver, "/2.3.0/clientinfo/{country_code}/{party_id}"), s.PutOcpiClientInfo)
	}

	if s.locationsSender != nil {
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.3.0/locations"), s.GetOcpiLocations)
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.3.0/locations/{location_id}"), s.GetOcpiLocation)
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.3.0/locations/{location_id}/{evse_uid}"), s.GetOcpiLocation)
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.3.0/locations/{location_id}/{evse_uid}/{connector_id}"), s.GetOcpiLocation)
	}
	if s.locationsReceiver != nil {
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.3.0/locations/{country_code}/{party_id}/{location_id}"), s.GetOcpiClientOwnedLocation)
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.3.0/locations/{country_code}/{party_id}/{location_id}/{evse_uid}"), s.GetOcpiClientOwnedLocation)
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.3.0/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id}"), s.GetOcpiClientOwnedLocation)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleReceiver, "/2.3.0/locations/{country_code}/{party_id}/{location_id}"), s.PutOcpiLocation)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleReceiver, "/2.3.0/locations/{country_code}/{party_id}/{location_id}/{evse_uid}"), s.PutOcpiLocation)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleReceiver, "/2.3.0/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id}"), s.PutOcpiLocation)
		router.HandleFunc("PATCH "+s.withRole(InterfaceRoleReceiver, "/2.3.0/locations/{country_code}/{party_id}/{location_id}"), s.PatchOcpiLocation)
		router.HandleFunc("PATCH "+s.withRole(InterfaceRoleReceiver, "/2.3.0/locations/{country_code}/{party_id}/{location_id}/{evse_uid}"), s.PatchOcpiLocation)
		router.HandleFunc("PATCH "+s.withRole(InterfaceRoleReceiver, "/2.3.0/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id}"), s.PatchOcpiLocation)
	}

	if s.sessionsSender != nil {
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.3.0/sessions"), s.GetOcpiSessions)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleSender, "/2.3.0/sessions/{session_id}/charging_preferences"), s.PutOcpiSesionChargingPreferences)
	}
	if s.sessionsReceiver != nil {
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.3.0/sessions/{country_code}/{party_id}/{session_id}"), s.GetOcpiSession)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleReceiver, "/2.3.0/sessions/{country_code}/{party_id}/{session_id}"), s.PutOcpiSession)
		router.HandleFunc("PATCH "+s.withRole(InterfaceRoleReceiver, "/2.3.0/sessions/{country_code}/{party_id}/{session_id}"), s.PatchOcpiSession)
	}

	if s.tariffsSender != nil {
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.3.0/tariffs"), s.GetOcpiTariffs)
	}
	if s.tariffsReceiver != nil {
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.3.0/tariffs/{country_code}/{party_id}/{tariff_id}"), s.GetOcpiTariff)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleReceiver, "/2.3.0/tariffs/{country_code}/{party_id}/{tariff_id}"), s.PutOcpiTariff)
		router.HandleFunc("DELETE "+s.withRole(InterfaceRoleReceiver, "/2.3.0/tariffs/{country_code}/{party_id}/{tariff_id}"), s.DeleteOcpiTariff)
	}

	if s.tokensSender != nil {
		router.HandleFunc(s.withRole(InterfaceRoleSender, "/2.3.0/tokens"), s.GetOcpiTokens)
		router.HandleFunc("POST "+s.withRole(InterfaceRoleSender, "/2.3.0/tokens/{token_uid}/authorize"), s.PostOcpiToken)
	}
	if s.tokensReceiver != nil {
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.3.0/tokens/{country_code}/{party_id}/{token_uid}"), s.GetOcpiToken)
		router.HandleFunc("PUT "+s.withRole(InterfaceRoleReceiver, "/2.3.0/tokens/{country_code}/{party_id}/{token_uid}"), s.PutOcpiToken)
		router.HandleFunc("PATCH "+s.withRole(InterfaceRoleReceiver, "/2.3.0/tokens/{country_code}/{party_id}/{token_uid}"), s.PatchOcpiToken)
	}

	if s.paymentsSender != nil {
		router.HandleFunc("/ptp/2.3.0/payments/terminals", s.GetOcpiPtpPaymentTerminals)
		router.HandleFunc("/ptp/2.3.0/payments/terminals/{terminal_id}", s.GetOcpiPtpPaymentTerminal)
		router.HandleFunc("POST /ptp/2.3.0/payments/terminals/activate", s.PostOcpiPtpActivatePaymentTerminal)
		router.HandleFunc("POST /ptp/2.3.0/payments/terminals/{terminal_id}/deactivate", s.PostOcpiPtpDeactivatePaymentTerminal)
		router.HandleFunc("PUT /ptp/2.3.0/payments/terminals/{terminal_id}", s.PutOcpiPtpPaymentTerminal)
		router.HandleFunc("PATCH /ptp/2.3.0/payments/terminals/{terminal_id}", s.PatchOcpiPtpPaymentTerminal)

		router.HandleFunc("/ptp/2.3.0/payments/financial-advice-confirmations", s.GetOcpiPtpFinancialAdviceConfirmations)
		router.HandleFunc("/ptp/2.3.0/payments/financial-advice-confirmations/{financial_advice_confirmation_id}", s.GetOcpiPtpFinancialAdviceConfirmation)
	}
	if s.paymentsReceiver != nil {
		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.3.0/payments/terminals/{terminal_id}"), s.GetOcpiPaymentTerminal)
		router.HandleFunc("POST "+s.withRole(InterfaceRoleReceiver, "/2.3.0/payments/terminals"), s.PostOcpiPaymentTerminal)

		router.HandleFunc(s.withRole(InterfaceRoleReceiver, "/2.3.0/payments/financial-advice-confirmations/{financial_advice_confirmation_id}"), s.GetOcpiFinancialAdviceConfirmation)
		router.HandleFunc("POST "+s.withRole(InterfaceRoleReceiver, "/2.3.0/payments/financial-advice-confirmations"), s.PostOcpiFinancialAdviceConfirmation)
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
	if !s.enabledRole {
		return path
	}
	switch role {
	case InterfaceRoleSender:
		return "/emsp" + path
	case InterfaceRoleReceiver:
		return "/cpo" + path
	default:
		panic("unreachable")
	}
}

func (s *Server) authorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")

		requestID := strings.TrimSpace(r.Header.Get(ocpi.HttpHeaderXRequestID))
		correlationID := strings.TrimSpace(r.Header.Get(ocpi.HttpHeaderXCorrelationID))
		defer func() {
			w.Header().Set(ocpi.HttpHeaderXRequestID, requestID)
			w.Header().Set(ocpi.HttpHeaderXCorrelationID, correlationID)
		}()

		token := strings.TrimSpace(r.Header.Get("Authorization"))
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token = strings.TrimPrefix(token, "Token ")
		b, err := base64.StdEncoding.DecodeString(token)
		if err == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token = string(b)
		if err := s.ocpi.VerifyCredentialsToken(r.Context(), token); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(ocpi.WithResponseContext(
			ocpi.WithRequestContext(
				r.Context(),
				&ocpi.RequestContext{
					Token:         token,
					RequestID:     requestID,
					RequestURI:    r.RequestURI,
					CorrelationID: correlationID,
				}), &ocpi.ResponseContext{
				Token:         token,
				RequestID:     requestID,
				CorrelationID: correlationID,
			})))
	})
}
