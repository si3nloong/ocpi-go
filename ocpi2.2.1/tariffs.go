package ocpi221

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/si3nloong/ocpi-go/ocpi"
)

type GetTariffsParams struct {
	DateFrom time.Time
	DateTo   time.Time
	Offset   uint32
	Limit    uint8
}

func (c *Client) GetTariffs(
	ctx context.Context,
	params ...GetTariffsParams,
) (ocpi.Result[[]Tariff], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDTariffs, InterfaceRoleSender)
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
		if p.Limit == 0 {
			p.Limit = 100
		}
		if !p.DateFrom.IsZero() {
			query.Set("date_from", p.DateFrom.Format(time.RFC3339))
		}
		if !p.DateTo.IsZero() {
			query.Set("date_to", p.DateTo.Format(time.RFC3339))
		}
		if p.Offset > 0 {
			query.Set("offset", strconv.FormatUint(uint64(p.Offset)*uint64(p.Limit), 10))
		}
		query.Set("limit", strconv.FormatUint(uint64(p.Limit), 10))
	}
	u.RawQuery = query.Encode()
	var o TariffsResponse
	if err := c.do(ctx, http.MethodGet, u.String(), nil, &o); err != nil {
		return nil, err
	}
	return ocpi.NewResult(o), nil
}

func (c *Client) GetTariff(ctx context.Context, countryCode, partyId, tariffId string) (any, error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDTariffs, InterfaceRoleSender)
	if err != nil {
		return nil, err
	}

	var o json.RawMessage
	if err := c.do(ctx, http.MethodGet, endpoint+"/"+countryCode+"/"+partyId+"/"+tariffId, nil, &o); err != nil {
		return nil, err
	}
	return string(o), nil
}
