package ocpi211

import (
	"context"

	"github.com/si3nloong/ocpi-go/ocpi"
)

type CPO interface {
	// (GET /ocpi/2.1.1/locations)
	GetLocations(ctx context.Context, params GetLocationsParams) (*ocpi.PaginationResponse[Location], error)
	// (GET /ocpi/2.1.1/cdrs)
	GetCDRs(ctx context.Context, params GetCdrsParams) (*ocpi.PaginationResponse[ChargeDetailRecord], error)
	// (GET /ocpi/2.1.1/sessions)
	GetSessions(ctx context.Context, params GetSessionsParams) (*ocpi.PaginationResponse[Session], error)
	// (GET /ocpi/2.1.1/tariffs)
	GetTariffs(ctx context.Context, params GetTariffsParams) (*ocpi.PaginationResponse[Session], error)
	// (POST /ocpi/2.1.1/commands/{command})
	PostCommand(ctx context.Context, commandType CommandType) (*CommandResponse, error)
}

type EMSP interface {
	// (GET /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id})
	GetLocation(ctx context.Context, countryCode string, partyID string, locationID string) (*Location, error)
	// (GET /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	GetLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string) (*EVSE, error)
	// (GET /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	GetLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string) (*Connector, error)
	// (PUT /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id})
	PutLocation(ctx context.Context, countryCode string, partyID string, locationID string, body ocpi.RawMessage[Location]) error
	// (PUT /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	PutLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, body ocpi.RawMessage[Location]) error
	// (PUT /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	PutLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string, body ocpi.RawMessage[Location]) error
	// (PATCH /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id})
	PatchLocation(ctx context.Context, countryCode string, partyID string, locationID string, body ocpi.RawMessage[PatchedLocation]) error
	// (PATCH /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	PatchLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, body ocpi.RawMessage[PatchedLocation]) error
	// (PATCH /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	PatchLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string, body ocpi.RawMessage[PatchedLocation]) error

	// (GET /ocpi/2.1.1/sessions/{country_code}/{party_id}/{session_id})
	GetSession(ctx context.Context, countryCode string, partyID string, sessionID string) (*Session, error)
	// (PUT /ocpi/2.1.1/sessions/{country_code}/{party_id}/{session_id})
	PutSession(ctx context.Context, countryCode string, partyID string, sessionID string, body ocpi.RawMessage[Session]) error
	// (PATCH /ocpi/2.1.1/sessions/{country_code}/{party_id}/{session_id})
	PatchSession(ctx context.Context, countryCode string, partyID string, sessionID string, body ocpi.RawMessage[PatchedSession]) error

	// (GET /ocpi/2.1.1/tariffs/{country_code}/{party_id}/{tariff_id})
	GetTariff(ctx context.Context, countryCode string, partyID string, sessionID string) (*Tariff, error)
	// (PUT /ocpi/2.1.1/tariffs/{country_code}/{party_id}/{tariff_id})
	PutTariff(ctx context.Context, countryCode string, partyID string, tariffID string, body ocpi.RawMessage[Tariff]) error
	// (PATCH /ocpi/2.1.1/tariffs/{country_code}/{party_id}/{tariff_id})
	PatchTariff(ctx context.Context, countryCode string, partyID string, tariffID string, body ocpi.RawMessage[Tariff]) error
	// (DELETE /ocpi/2.1.1/tariffs/{country_code}/{party_id}/{tariff_id})
	DeleteTariff(ctx context.Context, countryCode string, partyID string, tariffID string) error

	// (POST /ocpi/2.1.1/commands/{command}/{uid})
	PostAsyncCommand(ctx context.Context, commandType CommandType, uid string, body ocpi.RawMessage[CommandResponse]) error
}
