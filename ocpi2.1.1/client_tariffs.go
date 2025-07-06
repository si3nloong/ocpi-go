package ocpi211

import (
	"context"
	"net/http"

	"github.com/si3nloong/ocpi-go/ocpi"
)

func (c *ClientConn) GetClientOwnedTariff(ctx context.Context, countryCode, partyID, tariffID string) (*ocpi.Response[Tariff], error) {
	endpoint, err := c.getEndpoint(ctx, ModuleIDTariffs)
	if err != nil {
		return nil, err
	}

	var res ocpi.Response[Tariff]
	if err := c.do(ctx, http.MethodGet, endpoint+"/"+countryCode+"/"+partyID+"/"+tariffID, nil, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
