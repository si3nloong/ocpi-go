package ocpi221

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/si3nloong/ocpi-go/ocpi"
)

type Server struct {
	baseUrl string
	// httpServer *http.Server
	httpHandler http.Handler
	logger      *slog.Logger
	sender      Sender
	receiver    Receiver
}

func NewServer(sender Sender, receiver Receiver) *Server {
	s := new(Server)
	s.baseUrl = "/" + string(ocpi.VersionNumber221)
	s.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	s.sender = sender
	s.receiver = receiver
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.httpHandler == nil {
		s.httpHandler = s.handler()
	}
	s.httpHandler.ServeHTTP(w, r)
}

func (s *Server) handler() http.Handler {
	router := chi.NewRouter()

	router.HandleFunc(s.baseUrl, s.GetOcpiEndpoints)
	router.HandleFunc(s.baseUrl+"/credentials", s.GetOcpiCredentials)

	if s.sender != nil {
		router.HandleFunc(s.baseUrl+"/locations", s.GetOcpiLocations)
		router.HandleFunc(s.baseUrl+"/sessions", s.GetOcpiSessions)
		router.HandleFunc(s.baseUrl+"/tariffs", s.GetOcpiTariffs)
		router.HandleFunc(s.baseUrl+"/tokens", s.GetOcpiTokens)
		router.Post(s.baseUrl+"/tokens/{token_uid}/authorize", s.PostOcpiToken)
	}

	router.HandleFunc(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}(/{evse_uid}(/{connector_id}))", s.GetOcpiLocation)
	router.Put(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}(/{evse_uid}(/{connector_id}))", s.PutOcpiLocation)
	router.Patch(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}(/{evse_uid}(/{connector_id}))", s.PatchOcpiLocation)

	router.Get(s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.GetOcpiSession)
	router.Put(s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.PutOcpiSession)
	router.Patch(s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.PatchOcpiSession)

	router.Get(s.baseUrl+"/tokens/{country_code}/{party_id}/{token_uid}", s.GetOcpiToken)

	router.Get(s.baseUrl+"/cdrs/{id}", s.GetOcpiCDR)
	router.Post(s.baseUrl+"/cdrs", s.PostOcpiCDR)

	router.Post(s.baseUrl+"/commands/{command_type}/{session_uid}", s.PostOcpiCommandResponse)

	return router
}

// func (s *Server) Stop() error {
// 	if s.httpServer == nil {
// 		return nil
// 	}
// 	return s.httpServer.Shutdown(context.Background())
// }
