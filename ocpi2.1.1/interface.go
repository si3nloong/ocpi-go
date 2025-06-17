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
	// (POST /ocpi/2.1.1/commands/{command})
	PostCommand(ctx context.Context, commandType CommandType) (*CommandResponse, error)
}

type EMSP interface {
	// (GET /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id})
	GetLocation(ctx context.Context, countryCode string, partyID string, locationID string) (*Location, error)
	// (GET /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	GetLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUid string) (*EVSE, error)
	// (GET /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	GetLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUid string, connectorID string) (*Connector, error)
	// (PUT /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id})
	PutLocation(ctx context.Context, countryCode string, partyID string, locationID string, body ocpi.RawMessage[Location]) error
	// (PUT /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	PutLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUid string, body ocpi.RawMessage[Location]) error
	// (PUT /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	PutLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUid string, connectorID string, body ocpi.RawMessage[Location]) error
	// (PATCH /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id})
	PatchLocation(ctx context.Context, countryCode string, partyID string, locationID string, body ocpi.RawMessage[PatchedLocation]) error
	// (PATCH /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	PatchLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUid string, body ocpi.RawMessage[PatchedLocation]) error
	// (PATCH /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	PatchLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUid string, connectorID string, body ocpi.RawMessage[PatchedLocation]) error

	// (GET /ocpi/2.1.1/sessions/{country_code}/{party_id}/{session_id})
	GetSession(ctx context.Context, countryCode string, partyID string, sessionId string) (*Session, error)
	// (PUT /ocpi/2.1.1/sessions/{country_code}/{party_id}/{session_id})
	PutSession(ctx context.Context, countryCode string, partyID string, sessionId string, body ocpi.RawMessage[Session]) error
	// (PATCH /ocpi/2.1.1/sessions/{country_code}/{party_id}/{session_id})
	PatchSession(ctx context.Context, countryCode string, partyID string, sessionId string, body ocpi.RawMessage[PatchedSession]) error

	// (POST /ocpi/2.1.1/commands/{command}/{uid})
	PostAsyncCommand(ctx context.Context, commandType CommandType, uid string, body ocpi.RawMessage[CommandResponse]) error
}
