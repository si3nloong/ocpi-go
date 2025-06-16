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

type CDRsSender interface {
	// GetOcpiCDRs retrieves a list of charge detail records based on the provided parameters.
	// (GET /ocpi/2.2.1/cdrs)
	GetCDRs(ctx context.Context, params GetOcpiCdrsParams) ([]ChargeDetailRecord, error)
}
type CDRsReceiver interface {
	// (GET /ocpi/2.2.1/cdrs/{cdr_id})
	GetCDR(ctx context.Context, cdrId string) (*ChargeDetailRecord, error)
	// (POST /ocpi/2.2.1/cdrs)
	PostCDR(ctx context.Context, body ocpi.RawMessage[ChargeDetailRecord]) error
}

type ChargingProfilesSender interface {
}
type ChargingProfilesReceiver interface {
	// (GET /ocpi/2.2.1/chargingprofiles/{session_id})
	GetChargingProfile(ctx context.Context, sessionID string) (*ChargingProfile, error)
	// (PUT /ocpi/2.2.1/chargingprofiles/{session_id})
	PutChargingProfile(ctx context.Context, sessionID string) (*ChargingProfileResponse, error)
	// (DELETE /ocpi/2.2.1/chargingprofiles/{session_id})
	DeleteChargingProfile(ctx context.Context, sessionID string) (*ChargingProfileResponse, error)
}

type CommandsSender interface {
	// (POST /ocpi/2.2.1/commands/{command}/{uid})
	PostAsyncCommand(ctx context.Context, commandType CommandType, uid string, body ocpi.RawMessage[CommandResult]) (*CommandResponse, error)
}
type CommandsReceiver interface {
	// (POST /ocpi/2.2.1/commands/{command})
	PostCommand(ctx context.Context, commandType CommandType) (*CommandResponse, error)
}

type Credentials interface {
	VerifyToken(token string) bool
	GetCredential(ctx context.Context, tokenA string) (*Credential, error)
	PostCredential(ctx context.Context, tokenA string, body ocpi.RawMessage[Credential]) (*Credential, error)
	PutCredential(ctx context.Context, tokenA string, body ocpi.RawMessage[Credential]) (*Credential, error)
	DeleteCredential(ctx context.Context, tokenA string) (*Credential, error)
}

type HubClientInfoSender interface {
}
type HubClientInfoReceiver interface {
}

type LocationsSender interface {
	// GetOcpiLocations retrieves a list of locations based on the provided parameters.
	// (GET /ocpi/2.2.1/locations)
	GetLocations(ctx context.Context, params GetOcpiLocationsParams) ([]Location, error)
}
type LocationsReceiver interface {
	// (GET /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id})
	GetLocation(ctx context.Context, countryCode string, partyID string, locationID string) (*Location, error)
	// (GET /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	GetLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUid string) (*EVSE, error)
	// (GET /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	GetLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUid string, connectorID string) (*Connector, error)
	// (PUT /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id})
	PutLocation(ctx context.Context, countryCode string, partyID string, locationID string, body ocpi.RawMessage[Location]) error
	// (PUT /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	PutLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUid string, body ocpi.RawMessage[Location]) error
	// (PUT /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	PutLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUid string, connectorID string, body ocpi.RawMessage[Location]) error
	// (PATCH /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id})
	PatchLocation(ctx context.Context, countryCode string, partyID string, locationID string, body ocpi.RawMessage[PatchedLocation]) error
	// (PATCH /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	PatchLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUid string, body ocpi.RawMessage[PatchedLocation]) error
	// (PATCH /ocpi/2.2.1/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	PatchLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUid string, connectorID string, body ocpi.RawMessage[PatchedLocation]) error
}

type SessionsSender interface {
	// GetOcpiSessions retrieves a list of sessions based on the provided parameters.
	// (GET /ocpi/2.2.1/sessions)
	GetSessions(ctx context.Context, params GetOcpiSessionsParams) ([]Session, error)
}
type SessionsReceiver interface {
	// (GET /ocpi/2.2.1/sessions/{country_code}/{party_id}/{session_id})
	GetSession(ctx context.Context, countryCode string, partyID string, sessionId string) (*Session, error)
	// (PUT /ocpi/2.2.1/sessions/{country_code}/{party_id}/{session_id})
	PutSession(ctx context.Context, countryCode string, partyID string, sessionId string, body ocpi.RawMessage[Session]) error
	// (PATCH /ocpi/2.2.1/sessions/{country_code}/{party_id}/{session_id})
	PatchSession(ctx context.Context, countryCode string, partyID string, sessionId string, body ocpi.RawMessage[PatchedSession]) error
}

type TariffsSender interface {
}
type TariffsReceiver interface {
	// (GET /ocpi/2.2.1/tariffs/{country_code}/{party_id}/{tariff_id})
	GetTariff(ctx context.Context, countryCode string, partyID string, sessionId string) (*Tariff, error)
	// (PUT /ocpi/2.2.1/tariffs/{country_code}/{party_id}/{tariff_id})
	PutTariff(ctx context.Context, countryCode string, partyID string, tariffId string, body ocpi.RawMessage[Tariff]) error
	// (PATCH /ocpi/2.2.1/tariffs/{country_code}/{party_id}/{tariff_id})
	PatchTariff(ctx context.Context, countryCode string, partyID string, tariffId string, body ocpi.RawMessage[Tariff]) error
}

type TokensSender interface {
	// GetOcpiTokens retrieves a list of tokens based on the provided parameters.
	// (GET /ocpi/2.2.1/tokens)
	GetTokens(ctx context.Context, params GetOcpiTokensParams) ([]Token, error)
	// (POST /ocpi/2.2.1/tokens/{token_uid}/authorize[?type={type}])
	PostToken(ctx context.Context, tokenUID string, body ocpi.RawMessage[LocationReferences], tokenType ...TokenType) (*AuthorizationInfo, error)
}
type TokensReceiver interface {
	// (GET /ocpi/2.2.1/tokens/{country_code}/{party_id}/{token_uid}[?type={type}])
	GetToken(ctx context.Context, countryCode string, partyID string, tokenUID string, tokenType ...TokenType) (*Token, error)
	// (PUT /ocpi/2.2.1/tokens/{country_code}/{party_id}/{token_uid}[?type={type}])
	PutToken(ctx context.Context, countryCode string, partyID string, tokenUID string, body ocpi.RawMessage[Token], tokenType ...TokenType) error
	// (PATCH /ocpi/2.2.1/tokens/{country_code}/{party_id}/{token_uid}[?type={type}])
	PatchToken(ctx context.Context, countryCode string, partyID string, tokenUID string, body ocpi.RawMessage[PatchedToken], tokenType ...TokenType) error
}

type Versions interface {
	VersionsSender
	VersionsReceiver
}
type VersionsSender interface {
}
type VersionsReceiver interface {
}
