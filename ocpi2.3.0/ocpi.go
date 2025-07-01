package ocpi230

import (
	"encoding/json"
)

type Role string

// Defines values for Role.
const (
	RoleCPO   Role = "CPO"
	RoleEMSP  Role = "EMSP"
	RoleNAP   Role = "NAP"
	RoleNSP   Role = "NSP"
	RoleOther Role = "OTHER"
	RoleSCSP  Role = "SCSP"
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

// LocationsDataFacilities defines model for LocationsData.Facilities.
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
	FacilityBikeSharing    Facility = "BIKE_SHARING"
	FacilityBusStop        Facility = "BUS_STOP"
	FacilityTaxiStand      Facility = "TAXI_STAND"
	FacilityTramStop       Facility = "TRAM_STOP"
	FacilityMetroStation   Facility = "METRO_STATION"
	FacilityTrainStation   Facility = "TRAIN_STATION"
	FacilityAirport        Facility = "AIRPORT"
	FacilityParkingLot     Facility = "PARKING_LOT"
	FacilityCarpoolParking Facility = "CARPOOL_PARKING"
	FacilityFuelStation    Facility = "FUEL_STATION"
	FacilityWiFi           Facility = "WIFI"
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

type BookingLocation struct {
	CountryCode            string                   `json:"country_code" validate:"required,len=2"`
	PartyID                string                   `json:"party_id" validate:"required"`
	ID                     string                   `json:"id" validate:"required"`
	LocationID             string                   `json:"location_id" validate:"required"`
	EVSEUID                *string                  `json:"evse_uid,omitempty"`
	ConnectorID            *string                  `json:"connector_id,omitempty"`
	BookableParkingOptions []BookableParkingOptions `json:"bookable_parking_options,omitempty"`
	Bookable               *Bookable                `json:"bookable,omitempty"`
	TariffID               []string                 `json:"tariff_id,omitempty"`
	BookingTerms           []BookingTerms           `json:"booking_terms,omitempty"`
	Calendars              []Calendar               `json:"calendars,omitempty"`
	LastUpdated            DateTime                 `json:"last_updated" validate:"required"`
}

type BookableParkingOptions struct {
}

type BookingTerms struct {
}

type Calendar struct {
}

type Timeslot struct {
	StartFrom          DateTime     `json:"start_from" validate:"required"`
	EndBefore          DateTime     `json:"end_before" validate:"required"`
	MinPower           *json.Number `json:"min_power,omitempty"`
	MaxPower           *json.Number `json:"max_power,omitempty"`
	GreenEnergySupport *bool        `json:"green_energy_support,omitempty"`
}

type Bookable struct {
	ReservationRequired bool         `json:"reservation_required"`
	AdHoc               *json.Number `json:"ad_hoc,omitempty"`
}

type Location struct {
	CountryCode        string                  `json:"country_code" validate:"required"`
	PartyID            string                  `json:"party_id" validate:"required"`
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
	HelpPhone          *string                 `json:"help_phone,omitempty"`
}

type PartialLocation struct {
}

// PublishTokenType defines model for locations_data_publish_allowed_to.
type PublishTokenType struct {
	UID          *string    `json:"uid,omitempty"`
	Type         *TokenType `json:"type,omitempty"`
	VisualNumber *string    `json:"visual_number,omitempty"`
	Issuer       *string    `json:"issuer,omitempty"`
	GroupID      *string    `json:"group_id,omitempty"`
}

// GeoLocation defines model for cdrBody_cdr_location_coordinates.
type GeoLocation struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// AdditionalGeoLocation defines model for locations_data_related_locations.
type AdditionalGeoLocation struct {
	Latitude  string       `json:"latitude"`
	Longitude string       `json:"longitude"`
	Name      *DisplayText `json:"name,omitempty"`
}

// BusinessDetails defines model for businessDetails.
type BusinessDetails struct {
	Logo    *Image  `json:"logo,omitempty"`
	Name    string  `json:"name"`
	Website *string `json:"website,omitempty"`
}

// Image defines model for credentials_data_roles_business_details_logo.
type Image struct {
	URL       string        `json:"url" validate:"required"`
	Thumbnail *string       `json:"thumbnail,omitempty"`
	Category  ImageCategory `json:"category" validate:"required"`
	Type      string        `json:"type" validate:"required,max=4"`
	Width     *int          `json:"width,omitempty"`
	Height    *int          `json:"height,omitempty"`
}

// Hours defines model for locations_data_opening_times.
type Hours struct {
	ExceptionalClosings []ExceptionalPeriod `json:"exceptional_closings,omitempty"`
	ExceptionalOpenings []ExceptionalPeriod `json:"exceptional_openings,omitempty"`
	RegularHours        []RegularHours      `json:"regular_hours,omitempty"`
	Twentyfourseven     bool                `json:"twentyfourseven"`
}

// HoursRegularHours defines model for locations_data_opening_times_regular_hours.
type RegularHours struct {
	PeriodBegin string `json:"period_begin"`
	PeriodEnd   string `json:"period_end"`
	Weekday     int    `json:"weekday"`
}

// HoursExceptionalOpenings defines model for locations_data_opening_times_exceptional_openings.
type ExceptionalPeriod struct {
	PeriodBegin DateTime `json:"period_begin"`
	PeriodEnd   DateTime `json:"period_end"`
}

type EVSE struct {
}

type Parking struct {
}

type EnergyMix struct {
	IsGreenEnergy     bool                  `json:"is_green_energy"`
	EnergySources     []EnergySource        `json:"energy_sources,omitempty"`
	EnvironImpact     []EnvironmentalImpact `json:"environ_impact,omitempty"`
	SupplierName      *string               `json:"supplier_name,omitempty"`
	EnergyProductName *string               `json:"energy_product_name,omitempty"`
}

// EnergySource defines model for cdrBody_tariffs_energy_mix_energy_sources.
type EnergySource struct {
	Percentage json.Number          `json:"percentage"`
	Source     EnergySourceCategory `json:"source"`
}

// EnvironmentalImpact defines model for cdrBody_tariffs_energy_mix_environ_impact.
type EnvironmentalImpact struct {
	Amount   json.Number                 `json:"amount"`
	Category EnvironmentalImpactCategory `json:"category"`
}

type Connector struct {
}
