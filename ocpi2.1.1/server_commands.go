package ocpi211

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/si3nloong/ocpi-go/ocpi"
	ocpihttp "github.com/si3nloong/ocpi-go/ocpi/http"
)

func (s *Server) PostOcpiCommand(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	commandType := r.PathValue("command_type")
	resp, err := s.cpo.OnPostCommand(r.Context(), CommandType(commandType), body)
	if err != nil {
		ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, err)
		return
	}

	ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, resp)
}

func (s *Server) PostOcpiCommandResponse(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	commandType := r.PathValue("command_type")
	uid := r.PathValue("uid")
	if err := s.emsp.OnPostAsyncCommand(r.Context(), CommandType(commandType), uid, ocpi.RawMessage[CommandResponse](body)); err != nil {
		ocpihttp.Response(w, DateTime{Time: time.Now().UTC()}, err)
		return
	}

	ocpihttp.EmptyResponse(w, DateTime{Time: time.Now().UTC()})
}
