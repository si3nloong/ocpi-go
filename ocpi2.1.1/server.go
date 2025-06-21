package ocpi211

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/si3nloong/ocpi-go/ocpi"
)

type Server struct {
	baseUrl string
	roles   map[Role]struct{}
	cpo     CPO
	emsp    EMSP
	logger  *slog.Logger
}

func NewServer() *Server {
	s := new(Server)
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
	router := chi.NewRouter()

	router.Get(s.baseUrl, s.GetOcpiVersionDetails)
	router.Get(s.baseUrl+"/credentials", s.GetOcpiCredentials)
	router.Post(s.baseUrl+"/credentials", s.PostOcpiCredentials)
	router.Put(s.baseUrl+"/credentials", s.PutOcpiCredentials)
	router.Delete(s.baseUrl+"/credentials", s.DeleteOcpiCredentials)

	if s.cpo != nil {
		router.Get(s.baseUrl+"/locations", s.GetOcpiLocations)
		router.Get(s.baseUrl+"/sessions", s.GetOcpiLocations)
		router.Get(s.baseUrl+"/cdrs", s.GetOcpiCDRs)
		router.Get(s.baseUrl+"/tariffs", s.GetOcpiTariffs)
		router.Get(s.baseUrl+"/tokens/{country_code}/{party_id}/{token_uid}", s.GetOcpiToken)
		router.Put(s.baseUrl+"/tokens/{country_code}/{party_id}/{token_uid}", s.PutOcpiToken)
		router.Patch(s.baseUrl+"/tokens/{country_code}/{party_id}/{token_uid}", s.PatchOcpiToken)
		router.Post(s.baseUrl+"/commands/{command_type}", s.PostOcpiCommandResponse)
	}
	if s.emsp != nil {
		router.Get(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}(/{evse_uid}(/{connector_id}))", s.GetOcpiLocation)
		router.Put(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}(/{evse_uid}(/{connector_id}))", s.PutOcpiLocation)
		router.Patch(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}(/{evse_uid}(/{connector_id}))", s.PatchOcpiLocation)

		router.Get(s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.GetOcpiSession)
		router.Put(s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.PutOcpiSession)
		router.Patch(s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.PatchOcpiSession)

		router.Get(s.baseUrl+"/cdrs/{cdr_id}", s.GetOcpiCDR)
		router.Post(s.baseUrl+"/cdrs", s.PostOcpiCDR)

		router.Get(s.baseUrl+"/tariffs/{country_code}/{party_id}/{tariff_id}", s.GetOcpiTariff)
		router.Put(s.baseUrl+"/tariffs/{country_code}/{party_id}/{tariff_id}", s.PutOcpiTariff)
		router.Patch(s.baseUrl+"/tariffs/{country_code}/{party_id}/{tariff_id}", s.PatchOcpiTariff)
		router.Delete(s.baseUrl+"/tariffs/{country_code}/{party_id}/{tariff_id}", s.DeleteOcpiTariff)

		router.Get(s.baseUrl+"/tokens", s.GetOcpiTokens)
		router.Post(s.baseUrl+"/tokens/{token_uid}/authorize", s.PostOcpiToken)

		router.Post(s.baseUrl+"/commands/{command_type}/{uid}", s.PostOcpiCommandResponse)
	}

	return router
}
