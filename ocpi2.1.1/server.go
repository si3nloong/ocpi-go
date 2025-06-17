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

		router.Post(s.baseUrl+"/commands/{command_type}", s.PostOcpiCommandResponse)
	}
	if s.emsp != nil {
		router.Get(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}(/{evse_uid}(/{connector_id}))", s.GetOcpiLocation)
		router.Put(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}(/{evse_uid}(/{connector_id}))", s.PutOcpiLocation)
		router.Patch(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}(/{evse_uid}(/{connector_id}))", s.PatchOcpiLocation)

		router.Post(s.baseUrl+"/commands/{command_type}/{uid}", s.PostOcpiCommandResponse)
	}

	return router
}
