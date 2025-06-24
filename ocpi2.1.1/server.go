package ocpi211

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/si3nloong/ocpi-go/ocpi"
)

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
	s.serverConfig.baseUrl = fmt.Sprintf("/ocpi/%s", ocpi.VersionNumber211)
	for _, opt := range options {
		opt.apply(&s.serverConfig)
	}
	s.baseUrl = "/" + string(ocpi.VersionNumber211)
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

	router.HandleFunc(s.baseUrl, s.GetOcpiVersionDetails)
	router.HandleFunc(s.baseUrl+"/credentials", s.GetOcpiCredentials)
	router.HandleFunc("POST "+s.baseUrl+"/credentials", s.PostOcpiCredentials)
	router.HandleFunc("PUT "+s.baseUrl+"/credentials", s.PutOcpiCredentials)
	router.HandleFunc("DELETE "+s.baseUrl+"/credentials", s.DeleteOcpiCredentials)

	if s.cpo != nil {
		router.HandleFunc(s.baseUrl+"/locations", s.GetOcpiLocations)
		router.HandleFunc(s.baseUrl+"/sessions", s.GetOcpiLocations)
		router.HandleFunc(s.baseUrl+"/cdrs", s.GetOcpiCDRs)
		router.HandleFunc(s.baseUrl+"/tariffs", s.GetOcpiTariffs)
		router.HandleFunc(s.baseUrl+"/tokens/{country_code}/{party_id}/{token_uid}", s.GetOcpiToken)
		router.HandleFunc("PUT "+s.baseUrl+"/tokens/{country_code}/{party_id}/{token_uid}", s.PutOcpiToken)
		router.HandleFunc("PATCH "+s.baseUrl+"/tokens/{country_code}/{party_id}/{token_uid}", s.PatchOcpiToken)
		router.HandleFunc("POST "+s.baseUrl+"/commands/{command_type}", s.PostOcpiCommandResponse)
	}
	if s.emsp != nil {
		router.HandleFunc(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}(/{evse_uid}(/{connector_id}))", s.GetOcpiLocation)
		router.HandleFunc("PUT "+s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}(/{evse_uid}(/{connector_id}))", s.PutOcpiLocation)
		router.HandleFunc("PATCH "+s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}(/{evse_uid}(/{connector_id}))", s.PatchOcpiLocation)

		router.HandleFunc(s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.GetOcpiSession)
		router.HandleFunc("PUT "+s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.PutOcpiSession)
		router.HandleFunc("PATCH "+s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.PatchOcpiSession)

		router.HandleFunc(s.baseUrl+"/cdrs/{cdr_id}", s.GetOcpiCDR)
		router.HandleFunc("POST "+s.baseUrl+"/cdrs", s.PostOcpiCDR)

		router.HandleFunc(s.baseUrl+"/tariffs/{country_code}/{party_id}/{tariff_id}", s.GetOcpiTariff)
		router.HandleFunc("PUT "+s.baseUrl+"/tariffs/{country_code}/{party_id}/{tariff_id}", s.PutOcpiTariff)
		router.HandleFunc("PATCH "+s.baseUrl+"/tariffs/{country_code}/{party_id}/{tariff_id}", s.PatchOcpiTariff)
		router.HandleFunc("DELETE "+s.baseUrl+"/tariffs/{country_code}/{party_id}/{tariff_id}", s.DeleteOcpiTariff)

		router.HandleFunc(s.baseUrl+"/tokens", s.GetOcpiTokens)
		router.HandleFunc("POST "+s.baseUrl+"/tokens/{token_uid}/authorize", s.PostOcpiToken)

		router.HandleFunc("POST "+s.baseUrl+"/commands/{command_type}/{uid}", s.PostOcpiCommandResponse)
	}
	return router
}
