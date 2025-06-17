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

type VersionDetails struct {
	Endpoints []Endpoint         `json:"endpoints"`
	Version   ocpi.VersionNumber `json:"version"`
}

// Endpoint defines model for details_data_endpoints.
type Endpoint struct {
	// Identifier OCPI 2.1.1 modules
	Identifier ModuleID `json:"identifier"`

	// Url URL to the endpoint.
	URL string `json:"url"`
}

type Location struct {
	ID                 string                  `json:"id" validate:"required"`
	Type               LocationType            `json:"type" validate:"required"`
	Name               *string                 `json:"name"`
	Address            string                  `json:"address" validate:"required"`
	City               string                  `json:"city" validate:"required"`
	PostalCode         string                  `json:"postal_code" validate:"required,max=10"`
	Country            string                  `json:"country" validate:"required,len=3"`
	Coordinates        GeoLocation             `json:"coordinates"`
	RelatedLocations   []AdditionalGeoLocation `json:"related_locations,omitempty"`
	EVSEs              []EVSE                  `json:"evses,omitempty"`
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
	LastUpdated        ocpi.DateTime           `json:"last_updated"`
}

type Hours struct {
	RegularHours        []RegularHours      `json:"regular_hours,omitempty"`
	Twentyfourseven     bool                `json:"twentyfourseven"`
	ExceptionalOpenings []ExceptionalPeriod `json:"exceptional_openings,omitempty"`
	ExceptionalClosings []ExceptionalPeriod `json:"exceptional_closings,omitempty"`
}

// HoursRegularHours defines model for locations_data_opening_times_regular_hours.
type RegularHours struct {
	PeriodBegin string `json:"period_begin"`
	PeriodEnd   string `json:"period_end"`
	Weekday     int    `json:"weekday"`
}

// HoursExceptionalOpenings defines model for locations_data_opening_times_exceptional_openings.
type ExceptionalPeriod struct {
	PeriodBegin ocpi.DateTime `json:"period_begin"`
	PeriodEnd   ocpi.DateTime `json:"period_end"`
}

type PatchedLocation struct {
	ID                 *string                 `json:"id" validate:"required"`
	Type               *LocationType           `json:"type" validate:"required"`
	Name               *string                 `json:"name"`
	Address            *string                 `json:"address" validate:"required"`
	City               *string                 `json:"city" validate:"required"`
	PostalCode         *string                 `json:"postal_code" validate:"required,max=10"`
	Country            *string                 `json:"country" validate:"required,len=3"`
	Coordinates        *GeoLocation            `json:"coordinates"`
	RelatedLocations   []AdditionalGeoLocation `json:"related_locations,omitempty"`
	EVSEs              []EVSE                  `json:"evses,omitempty"`
	Directions         []DisplayText           `json:"directions,omitempty"`
	Operator           *BusinessDetails        `json:"operator,omitempty"`
	Suboperator        *BusinessDetails        `json:"suboperator,omitempty"`
	Owner              *BusinessDetails        `json:"owner,omitempty"`
	Facilities         []Facility              `json:"facilities,omitempty"`
	TimeZone           *string                 `json:"time_zone" validate:"required"`
	OpeningTimes       *Hours                  `json:"opening_times,omitempty"`
	ChargingWhenClosed *bool                   `json:"charging_when_closed,omitempty"`
	Images             []Image                 `json:"images,omitempty"`
	EnergyMix          *EnergyMix              `json:"energy_mix,omitempty"`
	LastUpdated        *ocpi.DateTime          `json:"last_updated"`
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
	LastUpdated         ocpi.DateTime        `json:"last_updated"`
}

// Connector defines model for connector.
type Connector struct {
	ID                 string          `json:"id" validate:"required"`
	Standard           ConnectorType   `json:"standard" validate:"required"`
	Format             ConnectorFormat `json:"format" validate:"required"`
	PowerType          PowerType       `json:"power_type" validate:"required"`
	Voltage            int             `json:"voltage"`
	Amperage           int             `json:"amperage"`
	TariffID           *string         `json:"tariff_id,omitempty"`
	TermsAndConditions *string         `json:"terms_and_conditions,omitempty"`
	LastUpdated        ocpi.DateTime   `json:"last_updated"`
}

type EnergyMix struct {
	IsGreenEnergy     bool                  `json:"is_green_energy"`
	EnergySources     []EnergySource        `json:"energy_sources,omitempty"`
	EnvironImpact     []EnvironmentalImpact `json:"environ_impact,omitempty"`
	SupplierName      *string               `json:"supplier_name,omitempty"`
	EnergyProductName *string               `json:"energy_product_name,omitempty"`
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

// GeoLocation defines model for cdrBody_cdr_location_coordinates.
type GeoLocation struct {
	Latitude  float64 `json:"latitude,string"`
	Longitude float64 `json:"longitude,string"`
}

// StatusSchedule defines model for evse_status_schedule.
type StatusSchedule struct {
	PeriodBegin ocpi.DateTime  `json:"period_begin"`
	PeriodEnd   *ocpi.DateTime `json:"period_end,omitempty"`
	Status      Status         `json:"status"`
}

// AdditionalGeoLocation defines model for locations_data_related_locations.
type AdditionalGeoLocation struct {
	Latitude  float64      `json:"latitude,string"`
	Longitude float64      `json:"longitude,string"`
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

// GetOcpiLocationsParams defines parameters for GetOcpiLocations.
type GetOcpiLocationsParams struct {
	// DateFrom Return Locations that have last_updated after or equal to this date time (inclusive).
	DateFrom *string `form:"date_from,omitempty" json:"date_from,omitempty"`

	// DateTo Return Locations that have last_updated up to this date time, but not including (exclusive).
	DateTo *string `form:"date_to,omitempty" json:"date_to,omitempty"`

	// Offset The offset of the first object returned. Default is 0.
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit Maximum number of objects to GET.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}
