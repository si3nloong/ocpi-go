package ocpi221

import (
	"context"

	"github.com/si3nloong/ocpi-go/ocpi"
)

type CPO interface {
	CDRsSender
	ChargingProfilesReceiver
	CommandsReceiver
	Credentials
	HubClientInfoReceiver
	LocationsSender
	SessionsSender
	TariffsSender
	TokensReceiver
	Versions
}

type EMSP interface {
	CDRsReceiver
	CommandsSender
	Credentials
	HubClientInfoReceiver
	LocationsReceiver
	SessionsReceiver
	TariffsReceiver
	TokensSender
	Versions
}

type Hub interface {
	CDRsSender
	CDRsReceiver
	ChargingProfilesSender
	ChargingProfilesReceiver
	CommandsSender
	CommandsReceiver
	Credentials
	HubClientInfoSender
	LocationsSender
	LocationsReceiver
	SessionsSender
	SessionsReceiver
	TariffsSender
	TariffsReceiver
	TokensSender
	TokensReceiver
	Versions
}

type NSP interface {
	Credentials
	HubClientInfoReceiver
	LocationsReceiver
	TariffsReceiver
	Versions
}

type NAP interface {
	Credentials
	HubClientInfoReceiver
	LocationsSender
	LocationsReceiver
	TariffsSender
	TariffsReceiver
	Versions
}

type SCSP interface {
	ChargingProfilesSender
	Credentials
	HubClientInfoReceiver
	SessionsReceiver
	Versions
}

type Credentials interface {
	VerifyToken(ctx context.Context, token string) bool
	StoreVersionDetails(ctx context.Context, endpoints VersionDetails) error

	// (GET /ocpi/2.2.1/credentials)
	OnGetCredential(ctx context.Context, token string) (*Credential, error)
	// (POST /ocpi/2.2.1/credentials)
	OnPostCredential(ctx context.Context, token string, body ocpi.RawMessage[Credential]) (*Credential, error)
	// (PUT /ocpi/2.2.1/credentials)
	OnPutCredential(ctx context.Context, token string, body ocpi.RawMessage[Credential]) (*Credential, error)
	// (DELETE /ocpi/2.2.1/credentials)
	OnDeleteCredential(ctx context.Context, token string) error
}

type CDRsSender interface {
	// (GET /ocpi/2.2.1/cdrs)
	OnGetCDRs(ctx context.Context, params GetCdrsParams) (*ocpi.PaginationResponse[ChargeDetailRecord], error)
}
type CDRsReceiver interface {
	// (GET /ocpi/2.2.1/cdrs/{cdr_id})
	OnGetCDR(ctx context.Context, cdrID string) (*ChargeDetailRecord, error)
	// (POST /ocpi/2.2.1/cdrs)
	OnPostCDR(ctx context.Context, body ocpi.RawMessage[ChargeDetailRecord]) (*ChargeDetailRecordResponse, error)
}

type ChargingProfilesSender interface {
}
type ChargingProfilesReceiver interface {
	// (GET /ocpi/2.2.1/chargingprofiles/{session_id})
	OnGetChargingProfile(ctx context.Context, sessionID string) (*ChargingProfile, error)
	// (PUT /ocpi/2.2.1/chargingprofiles/{session_id})
	OnPutChargingProfile(ctx context.Context, sessionID string) (*ChargingProfileResponse, error)
	// (DELETE /ocpi/2.2.1/chargingprofiles/{session_id})
	OnDeleteChargingProfile(ctx context.Context, sessionID string) (*ChargingProfileResponse, error)
}

type CommandsSender interface {
	// (POST /ocpi/2.2.1/commands/{command}/{uid})
	OnPostAsyncCommand(ctx context.Context, commandType CommandType, uid string, body ocpi.RawMessage[CommandResult]) error
}
type CommandsReceiver interface {
	// (POST /ocpi/2.2.1/commands/{command})
	OnPostCommand(ctx context.Context, commandType CommandType) (*CommandResponse, error)
}

type HubClientInfoSender interface {
	// (GET /ocpi/2.2.1/hubclientinfo)
	OnGetHubClientInfo(ctx context.Context, params GetHubClientInfoParams) (*ocpi.PaginationResponse[ClientInfo], error)
}
type HubClientInfoReceiver interface {
	// (GET /ocpi/2.2.1/clientinfo/{id})
	OnGetHubClientInfo(ctx context.Context, countryCode string, partyID string) (*ClientInfo, error)
	// (PUT /ocpi/2.2.1/clientinfo/{id})
	OnPutHubClientInfo(ctx context.Context, countryCode string, partyID string, body ocpi.RawMessage[ClientInfo]) error
}

type LocationsSender interface {
	// (GET /ocpi/2.2.1/locations)
	OnGetLocations(ctx context.Context, params GetLocationsParams) (*ocpi.PaginationResponse[Location], error)
	// (GET /ocpi/2.2.1/locations/{location_id})
	OnGetLocation(ctx context.Context, locationID string) (*Location, error)
	// (GET /ocpi/2.2.1/locations/{location_id}/{evse_uid})
	OnGetLocationEVSE(ctx context.Context, locationID string, evseUID string) (*EVSE, error)
	// (GET /ocpi/2.2.1/locations/{location_id}/{evse_uid}/{connector_id})
	OnGetLocationConnector(ctx context.Context, locationID string, evseUID string, connectorID string) (*Connector, error)
}
type LocationsReceiver interface {
	// (GET /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id})
	OnGetClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string) (*Location, error)
	// (GET /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	OnGetClientOwnedLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string) (*EVSE, error)
	// (GET /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	OnGetClientOwnedLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string) (*Connector, error)
	// (PUT /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id})
	OnPutClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string, body ocpi.RawMessage[Location]) error
	// (PUT /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	OnPutClientOwnedLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, body ocpi.RawMessage[Location]) error
	// (PUT /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	OnPutClientOwnedLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string, body ocpi.RawMessage[Location]) error
	// (PATCH /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id})
	OnPatchClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string, body ocpi.RawMessage[PatchedLocation]) error
	// (PATCH /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	OnPatchClientOwnedLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, body ocpi.RawMessage[PatchedLocation]) error
	// (PATCH /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	OnPatchClientOwnedLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string, body ocpi.RawMessage[PatchedLocation]) error
}

type SessionsSender interface {
	// GetOcpiSessions retrieves a list of sessions based on the provided parameters.
	// (GET /ocpi/2.2.1/sessions)
	OnGetSessions(ctx context.Context, params GetSessionsParams) (*ocpi.PaginationResponse[Session], error)
	// (PUT /ocpi/2.2.1/sessions/{session_id}/charging_preferences)
	OnPutSessionChargingPreferences(ctx context.Context, sessionID string, body ocpi.RawMessage[ChargingPreferences]) (*ChargingPreferences, error)
}
type SessionsReceiver interface {
	// (GET /ocpi/2.2.1/sessions/{country_code}/{party_id}/{session_id})
	OnGetClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string) (*Session, error)
	// (PUT /ocpi/2.2.1/sessions/{country_code}/{party_id}/{session_id})
	OnPutClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, body ocpi.RawMessage[Session]) error
	// (PATCH /ocpi/2.2.1/sessions/{country_code}/{party_id}/{session_id})
	OnPatchClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, body ocpi.RawMessage[PatchedSession]) error
}

type TariffsSender interface {
	// GetOcpiTariffs retrieves a list of tariffs based on the provided parameters.
	// (GET /ocpi/2.2.1/tariffs)
	OnGetTariffs(ctx context.Context, params GetTariffsParams) (*ocpi.PaginationResponse[Tariff], error)
}

type TariffsReceiver interface {
	// (GET /ocpi/2.2.1/tariffs/{country_code}/{party_id}/{tariff_id})
	OnGetClientOwnedTariff(ctx context.Context, countryCode string, partyID string, sessionID string) (*Tariff, error)
	// (PUT /ocpi/2.2.1/tariffs/{country_code}/{party_id}/{tariff_id})
	OnPutClientOwnedTariff(ctx context.Context, countryCode string, partyID string, tariffID string, body ocpi.RawMessage[Tariff]) error
	// (DELETE /ocpi/2.2.1/tariffs/{country_code}/{party_id}/{tariff_id})
	OnDeleteClientOwnedTariff(ctx context.Context, countryCode string, partyID string, tariffID string) error
}

type TokensSender interface {
	// (GET /ocpi/2.2.1/tokens)
	OnGetTokens(ctx context.Context, params GetTokensParams) (*ocpi.PaginationResponse[Token], error)
	// (POST /ocpi/2.2.1/tokens/{token_uid}/authorize[?type={type}])
	OnPostToken(ctx context.Context, tokenUID string, body ocpi.RawMessage[LocationReferences], tokenType ...TokenType) (*AuthorizationInfo, error)
}
type TokensReceiver interface {
	// (GET /ocpi/2.2.1/tokens/{country_code}/{party_id}/{token_uid}[?type={type}])
	OnGetClientOwnedToken(ctx context.Context, countryCode string, partyID string, tokenUID string, tokenType ...TokenType) (*Token, error)
	// (PUT /ocpi/2.2.1/tokens/{country_code}/{party_id}/{token_uid}[?type={type}])
	OnPutClientOwnedToken(ctx context.Context, countryCode string, partyID string, tokenUID string, body ocpi.RawMessage[Token], tokenType ...TokenType) error
	// (PATCH /ocpi/2.2.1/tokens/{country_code}/{party_id}/{token_uid}[?type={type}])
	OnPatchClientOwnedToken(ctx context.Context, countryCode string, partyID string, tokenUID string, body ocpi.RawMessage[PatchedToken], tokenType ...TokenType) error
}

type Versions interface {
	VersionsSender
	VersionsReceiver
}
type VersionsSender interface {
}
type VersionsReceiver interface {
}
