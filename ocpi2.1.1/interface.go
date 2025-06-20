package ocpi211

import (
	"context"

	"github.com/si3nloong/ocpi-go/ocpi"
)

type CPO interface {
	// (GET /ocpi/2.1.1/locations)
	OnGetLocations(ctx context.Context, params GetLocationsParams) (*ocpi.PaginationResponse[Location], error)
	// (GET /ocpi/2.1.1/cdrs)
	OnGetCDRs(ctx context.Context, params GetCdrsParams) (*ocpi.PaginationResponse[ChargeDetailRecord], error)
	// (GET /ocpi/2.1.1/sessions)
	OnGetSessions(ctx context.Context, params GetSessionsParams) (*ocpi.PaginationResponse[Session], error)
	// (GET /ocpi/2.1.1/tariffs)
	OnGetTariffs(ctx context.Context, params GetTariffsParams) (*ocpi.PaginationResponse[Session], error)
	// (GET /ocpi/2.1.1/tokens/{country_code}/{party_id}/{token_uid}[?type={type}])
	OnGetClientOwnedToken(ctx context.Context, countryCode string, partyID string, tokenUID string, tokenType ...TokenType) (*Token, error)
	// (PUT /ocpi/2.1.1/tokens/{country_code}/{party_id}/{token_uid})
	OnPutClientOwnedToken(ctx context.Context, countryCode string, partyID string, tokenUID string, body ocpi.RawMessage[Token], tokenType ...TokenType) error
	// (PATCH /ocpi/2.1.1/tokens/{country_code}/{party_id}/{token_uid})
	OnPatchClientOwnedToken(ctx context.Context, countryCode string, partyID string, tokenUID string, body ocpi.RawMessage[PatchedToken], tokenType ...TokenType) error
	// (POST /ocpi/2.1.1/commands/{command})
	OnPostCommand(ctx context.Context, commandType CommandType) (*CommandResponse, error)
}

type EMSP interface {
	// (GET /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id})
	OnGetClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string) (*Location, error)
	// (GET /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	OnGetClientOwnedLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string) (*EVSE, error)
	// (GET /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	OnGetClientOwnedLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string) (*Connector, error)
	// (PUT /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id})
	OnPutClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string, body ocpi.RawMessage[Location]) error
	// (PUT /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	OnPutClientOwnedLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, body ocpi.RawMessage[Location]) error
	// (PUT /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	OnPutClientOwnedLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string, body ocpi.RawMessage[Location]) error
	// (PATCH /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id})
	OnPatchClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string, body ocpi.RawMessage[PatchedLocation]) error
	// (PATCH /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	OnPatchClientOwnedLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, body ocpi.RawMessage[PatchedLocation]) error
	// (PATCH /ocpi/2.1.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	OnPatchClientOwnedLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string, body ocpi.RawMessage[PatchedLocation]) error

	// (GET /ocpi/2.1.1/sessions/{country_code}/{party_id}/{session_id})
	OnGetClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string) (*Session, error)
	// (PUT /ocpi/2.1.1/sessions/{country_code}/{party_id}/{session_id})
	OnPutClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, body ocpi.RawMessage[Session]) error
	// (PATCH /ocpi/2.1.1/sessions/{country_code}/{party_id}/{session_id})
	OnPatchClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, body ocpi.RawMessage[PatchedSession]) error

	// (GET /ocpi/2.1.1/cdrs/{cdr_id})
	OnGetCDR(ctx context.Context, cdrID string) (*ChargeDetailRecord, error)
	// (POST /ocpi/2.1.1/cdrs)
	OnPostCDR(ctx context.Context, body ocpi.RawMessage[ChargeDetailRecord]) (*ChargeDetailRecordResponse, error)

	// (GET /ocpi/2.1.1/tariffs/{country_code}/{party_id}/{tariff_id})
	OnGetClientOwnedTariff(ctx context.Context, countryCode string, partyID string, sessionID string) (*Tariff, error)
	// (PUT /ocpi/2.1.1/tariffs/{country_code}/{party_id}/{tariff_id})
	OnPutClientOwnedTariff(ctx context.Context, countryCode string, partyID string, tariffID string, body ocpi.RawMessage[Tariff]) error
	// (PATCH /ocpi/2.1.1/tariffs/{country_code}/{party_id}/{tariff_id})
	OnPatchClientOwnedTariff(ctx context.Context, countryCode string, partyID string, tariffID string, body ocpi.RawMessage[Tariff]) error
	// (DELETE /ocpi/2.1.1/tariffs/{country_code}/{party_id}/{tariff_id})
	OnDeleteClientOwnedTariff(ctx context.Context, countryCode string, partyID string, tariffID string) error

	// (GET /ocpi/2.1.1/tokens)
	OnGetTokens(ctx context.Context, params GetTokensParams) (*ocpi.PaginationResponse[Token], error)
	// (POST /ocpi/2.1.1/tokens/{token_uid}/authorize?{type=token_type})
	OnPostToken(ctx context.Context, tokenUID string, tokenType TokenType, body ocpi.RawMessage[*LocationReferences]) (*AuthorizationInfo, error)

	// (POST /ocpi/2.1.1/commands/{command}/{uid})
	OnPostAsyncCommand(ctx context.Context, commandType CommandType, uid string, body ocpi.RawMessage[CommandResponse]) error
}
