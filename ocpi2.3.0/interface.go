package ocpi230

import (
	"context"

	"github.com/si3nloong/ocpi-go/ocpi"
)

type CPO interface {
	BookingsSender
	CDRsSender
	CDRsReceiver
	ChargingProfilesReceiver
	CommandsReceiver
	Credentials
	HubClientInfoReceiver
	LocationsSender
	SessionsSender
	TariffsSender
	TokensReceiver
	PaymentsReceiver
	Versions
}

type EMSP interface {
	BookingsReceiver
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

type RoamingHub interface {
	BookingsSender
	BookingsReceiver
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

type PTP interface {
	CDRsReceiver
	CommandsSender
	Credentials
	LocationsReceiver
	SessionsReceiver
	TariffsReceiver
	PaymentsSender
	Versions
}

type BookingsSender interface {
	OnGetBooking(ctx context.Context) (*ocpi.PaginationResponse[BookingLocation], error)
}
type BookingsReceiver interface {
}

type CDRsSender interface {
}
type CDRsReceiver interface {
}

type ChargingProfilesSender interface {
}
type ChargingProfilesReceiver interface {
}

type CommandsSender interface {
}
type CommandsReceiver interface {
}

type Credentials interface {
}

type HubClientInfoSender interface {
}
type HubClientInfoReceiver interface {
}

type LocationsSender interface {
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
	OnPutClientOwnedLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, body ocpi.RawMessage[Location]) error
	// (PUT /ocpi/2.3.0/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	OnPutClientOwnedLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string, body ocpi.RawMessage[Location]) error
	// (PATCH /ocpi/2.3.0/locations/{country_code}/{party_id}/{location_id})
	OnPatchClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string, body ocpi.RawMessage[PartialLocation]) error
	// (PATCH /ocpi/2.3.0/locations/{country_code}/{party_id}/{location_id}/{evse_uid})
	OnPatchClientOwnedLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, body ocpi.RawMessage[PartialLocation]) error
	// (PATCH /ocpi/2.3.0/locations/{country_code}/{party_id}/{location_id}/{evse_uid}/{connector_id})
	OnPatchClientOwnedLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string, body ocpi.RawMessage[PartialLocation]) error
}

type SessionsSender interface {
}
type SessionsReceiver interface {
}

type TariffsSender interface {
}
type TariffsReceiver interface {
}

type TokensSender interface {
}
type TokensReceiver interface {
}

type PaymentsSender interface {
}
type PaymentsReceiver interface {
}

type Versions interface {
	VersionsSender
	VersionsReceiver
}
type VersionsSender interface {
}
type VersionsReceiver interface {
}
