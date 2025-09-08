package ocpi230

import (
	"context"

	"github.com/si3nloong/ocpi-go/ocpi"
)

type CPO interface {
	CDRsSender
	ChargingProfilesReceiver
	CommandsReceiver
	HubClientInfoReceiver
	LocationsSender
	SessionsSender
	TariffsSender
	TokensReceiver
	PaymentsReceiver
}

type EMSP interface {
	CDRsReceiver
	CommandsSender
	HubClientInfoReceiver
	LocationsReceiver
	SessionsReceiver
	TariffsReceiver
	TokensSender
}

type Hub interface {
	CDRsSender
	CDRsReceiver
	ChargingProfilesSender
	ChargingProfilesReceiver
	CommandsSender
	CommandsReceiver
	HubClientInfoSender
	LocationsSender
	LocationsReceiver
	SessionsSender
	SessionsReceiver
	TariffsSender
	TariffsReceiver
	TokensSender
	TokensReceiver
}

type NSP interface {
	HubClientInfoReceiver
	LocationsReceiver
	TariffsReceiver
}

type NAP interface {
	HubClientInfoReceiver
	LocationsSender
	LocationsReceiver
	TariffsSender
	TariffsReceiver
}

type SCSP interface {
	ChargingProfilesSender
	HubClientInfoReceiver
	SessionsReceiver
}

type PTP interface {
	CDRsReceiver
	CommandsSender
	LocationsReceiver
	SessionsReceiver
	TariffsReceiver
	PaymentsSender
}

type CredentialsReceiver interface {
	// (GET /ocpi/2.3.0/credentials)
	OnGetCredential(ctx context.Context, tokenC string) (*Credentials, error)
	// (POST /ocpi/2.3.0/credentials)
	OnPostCredential(ctx context.Context, tokenA string, body ocpi.RawMessage[Credentials]) (*Credentials, error)
	// (PUT /ocpi/2.3.0/credentials)
	OnPutCredential(ctx context.Context, tokenC string, body ocpi.RawMessage[Credentials]) (*Credentials, error)
	// (DELETE /ocpi/2.3.0/credentials)
	OnDeleteCredential(ctx context.Context, tokenC string) error
}

type CDRsSender interface {
	// (GET /ocpi/2.3.0/cdrs)
	OnGetCDRs(ctx context.Context, params GetCDRsParams) (*ocpi.PaginationResponse[CDR], error)
}
type CDRsReceiver interface {
	// (GET /ocpi/2.3.0/cdrs/{cdr_id})
	OnGetCDR(ctx context.Context, cdrID string) (*CDR, error)
	// (POST /ocpi/2.3.0/cdrs)
	OnPostCDR(ctx context.Context, body ocpi.RawMessage[CDR]) (*ChargeDetailRecordResponse, error)
}

type ChargingProfilesSender interface {
	// (POST /ocpi/2.3.0/activechargingprofile/{session_id})
	OnPostActiveChargingProfile(ctx context.Context, sessionID string, body ocpi.RawMessage[ActiveChargingProfileResult]) error
	// (POST /ocpi/2.3.0/chargingprofiles/chargingprofile/{session_id})
	OnPostChargingProfile(ctx context.Context, sessionID string, body ocpi.RawMessage[ChargingProfileResult]) error
	// (POST /ocpi/2.3.0/clearprofile/{session_id})
	OnPostClearProfile(ctx context.Context, sessionID string, body ocpi.RawMessage[ClearProfileResult]) error
	// (PUT /ocpi/2.3.0/clearprofile/{session_id})
	OnPutActiveChargingProfile(ctx context.Context, sessionID string, body ocpi.RawMessage[ActiveChargingProfile]) error
}
type ChargingProfilesReceiver interface {
	// (GET /ocpi/2.3.0/chargingprofiles/{session_id}?duration={duration}&response_url={url})
	OnGetActiveChargingProfile(ctx context.Context, sessionID string, duration int, responseURL string) (*ChargingProfileResponse, error)
	// (PUT /ocpi/2.3.0/chargingprofiles/{session_id})
	OnPutChargingProfile(ctx context.Context, sessionID string, body ocpi.RawMessage[SetChargingProfile]) (*ChargingProfileResponse, error)
	// (DELETE /ocpi/2.3.0/chargingprofiles/{session_id}?response_url={url})
	OnDeleteChargingProfile(ctx context.Context, sessionID string, responseURL string) (*ChargingProfileResponse, error)
}

type CommandsSender interface {
	// (POST /ocpi/2.3.0/commands/{command}/{uid})
	OnPostAsyncCommand(ctx context.Context, commandType CommandType, uid string, body ocpi.RawMessage[CommandResult]) error
}
type CommandsReceiver interface {
	// (POST /ocpi/2.3.0/commands/{command})
	OnPostCommand(ctx context.Context, commandType CommandType, body CommandRequest) (*CommandResponse, error)
}

type HubClientInfoSender interface {
	// (GET /ocpi/2.3.0/hubclientinfo)
	OnGetHubClientInfos(ctx context.Context, params GetHubClientInfoParams) (*ocpi.PaginationResponse[ClientInfo], error)
}
type HubClientInfoReceiver interface {
	// (GET /ocpi/2.3.0/clientinfo/{id})
	OnGetHubClientInfo(ctx context.Context, countryCode string, partyID string) (*ClientInfo, error)
	// (PUT /ocpi/2.3.0/clientinfo/{id})
	OnPutHubClientInfo(ctx context.Context, countryCode string, partyID string, body ocpi.RawMessage[ClientInfo]) error
}

type LocationsSender interface {
	// (GET /ocpi/2.3.0/locations)
	OnGetLocations(ctx context.Context, params GetLocationsParams) (*ocpi.PaginationResponse[Location], error)
	// (GET /ocpi/2.3.0/locations/{location_id})
	OnGetLocation(ctx context.Context, locationID string) (*Location, error)
	// (GET /ocpi/2.3.0/locations/{location_id}/{evse_uid})
	OnGetLocationEVSE(ctx context.Context, locationID string, evseUID string) (*EVSE, error)
	// (GET /ocpi/2.3.0/locations/{location_id}/{evse_uid}/{connector_id})
	OnGetLocationConnector(ctx context.Context, locationID string, evseUID string, connectorID string) (*Connector, error)
}
type LocationsReceiver interface {
	// (GET /ocpi/2.3.0/locations/{country_code}/{party_id}/{location_id})
	OnGetClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string) (*Location, error)
	// (GET /ocpi/2.3.0/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	OnGetClientOwnedLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string) (*EVSE, error)
	// (GET /ocpi/2.3.0/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	OnGetClientOwnedLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string) (*Connector, error)
	// (PUT /ocpi/2.3.0/locations/{country_code}/{party_id}/{location_id})
	OnPutClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string, body ocpi.RawMessage[Location]) error
	// (PUT /ocpi/2.3.0/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	OnPutClientOwnedLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, body ocpi.RawMessage[EVSE]) error
	// (PUT /ocpi/2.3.0/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	OnPutClientOwnedLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string, body ocpi.RawMessage[Connector]) error
	// (PATCH /ocpi/2.3.0/locations/{country_code}/{party_id}/{location_id})
	OnPatchClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string, body ocpi.RawMessage[PartialLocation]) error
	// (PATCH /ocpi/2.3.0/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	OnPatchClientOwnedLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, body ocpi.RawMessage[PartialEVSE]) error
	// (PATCH /ocpi/2.3.0/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	OnPatchClientOwnedLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string, body ocpi.RawMessage[PartialConnector]) error
}

type SessionsSender interface {
	// GetOcpiSessions retrieves a list of sessions based on the provided parameters.
	// (GET /ocpi/2.3.0/sessions)
	OnGetSessions(ctx context.Context, params GetSessionsParams) (*ocpi.PaginationResponse[Session], error)
	// (PUT /ocpi/2.3.0/sessions/{session_id}/charging_preferences)
	OnPutSessionChargingPreferences(ctx context.Context, sessionID string, body ocpi.RawMessage[ChargingPreferences]) (*ChargingPreferences, error)
}
type SessionsReceiver interface {
	// (GET /ocpi/2.3.0/sessions/{country_code}/{party_id}/{session_id})
	OnGetClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string) (*Session, error)
	// (PUT /ocpi/2.3.0/sessions/{country_code}/{party_id}/{session_id})
	OnPutClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, body ocpi.RawMessage[Session]) error
	// (PATCH /ocpi/2.3.0/sessions/{country_code}/{party_id}/{session_id})
	OnPatchClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, body ocpi.RawMessage[PartialSession]) error
}

type TariffsSender interface {
	// GetOcpiTariffs retrieves a list of tariffs based on the provided parameters.
	// (GET /ocpi/2.3.0/tariffs)
	OnGetTariffs(ctx context.Context, params GetTariffsParams) (*ocpi.PaginationResponse[Tariff], error)
}

type TariffsReceiver interface {
	// (GET /ocpi/2.3.0/tariffs/{country_code}/{party_id}/{tariff_id})
	OnGetClientOwnedTariff(ctx context.Context, countryCode string, partyID string, sessionID string) (*Tariff, error)
	// (PUT /ocpi/2.3.0/tariffs/{country_code}/{party_id}/{tariff_id})
	OnPutClientOwnedTariff(ctx context.Context, countryCode string, partyID string, tariffID string, body ocpi.RawMessage[Tariff]) error
	// (DELETE /ocpi/2.3.0/tariffs/{country_code}/{party_id}/{tariff_id})
	OnDeleteClientOwnedTariff(ctx context.Context, countryCode string, partyID string, tariffID string) error
}

type TokensSender interface {
	// (GET /ocpi/2.3.0/tokens)
	OnGetTokens(ctx context.Context, params GetTokensParams) (*ocpi.PaginationResponse[Token], error)
	// (POST /ocpi/2.3.0/tokens/{token_uid}/authorize[?type={type}])
	OnPostToken(ctx context.Context, tokenUID string, body ocpi.RawMessage[LocationReferences], tokenType ...TokenType) (*AuthorizationInfo, error)
}
type TokensReceiver interface {
	// (GET /ocpi/2.3.0/tokens/{country_code}/{party_id}/{token_uid}[?type={type}])
	OnGetClientOwnedToken(ctx context.Context, countryCode string, partyID string, tokenUID string, tokenType ...TokenType) (*Token, error)
	// (PUT /ocpi/2.3.0/tokens/{country_code}/{party_id}/{token_uid}[?type={type}])
	OnPutClientOwnedToken(ctx context.Context, countryCode string, partyID string, tokenUID string, body ocpi.RawMessage[Token], tokenType ...TokenType) error
	// (PATCH /ocpi/2.3.0/tokens/{country_code}/{party_id}/{token_uid}[?type={type}])
	OnPatchClientOwnedToken(ctx context.Context, countryCode string, partyID string, tokenUID string, body ocpi.RawMessage[PartialToken], tokenType ...TokenType) error
}

type VersionsSender interface {
}
type VersionsReceiver interface {
	// (GET /ocpi/2.3.0/details)
	OnVersionDetails(ctx context.Context, token string) ([]Endpoint, error)
}

type PaymentsSender interface {
	PaymentTerminalsSender
	FinancialAdviceConfirmationSender
}

type PaymentTerminalsSender interface {
	// (GET /ocpi/ptp/2.3.0/payments/terminals)
	OnGetPaymentTerminals(ctx context.Context, params GetPaymentTerminalsParams) (*ocpi.PaginationResponse[Terminal], error)
	// (GET /ocpi/ptp/2.3.0/payments/terminals/{terminal_id})
	OnGetPaymentTerminal(ctx context.Context, terminalID string) (*Terminal, error)
	// (POST /ocpi/ptp/2.3.0/payments/terminals/activate)
	OnPostActivatePaymentTerminal(ctx context.Context, body ocpi.RawMessage[Terminal]) error
	// (POST /ocpi/ptp/2.3.0/payments/terminals/deactivate)
	OnPostDeactivatePaymentTerminal(ctx context.Context, terminalID string) error
	// (PUT /ocpi/ptp/2.3.0/payments/terminals/{terminal_id})
	OnPutPaymentTerminal(ctx context.Context, body ocpi.RawMessage[Terminal]) error
	// (PATCH /ocpi/ptp/2.3.0/payments/terminals/{terminal_id})
	OnPatchPaymentTerminal(ctx context.Context, body ocpi.RawMessage[PartialTerminal]) error
}

type FinancialAdviceConfirmationSender interface {
	// (GET /ocpi/2.3.0/payments/financial-advice-confirmations)
	OnGetPaymentFinancialAdviceConfirmations(ctx context.Context, params GetPaymentFinancialAdviceConfirmationsParams) (*ocpi.PaginationResponse[FinancialAdviceConfirmation], error)
	// (GET /ocpi/2.3.0/payments/financial-advice-confirmations/{financial_advice_confirmation_id})
	OnGetPaymentFinancialAdviceConfirmation(ctx context.Context, financialAdviceConfirmationID string) (*FinancialAdviceConfirmation, error)
}

type PaymentsReceiver interface {
	PaymentTerminalsReceiver
	FinancialAdviceConfirmationReceiver
}

type PaymentTerminalsReceiver interface {
	// (GET /ocpi/2.3.0/payments/terminals/{terminal_id})
	OnGetTerminal(ctx context.Context, terminalID string) (*Terminal, error)
	// (POST /ocpi/2.3.0/payments/terminals)
	OnPostTerminal(ctx context.Context, body ocpi.RawMessage[Terminal]) error
}

type FinancialAdviceConfirmationReceiver interface {
	// (GET /ocpi/2.3.0/payments/financial-advice-confirmations/{financial_advice_confirmation_id})
	OnGetFinancialAdviceConfirmation(ctx context.Context, financialAdviceConfirmationID string) (*FinancialAdviceConfirmation, error)
	// (POST /ocpi/2.3.0/payments/financial-advice-confirmations/{financial_advice_confirmation_id})
	OnPostFinancialAdviceConfirmation(ctx context.Context, body ocpi.RawMessage[FinancialAdviceConfirmation]) error
}
