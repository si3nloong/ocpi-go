package ocpi211

import "context"

type EMSP interface {
	// (GET /ocpi/locations/{country_code}/{party_id}/{location_id})
	GetLocation(ctx context.Context, countryCode string, partyID string, locationID string) (*Location, error)
	// (GET /ocpi/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	GetLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string) (*Location, error)
	// (GET /ocpi/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	GetLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string) (*Location, error)
}
