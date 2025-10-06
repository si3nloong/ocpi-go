package ocpi211

import (
	"encoding/json"

	"github.com/si3nloong/ocpi-go/ocpi"
)

// ModuleID OCPI 2.1.1 modules
type ModuleID string

// Defines values for ModuleID.
const (
	ModuleIDCdrs        ModuleID = "cdrs"
	ModuleIDCommands    ModuleID = "commands"
	ModuleIDCredentials ModuleID = "credentials"
	ModuleIDLocations   ModuleID = "locations"
	ModuleIDSessions    ModuleID = "sessions"
	ModuleIDTariffs     ModuleID = "tariffs"
	ModuleIDTokens      ModuleID = "tokens"
)

// Role represents the role a party can have in OCPI
type Role string

const (
	// CPO (Charge Point Operator) role
	RoleCPO Role = "CPO"

	// EMSP (E-Mobility Service Provider) role
	RoleEMSP Role = "EMSP"
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

type LocationType string

const (
	LocationTypeOnStreet          LocationType = "ON_STREET"
	LocationTypeParkingGarage     LocationType = "PARKING_GARAGE"
	LocationTypeUndergroundGarage LocationType = "UNDERGROUND_GARAGE"
	LocationTypeParkingLot        LocationType = "PARKING_LOT"
	LocationTypeOther             LocationType = "OTHER"
	LocationTypeUnknown           LocationType = "UNKNOWN"
)

type Facility string

// Defines values for LocationsDataFacilities.
const (
	FacilityHotel          Facility = "HOTEL"
	FacilityRestaurant     Facility = "RESTAURANT"
	FacilityCafe           Facility = "CAFE"
	FacilityMall           Facility = "MALL"
	FacilitySuperMarket    Facility = "SUPERMARKET"
	FacilitySport          Facility = "SPORT"
	FacilityRecreationArea Facility = "RECREATION_AREA"
	FacilityNature         Facility = "NATURE"
	FacilityMuseum         Facility = "MUSEUM"
	FacilityBusStop        Facility = "BUS_STOP"
	FacilityTaxiStand      Facility = "TAXI_STAND"
	FacilityTrainStation   Facility = "TRAIN_STATION"
	FacilityAirport        Facility = "AIRPORT"
	FacilityCarpoolParking Facility = "CARPOOL_PARKING"
	FacilityFuelStation    Facility = "FUEL_STATION"
	FacilityWiFi           Facility = "WIFI"
)

type Capability string

// Defines values for EvseCapabilities.
const (
	CapabilityChargingProfileCapable Capability = "CHARGING_PROFILE_CAPABLE"
	CapabilityCreditCardPayable      Capability = "CREDIT_CARD_PAYABLE"
	CapabilityRemoteStartStopCapable Capability = "REMOTE_START_STOP_CAPABLE"
	CapabilityReservable             Capability = "RESERVABLE"
	CapabilityRFIDReader             Capability = "RFID_READER"
	CapabilityUnlockCapable          Capability = "UNLOCK_CAPABLE"
)

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

type PowerType string

// Defines values for PowerType.
const (
	PowerTypeAC1Phase PowerType = "AC_1_PHASE"
	PowerTypeAC3Phase PowerType = "AC_3_PHASE"
	PowerTypeDC       PowerType = "DC"
)

type ConnectorType string

// Defines values for ConnectorType.
const (
	ConnectorTypeCHAdeMO           ConnectorType = "CHADEMO"
	ConnectorTypeDomesticA         ConnectorType = "DOMESTIC_A"
	ConnectorTypeDomesticB         ConnectorType = "DOMESTIC_B"
	ConnectorTypeDomesticC         ConnectorType = "DOMESTIC_C"
	ConnectorTypeDomesticD         ConnectorType = "DOMESTIC_D"
	ConnectorTypeDomesticE         ConnectorType = "DOMESTIC_E"
	ConnectorTypeDomesticF         ConnectorType = "DOMESTIC_F"
	ConnectorTypeDomesticG         ConnectorType = "DOMESTIC_G"
	ConnectorTypeDomesticH         ConnectorType = "DOMESTIC_H"
	ConnectorTypeDomesticJ         ConnectorType = "DOMESTIC_J"
	ConnectorTypeDomesticK         ConnectorType = "DOMESTIC_K"
	ConnectorTypeDomesticL         ConnectorType = "DOMESTIC_L"
	ConnectorTypeIEC603092Single16 ConnectorType = "IEC_60309_2_single_16"
	ConnectorTypeIEC603092Three16  ConnectorType = "IEC_60309_2_three_16"
	ConnectorTypeIEC603092Three32  ConnectorType = "IEC_60309_2_three_32"
	ConnectorTypeIEC603092Three64  ConnectorType = "IEC_60309_2_three_64"
	ConnectorTypeIEC62196T1        ConnectorType = "IEC_62196_T1"
	ConnectorTypeIEC62196T1Combo   ConnectorType = "IEC_62196_T1_COMBO"
	ConnectorTypeIEC62196T2        ConnectorType = "IEC_62196_T2"
	ConnectorTypeIEC62196T2Combo   ConnectorType = "IEC_62196_T2_COMBO"
	ConnectorTypeIEC62196T3A       ConnectorType = "IEC_62196_T3A"
	ConnectorTypeIEC62196T3C       ConnectorType = "IEC_62196_T3C"
	ConnectorTypeTeslaR            ConnectorType = "TESLA_R"
	ConnectorTypeTeslaS            ConnectorType = "TESLA_S"
)

// ConnectorFormat defines model for Connector.Format.
type ConnectorFormat string

// Defines values for ConnectorFormat.
const (
	ConnectorFormatCable  ConnectorFormat = "CABLE"
	ConnectorFormatSocket ConnectorFormat = "SOCKET"
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
	EnvironmentalImpactCategoryNuclearWaste  EnvironmentalImpactCategory = "NUCLEAR_WASTE"
	EnvironmentalImpactCategoryCarbonDioxide EnvironmentalImpactCategory = "CARBON_DIOXIDE"
)

type AuthMethod string

const (
	AuthMethodAuthRequest AuthMethod = "AUTH_REQUEST"
	AuthMethodWhitelist   AuthMethod = "WHITELIST"
)

// CdrDimensionType defines model for CdrBodyChargingPeriodsDimensions.Type.
type CdrDimensionType string

// Defines values for CdrDimensionType.
const (
	CdrDimensionTypeEnergy      CdrDimensionType = "ENERGY"
	CdrDimensionTypeFlat        CdrDimensionType = "FLAT"
	CdrDimensionTypeMaxCurrent  CdrDimensionType = "MAX_CURRENT"
	CdrDimensionTypeMinCurrent  CdrDimensionType = "MIN_CURRENT"
	CdrDimensionTypeParkingTime CdrDimensionType = "PARKING_TIME"
	CdrDimensionTypeTime        CdrDimensionType = "TIME"
)

// SessionStatus defines model for Session.Status.
type SessionStatus string

// Defines values for SessionStatus.
const (
	SessionStatusActive    SessionStatus = "ACTIVE"
	SessionStatusCompleted SessionStatus = "COMPLETED"
	SessionStatusInvalid   SessionStatus = "INVALID"
	SessionStatusPending   SessionStatus = "PENDING"
)

// CommandType defines parameters for type of commands.
type CommandType string

// Defines values for PostOcpiCommandsCommandParamsCommand.
const (
	CommandTypeReserveNow      CommandType = "RESERVE_NOW"
	CommandTypeStartSession    CommandType = "START_SESSION"
	CommandTypeStopSession     CommandType = "STOP_SESSION"
	CommandTypeUnlockConnector CommandType = "UNLOCK_CONNECTOR"
)

// CommandResponseType defines model for CommandResponse.Result.
type CommandResponseType string

// Defines values for CommandResponseType.
const (
	CommandResponseTypeNotSupported   CommandResponseType = "NOT_SUPPORTED"
	CommandResponseTypeRejected       CommandResponseType = "REJECTED"
	CommandResponseTypeAccepted       CommandResponseType = "ACCEPTED"
	CommandResponseTypeTimeout        CommandResponseType = "TIMEOUT"
	CommandResponseTypeUnknownSession CommandResponseType = "UNKNOWN_SESSION"
)

// TokenType defines model for Token.Type.
type TokenType string

// Defines values for TokenType.
const (
	TokenTypeOther TokenType = "OTHER"
	TokenTypeRFID  TokenType = "RFID"
)

type WhitelistType string

// Defines values for WhitelistType.
const (
	WhitelistTypeAlways         WhitelistType = "ALWAYS"
	WhitelistTypeAllowed        WhitelistType = "ALLOWED"
	WhitelistTypeAllowedOffline WhitelistType = "ALLOWED_OFFLINE"
	WhitelistTypeNever          WhitelistType = "NEVER"
)

// Allowed defines model for Authorization.Allowed.
type Allowed string

// Defines values for Allowed.
const (
	AllowedAllowed    Allowed = "ALLOWED"
	AllowedBlocked    Allowed = "BLOCKED"
	AllowedExpired    Allowed = "EXPIRED"
	AllowedNoCredit   Allowed = "NO_CREDIT"
	AllowedNotAllowed Allowed = "NOT_ALLOWED"
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

type TariffDimensionType string

// Defines values for PriceComponent.
const (
	TariffDimensionTypeEnergy      TariffDimensionType = "ENERGY"
	TariffDimensionTypeFlat        TariffDimensionType = "FLAT"
	TariffDimensionTypeParkingTime TariffDimensionType = "PARKING_TIME"
	TariffDimensionTypeTime        TariffDimensionType = "TIME"
)

type VersionDetails struct {
	Endpoints []Endpoint         `json:"endpoints"`
	Version   ocpi.VersionNumber `json:"version"`
}

// Endpoint defines model for details_data_endpoints.
type Endpoint struct {
	// Identifier OCPI 2.1.1 modules
	Identifier ModuleID `json:"identifier" validate:"required"`

	// Url URL to the endpoint.
	URL string `json:"url" validate:"required,url"`
}

type Credentials struct {
	Token           string          `json:"token" validate:"required,max=64"`
	URL             string          `json:"url" validate:"required,url"`
	BusinessDetails BusinessDetails `json:"business_details" validate:"required"`
	PartyID         string          `json:"party_id" validate:"required,len=3"`
	CountryCode     string          `json:"country_code" validate:"required,len=2"`
}

type Location struct {
	ID                 string                  `json:"id" validate:"required"`
	Type               LocationType            `json:"type" validate:"required,locationType211"`
	Name               *string                 `json:"name,omitempty" validate:"omitempty,required"`
	Address            string                  `json:"address" validate:"required"`
	City               string                  `json:"city" validate:"required"`
	PostalCode         string                  `json:"postal_code" validate:"required,max=10"`
	Country            string                  `json:"country" validate:"required,len=3"`
	Coordinates        GeoLocation             `json:"coordinates" validate:"required"`
	RelatedLocations   []AdditionalGeoLocation `json:"related_locations,omitempty"`
	EVSEs              []EVSE                  `json:"evses,omitempty"`
	Directions         []DisplayText           `json:"directions,omitempty"`
	Operator           *BusinessDetails        `json:"operator,omitempty" validate:"omitempty,required"`
	Suboperator        *BusinessDetails        `json:"suboperator,omitempty" validate:"omitempty,required"`
	Owner              *BusinessDetails        `json:"owner,omitempty" validate:"omitempty,required"`
	Facilities         []Facility              `json:"facilities,omitempty"`
	TimeZone           *string                 `json:"time_zone,omitempty" validate:"omitempty,required"`
	OpeningTimes       *Hours                  `json:"opening_times,omitempty"`
	ChargingWhenClosed *bool                   `json:"charging_when_closed,omitempty"`
	Images             []Image                 `json:"images,omitempty" validate:"omitempty"`
	EnergyMix          *EnergyMix              `json:"energy_mix,omitempty"`
	LastUpdated        DateTime                `json:"last_updated" validate:"required"`
}

type Hours struct {
	RegularHours        []RegularHours      `json:"regular_hours,omitempty"`
	Twentyfourseven     bool                `json:"twentyfourseven"`
	ExceptionalOpenings []ExceptionalPeriod `json:"exceptional_openings,omitempty"`
	ExceptionalClosings []ExceptionalPeriod `json:"exceptional_closings,omitempty"`
}

// HoursRegularHours defines model for locations_data_opening_times_regular_hours.
type RegularHours struct {
	Weekday     int    `json:"weekday" validate:"required,gte=1,lte=7"`
	PeriodBegin string `json:"period_begin" validate:"required"`
	PeriodEnd   string `json:"period_end" validate:"required"`
}

// HoursExceptionalOpenings defines model for locations_data_opening_times_exceptional_openings.
type ExceptionalPeriod struct {
	PeriodBegin DateTime `json:"period_begin"`
	PeriodEnd   DateTime `json:"period_end"`
}

type PartialLocation struct {
	ID                 *string                 `json:"id" validate:"omitempty,required"`
	Type               *LocationType           `json:"type" validate:"omitempty,required,locationType211"`
	Name               *string                 `json:"name" validate:"omitempty,required"`
	Address            *string                 `json:"address" validate:"omitempty,required"`
	City               *string                 `json:"city" validate:"omitempty,required"`
	PostalCode         *string                 `json:"postal_code" validate:"omitempty,required,max=10"`
	Country            *string                 `json:"country" validate:"omitempty,required,len=3"`
	Coordinates        *GeoLocation            `json:"coordinates" validate:"omitempty,required"`
	RelatedLocations   []AdditionalGeoLocation `json:"related_locations,omitempty" validate:"omitempty,required"`
	EVSEs              []EVSE                  `json:"evses,omitempty" validate:"omitempty,required"`
	Directions         []DisplayText           `json:"directions,omitempty" validate:"omitempty,required"`
	Operator           *BusinessDetails        `json:"operator,omitempty" validate:"omitempty,required"`
	Suboperator        *BusinessDetails        `json:"suboperator,omitempty" validate:"omitempty,required"`
	Owner              *BusinessDetails        `json:"owner,omitempty" validate:"omitempty,required"`
	Facilities         []Facility              `json:"facilities,omitempty" validate:"omitempty,required"`
	TimeZone           *string                 `json:"time_zone" validate:"omitempty,required"`
	OpeningTimes       *Hours                  `json:"opening_times,omitempty" validate:"omitempty,required"`
	ChargingWhenClosed *bool                   `json:"charging_when_closed,omitempty" validate:"omitempty,required"`
	Images             []Image                 `json:"images,omitempty" validate:"omitempty,required"`
	EnergyMix          *EnergyMix              `json:"energy_mix,omitempty" validate:"omitempty,required"`
	LastUpdated        *DateTime               `json:"last_updated" validate:"omitempty,required"`
}

// EVSE defines model for evse.
type EVSE struct {
	UID                 string               `json:"uid" validate:"required"`
	EvseID              *string              `json:"evse_id,omitempty" validate:"omitempty,required"`
	Status              Status               `json:"status" validate:"required,status211"`
	StatusSchedule      []StatusSchedule     `json:"status_schedule,omitempty"`
	Capabilities        []Capability         `json:"capabilities,omitempty"`
	Connectors          []Connector          `json:"connectors" validate:"required"`
	FloorLevel          *string              `json:"floor_level,omitempty" validate:"omitempty,required"`
	Coordinates         *GeoLocation         `json:"coordinates,omitempty" validate:"omitempty,required"`
	PhysicalReference   *string              `json:"physical_reference,omitempty" validate:"omitempty,required"`
	Directions          []DisplayText        `json:"directions,omitempty" validate:"omitempty,required"`
	ParkingRestrictions []ParkingRestriction `json:"parking_restrictions,omitempty" validate:"omitempty,required"`
	Images              []Image              `json:"images,omitempty" validate:"omitempty,required"`
	LastUpdated         DateTime             `json:"last_updated" validate:"required"`
}

type PartialEVSE struct {
	UID                 *string              `json:"uid,omitempty" validate:"omitempty,required"`
	EvseID              *string              `json:"evse_id,omitempty" validate:"omitempty,required"`
	Status              *Status              `json:"status,omitempty" validate:"omitempty,required,status211"`
	StatusSchedule      []StatusSchedule     `json:"status_schedule,omitempty" validate:"omitempty,required"`
	Capabilities        []Capability         `json:"capabilities,omitempty" validate:"omitempty,required"`
	Connectors          []Connector          `json:"connectors,omitempty" validate:"omitempty,required"`
	FloorLevel          *string              `json:"floor_level,omitempty" validate:"omitempty,required"`
	Coordinates         *GeoLocation         `json:"coordinates,omitempty" validate:"omitempty,required"`
	PhysicalReference   *string              `json:"physical_reference,omitempty" validate:"omitempty,required"`
	Directions          []DisplayText        `json:"directions,omitempty" validate:"omitempty,required"`
	ParkingRestrictions []ParkingRestriction `json:"parking_restrictions,omitempty" validate:"omitempty,required"`
	Images              []Image              `json:"images,omitempty" validate:"omitempty,required"`
	LastUpdated         DateTime             `json:"last_updated" validate:"required"`
}

// Connector defines model for connector.
type Connector struct {
	ID                 string          `json:"id" validate:"required"`
	Standard           ConnectorType   `json:"standard" validate:"required"`
	Format             ConnectorFormat `json:"format" validate:"required,connectorFormat211"`
	PowerType          PowerType       `json:"power_type" validate:"required"`
	Voltage            int             `json:"voltage"`
	Amperage           int             `json:"amperage"`
	TariffID           *string         `json:"tariff_id,omitempty"`
	TermsAndConditions *string         `json:"terms_and_conditions,omitempty"`
	LastUpdated        DateTime        `json:"last_updated"`
}

type PartialConnector struct {
	ID                 *string          `json:"id,omitempty" validate:"omitempty,required"`
	Standard           *ConnectorType   `json:"standard,omitempty" validate:"omitempty"`
	Format             *ConnectorFormat `json:"format,omitempty" validate:"omitempty,connectorFormat211"`
	PowerType          *PowerType       `json:"power_type,omitempty" validate:"omitempty,powerType211"`
	Voltage            *int             `json:"voltage,omitempty" validate:"omitempty,required"`
	Amperage           *int             `json:"amperage,omitempty" validate:"omitempty,required"`
	TariffID           *string          `json:"tariff_id,omitempty" validate:"omitempty,required"`
	TermsAndConditions *string          `json:"terms_and_conditions,omitempty" validate:"omitempty,required"`
	LastUpdated        DateTime         `json:"last_updated" validate:"required"`
}

type CDR struct {
	ID               string           `json:"id" validate:"required"`
	StartDateTime    DateTime         `json:"start_date_time" validate:"required"`
	StopDateTime     DateTime         `json:"stop_date_time" validate:"required"`
	AuthID           string           `json:"auth_id" validate:"max=36"`
	AuthMethod       AuthMethod       `json:"auth_method" validate:"required"`
	Location         Location         `json:"location" validate:"required"`
	MeterID          *string          `json:"meter_id,omitempty" validate:"omitempty,required"`
	Currency         string           `json:"currency" validate:"required,len=3"`
	Tariffs          []Tariff         `json:"tariffs,omitempty"`
	ChargingPeriods  []ChargingPeriod `json:"charging_periods"`
	TotalCost        json.Number      `json:"total_cost" validate:"required"`
	TotalEnergy      json.Number      `json:"total_energy" validate:"required"`
	TotalTime        json.Number      `json:"total_time" validate:"required"`
	TotalParkingTime *json.Number     `json:"total_parking_time,omitempty" validate:"omitempty,required"`
	Remark           *string          `json:"remark,omitempty" validate:"omitempty,required"`
	LastUpdated      DateTime         `json:"last_updated" validate:"required"`
}

type EnergyMix struct {
	IsGreenEnergy     bool                  `json:"is_green_energy"`
	EnergySources     []EnergySource        `json:"energy_sources,omitempty"`
	EnvironImpact     []EnvironmentalImpact `json:"environ_impact,omitempty"`
	SupplierName      *string               `json:"supplier_name,omitempty" validate:"omitempty,required"`
	EnergyProductName *string               `json:"energy_product_name,omitempty" validate:"omitempty,required"`
}

type EnergySource struct {
	Percentage json.Number          `json:"percentage"`
	Source     EnergySourceCategory `json:"source"`
}

// EnvironmentalImpact defines model for cdrBody_tariffs_energy_mix_environ_impact.
type EnvironmentalImpact struct {
	Amount   json.Number                 `json:"amount"`
	Category EnvironmentalImpactCategory `json:"category"`
}

type Session struct {
	ID              string           `json:"id" validate:"required"`
	StartDateTime   DateTime         `json:"start_datetime" validate:"required"`
	EndDateTime     *DateTime        `json:"end_datetime"`
	Kwh             json.Number      `json:"kwh"`
	AuthID          string           `json:"auth_id" validate:"required"`
	AuthMethod      AuthMethod       `json:"auth_method" validate:"required"`
	Location        Location         `json:"location" validate:"required"`
	MeterID         *string          `json:"meter_id,omitempty"`
	Currency        string           `json:"currency" validate:"required,len=3"`
	ChargingPeriods []ChargingPeriod `json:"charging_periods,omitempty"`
	TotalCost       *json.Number     `json:"total_cost" validate:"required"`
	Status          SessionStatus    `json:"status" validate:"required,sessionStatus211"`
	LastUpdated     DateTime         `json:"last_updated" validate:"required"`
}

type PartialSession struct {
	ID              *string          `json:"id" validate:"omitempty,required"`
	StartDateTime   *DateTime        `json:"start_datetime" validate:"omitempty,required"`
	EndDateTime     *DateTime        `json:"end_datetime" validate:"omitempty,required"`
	Kwh             *json.Number     `json:"kwh" validate:"omitempty,required"`
	AuthID          *string          `json:"auth_id" validate:"omitempty,required"`
	AuthMethod      *AuthMethod      `json:"auth_method" validate:"omitempty,required"`
	Location        *PartialLocation `json:"location" validate:"omitempty,required"`
	MeterID         *string          `json:"meter_id,omitempty" validate:"omitempty,required"`
	Currency        *string          `json:"currency" validate:"omitempty,required,len=3"`
	ChargingPeriods []ChargingPeriod `json:"charging_periods,omitempty"`
	TotalCost       *json.Number     `json:"total_cost" validate:"omitempty,required"`
	Status          *SessionStatus   `json:"status" validate:"omitempty,required,sessionStatus211"`
	LastUpdated     *DateTime        `json:"last_updated" validate:"omitempty,required"`
}

type Tariff struct {
	ID            string          `json:"id" validate:"required"`
	Currency      string          `json:"currency" validate:"required,len=3"`
	TariffAltText []DisplayText   `json:"tariff_alt_text,omitempty"`
	TariffAltURL  *string         `json:"tariff_alt_url,omitempty"`
	Elements      []TariffElement `json:"elements"`
	EnergyMix     *EnergyMix      `json:"energy_mix" validate:"required"`
	LastUpdated   DateTime        `json:"last_updated" validate:"required"`
}

type PartialTariff struct {
	ID            *string         `json:"id,omitempty"`
	Currency      *string         `json:"currency,omitempty" validate:"omitempty,len=3"`
	TariffAltText []DisplayText   `json:"tariff_alt_text,omitempty"`
	TariffAltURL  *string         `json:"tariff_alt_url,omitempty"`
	Elements      []TariffElement `json:"elements,omitempty"`
	EnergyMix     *EnergyMix      `json:"energy_mix,omitempty"`
	LastUpdated   *DateTime       `json:"last_updated,omitempty"`
}

type ChargingPeriod struct {
	StartDateTime DateTime       `json:"start_date_time" validate:"required"`
	Dimensions    []CdrDimension `json:"dimensions"`
}

// CdrDimension defines model for session_charging_periods_dimensions.
type CdrDimension struct {
	Type   CdrDimensionType `json:"type"`
	Volume json.Number      `json:"volume"`
}

type TariffElement struct {
	PriceComponents []PriceComponent    `json:"price_components"`
	Restrictions    *TariffRestrictions `json:"restrictions,omitempty"`
}

type PriceComponent struct {
	Type     TariffDimensionType `json:"type" validate:"required"`
	Price    json.Number         `json:"price" validate:"required"`
	StepSize int                 `json:"step_size"`
}

type TariffRestrictions struct {
	StartTime   *string      `json:"start_time" validate:"required"`
	EndTime     *string      `json:"end_time,omitempty"`
	StartDate   *string      `json:"start_date" validate:"required"`
	EndDate     *string      `json:"end_date,omitempty"`
	MinKwh      *json.Number `json:"min_kwh,omitempty"`      // Minimum kWh for this restriction.
	MaxKwh      *json.Number `json:"max_kwh,omitempty"`      // Maximum kWh for this restriction.
	MinPower    *json.Number `json:"min_power,omitempty"`    // Minimum power for this restriction.
	MaxPower    *json.Number `json:"max_power,omitempty"`    // Maximum power for this restriction.
	MinDuration *int         `json:"min_duration,omitempty"` // Minimum duration for this restriction.
	MaxDuration *int         `json:"max_duration,omitempty"` // Maximum duration for this restriction.
	DayOfWeek   []DayOfWeek  `json:"day_of_week,omitempty" validate:"omitempty,dive,dayOfWeek211"`
}

// CommandResponse defines model for commandResponse.
type CommandResponse struct {
	Result CommandResponseType `json:"result"`
}

// GeoLocation defines model for cdrBody_cdr_location_coordinates.
type GeoLocation struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// StatusSchedule defines model for evse_status_schedule.
type StatusSchedule struct {
	PeriodBegin DateTime  `json:"period_begin"`
	PeriodEnd   *DateTime `json:"period_end,omitempty"`
	Status      Status    `json:"status" validate:"status211"`
}

// AdditionalGeoLocation defines model for locations_data_related_locations.
type AdditionalGeoLocation struct {
	Latitude  string       `json:"latitude"`
	Longitude string       `json:"longitude"`
	Name      *DisplayText `json:"name,omitempty"`
}

// BusinessDetails defines model for businessDetails.
type BusinessDetails struct {
	Name    string  `json:"name"`
	Website *string `json:"website,omitempty"`
	Logo    *Image  `json:"logo,omitempty"`
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

type StartSession struct {
	ResponseURL string  `json:"response_url" validate:"required"`
	Token       Token   `json:"token" validate:"required"`
	LocationID  string  `json:"location_id" validate:"required,max=39"`
	EvseUID     *string `json:"evse_uid,omitempty" validate:"omitempty,required,max=39"`
}

type StopSession struct {
	ResponseURL string `json:"response_url" validate:"required"`
	SessionID   string `json:"session_id" validate:"required,max=36"`
}

type ReserveNow struct {
	ResponseURL   string   `json:"response_url" validate:"required"`
	Token         Token    `json:"token" validate:"required"`
	ExpiryDate    DateTime `json:"expiry_date" validate:"required"`
	ReservationID int      `json:"reservation_id" validate:"required"`
	LocationID    string   `json:"location_id" validate:"required,max=39"`
	EvseUID       *string  `json:"evse_uid,omitempty" validate:"omitempty,required,max=39"`
}

type UnlockConnector struct {
	ResponseURL string `json:"response_url" validate:"required"`
	LocationID  string `json:"location_id" validate:"required,max=39"`
	EvseUID     string `json:"evse_uid" validate:"required,max=39"`
	ConnectorID string `json:"connector_id" validate:"required,max=36"`
}

type Token struct {
	UID          string        `json:"uid" validate:"required"`
	Type         TokenType     `json:"type" validate:"required"`
	AuthID       string        `json:"auth_id" validate:"required"`
	VisualNumber *string       `json:"visual_number"`
	Issuer       string        `json:"issuer"`
	Valid        bool          `json:"valid"`
	Whitelist    WhitelistType `json:"whitelist"`
	Language     *string       `json:"language"`
	LastUpdated  DateTime      `json:"last_updated"`
}

// Authorization Changed name of the object from official docs due to colliding naming of info property
type AuthorizationInfo struct {
	Allowed  Allowed             `json:"allowed"`
	Location *LocationReferences `json:"location,omitempty"`
	Info     *DisplayText        `json:"info,omitempty"`
}

// LocationReferences defines model for locationReferences.
type LocationReferences struct {
	LocationID   string  `json:"location_id"`
	EvseUIDs     *string `json:"evse_uids,omitempty"`
	ConnectorIDs *string `json:"connector_ids,omitempty"`
}

type ChargeDetailRecordResponse struct {
	Location string
}

type PartialToken struct {
	UID          *string        `json:"uid,omitempty" validate:"omitempty,required"`
	Type         *TokenType     `json:"type,omitempty"`
	AuthID       *string        `json:"auth_id,omitempty" validate:"omitempty,required"`
	VisualNumber *string        `json:"visual_number,omitempty"`
	Issuer       *string        `json:"issuer,omitempty"`
	Valid        *bool          `json:"valid,omitempty"`
	Whitelist    *WhitelistType `json:"whitelist,omitempty"`
	Language     *string        `json:"language,omitempty" validate:"omitempty,required,max=2"`
	LastUpdated  DateTime       `json:"last_updated" validate:"required"` // LastUpdated is required for PATCH request
}

// GetLocationsParams defines parameters for GetOcpiLocations.
type GetLocationsParams = ocpi.PaginatedRequest[DateTime]

type GetCDRsParams = ocpi.PaginatedRequest[DateTime]

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

type GetTokensParams = ocpi.PaginatedRequest[DateTime]
