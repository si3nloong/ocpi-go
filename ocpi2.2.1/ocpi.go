package ocpi221

import (
	"encoding/json"
	"net/url"

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
	RoleOTHER Role = "OTHER"
	RoleSCSP  Role = "SCSP"
)

// InterfaceRoleType Interface role endpoint implements.
type InterfaceRoleType string

// Defines values for InterfaceRoleType.
const (
	RoleReceiver InterfaceRoleType = "RECEIVER"
	RoleSender   InterfaceRoleType = "SENDER"
)

// ModuleIdentifier OCPI 2.2.1 modules
type ModuleIdentifier string

// Defines values for ModuleIdentifier.
const (
	ModuleIdentifierCdrs             ModuleIdentifier = "cdrs"
	ModuleIdentifierChargingProfiles ModuleIdentifier = "chargingprofiles"
	ModuleIdentifierCommands         ModuleIdentifier = "commands"
	ModuleIdentifierCredentials      ModuleIdentifier = "credentials"
	ModuleIdentifierHubClientInfo    ModuleIdentifier = "hubclientinfo"
	ModuleIdentifierLocations        ModuleIdentifier = "locations"
	ModuleIdentifierSessions         ModuleIdentifier = "sessions"
	ModuleIdentifierTariffs          ModuleIdentifier = "tariffs"
	ModuleIdentifierTokens           ModuleIdentifier = "tokens"
)

// ChargingProfileResultType defines model for ActiveChargingProfileResult.Result.
type ChargingProfileResultType string

// Defines values for ChargingProfileResultType.
const (
	ChargingProfileResultTypeAccepted ChargingProfileResultType = "ACCEPTED"
	ChargingProfileResultTypeRejected ChargingProfileResultType = "REJECTED"
	ChargingProfileResultTypeUknown   ChargingProfileResultType = "UNKNOWN"
)

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

// Defines values for EnergySourceCategory.
const (
	EnergySourceCategoryCoal          EnergySourceCategory = "COAL"
	EnergySourceCategoryGas           EnergySourceCategory = "GAS"
	EnergySourceCategoryGeneralFossil EnergySourceCategory = "GENERAL_FOSSIL"
	EnergySourceCategoryGeneralGreen  EnergySourceCategory = "GENERAL_GREEN"
	EnergySourceCategoryNuclear       EnergySourceCategory = "NUCLEAR"
	EnergySourceCategorySolar         EnergySourceCategory = "SOLAR"
	EnergySourceCategoryWater         EnergySourceCategory = "WATER"
	EnergySourceCategoryWind          EnergySourceCategory = "WIND"
)

// Defines values for EnvironmentalImpactCategory.
const (
	EnvironmentalImpactCategoryCarbonDioxide EnvironmentalImpactCategory = "CARBON_DIOXIDE"
	EnvironmentalImpactCategoryNuclearWaste  EnvironmentalImpactCategory = "NUCLEAR_WASTE"
)

// Defines values for ChargingPreferencesProfileType.
const (
	ChargingPreferencesProfileTypeCHEAP   ChargingPreferencesProfileType = "CHEAP"
	ChargingPreferencesProfileTypeFAST    ChargingPreferencesProfileType = "FAST"
	ChargingPreferencesProfileTypeGREEN   ChargingPreferencesProfileType = "GREEN"
	ChargingPreferencesProfileTypeREGULAR ChargingPreferencesProfileType = "REGULAR"
)

// Defines values for ChargingPreferencesResponseChargingPreferences.
const (
	ChargingPreferencesResponseChargingPreferencesACCEPTED                ChargingPreferencesResponseChargingPreferences = "ACCEPTED"
	ChargingPreferencesResponseChargingPreferencesDEPARTUREREQUIRED       ChargingPreferencesResponseChargingPreferences = "DEPARTURE_REQUIRED"
	ChargingPreferencesResponseChargingPreferencesENERGYNEEDREQUIRED      ChargingPreferencesResponseChargingPreferences = "ENERGY_NEED_REQUIRED"
	ChargingPreferencesResponseChargingPreferencesNotPossible             ChargingPreferencesResponseChargingPreferences = "NOT_POSSIBLE"
	ChargingPreferencesResponseChargingPreferencesPROFILETYPENOTSUPPORTED ChargingPreferencesResponseChargingPreferences = "PROFILE_TYPE_NOT_SUPPORTED"
)

// Defines values for ChargingRateUnit.
const (
	ChargingRateUnitAmperes ChargingRateUnit = "A"
	ChargingRateUnitWatts   ChargingRateUnit = "W"
)

// Defines values for ChargingProfileResponseType.
const (
	ChargingProfileResponseTypeAccepted       ChargingProfileResponseType = "ACCEPTED"
	ChargingProfileResponseTypeNotSupported   ChargingProfileResponseType = "NOT_SUPPORTED"
	ChargingProfileResponseTypeRejected       ChargingProfileResponseType = "REJECTED"
	ChargingProfileResponseTypeTooOften       ChargingProfileResponseType = "TOO_OFTEN"
	ChargingProfileResponseTypeUnknownSession ChargingProfileResponseType = "UNKNOWN_SESSION"
)

// Defines values for ChargingProfileResultResult.
const (
	ChargingProfileResultResultAccepted ChargingProfileResultResult = "ACCEPTED"
	ChargingProfileResultResultRejected ChargingProfileResultResult = "REJECTED"
	ChargingProfileResultResultUnknown  ChargingProfileResultResult = "UNKNOWN"
)

// Defines values for ClearProfileResultResult.
const (
	ClearProfileResultResultAccepted ClearProfileResultResult = "ACCEPTED"
	ClearProfileResultResultRejected ClearProfileResultResult = "REJECTED"
	ClearProfileResultResultUnknown  ClearProfileResultResult = "UNKNOWN"
)

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

// Defines values for ConnectorFormat.
const (
	ConnectorFormatCable  ConnectorFormat = "CABLE"
	ConnectorFormatSocket ConnectorFormat = "SOCKET"
)

// Defines values for PowerType.
const (
	PowerTypeAC1Phase      PowerType = "AC_1_PHASE"
	PowerTypeAC2Phase      PowerType = "AC_2_PHASE"
	PowerTypeAC2PhaseSplit PowerType = "AC_2_PHASE_SPLIT"
	PowerTypeAC3Phase      PowerType = "AC_3_PHASE"
	PowerTypeDC            PowerType = "DC"
)

// Defines values for ConnectorStandard.
const (
	ConnectorStandardCHAdeMO            ConnectorStandard = "CHADEMO"
	ConnectorStandardChaoJi             ConnectorStandard = "CHAOJI"
	ConnectorStandardDomesticA          ConnectorStandard = "DOMESTIC_A"
	ConnectorStandardDomesticB          ConnectorStandard = "DOMESTIC_B"
	ConnectorStandardDomesticC          ConnectorStandard = "DOMESTIC_C"
	ConnectorStandardDomesticD          ConnectorStandard = "DOMESTIC_D"
	ConnectorStandardDomesticE          ConnectorStandard = "DOMESTIC_E"
	ConnectorStandardDomesticF          ConnectorStandard = "DOMESTIC_F"
	ConnectorStandardDomesticG          ConnectorStandard = "DOMESTIC_G"
	ConnectorStandardDomesticH          ConnectorStandard = "DOMESTIC_H"
	ConnectorStandardDomesticJ          ConnectorStandard = "DOMESTIC_J"
	ConnectorStandardDomesticK          ConnectorStandard = "DOMESTIC_K"
	ConnectorStandardDomesticL          ConnectorStandard = "DOMESTIC_L"
	ConnectorStandardDomesticM          ConnectorStandard = "DOMESTIC_M"
	ConnectorStandardDomesticN          ConnectorStandard = "DOMESTIC_N"
	ConnectorStandardDomesticO          ConnectorStandard = "DOMESTIC_O"
	ConnectorStandardGBTAC              ConnectorStandard = "GBT_AC"
	ConnectorStandardGBTDC              ConnectorStandard = "GBT_DC"
	ConnectorStandardIEC603092Single16  ConnectorStandard = "IEC_60309_2_single_16"
	ConnectorStandardIEC603092Three16   ConnectorStandard = "IEC_60309_2_three_16"
	ConnectorStandardIEC603092Three32   ConnectorStandard = "IEC_60309_2_three_32"
	ConnectorStandardIEC603092Three64   ConnectorStandard = "IEC_60309_2_three_64"
	ConnectorStandardIEC62196T1         ConnectorStandard = "IEC_62196_T1"
	ConnectorStandardIEC62196T1Combo    ConnectorStandard = "IEC_62196_T1_COMBO"
	ConnectorStandardIEC62196T2         ConnectorStandard = "IEC_62196_T2"
	ConnectorStandardIEC62196T2Combo    ConnectorStandard = "IEC_62196_T2_COMBO"
	ConnectorStandardIEC62196T3A        ConnectorStandard = "IEC_62196_T3A"
	ConnectorStandardIEC62196T3C        ConnectorStandard = "IEC_62196_T3C"
	ConnectorStandardNema1030           ConnectorStandard = "NEMA_10_30"
	ConnectorStandardNema1050           ConnectorStandard = "NEMA_10_50"
	ConnectorStandardNema1430           ConnectorStandard = "NEMA_14_30"
	ConnectorStandardNema1450           ConnectorStandard = "NEMA_14_50"
	ConnectorStandardNema520            ConnectorStandard = "NEMA_5_20"
	ConnectorStandardNema630            ConnectorStandard = "NEMA_6_30"
	ConnectorStandardNema650            ConnectorStandard = "NEMA_6_50"
	ConnectorStandardPantographBottomUp ConnectorStandard = "PANTOGRAPH_BOTTOM_UP"
	ConnectorStandardPantographTopDown  ConnectorStandard = "PANTOGRAPH_TOP_DOWN"
	ConnectorStandardTeslaR             ConnectorStandard = "TESLA_R"
	ConnectorStandardTeslaS             ConnectorStandard = "TESLA_S"
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

// Defines values for EvseStatus.
const (
	EvseStatusAvailable   EvseStatus = "AVAILABLE"
	EvseStatusBlocked     EvseStatus = "BLOCKED"
	EvseStatusCharging    EvseStatus = "CHARGING"
	EvseStatusInOperative EvseStatus = "INOPERATIVE"
	EvseStatusOutOfOrder  EvseStatus = "OUTOFORDER"
	EvseStatusPlanned     EvseStatus = "PLANNED"
	EvseStatusRemoved     EvseStatus = "REMOVED"
	EvseStatusReserved    EvseStatus = "RESERVED"
	EvseStatusUnknown     EvseStatus = "UNKNOWN"
)

// Defines values for StatusScheduleStatus.
const (
	StatusScheduleStatusAvailable   StatusScheduleStatus = "AVAILABLE"
	StatusScheduleStatusBlocked     StatusScheduleStatus = "BLOCKED"
	StatusScheduleStatusCharging    StatusScheduleStatus = "CHARGING"
	StatusScheduleStatusInOperative StatusScheduleStatus = "INOPERATIVE"
	StatusScheduleStatusOutOfOrder  StatusScheduleStatus = "OUTOFORDER"
	StatusScheduleStatusPlanned     StatusScheduleStatus = "PLANNED"
	StatusScheduleStatusRemoved     StatusScheduleStatus = "REMOVED"
	StatusScheduleStatusReserved    StatusScheduleStatus = "RESERVED"
	StatusScheduleStatusUnknown     StatusScheduleStatus = "UNKNOWN"
)

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
	FacilityRECREATIONAREA Facility = "RECREATION_AREA"
	FacilityRestaurant     Facility = "RESTAURANT"
	FacilitySport          Facility = "SPORT"
	FacilitySuperMarket    Facility = "SUPERMARKET"
	FacilityTaxiStand      Facility = "TAXI_STAND"
	FacilityTrainStation   Facility = "TRAIN_STATION"
	FacilityTramStop       Facility = "TRAM_STOP"
	FacilityWiFi           Facility = "WIFI"
)

// Defines values for ParkingType.
const (
	ParkingTypeALONGMOTORWAY     ParkingType = "ALONG_MOTORWAY"
	ParkingTypeONDRIVEWAY        ParkingType = "ON_DRIVEWAY"
	ParkingTypeOnStreet          ParkingType = "ON_STREET"
	ParkingTypePARKINGGARAGE     ParkingType = "PARKING_GARAGE"
	ParkingTypeParkingLot        ParkingType = "PARKING_LOT"
	ParkingTypeUNDERGROUNDGARAGE ParkingType = "UNDERGROUND_GARAGE"
)

// Defines values for PublishTokenTypeType.
const (
	PublishTokenTypeTypeAdhocUser PublishTokenTypeType = "AD_HOC_USER"
	PublishTokenTypeTypeAppUser   PublishTokenTypeType = "APP_USER"
	PublishTokenTypeTypeOther     PublishTokenTypeType = "OTHER"
	PublishTokenTypeTypeRFID      PublishTokenTypeType = "RFID"
)

// Defines values for ReservationRestrictionType.
const (
	ReservationRestrictionTypeRESERVATION        ReservationRestrictionType = "RESERVATION"
	ReservationRestrictionTypeRESERVATIONEXPIRES ReservationRestrictionType = "RESERVATION_EXPIRES"
)

// Defines values for AuthMethod.
const (
	AuthMethodAuthRequest AuthMethod = "AUTH_REQUEST"
	AuthMethodCommand     AuthMethod = "COMMAND"
	AuthMethodWhitelist   AuthMethod = "WHITELIST"
)

// Defines values for SessionStatus.
const (
	ACTIVE      SessionStatus = "ACTIVE"
	COMPLETED   SessionStatus = "COMPLETED"
	INVALID     SessionStatus = "INVALID"
	PENDING     SessionStatus = "PENDING"
	RESERVATION SessionStatus = "RESERVATION"
)

// Defines values for TariffType.
const (
	TariffTypeAdHocPayment TariffType = "AD_HOC_PAYMENT"
	TariffTypeProfileCheap TariffType = "PROFILE_CHEAP"
	TariffTypeProfileFast  TariffType = "PROFILE_FAST"
	TariffTypeProfileGreen TariffType = "PROFILE_GREEN"
	TariffTypeRegular      TariffType = "REGULAR"
)

type PriceComponentType string

// Defines values for PriceComponent.
const (
	PriceComponentTypeEnergy      PriceComponentType = "ENERGY"
	PriceComponentTypeFlat        PriceComponentType = "FLAT"
	PriceComponentTypeParkingTime PriceComponentType = "PARKING_TIME"
	PriceComponentTypeTime        PriceComponentType = "TIME"
)

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

type VersionsResponse = ocpi.Response[[]Version]

// ActiveChargingProfileResult defines model for activeChargingProfileResult.
type ActiveChargingProfileResult struct {
	Profile *ActiveChargingProfile    `json:"profile,omitempty"`
	Result  ChargingProfileResultType `json:"result"`
}

// ActiveChargingProfile defines model for activeChargingProfile.
type ActiveChargingProfile struct {
	ChargingProfile ChargingProfile `json:"charging_profile"`
	StartDateTime   string          `json:"start_date_time"`
}

// Authorization Changed name of the object from official docs due to colliding naming of info property
type AuthorizationInfo struct {
	Allowed                AuthorizationAllowed    `json:"allowed"`
	AuthorizationReference *string                 `json:"authorization_reference,omitempty"`
	Info                   *CommandResponseMessage `json:"info,omitempty"`
	Location               *LocationReferences     `json:"location,omitempty"`
	Token                  Token                   `json:"token"`
}

// AuthorizationAllowed defines model for Authorization.Allowed.
type AuthorizationAllowed string

// BusinessDetails defines model for businessDetails.
type BusinessDetails struct {
	Logo    *Image  `json:"logo,omitempty"`
	Name    string  `json:"name"`
	Website *string `json:"website,omitempty"`
}

// CancelReservation defines model for cancelReservation.
type CancelReservation struct {
	ReservationID string `json:"reservation_id"`
	ResponseURL   string `json:"response_url"`
}

// CdrBody defines model for cdrBody.
type ChargeDetailRecord struct {
	CountryCode              string           `json:"country_code" validate:"required,len=2"`
	PartyID                  string           `json:"party_id" validate:"required,len=3"`
	ID                       string           `json:"id" validate:"required"`
	StartDateTime            ocpi.DateTime    `json:"start_date_time" validate:"required"`
	EndDateTime              ocpi.DateTime    `json:"end_date_time" validate:"required"`
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
	LastUpdated              ocpi.DateTime    `json:"last_updated" validate:"required"`
}

// AuthMethod defines model for CdrBody.AuthMethod.
type AuthMethod string

// CdrLocation defines model for cdrBody_cdr_location.
type CdrLocation struct {
	ID                 string            `json:"id" validate:"required"`
	Name               *string           `json:"name,omitempty"`
	Address            string            `json:"address" validate:"required"`
	City               string            `json:"city" validate:"required"`
	PostalCode         *string           `json:"postal_code,omitempty"`
	State              *string           `json:"state,omitempty"`
	Country            string            `json:"country" validate:"required,len=3"`
	Coordinates        GeoLocation       `json:"coordinates"`
	EvseUID            string            `json:"evse_uid" validate:"required,max=36"`
	EvseID             string            `json:"evse_id" validate:"required,max=48"`
	ConnectorID        string            `json:"connector_id"`
	ConnectorStandard  ConnectorStandard `json:"connector_standard" validate:"required"`
	ConnectorFormat    ConnectorFormat   `json:"connector_format" validate:"required"`
	ConnectorPowerType PowerType         `json:"connector_power_type" validate:"required"`
}

// GeoLocation defines model for cdrBody_cdr_location_coordinates.
type GeoLocation struct {
	Latitude  json.Number `json:"latitude"`
	Longitude json.Number `json:"longitude"`
}

// CdrToken defines model for cdrBody_cdr_token.
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

// CdrBodyTariffsElementsRestrictions defines model for cdrBody_tariffs_elements_restrictions.
type CdrBodyTariffsElementsRestrictions struct {
	DayOfWeek   *DayOfWeek                  `json:"day_of_week,omitempty"`
	EndDate     *string                     `json:"end_date,omitempty"`
	EndTime     *string                     `json:"end_time,omitempty"`
	MaxCurrent  *float32                    `json:"max_current,omitempty"`
	MaxDuration *int                        `json:"max_duration,omitempty"`
	MaxKwh      *float32                    `json:"max_kwh,omitempty"`
	MaxPower    *float32                    `json:"max_power,omitempty"`
	MinCurrent  *float32                    `json:"min_current,omitempty"`
	MinDuration *int                        `json:"min_duration,omitempty"`
	MinKwh      *float32                    `json:"min_kwh,omitempty"`
	MinPower    *float32                    `json:"min_power,omitempty"`
	Reservation *ReservationRestrictionType `json:"reservation,omitempty"`
	StartDate   *string                     `json:"start_date,omitempty"`
	StartTime   *string                     `json:"start_time,omitempty"`
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
	Percentage json.Number          `json:"percentage"`
	Source     EnergySourceCategory `json:"source"`
}

// EnergySourceCategory defines model for EnergySource.Source.
type EnergySourceCategory string

// EnvironmentalImpact defines model for cdrBody_tariffs_energy_mix_environ_impact.
type EnvironmentalImpact struct {
	Amount   json.Number                 `json:"amount"`
	Category EnvironmentalImpactCategory `json:"category"`
}

// EnvironmentalImpactCategory defines model for EnvironmentalImpact.Category.
type EnvironmentalImpactCategory string

// DisplayText defines model for cdrBody_tariffs_tariff_alt_text.
type DisplayText struct {
	Language string `json:"language"`
	Text     string `json:"text"`
}

// CdrResponse defines model for cdrResponse.
type CdrResponse struct {
	StatusCode    float32 `json:"status_code"`
	StatusMessage *string `json:"status_message,omitempty"`
	TimeStamp     *string `json:"timeStamp,omitempty"`
	URL           string  `json:"url"`
}

// CdrsResponse defines model for cdrsResponse.
type CdrsResponse struct {
	Cdrs          *ChargeDetailRecord `json:"cdrs,omitempty"`
	StatusCode    float32             `json:"status_code"`
	StatusMessage *string             `json:"status_message,omitempty"`
	TimeStamp     *string             `json:"timeStamp,omitempty"`
}

// ChargingPreferences defines model for chargingPreferences.
type ChargingPreferences struct {
	DepartureTime *string                        `json:"departure_time,omitempty"`
	EnergyNeed    *float32                       `json:"energy_need,omitempty"`
	ProfileType   ChargingPreferencesProfileType `json:"profile_type"`
}

// ChargingPreferencesProfileType defines model for ChargingPreferences.ProfileType.
type ChargingPreferencesProfileType string

// ChargingPreferencesResponse defines model for chargingPreferencesResponse.
type ChargingPreferencesResponse struct {
	ChargingPreferences ChargingPreferencesResponseChargingPreferences `json:"charging_preferences"`
	StatusCode          float32                                        `json:"status_code"`
	StatusMessage       *string                                        `json:"status_message,omitempty"`
	TimeStamp           *string                                        `json:"timeStamp,omitempty"`
}

// ChargingPreferencesResponseChargingPreferences defines model for ChargingPreferencesResponse.ChargingPreferences.
type ChargingPreferencesResponseChargingPreferences string

// ChargingProfile defines model for chargingProfile.
type ChargingProfile struct {
	ChargingProfilePeriod []ChargingProfilePeriod `json:"charging_profile_period,omitempty"`
	ChargingRateUnit      ChargingRateUnit        `json:"charging_rate_unit"`
	Duration              *int                    `json:"duration,omitempty"`
	MinChargingRate       *json.Number            `json:"min_charging_rate,omitempty"`
	StartDateTime         *ocpi.DateTime          `json:"start_date_time,omitempty"`
}

// ChargingRateUnit defines model for ChargingProfile.ChargingRateUnit.
type ChargingRateUnit string

// ChargingProfileResponse defines model for chargingProfileResponse.
type ChargingProfileResponse struct {
	// Result Response to the ChargingProfile request from the eMSP to the CPO.
	Result  ChargingProfileResponseType `json:"result"`
	Timeout int                         `json:"timeout"`
}

// ChargingProfileResponseType Response to the ChargingProfile request from the eMSP to the CPO.
type ChargingProfileResponseType string

// ChargingProfileResult defines model for chargingProfileResult.
type ChargingProfileResult struct {
	Result ChargingProfileResultResult `json:"result"`
}

// ChargingProfileResultResult defines model for ChargingProfileResult.Result.
type ChargingProfileResultResult string

// ChargingProfilePeriod defines model for chargingProfile_charging_profile_period.
type ChargingProfilePeriod struct {
	StartPeriod int         `json:"start_period"`
	Limit       json.Number `json:"limit"`
}

// ChargingProfilesResponse defines model for chargingProfilesResponse.
type ChargingProfilesResponse struct {
	ChargingProfile *ChargingProfileResponse `json:"chargingProfile,omitempty"`
	StatusCode      float32                  `json:"status_code"`
	StatusMessage   *string                  `json:"status_message,omitempty"`
	TimeStamp       *string                  `json:"timeStamp,omitempty"`
}

// ClearProfileResult defines model for clearProfileResult.
type ClearProfileResult struct {
	Result ClearProfileResultResult `json:"result"`
}

// ClearProfileResultResult defines model for ClearProfileResult.Result.
type ClearProfileResultResult string

// ClientInfo defines model for clientInfo.
type ClientInfo struct {
	PartyID     string           `json:"party_id"`
	CountryCode string           `json:"country_code"`
	Role        Role             `json:"role"`
	Status      ConnectionStatus `json:"status"`
	LastUpdated ocpi.DateTime    `json:"last_updated"`
}

// ConnectionStatus defines model for ClientInfo.Status.
type ConnectionStatus string

// ClientInfoResponse defines model for clientInfoResponse.
type ClientInfoResponse struct {
	ClientsInfo   *ClientInfo `json:"clients_info,omitempty"`
	StatusCode    float32     `json:"status_code"`
	StatusMessage *string     `json:"status_message,omitempty"`
	TimeStamp     *string     `json:"timeStamp,omitempty"`
}

// ClientsInfoResponse defines model for clientsInfoResponse.
type ClientsInfoResponse struct {
	ClientInfo    ClientInfo `json:"client_info"`
	StatusCode    float32    `json:"status_code"`
	StatusMessage *string    `json:"status_message,omitempty"`
	TimeStamp     *string    `json:"timeStamp,omitempty"`
}

// CommandResponse defines model for commandResponse.
type CommandResponse struct {
	Message *CommandResponseMessage `json:"message,omitempty"`
	Result  CommandResponseType     `json:"result"`
	Timeout int                     `json:"timeout"`
}

// CommandResponseType defines model for CommandResponse.Result.
type CommandResponseType string

// CommandResponseMessage defines model for commandResponse_message.
type CommandResponseMessage struct {
	Language string `json:"language"`
	Text     string `json:"text"`
}

// CommandResult defines model for commandResult.
type CommandResult struct {
	Message *CommandResponseMessage `json:"message,omitempty"`
	Result  CommandResultType       `json:"result"`
}

// CommandResultType defines model for CommandResult.Result.
type CommandResultType string

type ConnectorCapability string

// Connector defines model for connector.
type Connector struct {
	Format             ConnectorFormat       `json:"format"`
	Id                 string                `json:"id"`
	LastUpdated        ocpi.DateTime         `json:"last_updated"`
	MaxAmperage        int                   `json:"max_amperage"`
	MaxElectricPower   *int                  `json:"max_electric_power,omitempty"`
	MaxVoltage         int                   `json:"max_voltage"`
	PowerType          PowerType             `json:"power_type"`
	Standard           ConnectorStandard     `json:"standard"`
	Capabilities       []ConnectorCapability `json:"capabilities"`
	TariffIds          []string              `json:"tariff_ids,omitempty"`
	TermsAndConditions *string               `json:"terms_and_conditions,omitempty"`
}

// ConnectorFormat defines model for Connector.Format.
type ConnectorFormat string

// PowerType defines model for Connector.PowerType.
type PowerType string

// ConnectorStandard defines model for Connector.Standard.
type ConnectorStandard string

// CredentialsData defines model for credentials_data.
type CredentialsData struct {
	Roles []CredentialsDataRoles `json:"roles,omitempty"`
	Token string                 `json:"token"`
	Url   string                 `json:"url"`
}

// CredentialsDataRoles defines model for credentials_data_roles.
type CredentialsDataRoles struct {
	BusinessDetails CredentialsDataRolesBusinessDetails `json:"business_details"`
	CountryCode     string                              `json:"country_code"`
	PartyID         string                              `json:"party_id"`
	Role            Role                                `json:"role"`
}

// CredentialsDataRolesBusinessDetails defines model for credentials_data_roles_business_details.
type CredentialsDataRolesBusinessDetails struct {
	Logo    *Image  `json:"logo,omitempty"`
	Name    string  `json:"name"`
	Website *string `json:"website,omitempty"`
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

// ImageCategory defines model for Image.Category.
type ImageCategory string

// Details defines model for details.
type Details struct {
	Data          *DetailsData `json:"data,omitempty"`
	StatusCode    float32      `json:"status_code"`
	StatusMessage *string      `json:"status_message,omitempty"`
	TimeStamp     *string      `json:"timeStamp,omitempty"`
}

// DetailsData defines model for details_data.
type DetailsData struct {
	Endpoints []DetailsDataEndpoints `json:"endpoints,omitempty"`
	Version   DetailsDataVersion     `json:"version"`
}

// DetailsDataVersion defines model for DetailsData.Version.
type DetailsDataVersion string

// DetailsDataEndpoints defines model for details_data_endpoints.
type DetailsDataEndpoints struct {
	// Identifier OCPI 2.2.1 modules
	Identifier ModuleIdentifier `json:"identifier"`

	// Role Interface role endpoint implements.
	Role InterfaceRoleType `json:"role"`

	// Url URL to the endpoint.
	Url string `json:"url"`
}

// EVSE defines model for evse.
type EVSE struct {
	Uid                      string               `json:"uid" validate:"required"`
	EvseId                   *string              `json:"evse_id,omitempty"`
	Status                   EvseStatus           `json:"status" validate:"required"`
	StatusSchedule           []StatusSchedule     `json:"status_schedule,omitempty"`
	Capabilities             []Capability         `json:"capabilities,omitempty"`
	Connectors               []Connector          `json:"connectors" validate:"required"`
	FloorLevel               *string              `json:"floor_level,omitempty"`
	Coordinates              *GeoLocation         `json:"coordinates,omitempty"`
	PhysicalReference        *string              `json:"physical_reference,omitempty"`
	Directions               []DisplayText        `json:"directions,omitempty"`
	ParkingRestrictions      []ParkingRestriction `json:"parking_restrictions,omitempty"`
	Parking                  []EVSEParking        `json:"parking,omitempty"`
	Images                   []Image              `json:"images,omitempty"`
	AcceptedServiceProviders *string              `json:"accepted_service_providers,omitempty"`
	LastUpdated              ocpi.DateTime        `json:"last_updated"`
}

type EVSEParking struct {
	ParkingID    string        `json:"parking_id" validate:"required"`
	EVSEPosition *EVSEPosition `json:"evse_position"`
}

type EVSEPosition string

const (
	EVSEPositionLeft   EVSEPosition = "LEFT"
	EVSEPositionRight  EVSEPosition = "RIGHT"
	EVSEPositionCenter EVSEPosition = "CENTER"
)

// Capability defines model for Evse.Capabilities.
type Capability string

// ParkingRestriction defines model for Evse.ParkingRestriction.
type ParkingRestriction string

// EvseStatus defines model for Evse.Status.
type EvseStatus string

// StatusSchedule defines model for evse_status_schedule.
type StatusSchedule struct {
	PeriodBegin string               `json:"period_begin"`
	PeriodEnd   *string              `json:"period_end,omitempty"`
	Status      StatusScheduleStatus `json:"status"`
}

// StatusScheduleStatus defines model for StatusSchedule.Status.
type StatusScheduleStatus string

// LocationReferences defines model for locationReferences.
type LocationReferences struct {
	EvseUids   *string `json:"evse_uids,omitempty"`
	LocationId string  `json:"location_id"`
}

// Locations defines model for locations.
type LocationsResponse = ocpi.Response[[]Location]

// Locations defines model for locations.
type LocationResponse = ocpi.Response[Location]

// LocationsData defines model for locations_data.
type Location struct {
	CountryCode        string                  `json:"country_code" validate:"required,len=2"`
	PartyId            string                  `json:"party_id" validate:"required,len=3"`
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
	Evses              []EVSE                  `json:"evses,omitempty"`
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
	LastUpdated        ocpi.DateTime           `json:"last_updated" validate:"required"`
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

type PatchedLocation struct {
	CountryCode        *string                 `json:"country_code,omitempty"`
	PartyId            *string                 `json:"party_id,omitempty"`
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
	Evses              []EVSE                  `json:"evses,omitempty"`
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
	HelpPhone          *string                 `json:"help_phone,omitempty"`
	LastUpdated        ocpi.DateTime           `json:"last_updated"`
}

// LocationsDataFacilities defines model for LocationsData.Facilities.
type Facility string

// ParkingType defines model for LocationsData.ParkingType.
type ParkingType string

// Hours defines model for locations_data_opening_times.
type Hours struct {
	ExceptionalClosings []ExceptionalPeriod `json:"exceptional_closings,omitempty"`
	ExceptionalOpenings []ExceptionalPeriod `json:"exceptional_openings,omitempty"`
	RegularHours        []RegularHours      `json:"regular_hours,omitempty"`
	Twentyfourseven     bool                `json:"twentyfourseven"`
}

// HoursExceptionalOpenings defines model for locations_data_opening_times_exceptional_openings.
type ExceptionalPeriod struct {
	PeriodBegin ocpi.DateTime `json:"period_begin"`
	PeriodEnd   ocpi.DateTime `json:"period_end"`
}

// HoursRegularHours defines model for locations_data_opening_times_regular_hours.
type RegularHours struct {
	PeriodBegin string `json:"period_begin"`
	PeriodEnd   string `json:"period_end"`
	Weekday     int    `json:"weekday"`
}

// PublishTokenType defines model for locations_data_publish_allowed_to.
type PublishTokenType struct {
	GroupID      *string               `json:"group_id,omitempty"`
	Issuer       *string               `json:"issuer,omitempty"`
	Type         *PublishTokenTypeType `json:"type,omitempty"`
	UID          *string               `json:"uid,omitempty"`
	VisualNumber *string               `json:"visual_number,omitempty"`
}

// PublishTokenTypeType defines model for PublishTokenType.Type.
type PublishTokenTypeType string

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

// ReservationRestrictionType defines model for reservationRestrictionType.
type ReservationRestrictionType string

// ReserveNow defines model for reserveNow.
type ReserveNowRequest struct {
	AuthorizationReference *string `json:"authorization_reference,omitempty"`
	EvseUid                *string `json:"evse_uid,omitempty"`
	ExpiryDate             string  `json:"expiry_date" validate:"required"`
	LocationID             string  `json:"location_id" validate:"required"`
	ReservationID          string  `json:"reservation_id" validate:"required"`
	ResponseURL            string  `json:"response_url" validate:"required"`
	Token                  Token   `json:"token"`
}

// Session defines model for session.
type Session struct {
	CountryCode            string           `json:"country_code" validate:"required,len=2"`
	PartyID                string           `json:"party_id" validate:"required,len=3"`
	ID                     string           `json:"id" validate:"required"`
	StartDateTime          ocpi.DateTime    `json:"start_date_time" validate:"required"`
	EndDateTime            *ocpi.DateTime   `json:"end_date_time,omitempty"`
	Kwh                    json.Number      `json:"kwh" validate:"required"`
	CdrToken               CdrToken         `json:"cdr_token"`
	AuthMethod             AuthMethod       `json:"auth_method" validate:"required"`
	AuthorizationReference *string          `json:"authorization_reference,omitempty"`
	LocationID             string           `json:"location_id" validate:"required"`
	EvseUid                string           `json:"evse_uid" validate:"required"`
	ConnectorID            string           `json:"connector_id" validate:"required"`
	MeterID                *string          `json:"meter_id,omitempty"`
	Currency               string           `json:"currency" validate:"required,len=3"`
	ChargingPeriods        []ChargingPeriod `json:"charging_periods,omitempty"`
	TotalCost              *Price           `json:"total_cost,omitempty"`
	Status                 SessionStatus    `json:"status" validate:"required"`
	LastUpdated            ocpi.DateTime    `json:"last_updated" validate:"required"`
}

type PatchedSession struct {
	LastUpdated            ocpi.DateTime   `json:"last_updated" validate:"required"`
	AuthMethod             *AuthMethod     `json:"auth_method,omitempty"`
	AuthorizationReference *string         `json:"authorization_reference,omitempty"`
	CdrToken               *CdrToken       `json:"cdr_token,omitempty"`
	ChargingPeriods        *ChargingPeriod `json:"charging_periods,omitempty"`
	ConnectorID            *string         `json:"connector_id,omitempty"`
	CountryCode            *string         `json:"country_code,omitempty"`
	Currency               *string         `json:"currency,omitempty"`
	EndDateTime            *ocpi.DateTime  `json:"end_date_time,omitempty"`
	EvseUid                *string         `json:"evse_uid,omitempty"`
	ID                     *string         `json:"id,omitempty"`
	Kwh                    *json.Number    `json:"kwh,omitempty"`
	LocationID             *string         `json:"location_id,omitempty"`
	MeterID                *string         `json:"meter_id,omitempty"`
	PartyID                *string         `json:"party_id,omitempty"`
	StartDateTime          *ocpi.DateTime  `json:"start_date_time,omitempty"`
	Status                 *SessionStatus  `json:"status,omitempty"`
	TotalCost              *Price          `json:"total_cost,omitempty"`
}

// SessionStatus defines model for Session.Status.
type SessionStatus string

// SessionResponse defines model for sessionResponse.
type SessionsResponse = ocpi.Response[[]Session]

// SessionResponse defines model for sessionResponse.
type SessionResponse = ocpi.Response[Session]

// ChargingPeriod defines model for session_charging_periods.
type ChargingPeriod struct {
	StartDateTime ocpi.DateTime  `json:"start_date_time" validate:"required"`
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
	ResponseUrl     string          `json:"response_url"`
}

// StartSessionRequest defines model for startSession.
type StartSessionRequest struct {
	AuthorizationReference *string `json:"authorization_reference,omitempty"`
	ConnectorID            *string `json:"connector_id,omitempty"`
	EvseUid                *string `json:"evse_uid,omitempty"`
	LocationID             string  `json:"location_id" validate:"required"`
	ResponseURL            string  `json:"response_url" validate:"required"`
	Token                  Token   `json:"token"`
}

// StopSessionRequest defines model for stopSession.
type StopSessionRequest struct {
	ResponseURL string  `json:"response_url"`
	SessionID   *string `json:"session_id,omitempty"`
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
	StartDateTime *ocpi.DateTime  `json:"start_date_time,omitempty"`
	EndDateTime   *ocpi.DateTime  `json:"end_date_time,omitempty"`
	EnergyMix     *EnergyMix      `json:"energy_mix,omitempty"`
	LastUpdated   ocpi.DateTime   `json:"last_updated" validate:"required"`
}

// TariffType defines model for Tariff.Type.
type TariffType string

// TariffDeleteResponse defines model for tariffDeleteResponse.
type TariffDeleteResponse struct {
	StatusCode    float32 `json:"status_code"`
	StatusMessage *string `json:"status_message,omitempty"`
	TimeStamp     string  `json:"timeStamp"`
}

// TariffResponse defines model for tariffResponse.
type TariffResponse struct {
	StatusCode    float32 `json:"status_code"`
	StatusMessage *string `json:"status_message,omitempty"`
	Tariff        Tariff  `json:"tariff"`
	TimeStamp     *string `json:"timeStamp,omitempty"`
}

// TariffElement defines model for tariff_elements.
type TariffElement struct {
	PriceComponents []PriceComponent    `json:"price_components,omitempty"`
	Restrictions    *TariffRestrictions `json:"restrictions,omitempty"`
}

// PriceComponent defines model for TariffElement.PriceComponents.
type PriceComponent struct {
	Type     PriceComponentType `json:"type"`
	StepSize uint16             `json:"step_size"`
	Price    json.Number        `json:"price"`
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

// DayOfWeek defines model for TariffRestrictions.DayOfWeek.
type DayOfWeek string

// TariffsResponse defines model for tariffsResponse.
type TariffsResponse = ocpi.Response[[]Tariff]

type InvoiceCreator string

const (
	InvoiceCreatorCPO InvoiceCreator = "CPO"
	InvoiceCreatorPTP InvoiceCreator = "PTP"
)

type Terminal struct {
	TerminalId        string          `json:"terminal_id"`
	CustomerReference *string         `json:"customer_reference,omitempty"`
	PartyId           *string         `json:"party_id,omitempty"`
	CountryCode       *string         `json:"country_code,omitempty"`
	Address           *string         `json:"address,omitempty"`
	City              *string         `json:"city,omitempty"`
	PostalCode        *string         `json:"postal_code,omitempty"`
	State             *string         `json:"state,omitempty"`
	Country           *string         `json:"country,omitempty"`
	Coordinates       *GeoLocation    `json:"coordinates,omitempty"`
	InvoiceBaseUrl    *url.URL        `json:"invoice_base_url,omitempty"`
	InvoiceCreator    *InvoiceCreator `json:"invoice_creator,omitempty"`
	Reference         *string         `json:"reference,omitempty"`
	LocationIds       []string        `json:"location_ids,omitempty"`
	EvseUids          []string        `json:"evse_uids,omitempty"`
	LastUpdated       ocpi.DateTime   `json:"last_updated"`
}

type CaptureStatusCode string

const (
	CaptureStatusCodeSuccess        CaptureStatusCode = "SUCCESS"
	CaptureStatusCodePartialSuccess CaptureStatusCode = "PARTIAL_SUCCESS"
	CaptureStatusCodeFailed         CaptureStatusCode = "FAILED"
)

type FinancialAdviceConfirmation struct {
	ID                     string            `json:"id"`
	AuthorizationReference string            `json:"authorization_reference"`
	TotalCosts             Price             `json:"total_costs"`
	Currency               string            `json:"currency"`
	EftData                string            `json:"eft_data"`
	CaptureStatusCode      CaptureStatusCode `json:"capture_status_code"`
	CaptureStatusMessage   *string           `json:"capture_status_message,omitempty"`
	LastUpdated            ocpi.DateTime     `json:"last_updated"`
}

// Token defines model for token.
type Token struct {
	ContractId         string               `json:"contract_id"`
	CountryCode        string               `json:"country_code"`
	DefaultProfileType *ProfileType         `json:"default_profile_type,omitempty"`
	EnergyContract     *TokenEnergyContract `json:"energy_contract,omitempty"`
	GroupId            *string              `json:"group_id,omitempty"`
	Issuer             string               `json:"issuer"`
	Language           *string              `json:"language,omitempty"`
	LastUpdated        ocpi.DateTime        `json:"last_updated"`
	PartyId            string               `json:"party_id"`
	Type               TokenType            `json:"type"`
	Uid                string               `json:"uid"`
	Valid              bool                 `json:"valid"`
	VisualNumber       *string              `json:"visual_number,omitempty"`
	Whitelist          WhitelistType        `json:"whitelist"`
}

type PatchedToken struct {
	ContractId         string               `json:"contract_id"`
	CountryCode        string               `json:"country_code"`
	DefaultProfileType *ProfileType         `json:"default_profile_type,omitempty"`
	EnergyContract     *TokenEnergyContract `json:"energy_contract,omitempty"`
	GroupId            *string              `json:"group_id,omitempty"`
	Issuer             string               `json:"issuer"`
	Language           *string              `json:"language,omitempty"`
	LastUpdated        ocpi.DateTime        `json:"last_updated"`
	PartyId            string               `json:"party_id"`
	Type               TokenType            `json:"type"`
	Uid                string               `json:"uid"`
	Valid              bool                 `json:"valid"`
	VisualNumber       *string              `json:"visual_number,omitempty"`
	Whitelist          WhitelistType        `json:"whitelist"`
}

// TokenResponse defines model for tokenResponse.
type TokenResponse struct {
	StatusCode    float32 `json:"status_code"`
	StatusMessage *string `json:"status_message,omitempty"`
	TimeStamp     *string `json:"timeStamp,omitempty"`
	Token         Token   `json:"token"`
}

// TokenEnergyContract defines model for token_energy_contract.
type TokenEnergyContract struct {
	ContractId   *string `json:"contract_id,omitempty"`
	SupplierName string  `json:"supplier_name"`
}

// TokensResponse defines model for tokensResponse.
type TokensResponse struct {
	StatusCode    float32 `json:"status_code"`
	StatusMessage *string `json:"status_message,omitempty"`
	TimeStamp     *string `json:"timeStamp,omitempty"`
	Tokens        *Token  `json:"tokens,omitempty"`
}

// UnlockConnector defines model for unlockConnector.
type UnlockConnector struct {
	ConnectorID string `json:"connector_id"`
	EvseUID     string `json:"evse_uid"`
	LocationID  string `json:"location_id"`
	ResponseURL string `json:"response_url"`
}

// VersionsData defines model for versions_data.
type Version struct {
	Url     string       `json:"url"`
	Version ocpi.Version `json:"version"`
}

// GetOcpiCdrsParams defines parameters for GetOcpiCdrs.
type GetOcpiCdrsParams struct {
	// DateFrom Return CDRs that have last_updated after or equal to this Date/Time (inclusive).
	DateFrom *string `form:"date_from,omitempty" json:"date_from,omitempty"`

	// DateTo Return CDRs that have last_updated up to this Date/Time, but not including (exclusive).
	DateTo *string `form:"date_to,omitempty" json:"date_to,omitempty"`

	// Offset The offset of the first object returned. Default is 0.
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit Maximum number of objects to GET.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// GetOcpiHubclientinfoParams defines parameters for GetOcpiHubclientinfo.
type GetOcpiHubclientinfoParams struct {
	// DateFrom Return ClientInfo that have last_updated after or equal to Date/Time (inclusive).
	DateFrom *string `form:"date_from,omitempty" json:"date_from,omitempty"`

	// DateTo Return ClientInfo that have last_updated up to Date/Time, but not including (exclusive).
	DateTo *string `form:"date_to,omitempty" json:"date_to,omitempty"`

	// Offset The offset of the first object returned. Default is 0.
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit Maximum number of objects to GET.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
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

// GetOcpiSessionsParams defines parameters for GetOcpiSessions.
type GetOcpiSessionsParams struct {
	// DateFrom Return Sessions that have last_updated after or equal to this date time (inclusive).
	DateFrom *string `form:"date_from,omitempty" json:"date_from,omitempty"`

	// DateTo Return Sessions that have last_updated up to this date time, but not including (exclusive).
	DateTo *string `form:"date_to,omitempty" json:"date_to,omitempty"`

	// Offset The offset of the first object returned. Default is 0.
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit Maximum number of objects to GET.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// GetOcpiTariffsParams defines parameters for GetOcpiTariffs.
type GetOcpiTariffsParams struct {
	// DateFrom Return Tariffs that have last_updated after or equal to Date/Time (inclusive).
	DateFrom *string `form:"date_from,omitempty" json:"date_from,omitempty"`

	// DateTo Return Tariffs that have last_updated up to Date/Time, but not including (exclusive).
	DateTo *string `form:"date_to,omitempty" json:"date_to,omitempty"`

	// Offset The offset of the first object returned. Default is 0.
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit Maximum number of objects to GET.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// GetOcpiTokensParams defines parameters for GetOcpiTokens.
type GetOcpiTokensParams struct {
	// DateFrom Return tokens that have last_updated after or equal to this Date/Time (inclusive).
	DateFrom *string `form:"date_from,omitempty" json:"date_from,omitempty"`

	// DateTo Return tokens that have last_updated up to Date/Time, but not including (exclusive).
	DateTo *string `form:"date_to,omitempty" json:"date_to,omitempty"`

	// Offset The offset of the first object returned. Default is 0.
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit Maximum number of objects to GET.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}
