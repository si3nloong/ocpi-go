package ocpi221

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	"github.com/si3nloong/ocpi-go/ocpi"
)

type Server struct {
	baseUrl    string
	httpServer *http.Server
	logger     *slog.Logger
	sender     Sender
	receiver   Receiver
}

func NewServer(basePath string, sender Sender, receiver Receiver) *Server {
	s := new(Server)
	s.baseUrl = path.Join(basePath, string(ocpi.VersionN221))
	s.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	s.sender = sender
	s.receiver = receiver
	return s
}

func (s *Server) Handler() http.Handler {
	router := chi.NewRouter()

	// router.Use(func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		b, _ := httputil.DumpRequest(r, true)
	// 		s.logger.DebugContext(r.Context(), `Client Request`, "request", unsafe.String(unsafe.SliceData(b), len(b)))
	// 		recorder := httptest.NewRecorder()
	// 		next.ServeHTTP(recorder, r)
	// 		result := recorder.Result()
	// 		b, _ = httputil.DumpResponse(result, true)
	// 		s.logger.DebugContext(r.Context(), `Server Response`, "response", unsafe.String(unsafe.SliceData(b), len(b)))
	// 		w.WriteHeader(result.StatusCode)
	// 		w.Write(recorder.Body.Bytes())
	// 		for k, headers := range result.Header.Clone() {
	// 			for len(headers) > 0 {
	// 				h := headers[0]
	// 				w.Header().Set(k, h)
	// 				headers = headers[1:]
	// 			}
	// 		}
	// 	})
	// })

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
	router.With(
		httpin.NewInput(PutOcpiSessionRequestBody{}),
	).Put(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}(/{evse_uid}(/{connector_id}))", s.PutOcpiLocation)

	router.With(
		httpin.NewInput(PatchOcpiSessionRequestBody{}),
	).Patch(s.baseUrl+"/locations/{country_code}/{party_id}/{location_id}(/{evse_uid}(/{connector_id}))", s.PatchOcpiLocation)

	router.Get(s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.GetOcpiSession)
	router.With(
		httpin.NewInput(PutOcpiSessionRequestBody{}),
	).Put(s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.PutOcpiSession)
	router.With(
		httpin.NewInput(PatchOcpiSessionRequestBody{}),
	).Patch(s.baseUrl+"/sessions/{country_code}/{party_id}/{session_id}", s.PatchOcpiSession)

	router.Get(s.baseUrl+"/tokens/{country_code}/{party_id}/{token_uid}", s.GetOcpiToken)

	router.Get(s.baseUrl+"/cdrs/{id}", s.GetOcpiCDR)
	router.With(
		httpin.NewInput(PostOcpiCdrRequestBody{}),
	).Post(s.baseUrl+"/cdrs", s.PostOcpiCDR)

	router.Post(s.baseUrl+"/commands/{command_type}/{session_uid}", s.PostOcpiCommandResponse)

	return router
}

func (s *Server) Start(listenPort int) error {
	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", listenPort),
		Handler: s.Handler(),
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	if s.httpServer == nil {
		return nil
	}
	return s.httpServer.Shutdown(context.Background())
}

func getHostname(r *http.Request) string {
	hostname := "http://" + r.Host
	if r.TLS != nil {
		hostname = "https://" + r.Host
	}
	return hostname
}

func toResponse[T any](value T) ([]byte, error) {
	return json.Marshal(ocpi.Response[T]{
		Data:          value,
		StatusCode:    ocpi.GenericSuccessCode,
		StatusMessage: ocpi.GenericSuccessCode.String(),
		Timestamp:     time.Now().UTC(),
	})
}

func httpError(w http.ResponseWriter, err error, statusCode ocpi.StatusCode) {
	b, _ := json.Marshal(ocpi.Response[any]{
		StatusCode:    statusCode,
		StatusMessage: err.Error(),
		Timestamp:     time.Now(),
	})
	h := w.Header()
	h.Del("Content-Length")
	h.Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
