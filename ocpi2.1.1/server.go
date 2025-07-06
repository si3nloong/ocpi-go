package ocpi211

import (
	"log/slog"
	"net/http"
	"os"
)

type Handler func(http.ResponseWriter, *http.Request)

type ServerOption interface {
	apply(*serverConfig)
}

type serverConfig struct {
	baseUrl string
}

type Server struct {
	serverConfig
	roles  map[Role]struct{}
	cpo    CPO
	emsp   EMSP
	logger *slog.Logger
}

func NewServer(options ...ServerOption) *Server {
	s := new(Server)
	s.roles = make(map[Role]struct{})
	s.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	return s
}

func (s *Server) SetCPO(cpo CPO) {
	s.roles[RoleCPO] = struct{}{}
	s.cpo = cpo
}

func (s *Server) SetEMSP(emsp EMSP) {
	s.roles[RoleEMSP] = struct{}{}
	s.emsp = emsp
}

func (s *Server) Handler() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/cpo/2.1.1/details", s.GetOcpiVersionDetails(RoleCPO))
	router.HandleFunc("/cpo/2.1.1/credentials", s.GetOcpiCredentials)
	router.HandleFunc("POST /cpo/2.1.1/credentials", s.PostOcpiCredentials)
	router.HandleFunc("PUT /cpo/2.1.1/credentials", s.PutOcpiCredentials)
	router.HandleFunc("DELETE /cpo/2.1.1/credentials", s.DeleteOcpiCredentials)

	router.HandleFunc("/emsp/2.1.1/details", s.GetOcpiVersionDetails(RoleEMSP))
	router.HandleFunc("/emsp/2.1.1/credentials", s.GetOcpiCredentials)
	router.HandleFunc("POST /emsp/2.1.1/credentials", s.PostOcpiCredentials)
	router.HandleFunc("PUT /emsp/2.1.1/credentials", s.PutOcpiCredentials)
	router.HandleFunc("DELETE /emsp/2.1.1/credentials", s.DeleteOcpiCredentials)

	if s.cpo != nil {
		router.HandleFunc("/cpo/2.1.1/locations", s.GetOcpiLocations)
		router.HandleFunc("/cpo/2.1.1/sessions", s.GetOcpiLocations)
		router.HandleFunc("/cpo/2.1.1/cdrs", s.GetOcpiCDRs)
		router.HandleFunc("/cpo/2.1.1/tariffs", s.GetOcpiTariffs)
		router.HandleFunc("/cpo/2.1.1/tokens/{country_code}/{party_id}/{token_uid}", s.GetOcpiToken)
		router.HandleFunc("PUT /cpo/2.1.1/tokens/{country_code}/{party_id}/{token_uid}", s.PutOcpiToken)
		router.HandleFunc("PATCH /cpo/2.1.1/tokens/{country_code}/{party_id}/{token_uid}", s.PatchOcpiToken)
		router.HandleFunc("POST /cpo/2.1.1/commands/{command_type}", s.PostOcpiCommandResponse)
	}

	if s.emsp != nil {
		router.HandleFunc("/emsp/2.1.1/locations/{country_code}/{party_id}/{location_id}", s.GetOcpiLocation)
		router.HandleFunc("/emsp/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}", s.GetOcpiLocation)
		router.HandleFunc("/emsp/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id}", s.GetOcpiLocation)
		router.HandleFunc("PUT /emsp/2.1.1/locations/{country_code}/{party_id}/{location_id}", s.PutOcpiLocation)
		router.HandleFunc("PUT /emsp/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}", s.PutOcpiLocation)
		router.HandleFunc("PUT /emsp/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id}", s.PutOcpiLocation)
		router.HandleFunc("PATCH /emsp/2.1.1/locations/{country_code}/{party_id}/{location_id}", s.PatchOcpiLocation)
		router.HandleFunc("PATCH /emsp/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}", s.PatchOcpiLocation)
		router.HandleFunc("PATCH /emsp/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id}", s.PatchOcpiLocation)

		router.HandleFunc("/emsp/2.1.1/sessions/{country_code}/{party_id}/{session_id}", s.GetOcpiSession)
		router.HandleFunc("PUT /emsp/2.1.1/sessions/{country_code}/{party_id}/{session_id}", s.PutOcpiSession)
		router.HandleFunc("PATCH /emsp/2.1.1/sessions/{country_code}/{party_id}/{session_id}", s.PatchOcpiSession)

		router.HandleFunc("/emsp/2.1.1/cdrs/{cdr_id}", s.GetOcpiCDR)
		router.HandleFunc("POST /emsp/2.1.1/cdrs", s.PostOcpiCDR)

		router.HandleFunc("/emsp/2.1.1/tariffs/{country_code}/{party_id}/{tariff_id}", s.GetOcpiTariff)
		router.HandleFunc("PUT /emsp/2.1.1/tariffs/{country_code}/{party_id}/{tariff_id}", s.PutOcpiTariff)
		router.HandleFunc("PATCH /emsp/2.1.1/tariffs/{country_code}/{party_id}/{tariff_id}", s.PatchOcpiTariff)
		router.HandleFunc("DELETE /emsp/2.1.1/tariffs/{country_code}/{party_id}/{tariff_id}", s.DeleteOcpiTariff)

		router.HandleFunc("/emsp/2.1.1/tokens", s.GetOcpiTokens)
		router.HandleFunc("POST /emsp/2.1.1/tokens/{token_uid}/authorize", s.PostOcpiToken)

		router.HandleFunc("POST /emsp/2.1.1/commands/{command_type}/{uid}", s.PostOcpiCommandResponse)
	}
	return s.authorizeMiddleware(router)
}
