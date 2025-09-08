package ocpi230

import (
	"encoding/json"
	"net/http"

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
	resp, err := s.commandsReceiver.OnPostCommand(r.Context(), CommandType(commandType), CommandRequest(body))
	if err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, resp)
}

func (s *Server) PostOcpiCommandResponse(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ocpihttp.BadRequest(w, r, err.Error())
		return
	}

	commandType := r.PathValue("command_type")
	uid := r.PathValue("uid")
	if err := s.commandsSender.OnPostAsyncCommand(r.Context(), CommandType(commandType), uid, ocpi.RawMessage[CommandResult](body)); err != nil {
		ocpihttp.Response(w, err)
		return
	}

	ocpihttp.Response(w, ocpi.NewEmptyResponse())
}
