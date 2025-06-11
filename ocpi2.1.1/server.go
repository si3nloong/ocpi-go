package ocpi211

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/si3nloong/ocpi-go/ocpi"
)

type Server struct {
	baseUrl     string
	cpo         CPO
	emsp        EMSP
	httpHandler http.Handler
	logger      *slog.Logger
}

func NewServer() *Server {
	s := new(Server)
	s.baseUrl = "/" + string(ocpi.VersionNumber211)
	s.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := chi.NewRouter()

	router.HandleFunc(s.baseUrl+"/credentials", s.GetOcpiCredentials)

	router.ServeHTTP(w, r)
}
