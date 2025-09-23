package ocpi221

import (
	"encoding/json"

	"github.com/si3nloong/ocpi-go/ocpi"
)

type Role string

// Defines values for Role.
const (
	RoleCPO   Role = "CPO"
	RoleEMSP  Role = "EMSP"
	RoleHUB   Role = "HUB"
	RoleNAP   Role = "NAP"
	RoleNSP   Role = "NSP"
	RoleOther Role = "OTHER"
	RoleSCSP  Role = "SCSP"
)

// InterfaceRole Interface role endpoint implements.
type InterfaceRole string

// Defines values for InterfaceRole.
const (
	InterfaceRoleReceiver InterfaceRole = "RECEIVER"
	InterfaceRoleSender   InterfaceRole = "SENDER"
)

// ModuleID OCPI 2.2.1 modules
type ModuleID string

// Defines values for ModuleID.
const (
	ModuleIDCdrs             ModuleID = "cdrs"
	ModuleIDChargingProfiles ModuleID = "chargingprofiles"
	ModuleIDCommands         ModuleID = "commands"
	ModuleIDCredentials      ModuleID = "credentials"
	ModuleIDHubClientInfo    ModuleID = "hubclientinfo"
	ModuleIDLocations        ModuleID = "locations"
	ModuleIDSessions         ModuleID = "sessions"
	ModuleIDTariffs          ModuleID = "tariffs"
	ModuleIDTokens           ModuleID = "tokens"
)

// ChargingProfileResultType defines model for ActiveChargingProfileResult.Result.
type ChargingProfileResultType string

// Defines values for ChargingProfileResultType.
const (
	ChargingProfileResultTypeAccepted ChargingProfileResultType = "ACCEPTED"
	ChargingProfileResultTypeRejected ChargingProfileResultType = "REJECTED"
	ChargingProfileResultTypeUknown   ChargingProfileResultType = "UNKNOWN"
)

// AuthorizationAllowed defines model for Authorization.Allowed.
type AuthorizationAllowed string

// Defines values for AuthorizationAllowed.
const (
	AuthorizationAllowedAllowed    AuthorizationAllowed = "ALLOWED"
	AuthorizationAllowedBlocked    AuthorizationAllowed = "BLOCKED"
	AuthorizationAllowedExpired    AuthorizationAllowed = "EXPIRED"
	AuthorizationAllowedNoCredit   AuthorizationAllowed = "NO_CREDIT"
	AuthorizationAllowedNotAllowed AuthorizationAllowed = "NOT_ALLOWED"
)

// CdrDimensionType defines model for CdrBodyChargingPeriodsDimensions.Type.
type CdrDimensionType string

// Defines values for CdrDimensionType.
const (
	CdrDimensionTypeCurrent         CdrDimensionType = "CURRENT"
	CdrDimensionTypeEnergy          CdrDimensionType = "ENERGY"
	CdrDimensionTypeEnergyExport    CdrDimensionType = "ENERGY_EXPORT"
	CdrDimensionTypeEnergyImport    CdrDimensionType = "ENERGY_IMPORT"
	CdrDimensionTypeMaxCurrent      CdrDimensionType = "MAX_CURRENT"
	CdrDimensionTypeMaxPower        CdrDimensionType = "MAX_POWER"
	CdrDimensionTypeMinCurrent      CdrDimensionType = "MIN_CURRENT"
	CdrDimensionTypeMinPower        CdrDimensionType = "MIN_POWER"
	CdrDimensionTypeParkingTime     CdrDimensionType = "PARKING_TIME"
	CdrDimensionTypePower           CdrDimensionType = "POWER"
	CdrDimensionTypeReservationTime CdrDimensionType = "RESERVATION_TIME"
	CdrDimensionTypeStateOfCharge   CdrDimensionType = "STATE_OF_CHARGE"
	CdrDimensionTypeTime            CdrDimensionType = "TIME"
)

// EnergySourceCategory defines model for EnergySource.Source.
type EnergySourceCategory string

// Defines values for EnergySourceCategory.
const (
	EnergySourceCategoryNuclear       EnergySourceCategory = "NUCLEAR"
	EnergySourceCategoryGeneralFossil EnergySourceCategory = "GENERAL_FOSSIL"
	EnergySourceCategoryCoal          EnergySourceCategory = "COAL"
	EnergySourceCategoryGas           EnergySourceCategory = "GAS"
	EnergySourceCategoryGeneralGreen  EnergySourceCategory = "GENERAL_GREEN"
	EnergySourceCategorySolar         EnergySourceCategory = "SOLAR"
	EnergySourceCategoryWind          EnergySourceCategory = "WIND"
	EnergySourceCategoryWater         EnergySourceCategory = "WATER"
)

// EnvironmentalImpactCategory defines model for EnvironmentalImpact.Category.
type EnvironmentalImpactCategory string

// Defines values for EnvironmentalImpactCategory.
const (
	EnvironmentalImpactCategoryCarbonDioxide EnvironmentalImpactCategory = "CARBON_DIOXIDE"
	EnvironmentalImpactCategoryNuclearWaste  EnvironmentalImpactCategory = "NUCLEAR_WASTE"
)

// ChargingRateUnit defines model for ChargingProfile.ChargingRateUnit.
type ChargingRateUnit string

// Defines values for ChargingRateUnit.
const (
	ChargingRateUnitAmperes ChargingRateUnit = "A"
	ChargingRateUnitWatts   ChargingRateUnit = "W"
)

// ChargingProfileResponseType Response to the ChargingProfile request from the eMSP to the CPO.
type ChargingProfileResponseType string

// Defines values for ChargingProfileResponseType.
const (
	ChargingProfileResponseTypeAccepted       ChargingProfileResponseType = "ACCEPTED"
	ChargingProfileResponseTypeNotSupported   ChargingProfileResponseType = "NOT_SUPPORTED"
	ChargingProfileResponseTypeRejected       ChargingProfileResponseType = "REJECTED"
	ChargingProfileResponseTypeTooOften       ChargingProfileResponseType = "TOO_OFTEN"
	ChargingProfileResponseTypeUnknownSession ChargingProfileResponseType = "UNKNOWN_SESSION"
)

// ConnectionStatus defines model for ClientInfo.Status.
type ConnectionStatus string

// Defines values for ConnectionStatus.
const (
	ConnectionStatusConnected ConnectionStatus = "CONNECTED"
	ConnectionStatusOffline   ConnectionStatus = "OFFLINE"
	ConnectionStatusPlanned   ConnectionStatus = "PLANNED"
	ConnectionStatusSuspended ConnectionStatus = "SUSPENDED"
)

// CommandType defines parameters for type of commands.
type CommandType string

// Defines values for PostOcpiCommandsCommandParamsCommand.
const (
	CommandTypeCancelReservation CommandType = "CANCEL_RESERVATION"
	CommandTypeReserveNow        CommandType = "RESERVE_NOW"
	CommandTypeStartSession      CommandType = "START_SESSION"
	CommandTypeStopSession       CommandType = "STOP_SESSION"
	CommandTypeUnlockConnector   CommandType = "UNLOCK_CONNECTOR"
)

// CommandResponseType defines model for CommandResponse.Result.
type CommandResponseType string

// Defines values for CommandResponseType.
const (
	CommandResponseTypeAccepted       CommandResponseType = "ACCEPTED"
	CommandResponseTypeNotSupported   CommandResponseType = "NOT_SUPPORTED"
	CommandResponseTypeRejected       CommandResponseType = "REJECTED"
	CommandResponseTypeUnknownSession CommandResponseType = "UNKNOWN_SESSION"
)

// Defines values for CommandResultType.
const (
	CommandResultTypeAccepted            CommandResultType = "ACCEPTED"
	CommandResultTypeCanceledReservation CommandResultType = "CANCELED_RESERVATION"
	CommandResultTypeEVSEInOperative     CommandResultType = "EVSE_INOPERATIVE"
	CommandResultTypeEVSEOccupied        CommandResultType = "EVSE_OCCUPIED"
	CommandResultTypeFailed              CommandResultType = "FAILED"
	CommandResultTypeNotSupported        CommandResultType = "NOT_SUPPORTED"
	CommandResultTypeRejected            CommandResultType = "REJECTED"
	CommandResultTypeTimeout             CommandResultType = "TIMEOUT"
	CommandResultTypeUnknownReservation  CommandResultType = "UNKNOWN_RESERVATION"
)

// ConnectorFormat defines model for Connector.Format.
type ConnectorFormat string

// Defines values for ConnectorFormat.
const (
	ConnectorFormatCable  ConnectorFormat = "CABLE"
	ConnectorFormatSocket ConnectorFormat = "SOCKET"
)

// PowerType defines model for Connector.PowerType.
type PowerType string

// Defines values for PowerType.
const (
	PowerTypeAC1Phase      PowerType = "AC_1_PHASE"
	PowerTypeAC2Phase      PowerType = "AC_2_PHASE"
	PowerTypeAC2PhaseSplit PowerType = "AC_2_PHASE_SPLIT"
	PowerTypeAC3Phase      PowerType = "AC_3_PHASE"
	PowerTypeDC            PowerType = "DC"
)

// ConnectorType defines model for Connector.Standard.
type ConnectorType string

// Defines values for ConnectorType.
const (
	ConnectorTypeCHAdeMO            ConnectorType = "CHADEMO"
	ConnectorTypeChaoJi             ConnectorType = "CHAOJI"
	ConnectorTypeDomesticA          ConnectorType = "DOMESTIC_A"
	ConnectorTypeDomesticB          ConnectorType = "DOMESTIC_B"
	ConnectorTypeDomesticC          ConnectorType = "DOMESTIC_C"
	ConnectorTypeDomesticD          ConnectorType = "DOMESTIC_D"
	ConnectorTypeDomesticE          ConnectorType = "DOMESTIC_E"
	ConnectorTypeDomesticF          ConnectorType = "DOMESTIC_F"
	ConnectorTypeDomesticG          ConnectorType = "DOMESTIC_G"
	ConnectorTypeDomesticH          ConnectorType = "DOMESTIC_H"
	ConnectorTypeDomesticJ          ConnectorType = "DOMESTIC_J"
	ConnectorTypeDomesticK          ConnectorType = "DOMESTIC_K"
	ConnectorTypeDomesticL          ConnectorType = "DOMESTIC_L"
	ConnectorTypeDomesticM          ConnectorType = "DOMESTIC_M"
	ConnectorTypeDomesticN          ConnectorType = "DOMESTIC_N"
	ConnectorTypeDomesticO          ConnectorType = "DOMESTIC_O"
	ConnectorTypeGBTAC              ConnectorType = "GBT_AC"
	ConnectorTypeGBTDC              ConnectorType = "GBT_DC"
	ConnectorTypeIEC603092Single16  ConnectorType = "IEC_60309_2_single_16"
	ConnectorTypeIEC603092Three16   ConnectorType = "IEC_60309_2_three_16"
	ConnectorTypeIEC603092Three32   ConnectorType = "IEC_60309_2_three_32"
	ConnectorTypeIEC603092Three64   ConnectorType = "IEC_60309_2_three_64"
	ConnectorTypeIEC62196T1         ConnectorType = "IEC_62196_T1"
	ConnectorTypeIEC62196T1Combo    ConnectorType = "IEC_62196_T1_COMBO"
	ConnectorTypeIEC62196T2         ConnectorType = "IEC_62196_T2"
	ConnectorTypeIEC62196T2Combo    ConnectorType = "IEC_62196_T2_COMBO"
	ConnectorTypeIEC62196T3A        ConnectorType = "IEC_62196_T3A"
	ConnectorTypeIEC62196T3C        ConnectorType = "IEC_62196_T3C"
	ConnectorTypeNema1030           ConnectorType = "NEMA_10_30"
	ConnectorTypeNema1050           ConnectorType = "NEMA_10_50"
	ConnectorTypeNema1430           ConnectorType = "NEMA_14_30"
	ConnectorTypeNema1450           ConnectorType = "NEMA_14_50"
	ConnectorTypeNema520            ConnectorType = "NEMA_5_20"
	ConnectorTypeNema630            ConnectorType = "NEMA_6_30"
	ConnectorTypeNema650            ConnectorType = "NEMA_6_50"
	ConnectorTypePantographBottomUp ConnectorType = "PANTOGRAPH_BOTTOM_UP"
	ConnectorTypePantographTopDown  ConnectorType = "PANTOGRAPH_TOP_DOWN"
	ConnectorTypeTeslaR             ConnectorType = "TESLA_R"
	ConnectorTypeTeslaS             ConnectorType = "TESLA_S"
)

// Defines values for EvseCapabilities.
const (
	CapabilityChargingPreferencesCapable    Capability = "CHARGING_PREFERENCES_CAPABLE"
	CapabilityChargingProfileCapable        Capability = "CHARGING_PROFILE_CAPABLE"
	CapabilityChipCardSupport               Capability = "CHIP_CARD_SUPPORT"
	CapabilityContactlessCardSupport        Capability = "CONTACTLESS_CARD_SUPPORT"
	CapabilityCreditCardPayable             Capability = "CREDIT_CARD_PAYABLE"
	CapabilityDebitCardPayable              Capability = "DEBIT_CARD_PAYABLE"
	CapabilityPEDTerminal                   Capability = "PED_TERMINAL"
	CapabilityRemoteStartStopCapable        Capability = "REMOTE_START_STOP_CAPABLE"
	CapabilityReservable                    Capability = "RESERVABLE"
	CapabilityRFIDReader                    Capability = "RFID_READER"
	CapabilityStartSessionConnectorRequired Capability = "START_SESSION_CONNECTOR_REQUIRED"
	CapabilityTokenGroupCapable             Capability = "TOKEN_GROUP_CAPABLE"
	CapabilityUnlockCapable                 Capability = "UNLOCK_CAPABLE"
)

// ParkingRestriction defines model for Evse.ParkingRestriction.
type ParkingRestriction string

// Defines values for ParkingRestriction.
const (
	ParkingRestrictionCustomers   ParkingRestriction = "CUSTOMERS"
	ParkingRestrictionDisabled    ParkingRestriction = "DISABLED"
	ParkingRestrictionEmployees   ParkingRestriction = "EMPLOYEES"
	ParkingRestrictionEVOnly      ParkingRestriction = "EV_ONLY"
	ParkingRestrictionMotorcycles ParkingRestriction = "MOTORCYCLES"
	ParkingRestrictionPlugged     ParkingRestriction = "PLUGGED"
	ParkingRestrictionTaxis       ParkingRestriction = "TAXIS"
	ParkingRestrictionTenants     ParkingRestriction = "TENANTS"
)

// Status defines model for Evse.Status.
type Status string

// Defines values for Status.
const (
	StatusAvailable   Status = "AVAILABLE"
	StatusBlocked     Status = "BLOCKED"
	StatusCharging    Status = "CHARGING"
	StatusInOperative Status = "INOPERATIVE"
	StatusOutOfOrder  Status = "OUTOFORDER"
	StatusPlanned     Status = "PLANNED"
	StatusRemoved     Status = "REMOVED"
	StatusReserved    Status = "RESERVED"
	StatusUnknown     Status = "UNKNOWN"
)

// ImageCategory defines model for Image.Category.
type ImageCategory string

// Defines values for ImageCategory.
const (
	ImageCategoryCharger  ImageCategory = "CHARGER"
	ImageCategoryEntrance ImageCategory = "ENTRANCE"
	ImageCategoryLocation ImageCategory = "LOCATION"
	ImageCategoryNetwork  ImageCategory = "NETWORK"
	ImageCategoryOperator ImageCategory = "OPERATOR"
	ImageCategoryOther    ImageCategory = "OTHER"
	ImageCategoryOwner    ImageCategory = "OWNER"
)

// LocationsDataFacilities defines model for LocationsData.Facilities.
type Facility string

// Defines values for LocationsDataFacilities.
const (
	FacilityAirport        Facility = "AIRPORT"
	FacilityBikeSharing    Facility = "BIKE_SHARING"
	FacilityBusStop        Facility = "BUS_STOP"
	FacilityCafe           Facility = "CAFE"
	FacilityCarpoolParking Facility = "CARPOOL_PARKING"
	FacilityFuelStation    Facility = "FUEL_STATION"
	FacilityHotel          Facility = "HOTEL"
	FacilityMall           Facility = "MALL"
	FacilityMetroStation   Facility = "METRO_STATION"
	FacilityMuseum         Facility = "MUSEUM"
	FacilityNature         Facility = "NATURE"
	FacilityParkingLot     Facility = "PARKING_LOT"
	FacilityRecreationArea Facility = "RECREATION_AREA"
	FacilityRestaurant     Facility = "RESTAURANT"
	FacilitySport          Facility = "SPORT"
	FacilitySuperMarket    Facility = "SUPERMARKET"
	FacilityTaxiStand      Facility = "TAXI_STAND"
	FacilityTrainStation   Facility = "TRAIN_STATION"
	FacilityTramStop       Facility = "TRAM_STOP"
	FacilityWiFi           Facility = "WIFI"
)

// ParkingType defines model for LocationsData.ParkingType.
type ParkingType string

// Defines values for ParkingType.
const (
	ParkingTypeAlongMotorway     ParkingType = "ALONG_MOTORWAY"
	ParkingTypeOnDriveway        ParkingType = "ON_DRIVEWAY"
	ParkingTypeOnStreet          ParkingType = "ON_STREET"
	ParkingTypeParkingGarage     ParkingType = "PARKING_GARAGE"
	ParkingTypeParkingLot        ParkingType = "PARKING_LOT"
	ParkingTypeUndergroundGarage ParkingType = "UNDERGROUND_GARAGE"
)

// ReservationRestrictionType defines model for reservationRestrictionType.
type ReservationRestrictionType string

// Defines values for ReservationRestrictionType.
const (
	ReservationRestrictionTypeReservation        ReservationRestrictionType = "RESERVATION"
	ReservationRestrictionTypeReservationExpires ReservationRestrictionType = "RESERVATION_EXPIRES"
)

// AuthMethod defines model for CdrBody.AuthMethod.
type AuthMethod string

// Defines values for AuthMethod.
const (
	AuthMethodAuthRequest AuthMethod = "AUTH_REQUEST"
	AuthMethodCommand     AuthMethod = "COMMAND"
	AuthMethodWhitelist   AuthMethod = "WHITELIST"
)

// SessionStatus defines model for Session.Status.
type SessionStatus string

// Defines values for SessionStatus.
const (
	SessionStatusActive      SessionStatus = "ACTIVE"
	SessionStatusCompleted   SessionStatus = "COMPLETED"
	SessionStatusInvalid     SessionStatus = "INVALID"
	SessionStatusPending     SessionStatus = "PENDING"
	SessionStatusReservation SessionStatus = "RESERVATION"
)

// TariffType defines model for Tariff.Type.
type TariffType string

// Defines values for TariffType.
const (
	TariffTypeAdHocPayment TariffType = "AD_HOC_PAYMENT"
	TariffTypeProfileCheap TariffType = "PROFILE_CHEAP"
	TariffTypeProfileFast  TariffType = "PROFILE_FAST"
	TariffTypeProfileGreen TariffType = "PROFILE_GREEN"
	TariffTypeRegular      TariffType = "REGULAR"
)

type TariffDimensionType string

// Defines values for PriceComponent.
const (
	TariffDimensionTypeEnergy      TariffDimensionType = "ENERGY"
	TariffDimensionTypeFlat        TariffDimensionType = "FLAT"
	TariffDimensionTypeParkingTime TariffDimensionType = "PARKING_TIME"
	TariffDimensionTypeTime        TariffDimensionType = "TIME"
)

// DayOfWeek defines model for TariffRestrictions.DayOfWeek.
type DayOfWeek string

// Defines values for DayOfWeek.
const (
	DayOfWeekMonday    DayOfWeek = "MONDAY"
	DayOfWeekTuesday   DayOfWeek = "TUESDAY"
	DayOfWeekWednesday DayOfWeek = "WEDNESDAY"
	DayOfWeekThursday  DayOfWeek = "THURSDAY"
	DayOfWeekFriday    DayOfWeek = "FRIDAY"
	DayOfWeekSaturday  DayOfWeek = "SATURDAY"
	DayOfWeekSunday    DayOfWeek = "SUNDAY"
)

// ProfileType defines model for Token.DefaultProfileType.
type ProfileType string

// Defines values for ProfileType.
const (
	ProfileTypeCheap   ProfileType = "CHEAP"
	ProfileTypeFast    ProfileType = "FAST"
	ProfileTypeGreen   ProfileType = "GREEN"
	ProfileTypeRegular ProfileType = "REGULAR"
)

// TokenType defines model for Token.Type.
type TokenType string

// Defines values for TokenType.
const (
	TokenTypeAdHocUser TokenType = "AD_HOC_USER"
	TokenTypeAppUser   TokenType = "APP_USER"
	TokenTypeOther     TokenType = "OTHER"
	TokenTypeRFID      TokenType = "RFID"
)

// WhitelistType defines model for Token.Whitelist.
type WhitelistType string

// Defines values for WhitelistType.
const (
	WhitelistTypeAllowed        WhitelistType = "ALLOWED"
	WhitelistTypeAllowedOffline WhitelistType = "ALLOWED_OFFLINE"
	WhitelistTypeAlways         WhitelistType = "ALWAYS"
	WhitelistTypeNever          WhitelistType = "NEVER"
)

// ChargingPreferencesResponse defines model for ChargingPreferencesResponse.ChargingPreferences.
type ChargingPreferencesResponse string

const (
	ChargingPreferencesResponseAccepted                ChargingPreferencesResponse = "ACCEPTED"
	ChargingPreferencesResponseDepartureRequired       ChargingPreferencesResponse = "DEPARTURE_REQUIRED"
	ChargingPreferencesResponseEnergyNeedRequired      ChargingPreferencesResponse = "ENERGY_NEED_REQUIRED"
	ChargingPreferencesResponseNotPossible             ChargingPreferencesResponse = "NOT_POSSIBLE"
	ChargingPreferencesResponseProfileTypeNotSupported ChargingPreferencesResponse = "PROFILE_TYPE_NOT_SUPPORTED"
)

type Credentials struct {
	Token string            `json:"token"`
	URL   string            `json:"url"`
	Roles []CredentialsRole `json:"roles"`
}

type CredentialsRole struct {
	Role            Role            `json:"role"`
	BusinessDetails BusinessDetails `json:"business_details"`
	PartyID         string          `json:"party_id" validate:"required,len=3"`
	CountryCode     string          `json:"country_code" validate:"required,len=2"`
}

// ActiveChargingProfileResult defines model for activeChargingProfileResult.
type ActiveChargingProfileResult struct {
	Result  ChargingProfileResultType `json:"result" validate:"required"`
	Profile *ActiveChargingProfile    `json:"profile,omitempty"`
}

// ActiveChargingProfile defines model for activeChargingProfile.
type ActiveChargingProfile struct {
	StartDateTime   DateTime        `json:"start_date_time" validate:"required"`
	ChargingProfile ChargingProfile `json:"charging_profile" validate:"required"`
}

// Authorization Changed name of the object from official docs due to colliding naming of info property
type AuthorizationInfo struct {
	Allowed                AuthorizationAllowed `json:"allowed"`
	AuthorizationReference *string              `json:"authorization_reference,omitempty"`
	Info                   *DisplayText         `json:"info,omitempty"`
	Location               *LocationReferences  `json:"location,omitempty"`
	Token                  Token                `json:"token"`
}

// BusinessDetails defines model for businessDetails.
type BusinessDetails struct {
	Logo    *Image  `json:"logo,omitempty"`
	Name    string  `json:"name"`
	Website *string `json:"website,omitempty"`
}

// CancelReservation defines model for cancelReservation.
type CancelReservation struct {
	ResponseURL   string `json:"response_url"`
	ReservationID string `json:"reservation_id"`
}

// CDR defines model for ChargeDetailRecord.
type CDR struct {
	CountryCode              string           `json:"country_code" validate:"required,len=2"`
	PartyID                  string           `json:"party_id" validate:"required,len=3"`
	ID                       string           `json:"id" validate:"required"`
	StartDateTime            DateTime         `json:"start_date_time" validate:"required"`
	EndDateTime              DateTime         `json:"end_date_time" validate:"required"`
	SessionID                *string          `json:"session_id,omitempty"`
	CdrToken                 CdrToken         `json:"cdr_token" validate:"required"`
	AuthMethod               AuthMethod       `json:"auth_method" validate:"required"`
	AuthorizationReference   *string          `json:"authorization_reference,omitempty"`
	CdrLocation              CdrLocation      `json:"cdr_location" validate:"required"`
	MeterID                  *string          `json:"meter_id,omitempty"`
	Currency                 string           `json:"currency" validate:"required,len=3"`
	Tariffs                  []Tariff         `json:"tariffs,omitempty"`
	ChargingPeriods          []ChargingPeriod `json:"charging_periods,omitempty"`
	SignedData               *SignedData      `json:"signed_data,omitempty"`
	TotalCost                Price            `json:"total_cost" validate:"required"`
	TotalFixedCost           *Price           `json:"total_fixed_cost,omitempty"`
	TotalEnergy              json.Number      `json:"total_energy" validate:"required"`
	TotalEnergyCost          *Price           `json:"total_energy_cost,omitempty"`
	TotalTime                json.Number      `json:"total_time" validate:"required"`
	TotalTimeCost            *Price           `json:"total_time_cost,omitempty"`
	TotalParkingTime         *json.Number     `json:"total_parking_time,omitempty"`
	TotalParkingCost         *Price           `json:"total_parking_cost,omitempty"`
	TotalReservationCost     *Price           `json:"total_reservation_cost,omitempty"`
	Remark                   *string          `json:"remark,omitempty"`
	InvoiceReferenceID       *string          `json:"invoice_reference_id,omitempty"`
	Credit                   *bool            `json:"credit,omitempty"`
	CreditReferenceID        *string          `json:"credit_reference_id,omitempty"`
	HomeChargingCompensation *bool            `json:"home_charging_compensation,omitempty"`
	LastUpdated              DateTime         `json:"last_updated" validate:"required"`
}

// CdrLocation defines model for cdrBody_cdr_location.
type CdrLocation struct {
	ID                 string          `json:"id" validate:"required"`
	Name               *string         `json:"name,omitempty"`
	Address            string          `json:"address" validate:"required"`
	City               string          `json:"city" validate:"required"`
	PostalCode         *string         `json:"postal_code,omitempty"`
	State              *string         `json:"state,omitempty"`
	Country            string          `json:"country" validate:"required,len=3"`
	Coordinates        GeoLocation     `json:"coordinates"`
	EvseUID            string          `json:"evse_uid" validate:"required,max=36"`
	EvseID             string          `json:"evse_id" validate:"required,max=48"`
	ConnectorID        string          `json:"connector_id"`
	ConnectorStandard  ConnectorType   `json:"connector_standard" validate:"required"`
	ConnectorFormat    ConnectorFormat `json:"connector_format" validate:"required"`
	ConnectorPowerType PowerType       `json:"connector_power_type" validate:"required"`
}

// GeoLocation defines model for cdrBody_cdr_location_coordinates.
type GeoLocation struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// CdrTokendefines model for cdrBody_cdr_token.
type CdrToken struct {
	CountryCode string    `json:"country_code"`
	PartyID     string    `json:"party_id"`
	UID         string    `json:"uid"`
	Type        TokenType `json:"type"`
	ContractID  string    `json:"contract_id"`
}

// SignedData defines model for cdrBody_signed_data.
type SignedData struct {
	EncodingMethod        string        `json:"encoding_method" validate:"required"`
	EncodingMethodVersion *int          `json:"encoding_method_version,omitempty"`
	PublicKey             *string       `json:"public_key,omitempty"`
	SignedValues          []SignedValue `json:"signed_values"`
	URL                   *string       `json:"url,omitempty"`
}

// SignedValue defines model for cdrBody_signed_data_signed_values.
type SignedValue struct {
	Nature     string `json:"nature" validate:"required,max=32"`
	PlainData  string `json:"plain_data" validate:"required,max=512"`
	SignedData string `json:"signed_data" validate:"required,max=5000"`
}

// EnergyMix defines model for cdrBody_tariffs_energy_mix.
type EnergyMix struct {
	IsGreenEnergy     bool                  `json:"is_green_energy"`
	EnergySources     []EnergySource        `json:"energy_sources,omitempty"`
	EnvironImpact     []EnvironmentalImpact `json:"environ_impact,omitempty"`
	SupplierName      *string               `json:"supplier_name,omitempty"`
	EnergyProductName *string               `json:"energy_product_name,omitempty"`
}

// EnergySource defines model for cdrBody_tariffs_energy_mix_energy_sources.
type EnergySource struct {
	Percentage json.Number          `json:"percentage" validate:"required,number"`
	Source     EnergySourceCategory `json:"source"`
}

// EnvironmentalImpact defines model for cdrBody_tariffs_energy_mix_environ_impact.
type EnvironmentalImpact struct {
	Amount   json.Number                 `json:"amount" validate:"required,number"`
	Category EnvironmentalImpactCategory `json:"category"`
}

// ChargingPreferences defines model for chargingPreferences.
type ChargingPreferences struct {
	ProfileType      ProfileType  `json:"profile_type"`
	DepartureTime    *DateTime    `json:"departure_time,omitempty"`
	EnergyNeed       *json.Number `json:"energy_need,omitempty"  validate:"omitempty,number"`
	DischargeAllowed *bool        `json:"discharged_allowed,omitempty"`
}

// ChargingProfile defines model for chargingProfile.
type ChargingProfile struct {
	StartDateTime         *DateTime               `json:"start_date_time,omitempty"`
	Duration              *int                    `json:"duration,omitempty"`
	ChargingRateUnit      ChargingRateUnit        `json:"charging_rate_unit" validate:"required"`
	MinChargingRate       *json.Number            `json:"min_charging_rate,omitempty"`
	ChargingProfilePeriod []ChargingProfilePeriod `json:"charging_profile_period,omitempty"`
}

// ChargingProfileResponse defines model for chargingProfileResponse.
type ChargingProfileResponse struct {
	// Result Response to the ChargingProfile request from the eMSP to the CPO.
	Result  ChargingProfileResponseType `json:"result" validate:"required"`
	Timeout int                         `json:"timeout" validate:"required"`
}

// ChargingProfileResult defines model for chargingProfileResult.
type ChargingProfileResult struct {
	Result ChargingProfileResultResult `json:"result"`
}

// ChargingProfileResultResult defines model for ChargingProfileResult.Result.
type ChargingProfileResultResult string

// ChargingProfilePeriod defines model for chargingProfile_charging_profile_period.
type ChargingProfilePeriod struct {
	StartPeriod int         `json:"start_period" validate:"required"`
	Limit       json.Number `json:"limit" validate:"required"`
}

// ClearProfileResult defines model for clearProfileResult.
type ClearProfileResult struct {
	Result ClearProfileResultResult `json:"result"`
}

// ClearProfileResultResult defines model for ClearProfileResult.Result.
type ClearProfileResultResult string

// ClientInfo defines model for clientInfo.
type ClientInfo struct {
	PartyID     string           `json:"party_id" validate:"required,len=3"`
	CountryCode string           `json:"country_code" validate:"required,len=2"`
	Role        Role             `json:"role" valdate:"required"`
	Status      ConnectionStatus `json:"status" valdate:"required"`
	LastUpdated DateTime         `json:"last_updated" valdate:"required"`
}

// CommandResponse defines model for commandResponse.
type CommandResponse struct {
	Result  CommandResponseType `json:"result"`
	Timeout int                 `json:"timeout"`
	Message []DisplayText       `json:"message,omitempty"`
}

// CommandResult defines model for commandResult.
type CommandResult struct {
	Result  CommandResultType `json:"result"`
	Message []DisplayText     `json:"message,omitempty"`
}

// CommandResultType defines model for CommandResult.Result.
type CommandResultType string

// Connector defines model for connector.
type Connector struct {
	ID                 string          `json:"id"`
	Standard           ConnectorType   `json:"standard"`
	Format             ConnectorFormat `json:"format"`
	PowerType          PowerType       `json:"power_type"`
	MaxVoltage         int             `json:"max_voltage"`
	MaxAmperage        int             `json:"max_amperage"`
	MaxElectricPower   *int            `json:"max_electric_power,omitempty"`
	TariffIDs          []string        `json:"tariff_ids,omitempty"`
	TermsAndConditions *string         `json:"terms_and_conditions,omitempty"`
	LastUpdated        DateTime        `json:"last_updated"`
}

type PartialConnector struct {
	ID                 *string          `json:"id,omitempty"`
	Standard           *ConnectorType   `json:"standard,omitempty"`
	Format             *ConnectorFormat `json:"format,omitempty"`
	PowerType          *PowerType       `json:"power_type,omitempty"`
	MaxVoltage         *int             `json:"max_voltage,omitempty"`
	MaxAmperage        *int             `json:"max_amperage,omitempty"`
	MaxElectricPower   *int             `json:"max_electric_power,omitempty"`
	TariffIDs          []string         `json:"tariff_ids,omitempty"`
	TermsAndConditions *string          `json:"terms_and_conditions,omitempty"`
	LastUpdated        DateTime         `json:"last_updated"`
}

// Image defines model for credentials_data_roles_business_details_logo.
type Image struct {
	URL       string        `json:"url" validate:"required"`
	Thumbnail *string       `json:"thumbnail,omitempty"`
	Category  ImageCategory `json:"category" validate:"required"`
	Type      string        `json:"type" validate:"required"`
	Width     *int          `json:"width,omitempty"`
	Height    *int          `json:"height,omitempty"`
}

// DetailsData defines model for details_data.
type VersionDetails struct {
	Endpoints []Endpoint         `json:"endpoints"`
	Version   ocpi.VersionNumber `json:"version"`
}

// Endpoint defines model for details_data_endpoints.
type Endpoint struct {
	// Identifier OCPI 2.2.1 modules
	Identifier ModuleID `json:"identifier"`

	// Role Interface role endpoint implements.
	Role InterfaceRole `json:"role"`

	// Url URL to the endpoint.
	URL string `json:"url"`
}

// EVSE defines model for evse.
type EVSE struct {
	UID                 string               `json:"uid" validate:"required"`
	EvseID              *string              `json:"evse_id,omitempty"`
	Status              Status               `json:"status" validate:"required"`
	StatusSchedule      []StatusSchedule     `json:"status_schedule,omitempty"`
	Capabilities        []Capability         `json:"capabilities,omitempty"`
	Connectors          []Connector          `json:"connectors" validate:"required"`
	FloorLevel          *string              `json:"floor_level,omitempty"`
	Coordinates         *GeoLocation         `json:"coordinates,omitempty"`
	PhysicalReference   *string              `json:"physical_reference,omitempty"`
	Directions          []DisplayText        `json:"directions,omitempty"`
	ParkingRestrictions []ParkingRestriction `json:"parking_restrictions,omitempty"`
	Images              []Image              `json:"images,omitempty"`
	LastUpdated         DateTime             `json:"last_updated"`
}

type PartialEVSE struct {
	UID                 *string              `json:"uid,omitempty"`
	EvseID              *string              `json:"evse_id,omitempty"`
	Status              *Status              `json:"status,omitempty"`
	StatusSchedule      []StatusSchedule     `json:"status_schedule,omitempty"`
	Capabilities        []Capability         `json:"capabilities,omitempty"`
	Connectors          []Connector          `json:"connectors"`
	FloorLevel          *string              `json:"floor_level,omitempty"`
	Coordinates         *GeoLocation         `json:"coordinates,omitempty"`
	PhysicalReference   *string              `json:"physical_reference,omitempty"`
	Directions          []DisplayText        `json:"directions,omitempty"`
	ParkingRestrictions []ParkingRestriction `json:"parking_restrictions,omitempty"`
	Images              []Image              `json:"images,omitempty"`
	LastUpdated         DateTime             `json:"last_updated"`
}

type EVSEPosition string

const (
	EVSEPositionLeft   EVSEPosition = "LEFT"
	EVSEPositionRight  EVSEPosition = "RIGHT"
	EVSEPositionCenter EVSEPosition = "CENTER"
)

// Capability defines model for Evse.Capabilities.
type Capability string

// StatusSchedule defines model for evse_status_schedule.
type StatusSchedule struct {
	PeriodBegin DateTime  `json:"period_begin"`
	PeriodEnd   *DateTime `json:"period_end,omitempty"`
	Status      Status    `json:"status"`
}

// LocationReferences defines model for locationReferences.
type LocationReferences struct {
	EvseUIDs   *string `json:"evse_uids,omitempty"`
	LocationID string  `json:"location_id"`
}

// LocationsData defines model for locations_data.
type Location struct {
	CountryCode        string                  `json:"country_code" validate:"required,len=2"`
	PartyID            string                  `json:"party_id" validate:"required,len=3"`
	ID                 string                  `json:"id" validate:"required"`
	Publish            bool                    `json:"publish"`
	PublishAllowedTo   []PublishTokenType      `json:"publish_allowed_to,omitempty"`
	Name               *string                 `json:"name,omitempty"`
	Address            string                  `json:"address"`
	City               string                  `json:"city"`
	PostalCode         *string                 `json:"postal_code,omitempty"`
	State              *string                 `json:"state,omitempty" validate:"omitempty,required,max=20"`
	Country            string                  `json:"country" validate:"required,len=3"`
	Coordinates        GeoLocation             `json:"coordinates"`
	RelatedLocations   []AdditionalGeoLocation `json:"related_locations,omitempty"`
	ParkingType        *ParkingType            `json:"parking_type,omitempty"`
	EVSEs              []EVSE                  `json:"evses,omitempty"`
	ParkingPlaces      []Parking               `json:"parking_places,omitempty"`
	Directions         []DisplayText           `json:"directions,omitempty"`
	Operator           *BusinessDetails        `json:"operator,omitempty"`
	Suboperator        *BusinessDetails        `json:"suboperator,omitempty"`
	Owner              *BusinessDetails        `json:"owner,omitempty"`
	Facilities         []Facility              `json:"facilities,omitempty"`
	TimeZone           string                  `json:"time_zone" validate:"required"`
	OpeningTimes       *Hours                  `json:"opening_times,omitempty"`
	ChargingWhenClosed *bool                   `json:"charging_when_closed,omitempty"`
	Images             []Image                 `json:"images,omitempty"`
	EnergyMix          *EnergyMix              `json:"energy_mix,omitempty"`
	LastUpdated        DateTime                `json:"last_updated" validate:"required"`
}

type Parking struct {
	ID                           string        `json:"id"`
	PhysicalReference            *string       `json:"physical_reference"`
	VehicleTypes                 []VehicleType `json:"vehicle_types"`
	MaxVehicleWeight             *json.Number  `json:"max_vehicle_weight"`
	MaxVehicleHeight             *json.Number  `json:"max_vehicle_height"`
	MaxVehicleLength             *json.Number  `json:"max_vehicle_length"`
	MaxVehicleWidth              *json.Number  `json:"max_vehicle_width"`
	MaxVehicleParkingSpaceLength *json.Number  `json:"parking_space_length"`
	MaxVehicleParkingSpaceWidth  *json.Number  `json:"parking_space_width"`
	DangerousGoodsAllowed        *bool         `json:"dangerous_goods_allowed"`
}

type VehicleType string

const (
	VehicleTypeMotorcycle                 VehicleType = "MOTORCYCLE"
	VehicleTypePersonalVehicle            VehicleType = "PERSONAL_VEHICLE"
	VehicleTypePersonalVehicleWithTrailer VehicleType = "PERSONAL_VEHICLE_WITH_TRAILER"
	VehicleTypeVan                        VehicleType = "VAN"
	VehicleTypeSemiTractor                VehicleType = "SEMI_TRACTOR"
	VehicleTypeRIGID                      VehicleType = "RIGID"
	VehicleTypeTruckWithTrailer           VehicleType = "TRUCK_WITH_TRAILER"
	VehicleTypeBus                        VehicleType = "BUS"
	VehicleTypeDisabled                   VehicleType = "DISABLED"
)

type PartialLocation struct {
	CountryCode        *string                 `json:"country_code,omitempty"`
	PartyID            *string                 `json:"party_id,omitempty"`
	ID                 *string                 `json:"id,omitempty"`
	Publish            *bool                   `json:"publish,omitempty"`
	PublishAllowedTo   []PublishTokenType      `json:"publish_allowed_to,omitempty"`
	Name               *string                 `json:"name,omitempty"`
	Address            *string                 `json:"address,omitempty"`
	City               *string                 `json:"city,omitempty"`
	PostalCode         *string                 `json:"postal_code,omitempty"`
	State              *string                 `json:"state,omitempty"`
	Country            *string                 `json:"country,omitempty"`
	Coordinates        *GeoLocation            `json:"coordinates,omitempty"`
	RelatedLocations   []AdditionalGeoLocation `json:"related_locations,omitempty"`
	ParkingType        *ParkingType            `json:"parking_type,omitempty"`
	EVSEs              []EVSE                  `json:"evses,omitempty"`
	ParkingPlaces      []Parking               `json:"parking_places,omitempty"`
	Directions         []DisplayText           `json:"directions,omitempty"`
	Operator           *BusinessDetails        `json:"operator,omitempty"`
	Suboperator        *BusinessDetails        `json:"suboperator,omitempty"`
	Owner              *BusinessDetails        `json:"owner,omitempty"`
	Facilities         []Facility              `json:"facilities,omitempty"`
	TimeZone           *string                 `json:"time_zone,omitempty"`
	OpeningTimes       *Hours                  `json:"opening_times,omitempty"`
	ChargingWhenClosed *bool                   `json:"charging_when_closed,omitempty"`
	Images             []Image                 `json:"images,omitempty"`
	EnergyMix          *EnergyMix              `json:"energy_mix,omitempty"`
	LastUpdated        DateTime                `json:"last_updated"`
}

// Hours defines model for locations_data_opening_times.
type Hours struct {
	ExceptionalClosings []ExceptionalPeriod `json:"exceptional_closings,omitempty"`
	ExceptionalOpenings []ExceptionalPeriod `json:"exceptional_openings,omitempty"`
	RegularHours        []RegularHours      `json:"regular_hours,omitempty"`
	Twentyfourseven     bool                `json:"twentyfourseven"`
}

// HoursExceptionalOpenings defines model for locations_data_opening_times_exceptional_openings.
type ExceptionalPeriod struct {
	PeriodBegin DateTime `json:"period_begin"`
	PeriodEnd   DateTime `json:"period_end"`
}

// HoursRegularHours defines model for locations_data_opening_times_regular_hours.
type RegularHours struct {
	PeriodBegin string `json:"period_begin"`
	PeriodEnd   string `json:"period_end"`
	Weekday     int    `json:"weekday"`
}

// PublishTokenType defines model for locations_data_publish_allowed_to.
type PublishTokenType struct {
	UID          *string    `json:"uid,omitempty"`
	Type         *TokenType `json:"type,omitempty"`
	VisualNumber *string    `json:"visual_number,omitempty"`
	Issuer       *string    `json:"issuer,omitempty"`
	GroupID      *string    `json:"group_id,omitempty"`
}

// AdditionalGeoLocation defines model for locations_data_related_locations.
type AdditionalGeoLocation struct {
	Latitude  string       `json:"latitude"`
	Longitude string       `json:"longitude"`
	Name      *DisplayText `json:"name,omitempty"`
}

// Price defines model for price.
type Price struct {
	ExclVat json.Number  `json:"excl_vat"`
	InclVat *json.Number `json:"incl_vat,omitempty"`
}

// ReserveNow defines model for reserveNow.
type ReserveNow struct {
	ResponseURL            string   `json:"response_url" validate:"required"`
	Token                  Token    `json:"token"`
	ExpiryDate             DateTime `json:"expiry_date" validate:"required"`
	ReservationID          string   `json:"reservation_id" validate:"required"`
	LocationID             string   `json:"location_id" validate:"required"`
	EvseUID                *string  `json:"evse_uid,omitempty"`
	AuthorizationReference *string  `json:"authorization_reference,omitempty"`
}

// Session defines model for session.
type Session struct {
	CountryCode            string           `json:"country_code" validate:"required,len=2"`
	PartyID                string           `json:"party_id" validate:"required,len=3"`
	ID                     string           `json:"id" validate:"required"`
	StartDateTime          DateTime         `json:"start_date_time" validate:"required"`
	EndDateTime            *DateTime        `json:"end_date_time,omitempty"`
	Kwh                    json.Number      `json:"kwh" validate:"required"`
	CdrToken               CdrToken         `json:"cdr_token"`
	AuthMethod             AuthMethod       `json:"auth_method" validate:"required"`
	AuthorizationReference *string          `json:"authorization_reference,omitempty"`
	LocationID             string           `json:"location_id" validate:"required"`
	EvseUID                string           `json:"evse_uid" validate:"required"`
	ConnectorID            string           `json:"connector_id" validate:"required"`
	MeterID                *string          `json:"meter_id,omitempty"`
	Currency               string           `json:"currency" validate:"required,len=3"`
	ChargingPeriods        []ChargingPeriod `json:"charging_periods,omitempty"`
	TotalCost              *Price           `json:"total_cost,omitempty"`
	Status                 SessionStatus    `json:"status" validate:"required"`
	LastUpdated            DateTime         `json:"last_updated" validate:"required"`
}

type PartialSession struct {
	AuthMethod             *AuthMethod      `json:"auth_method,omitempty"`
	AuthorizationReference *string          `json:"authorization_reference,omitempty"`
	CdrToken               *CdrToken        `json:"cdr_token,omitempty"`
	ChargingPeriods        []ChargingPeriod `json:"charging_periods,omitempty"`
	ConnectorID            *string          `json:"connector_id,omitempty"`
	CountryCode            *string          `json:"country_code,omitempty"`
	Currency               *string          `json:"currency,omitempty"`
	EndDateTime            *DateTime        `json:"end_date_time,omitempty"`
	EvseUID                *string          `json:"evse_uid,omitempty"`
	ID                     *string          `json:"id,omitempty"`
	Kwh                    *json.Number     `json:"kwh,omitempty"`
	LocationID             *string          `json:"location_id,omitempty"`
	MeterID                *string          `json:"meter_id,omitempty"`
	PartyID                *string          `json:"party_id,omitempty"`
	StartDateTime          *DateTime        `json:"start_date_time,omitempty"`
	Status                 *SessionStatus   `json:"status,omitempty"`
	TotalCost              *Price           `json:"total_cost,omitempty"`
	LastUpdated            DateTime         `json:"last_updated" validate:"required"`
}

// ChargingPeriod defines model for session_charging_periods.
type ChargingPeriod struct {
	StartDateTime DateTime       `json:"start_date_time" validate:"required"`
	Dimensions    []CdrDimension `json:"dimensions"`
	TariffID      *string        `json:"tariff_id,omitempty"`
}

// CdrDimension defines model for session_charging_periods_dimensions.
type CdrDimension struct {
	Type   CdrDimensionType `json:"type"`
	Volume json.Number      `json:"volume"`
}

// SetChargingProfile defines model for setChargingProfile.
type SetChargingProfile struct {
	ChargingProfile ChargingProfile `json:"charging_profile"`
	ResponseURL     string          `json:"response_url"`
}

// StartSessionRequest defines model for startSession.
type StartSession struct {
	AuthorizationReference *string `json:"authorization_reference,omitempty"`
	ConnectorID            *string `json:"connector_id,omitempty"`
	EvseUID                *string `json:"evse_uid,omitempty"`
	LocationID             string  `json:"location_id" validate:"required"`
	ResponseURL            string  `json:"response_url" validate:"required"`
	Token                  Token   `json:"token"`
}

// StopSession defines model for stopSession.
type StopSession struct {
	ResponseURL string `json:"response_url"`
	SessionID   string `json:"session_id"`
}

// Tariff defines model for tariff.
type Tariff struct {
	CountryCode   string          `json:"country_code" validate:"required,len=2"`
	PartyID       string          `json:"party_id" validate:"required,len=3"`
	ID            string          `json:"id" validate:"required"`
	Currency      string          `json:"currency" validate:"required,len=3"`
	Type          *TariffType     `json:"type,omitempty"`
	TariffAltText []DisplayText   `json:"tariff_alt_text,omitempty"`
	TariffAltURL  *string         `json:"tariff_alt_url,omitempty"`
	MinPrice      *Price          `json:"min_price,omitempty"`
	MaxPrice      *Price          `json:"max_price,omitempty"`
	Elements      []TariffElement `json:"elements"`
	StartDateTime *DateTime       `json:"start_date_time,omitempty"`
	EndDateTime   *DateTime       `json:"end_date_time,omitempty"`
	EnergyMix     *EnergyMix      `json:"energy_mix,omitempty"`
	LastUpdated   DateTime        `json:"last_updated" validate:"required"`
}

// TariffElement defines model for tariff_elements.
type TariffElement struct {
	PriceComponents []PriceComponent    `json:"price_components,omitempty"`
	Restrictions    *TariffRestrictions `json:"restrictions,omitempty"`
}

// PriceComponent defines model for TariffElement.PriceComponents.
type PriceComponent struct {
	Type     TariffDimensionType `json:"type" validate:"required"`
	Price    json.Number         `json:"price" validate:"required"`
	Vat      *json.Number        `json:"vat,omitempty"`
	StepSize int                 `json:"step_size" validate:"required"`
}

// TariffRestrictions defines model for tariff_elements_restrictions.
type TariffRestrictions struct {
	StartTime   *string                     `json:"start_time,omitempty"`
	EndTime     *string                     `json:"end_time,omitempty"`
	StartDate   *string                     `json:"start_date,omitempty"`
	EndDate     *string                     `json:"end_date,omitempty"`
	MinKwh      *json.Number                `json:"min_kwh,omitempty"`
	MaxKwh      *json.Number                `json:"max_kwh,omitempty"`
	MinCurrent  *json.Number                `json:"min_current,omitempty"`
	MaxCurrent  *json.Number                `json:"max_current,omitempty"`
	MinPower    *json.Number                `json:"min_power,omitempty"`
	MaxPower    *json.Number                `json:"max_power,omitempty"`
	MinDuration *int                        `json:"min_duration,omitempty"`
	MaxDuration *int                        `json:"max_duration,omitempty"`
	DayOfWeek   []DayOfWeek                 `json:"day_of_week,omitempty"`
	Reservation *ReservationRestrictionType `json:"reservation,omitempty"`
}

// Token defines model for token.
type Token struct {
	ContractID         string               `json:"contract_id"`
	CountryCode        string               `json:"country_code"`
	DefaultProfileType *ProfileType         `json:"default_profile_type,omitempty"`
	EnergyContract     *TokenEnergyContract `json:"energy_contract,omitempty"`
	GroupID            *string              `json:"group_id,omitempty"`
	Issuer             string               `json:"issuer"`
	Language           *string              `json:"language,omitempty"`
	LastUpdated        DateTime             `json:"last_updated"`
	PartyID            string               `json:"party_id"`
	Type               TokenType            `json:"type"`
	UID                string               `json:"uid"`
	Valid              bool                 `json:"valid"`
	VisualNumber       *string              `json:"visual_number,omitempty"`
	Whitelist          WhitelistType        `json:"whitelist"`
}

type PartialToken struct {
	ContractID         string               `json:"contract_id"`
	CountryCode        string               `json:"country_code"`
	DefaultProfileType *ProfileType         `json:"default_profile_type,omitempty"`
	EnergyContract     *TokenEnergyContract `json:"energy_contract,omitempty"`
	GroupID            *string              `json:"group_id,omitempty"`
	Issuer             string               `json:"issuer"`
	Language           *string              `json:"language,omitempty"`
	LastUpdated        DateTime             `json:"last_updated"`
	PartyID            string               `json:"party_id"`
	Type               TokenType            `json:"type"`
	UID                string               `json:"uid"`
	Valid              bool                 `json:"valid"`
	VisualNumber       *string              `json:"visual_number,omitempty"`
	Whitelist          WhitelistType        `json:"whitelist"`
}

// TokenEnergyContract defines model for token_energy_contract.
type TokenEnergyContract struct {
	ContractID   *string `json:"contract_id,omitempty"`
	SupplierName string  `json:"supplier_name"`
}

// UnlockConnector defines model for unlockConnector.
type UnlockConnector struct {
	ResponseURL string `json:"response_url"`
	LocationID  string `json:"location_id"`
	EvseUID     string `json:"evse_uid"`
	ConnectorID string `json:"connector_id"`
}

type ChargeDetailRecordResponse struct {
	Location string
}

// GetCDRsParams defines parameters for GetOcpiCdrs.
type GetCDRsParams = ocpi.PaginatedRequest[DateTime]

// GetHubClientInfoParams defines parameters for GetOcpiHubclientinfo.
type GetHubClientInfoParams = ocpi.PaginatedRequest[DateTime]

// GetLocationsParams defines parameters for GetOcpiLocations.
type GetLocationsParams = ocpi.PaginatedRequest[DateTime]

// GetSessionsParams defines parameters for GetOcpiSessions.
type GetSessionsParams struct {
	// DateTo Return tokens that have last_updated up to Date/Time, but not including (exclusive).
	DateTo *DateTime `form:"date_to,omitempty" json:"date_to,omitempty"`

	// Offset The offset of the first object returned. Default is 0.
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit Maximum number of objects to GET.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// GetTariffsParams defines parameters for GetOcpiTariffs.
type GetTariffsParams = ocpi.PaginatedRequest[DateTime]

// GetTokensParams defines parameters for GetOcpiTokens.
type GetTokensParams = ocpi.PaginatedRequest[DateTime]
