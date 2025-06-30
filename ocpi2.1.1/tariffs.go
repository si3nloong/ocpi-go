package ocpi211

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/si3nloong/ocpi-go/internal/httputil"
	"github.com/si3nloong/ocpi-go/ocpi"
)

func (s *Server) GetOcpiTariffs(w http.ResponseWriter, r *http.Request) {

	params := GetTariffsParams{}
	response, err := s.cpo.OnGetTariffs(r.Context(), params)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	httputil.ResponsePagination(w, r, response)
}

func (s *Server) GetOcpiTariff(w http.ResponseWriter, r *http.Request) {

	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	tariffID := r.PathValue("tariff_id")

	tariff, err := s.emsp.OnGetClientOwnedTariff(r.Context(), countryCode, partyID, tariffID)
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewResponse(tariff))
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PutOcpiTariff(w http.ResponseWriter, r *http.Request) {

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	tariffID := r.PathValue("tariff_id")

	if err := s.emsp.OnPutClientOwnedTariff(
		r.Context(),
		countryCode,
		partyID,
		tariffID,
		ocpi.RawMessage[Tariff](body),
	); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) PatchOcpiTariff(w http.ResponseWriter, r *http.Request) {

	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	tariffID := r.PathValue("tariff_id")

	if err := s.emsp.OnPatchClientOwnedTariff(
		r.Context(),
		countryCode,
		partyID,
		tariffID,
		ocpi.RawMessage[Tariff](body),
	); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (s *Server) DeleteOcpiTariff(w http.ResponseWriter, r *http.Request) {

	countryCode := r.PathValue("country_code")
	partyID := r.PathValue("party_id")
	tariffID := r.PathValue("tariff_id")

	if err := s.emsp.OnDeleteClientOwnedTariff(
		r.Context(),
		countryCode,
		partyID,
		tariffID,
	); err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	b, err := json.Marshal(ocpi.NewEmptyResponse())
	if err != nil {
		httputil.ResponseError(w, err, ocpi.StatusCodeServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (c *ClientConn) GetTariffs(
	ctx context.Context,
	params ...GetTariffsParams,
) (ocpi.Result[[]Tariff], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDTariffs)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("limit", "100")
	if len(params) > 0 {
		p := params[0]
		if p.DateFrom != nil && p.DateFrom.IsZero() {
			query.Set("date_from", p.DateFrom.Format(time.RFC3339))
		}
		if p.DateTo != nil && p.DateTo.IsZero() {
			query.Set("date_to", p.DateTo.Format(time.RFC3339))
		}
		if p.Offset != nil && *p.Offset > 0 {
			query.Set("offset", strconv.FormatUint(uint64(*p.Offset), 10))
		}
		if p.Limit != nil && *p.Limit > 0 {
			query.Set("limit", strconv.FormatUint(uint64(*p.Limit), 10))
		}
	}
	u.RawQuery = query.Encode()
	var o TariffsResponse
	if err := c.do(ctx, http.MethodGet, u.String(), nil, &o); err != nil {
		return nil, err
	}
	return ocpi.NewResult(o), nil
}

func (c *ClientConn) GetTariff(ctx context.Context, countryCode, partyID, tariffID string) (any, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDTariffs)
	if err != nil {
		return nil, err
	}

	var o json.RawMessage
	if err := c.do(ctx, http.MethodGet, endpoint+"/"+countryCode+"/"+partyID+"/"+tariffID, nil, &o); err != nil {
		return nil, err
	}
	return string(o), nil
}
