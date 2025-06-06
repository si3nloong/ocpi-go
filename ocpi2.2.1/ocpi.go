package ocpi221

import (
	"encoding/json"
	"net/url"

	"github.com/si3nloong/ocpi-go/ocpi"
)

// Defines values for ActiveChargingProfileResultResult.
const (
	ActiveChargingProfileResultResultAccepted ActiveChargingProfileResultResult = "ACCEPTED"
	ActiveChargingProfileResultResultRejected ActiveChargingProfileResultResult = "REJECTED"
	ActiveChargingProfileResultResultUknown   ActiveChargingProfileResultResult = "UNKNOWN"
)

// Defines values for AuthorizationAllowed.
const (
	AuthorizationAllowedALLOWED    AuthorizationAllowed = "ALLOWED"
	AuthorizationAllowedBLOCKED    AuthorizationAllowed = "BLOCKED"
	AuthorizationAllowedEXPIRED    AuthorizationAllowed = "EXPIRED"
	AuthorizationAllowedNOCREDIT   AuthorizationAllowed = "NO_CREDIT"
	AuthorizationAllowedNOTALLOWED AuthorizationAllowed = "NOT_ALLOWED"
)

// Defines values for CdrBodyAuthMethod.
const (
	CdrBodyAuthMethodAUTHREQUEST CdrBodyAuthMethod = "AUTH_REQUEST"
	CdrBodyAuthMethodCOMMAND     CdrBodyAuthMethod = "COMMAND"
	CdrBodyAuthMethodWHITELIST   CdrBodyAuthMethod = "WHITELIST"
)

// Defines values for CdrBodyCdrLocationConnectorFormat.
const (
	CdrBodyCdrLocationConnectorFormatCABLE  CdrBodyCdrLocationConnectorFormat = "CABLE"
	CdrBodyCdrLocationConnectorFormatSOCKET CdrBodyCdrLocationConnectorFormat = "SOCKET"
)

// Defines values for CdrBodyChargingPeriodsDimensionsType.
const (
	CdrBodyChargingPeriodsDimensionsTypeCURRENT         CdrBodyChargingPeriodsDimensionsType = "CURRENT"
	CdrBodyChargingPeriodsDimensionsTypeENERGY          CdrBodyChargingPeriodsDimensionsType = "ENERGY"
	CdrBodyChargingPeriodsDimensionsTypeENERGYEXPORT    CdrBodyChargingPeriodsDimensionsType = "ENERGY_EXPORT"
	CdrBodyChargingPeriodsDimensionsTypeENERGYIMPORT    CdrBodyChargingPeriodsDimensionsType = "ENERGY_IMPORT"
	CdrBodyChargingPeriodsDimensionsTypeMAXCURRENT      CdrBodyChargingPeriodsDimensionsType = "MAX_CURRENT"
	CdrBodyChargingPeriodsDimensionsTypeMAXPOWER        CdrBodyChargingPeriodsDimensionsType = "MAX_POWER"
	CdrBodyChargingPeriodsDimensionsTypeMINCURRENT      CdrBodyChargingPeriodsDimensionsType = "MIN_CURRENT"
	CdrBodyChargingPeriodsDimensionsTypeMINPOWER        CdrBodyChargingPeriodsDimensionsType = "MIN_POWER"
	CdrBodyChargingPeriodsDimensionsTypePARKINGTIME     CdrBodyChargingPeriodsDimensionsType = "PARKING_TIME"
	CdrBodyChargingPeriodsDimensionsTypePOWER           CdrBodyChargingPeriodsDimensionsType = "POWER"
	CdrBodyChargingPeriodsDimensionsTypeRESERVATIONTIME CdrBodyChargingPeriodsDimensionsType = "RESERVATION_TIME"
	CdrBodyChargingPeriodsDimensionsTypeSTATEOFCHARGE   CdrBodyChargingPeriodsDimensionsType = "STATE_OF_CHARGE"
	CdrBodyChargingPeriodsDimensionsTypeTIME            CdrBodyChargingPeriodsDimensionsType = "TIME"
)

// Defines values for CdrBodyTariffsElementsPriceComponents.
const (
	CdrBodyTariffsElementsPriceComponentsENERGY      CdrBodyTariffsElementsPriceComponents = "ENERGY"
	CdrBodyTariffsElementsPriceComponentsFLAT        CdrBodyTariffsElementsPriceComponents = "FLAT"
	CdrBodyTariffsElementsPriceComponentsPARKINGTIME CdrBodyTariffsElementsPriceComponents = "PARKING_TIME"
	CdrBodyTariffsElementsPriceComponentsTIME        CdrBodyTariffsElementsPriceComponents = "TIME"
)

// Defines values for CdrBodyTariffsElementsRestrictionsDayOfWeek.
const (
	CdrBodyTariffsElementsRestrictionsDayOfWeekFRIDAY    CdrBodyTariffsElementsRestrictionsDayOfWeek = "FRIDAY"
	CdrBodyTariffsElementsRestrictionsDayOfWeekMONDAY    CdrBodyTariffsElementsRestrictionsDayOfWeek = "MONDAY"
	CdrBodyTariffsElementsRestrictionsDayOfWeekSATURDAY  CdrBodyTariffsElementsRestrictionsDayOfWeek = "SATURDAY"
	CdrBodyTariffsElementsRestrictionsDayOfWeekSUNDAY    CdrBodyTariffsElementsRestrictionsDayOfWeek = "SUNDAY"
	CdrBodyTariffsElementsRestrictionsDayOfWeekTHURSDAY  CdrBodyTariffsElementsRestrictionsDayOfWeek = "THURSDAY"
	CdrBodyTariffsElementsRestrictionsDayOfWeekTUESDAY   CdrBodyTariffsElementsRestrictionsDayOfWeek = "TUESDAY"
	CdrBodyTariffsElementsRestrictionsDayOfWeekWEDNESDAY CdrBodyTariffsElementsRestrictionsDayOfWeek = "WEDNESDAY"
)

// Defines values for CdrBodyTariffsElementsRestrictionsReservation.
const (
	CdrBodyTariffsElementsRestrictionsReservationRESERVATION        CdrBodyTariffsElementsRestrictionsReservation = "RESERVATION"
	CdrBodyTariffsElementsRestrictionsReservationRESERVATIONEXPIRES CdrBodyTariffsElementsRestrictionsReservation = "RESERVATION_EXPIRES"
)

// Defines values for EnergySourceCategory.
const (
	COAL          EnergySourceCategory = "COAL"
	GAS           EnergySourceCategory = "GAS"
	GENERALFOSSIL EnergySourceCategory = "GENERAL_FOSSIL"
	GENERALGREEN  EnergySourceCategory = "GENERAL_GREEN"
	NUCLEAR       EnergySourceCategory = "NUCLEAR"
	SOLAR         EnergySourceCategory = "SOLAR"
	WATER         EnergySourceCategory = "WATER"
	WIND          EnergySourceCategory = "WIND"
)

// Defines values for EnvironmentalImpactCategory.
const (
	EnvironmentalImpactCategoryCARBONDIOXIDE EnvironmentalImpactCategory = "CARBON_DIOXIDE"
	EnvironmentalImpactCategoryNUCLEARWASTE  EnvironmentalImpactCategory = "NUCLEAR_WASTE"
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
	ChargingPreferencesResponseChargingPreferencesNOTPOSSIBLE             ChargingPreferencesResponseChargingPreferences = "NOT_POSSIBLE"
	ChargingPreferencesResponseChargingPreferencesPROFILETYPENOTSUPPORTED ChargingPreferencesResponseChargingPreferences = "PROFILE_TYPE_NOT_SUPPORTED"
)

// Defines values for ChargingProfileChargingRateUnit.
const (
	A ChargingProfileChargingRateUnit = "A"
	W ChargingProfileChargingRateUnit = "W"
)

// Defines values for ChargingProfileResponseType.
const (
	ChargingProfileResponseTypeAccepted       ChargingProfileResponseType = "ACCEPTED"
	ChargingProfileResponseTypeNOTSUPPORTED   ChargingProfileResponseType = "NOT_SUPPORTED"
	ChargingProfileResponseTypeREJECTED       ChargingProfileResponseType = "REJECTED"
	ChargingProfileResponseTypeTOOOFTEN       ChargingProfileResponseType = "TOO_OFTEN"
	ChargingProfileResponseTypeUNKNOWNSESSION ChargingProfileResponseType = "UNKNOWN_SESSION"
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

// Defines values for ClientInfoStatus.
const (
	ClientInfoStatusCONNECTED ClientInfoStatus = "CONNECTED"
	ClientInfoStatusOFFLINE   ClientInfoStatus = "OFFLINE"
	ClientInfoStatusPLANNED   ClientInfoStatus = "PLANNED"
	ClientInfoStatusSUSPENDED ClientInfoStatus = "SUSPENDED"
)

// Defines values for CommandResponseResult.
const (
	CommandResponseResultAccepted       CommandResponseResult = "ACCEPTED"
	CommandResponseResultNotSupported   CommandResponseResult = "NOT_SUPPORTED"
	CommandResponseResultRejected       CommandResponseResult = "REJECTED"
	CommandResponseResultUnknownSession CommandResponseResult = "UNKNOWN_SESSION"
)

// Defines values for CommandResultResult.
const (
	ACCEPTED            CommandResultResult = "ACCEPTED"
	CANCELEDRESERVATION CommandResultResult = "CANCELED_RESERVATION"
	EVSEINOPERATIVE     CommandResultResult = "EVSE_INOPERATIVE"
	EVSEOCCUPIED        CommandResultResult = "EVSE_OCCUPIED"
	FAILED              CommandResultResult = "FAILED"
	NOTSUPPORTED        CommandResultResult = "NOT_SUPPORTED"
	REJECTED            CommandResultResult = "REJECTED"
	TIMEOUT             CommandResultResult = "TIMEOUT"
	UNKNOWNRESERVATION  CommandResultResult = "UNKNOWN_RESERVATION"
)

// Defines values for ConnectorFormat.
const (
	ConnectorFormatCABLE  ConnectorFormat = "CABLE"
	ConnectorFormatSOCKET ConnectorFormat = "SOCKET"
)

// Defines values for PowerType.
const (
	PowerTypeAC1PHASE      PowerType = "AC_1_PHASE"
	PowerTypeAC2PHASE      PowerType = "AC_2_PHASE"
	PowerTypeAC2PHASESPLIT PowerType = "AC_2_PHASE_SPLIT"
	PowerTypeAC3PHASE      PowerType = "AC_3_PHASE"
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

// Defines values for CredentialsDataRolesBusinessDetailsLogoCategory.
const (
	CredentialsDataRolesBusinessDetailsLogoCategoryCHARGER  CredentialsDataRolesBusinessDetailsLogoCategory = "CHARGER"
	CredentialsDataRolesBusinessDetailsLogoCategoryENTRANCE CredentialsDataRolesBusinessDetailsLogoCategory = "ENTRANCE"
	CredentialsDataRolesBusinessDetailsLogoCategoryLOCATION CredentialsDataRolesBusinessDetailsLogoCategory = "LOCATION"
	CredentialsDataRolesBusinessDetailsLogoCategoryNETWORK  CredentialsDataRolesBusinessDetailsLogoCategory = "NETWORK"
	CredentialsDataRolesBusinessDetailsLogoCategoryOPERATOR CredentialsDataRolesBusinessDetailsLogoCategory = "OPERATOR"
	CredentialsDataRolesBusinessDetailsLogoCategoryOTHER    CredentialsDataRolesBusinessDetailsLogoCategory = "OTHER"
	CredentialsDataRolesBusinessDetailsLogoCategoryOWNER    CredentialsDataRolesBusinessDetailsLogoCategory = "OWNER"
)

// Defines values for EnvironmentalImpactCategoryType.
const (
	EnvironmentalImpactCategoryTypeCARBONDIOXIDE EnvironmentalImpactCategoryType = "CARBON_DIOXIDE"
	EnvironmentalImpactCategoryTypeNUCLEARWASTE  EnvironmentalImpactCategoryType = "NUCLEAR_WASTE"
)

// Defines values for EvseCapabilities.
const (
	CapabilityCHARGINGPREFERENCESCAPABLE    Capability = "CHARGING_PREFERENCES_CAPABLE"
	CapabilityCHARGINGPROFILECAPABLE        Capability = "CHARGING_PROFILE_CAPABLE"
	CapabilityCHIPCARDSUPPORT               Capability = "CHIP_CARD_SUPPORT"
	CapabilityCONTACTLESSCARDSUPPORT        Capability = "CONTACTLESS_CARD_SUPPORT"
	CapabilityCREDITCARDPAYABLE             Capability = "CREDIT_CARD_PAYABLE"
	CapabilityDEBITCARDPAYABLE              Capability = "DEBIT_CARD_PAYABLE"
	CapabilityPEDTERMINAL                   Capability = "PED_TERMINAL"
	CapabilityREMOTESTARTSTOPCAPABLE        Capability = "REMOTE_START_STOP_CAPABLE"
	CapabilityRESERVABLE                    Capability = "RESERVABLE"
	CapabilityRFIDREADER                    Capability = "RFID_READER"
	CapabilitySTARTSESSIONCONNECTORREQUIRED Capability = "START_SESSION_CONNECTOR_REQUIRED"
	CapabilityTOKENGROUPCAPABLE             Capability = "TOKEN_GROUP_CAPABLE"
	CapabilityUNLOCKCAPABLE                 Capability = "UNLOCK_CAPABLE"
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

// Defines values for InterfaceRoleType.
const (
	RoleReceiver InterfaceRoleType = "RECEIVER"
	RoleSender   InterfaceRoleType = "SENDER"
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

// Defines values for ModuleIDType.
const (
	ModuleIDTypeCdrs             ModuleIDType = "cdrs"
	ModuleIDTypeChargingProfiles ModuleIDType = "chargingprofiles"
	ModuleIDTypeCommands         ModuleIDType = "commands"
	ModuleIDTypeCredentials      ModuleIDType = "credentials"
	ModuleIDTypeHubClientInfo    ModuleIDType = "hubclientinfo"
	ModuleIDTypeLocations        ModuleIDType = "locations"
	ModuleIDTypeSessions         ModuleIDType = "sessions"
	ModuleIDTypeTariffs          ModuleIDType = "tariffs"
	ModuleIDTypeTokens           ModuleIDType = "tokens"
)

// Defines values for ReservationRestrictionType.
const (
	ReservationRestrictionTypeRESERVATION        ReservationRestrictionType = "RESERVATION"
	ReservationRestrictionTypeRESERVATIONEXPIRES ReservationRestrictionType = "RESERVATION_EXPIRES"
)

// Defines values for SessionAuthMethod.
const (
	SessionAuthMethodAuthRequest SessionAuthMethod = "AUTH_REQUEST"
	SessionAuthMethodCommand     SessionAuthMethod = "COMMAND"
	SessionAuthMethodWhitelist   SessionAuthMethod = "WHITELIST"
)

// Defines values for SessionStatus.
const (
	ACTIVE      SessionStatus = "ACTIVE"
	COMPLETED   SessionStatus = "COMPLETED"
	INVALID     SessionStatus = "INVALID"
	PENDING     SessionStatus = "PENDING"
	RESERVATION SessionStatus = "RESERVATION"
)

// Defines values for SessionChargingPeriodsDimensionsType.
const (
	AUTHREQUEST SessionChargingPeriodsDimensionsType = "AUTH_REQUEST"
	COMMAND     SessionChargingPeriodsDimensionsType = "COMMAND"
	WHITELIST   SessionChargingPeriodsDimensionsType = "WHITELIST"
)

// Defines values for TariffType.
const (
	TariffTypeAdHocPayment TariffType = "AD_HOC_PAYMENT"
	TariffTypeProfileCheap TariffType = "PROFILE_CHEAP"
	TariffTypeProfileFast  TariffType = "PROFILE_FAST"
	TariffTypeProfileGreen TariffType = "PROFILE_GREEN"
	TariffTypeRegular      TariffType = "REGULAR"
)

type TariffElementsPriceComponentsType string

// Defines values for TariffElementsPriceComponents.
const (
	TariffElementsPriceComponentsTypeEnergy      TariffElementsPriceComponentsType = "ENERGY"
	TariffElementsPriceComponentsTypeFlat        TariffElementsPriceComponentsType = "FLAT"
	TariffElementsPriceComponentsTypeParkingTime TariffElementsPriceComponentsType = "PARKING_TIME"
	TariffElementsPriceComponentsTypeTime        TariffElementsPriceComponentsType = "TIME"
)

// Defines values for TariffElementsRestrictionsDayOfWeek.
const (
	TariffElementsRestrictionsDayOfWeekMonday    TariffElementsRestrictionsDayOfWeek = "MONDAY"
	TariffElementsRestrictionsDayOfWeekTUESDAY   TariffElementsRestrictionsDayOfWeek = "TUESDAY"
	TariffElementsRestrictionsDayOfWeekWEDNESDAY TariffElementsRestrictionsDayOfWeek = "WEDNESDAY"
	TariffElementsRestrictionsDayOfWeekTHURSDAY  TariffElementsRestrictionsDayOfWeek = "THURSDAY"
	TariffElementsRestrictionsDayOfWeekFRIDAY    TariffElementsRestrictionsDayOfWeek = "FRIDAY"
	TariffElementsRestrictionsDayOfWeekSaturday  TariffElementsRestrictionsDayOfWeek = "SATURDAY"
	TariffElementsRestrictionsDayOfWeekSunday    TariffElementsRestrictionsDayOfWeek = "SUNDAY"
)

// Defines values for TokenDefaultProfileType.
const (
	TokenDefaultProfileTypeCheap   TokenDefaultProfileType = "CHEAP"
	TokenDefaultProfileTypeFast    TokenDefaultProfileType = "FAST"
	TokenDefaultProfileTypeGreen   TokenDefaultProfileType = "GREEN"
	TokenDefaultProfileTypeRegular TokenDefaultProfileType = "REGULAR"
)

// Defines values for TokenType.
const (
	TokenTypeAdHocUser TokenType = "AD_HOC_USER"
	TokenTypeAppUser   TokenType = "APP_USER"
	TokenTypeOther     TokenType = "OTHER"
	TokenTypeRFID      TokenType = "RFID"
)

// Defines values for WhitelistType.
const (
	WhitelistTypeAllowed        WhitelistType = "ALLOWED"
	WhitelistTypeAllowedOffline WhitelistType = "ALLOWED_OFFLINE"
	WhitelistTypeAlways         WhitelistType = "ALWAYS"
	WhitelistTypeNever          WhitelistType = "NEVER"
)

// Defines values for PostOcpiCommandsCommandParamsCommand.
const (
	CommandTypeCancelReservation CommandType = "CANCEL_RESERVATION"
	CommandTypeReserveNow        CommandType = "RESERVE_NOW"
	CommandTypeStartSession      CommandType = "START_SESSION"
	CommandTypeStopSession       CommandType = "STOP_SESSION"
	CommandTypeUnlockConnector   CommandType = "UNLOCK_CONNECTOR"
)

type VersionsResponse = ocpi.Response[[]Version]

// ActiveChargingProfile defines model for activeChargingProfile.
type ActiveChargingProfile struct {
	ChargingProfile ChargingProfile `json:"charging_profile"`
	StartDateTime   string          `json:"start_date_time"`
}

// ActiveChargingProfileResult defines model for activeChargingProfileResult.
type ActiveChargingProfileResult struct {
	Profile *ActiveChargingProfileResultProfile `json:"profile,omitempty"`
	Result  ActiveChargingProfileResultResult   `json:"result"`
}

// ActiveChargingProfileResultResult defines model for ActiveChargingProfileResult.Result.
type ActiveChargingProfileResultResult string

// ActiveChargingProfileResultProfile defines model for activeChargingProfileResult_profile.
type ActiveChargingProfileResultProfile struct {
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
	Logo    *CredentialsDataRolesBusinessDetailsLogo `json:"logo,omitempty"`
	Name    string                                   `json:"name"`
	Website *string                                  `json:"website,omitempty"`
}

// CancelReservation defines model for cancelReservation.
type CancelReservation struct {
	ReservationId string `json:"reservation_id"`
	ResponseUrl   string `json:"response_url"`
}

// CdrBody defines model for cdrBody.
type ChargeDetailRecord struct {
	AuthMethod               CdrBodyAuthMethod       `json:"auth_method"`
	AuthorizationReference   *string                 `json:"authorization_reference,omitempty"`
	CdrLocation              CdrBodyCdrLocation      `json:"cdr_location"`
	CdrToken                 CdrBodyCdrToken         `json:"cdr_token"`
	ChargingPeriods          *CdrBodyChargingPeriods `json:"charging_periods,omitempty"`
	CountryCode              string                  `json:"country_code"`
	Credit                   *bool                   `json:"credit,omitempty"`
	CreditReferenceId        *string                 `json:"credit_reference_id,omitempty"`
	Currency                 string                  `json:"currency"`
	EndDateTime              string                  `json:"end_date_time"`
	HomeChargingCompensation *bool                   `json:"home_charging_compensation,omitempty"`
	Id                       string                  `json:"id"`
	InvoiceReferenceId       *string                 `json:"invoice_reference_id,omitempty"`
	LastUpdated              string                  `json:"last_updated"`
	MeterId                  *string                 `json:"meter_id,omitempty"`
	PartyId                  string                  `json:"party_id"`
	Remark                   *string                 `json:"remark,omitempty"`
	SessionId                *string                 `json:"session_id,omitempty"`
	SignedData               *CdrBodySignedData      `json:"signed_data,omitempty"`
	StartDateTime            string                  `json:"start_date_time"`
	Tariffs                  *CdrBodyTariffs         `json:"tariffs,omitempty"`
	TotalCost                Price                   `json:"total_cost"`
	TotalEnergy              float32                 `json:"total_energy"`
	TotalEnergyCost          *Price                  `json:"total_energy_cost,omitempty"`
	TotalFixedCost           *Price                  `json:"total_fixed_cost,omitempty"`
	TotalParkingCost         *Price                  `json:"total_parking_cost,omitempty"`
	TotalParkingTime         *float32                `json:"total_parking_time,omitempty"`
	TotalReservationCost     *Price                  `json:"total_reservation_cost,omitempty"`
	TotalTime                float32                 `json:"total_time"`
	TotalTimeCost            *Price                  `json:"total_time_cost,omitempty"`
}

// CdrBodyAuthMethod defines model for CdrBody.AuthMethod.
type CdrBodyAuthMethod string

// CdrBodyCdrLocation defines model for cdrBody_cdr_location.
type CdrBodyCdrLocation struct {
	Address           *string                            `json:"address,omitempty"`
	City              *string                            `json:"city,omitempty"`
	ConnectorFormat   *CdrBodyCdrLocationConnectorFormat `json:"connector_format,omitempty"`
	ConnectorId       *string                            `json:"connector_id,omitempty"`
	PowerType         *CdrBodyCdrLocationPowerType       `json:"connector_power_type,omitempty"`
	ConnectorStandard *ConnectorStandard                 `json:"connector_standard,omitempty"`
	Coordinates       *GeoLocation                       `json:"coordinates,omitempty"`
	Country           *string                            `json:"country,omitempty"`
	EvseId            *string                            `json:"evse_id,omitempty"`
	EvseUid           *string                            `json:"evse_uid,omitempty"`
	Id                *string                            `json:"id,omitempty"`
	Name              *string                            `json:"name,omitempty"`
	PostalCode        *string                            `json:"postal_code,omitempty"`
	State             *string                            `json:"state,omitempty"`
}

// CdrBodyCdrLocationConnectorFormat defines model for CdrBodyCdrLocation.ConnectorFormat.
type CdrBodyCdrLocationConnectorFormat string

// CdrBodyCdrLocationPowerType defines model for CdrBodyCdrLocation.PowerType.
type CdrBodyCdrLocationPowerType string

// GeoLocation defines model for cdrBody_cdr_location_coordinates.
type GeoLocation struct {
	Latitude  json.Number `json:"latitude"`
	Longitude json.Number `json:"longitude"`
}

// CdrBodyCdrToken defines model for cdrBody_cdr_token.
type CdrBodyCdrToken struct {
	ContractId  string              `json:"contract_id"`
	CountryCode string              `json:"country_code"`
	PartyId     string              `json:"party_id"`
	Type        CdrBodyCdrTokenType `json:"type"`
	Uid         string              `json:"uid"`
}

// CdrBodyCdrTokenType defines model for CdrBodyCdrToken.Type.
type CdrBodyCdrTokenType string

// CdrBodyChargingPeriods defines model for cdrBody_charging_periods.
type CdrBodyChargingPeriods struct {
	Dimensions    *CdrBodyChargingPeriodsDimensions `json:"dimensions,omitempty"`
	StartDateTime string                            `json:"start_date_time"`
	TariffId      *string                           `json:"tariff_id,omitempty"`
}

// CdrBodyChargingPeriodsDimensions defines model for cdrBody_charging_periods_dimensions.
type CdrBodyChargingPeriodsDimensions struct {
	Type   CdrBodyChargingPeriodsDimensionsType `json:"type"`
	Volume float32                              `json:"volume"`
}

// CdrBodyChargingPeriodsDimensionsType defines model for CdrBodyChargingPeriodsDimensions.Type.
type CdrBodyChargingPeriodsDimensionsType string

// CdrBodySignedData defines model for cdrBody_signed_data.
type CdrBodySignedData struct {
	EncodingMethod        string                         `json:"encoding_method"`
	EncodingMethodVersion *int                           `json:"encoding_method_version,omitempty"`
	PublicKey             *string                        `json:"public_key,omitempty"`
	SignedValues          *CdrBodySignedDataSignedValues `json:"signed_values,omitempty"`
	Url                   *string                        `json:"url,omitempty"`
}

// CdrBodySignedDataSignedValues defines model for cdrBody_signed_data_signed_values.
type CdrBodySignedDataSignedValues struct {
	Nature     string `json:"nature"`
	PlainData  string `json:"plain_data"`
	SignedData string `json:"signed_data"`
}

// CdrBodyTariffs defines model for cdrBody_tariffs.
type CdrBodyTariffs struct {
	CountryCode   string                  `json:"country_code"`
	Currency      string                  `json:"currency"`
	Elements      *CdrBodyTariffsElements `json:"elements,omitempty"`
	EndDateTime   *string                 `json:"end_date_time,omitempty"`
	EnergyMix     *EnergyMix              `json:"energy_mix,omitempty"`
	Id            string                  `json:"id"`
	LastUpdated   string                  `json:"last_updated"`
	MaxPrice      *Price                  `json:"max_price,omitempty"`
	MinPrice      *Price                  `json:"min_price,omitempty"`
	PartyId       string                  `json:"party_id"`
	StartDateTime *string                 `json:"start_date_time,omitempty"`
	TariffAltText *DisplayText            `json:"tariff_alt_text,omitempty"`
	TariffAltUrl  *string                 `json:"tariff_alt_url,omitempty"`
	Type          *TariffType             `json:"type,omitempty"`
}

// CdrBodyTariffsElements defines model for cdrBody_tariffs_elements.
type CdrBodyTariffsElements struct {
	PriceComponents *CdrBodyTariffsElementsPriceComponents `json:"price_components,omitempty"`
	Restrictions    *CdrBodyTariffsElementsRestrictions    `json:"restrictions,omitempty"`
}

// CdrBodyTariffsElementsPriceComponents defines model for CdrBodyTariffsElements.PriceComponents.
type CdrBodyTariffsElementsPriceComponents string

// CdrBodyTariffsElementsRestrictions defines model for cdrBody_tariffs_elements_restrictions.
type CdrBodyTariffsElementsRestrictions struct {
	DayOfWeek   *CdrBodyTariffsElementsRestrictionsDayOfWeek   `json:"day_of_week,omitempty"`
	EndDate     *string                                        `json:"end_date,omitempty"`
	EndTime     *string                                        `json:"end_time,omitempty"`
	MaxCurrent  *float32                                       `json:"max_current,omitempty"`
	MaxDuration *int                                           `json:"max_duration,omitempty"`
	MaxKwh      *float32                                       `json:"max_kwh,omitempty"`
	MaxPower    *float32                                       `json:"max_power,omitempty"`
	MinCurrent  *float32                                       `json:"min_current,omitempty"`
	MinDuration *int                                           `json:"min_duration,omitempty"`
	MinKwh      *float32                                       `json:"min_kwh,omitempty"`
	MinPower    *float32                                       `json:"min_power,omitempty"`
	Reservation *CdrBodyTariffsElementsRestrictionsReservation `json:"reservation,omitempty"`
	StartDate   *string                                        `json:"start_date,omitempty"`
	StartTime   *string                                        `json:"start_time,omitempty"`
}

// CdrBodyTariffsElementsRestrictionsDayOfWeek defines model for CdrBodyTariffsElementsRestrictions.DayOfWeek.
type CdrBodyTariffsElementsRestrictionsDayOfWeek string

// CdrBodyTariffsElementsRestrictionsReservation defines model for CdrBodyTariffsElementsRestrictions.Reservation.
type CdrBodyTariffsElementsRestrictionsReservation string

// EnergyMix defines model for cdrBody_tariffs_energy_mix.
type EnergyMix struct {
	IsGreenEnergy     bool                 `json:"is_green_energy"`
	EnergySources     *EnergySource        `json:"energy_sources,omitempty"`
	EnvironImpact     *EnvironmentalImpact `json:"environ_impact,omitempty"`
	SupplierName      *string              `json:"supplier_name,omitempty"`
	EnergyProductName *string              `json:"energy_product_name,omitempty"`
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
	Url           string  `json:"url"`
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
	ChargingProfilePeriod *ChargingProfileChargingProfilePeriod `json:"charging_profile_period,omitempty"`
	ChargingRateUnit      ChargingProfileChargingRateUnit       `json:"charging_rate_unit"`
	Duration              *int                                  `json:"duration,omitempty"`
	MinChargingRate       *float32                              `json:"min_charging_rate,omitempty"`
	StartDateTime         *string                               `json:"start_date_time,omitempty"`
}

// ChargingProfileChargingRateUnit defines model for ChargingProfile.ChargingRateUnit.
type ChargingProfileChargingRateUnit string

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

// ChargingProfileChargingProfilePeriod defines model for chargingProfile_charging_profile_period.
type ChargingProfileChargingProfilePeriod struct {
	Limit       float32 `json:"limit"`
	StartPeriod int     `json:"start_period"`
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
	CountryCode string           `json:"country_code"`
	LastUpdated string           `json:"last_updated"`
	PartyId     string           `json:"party_id"`
	Role        ClientInfoRole   `json:"role"`
	Status      ClientInfoStatus `json:"status"`
}

// ClientInfoRole defines model for ClientInfo.Role.
type ClientInfoRole string

// ClientInfoStatus defines model for ClientInfo.Status.
type ClientInfoStatus string

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
	Result  CommandResponseResult   `json:"result"`
	Timeout int                     `json:"timeout"`
}

// CommandResponseResult defines model for CommandResponse.Result.
type CommandResponseResult string

// CommandResponseMessage defines model for commandResponse_message.
type CommandResponseMessage struct {
	Language string `json:"language"`
	Text     string `json:"text"`
}

// CommandResult defines model for commandResult.
type CommandResult struct {
	Message *CommandResponseMessage `json:"message,omitempty"`
	Result  CommandResultResult     `json:"result"`
}

// CommandResultResult defines model for CommandResult.Result.
type CommandResultResult string

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
	PartyId         string                              `json:"party_id"`
	Role            CredentialsDataRolesRole            `json:"role"`
}

// CredentialsDataRolesRole defines model for CredentialsDataRoles.Role.
type CredentialsDataRolesRole string

// CredentialsDataRolesBusinessDetails defines model for credentials_data_roles_business_details.
type CredentialsDataRolesBusinessDetails struct {
	Logo    *CredentialsDataRolesBusinessDetailsLogo `json:"logo,omitempty"`
	Name    string                                   `json:"name"`
	Website *string                                  `json:"website,omitempty"`
}

// CredentialsDataRolesBusinessDetailsLogo defines model for credentials_data_roles_business_details_logo.
type CredentialsDataRolesBusinessDetailsLogo struct {
	Category  CredentialsDataRolesBusinessDetailsLogoCategory `json:"category"`
	Height    *float32                                        `json:"height,omitempty"`
	Thumbnail *string                                         `json:"thumbnail,omitempty"`
	Type      string                                          `json:"type"`
	Url       string                                          `json:"url"`
	Width     *float32                                        `json:"width,omitempty"`
}

// CredentialsDataRolesBusinessDetailsLogoCategory defines model for CredentialsDataRolesBusinessDetailsLogo.Category.
type CredentialsDataRolesBusinessDetailsLogoCategory string

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
	Identifier ModuleIDType `json:"identifier"`

	// Role Interface role endpoint implements.
	Role InterfaceRoleType `json:"role"`

	// Url URL to the endpoint.
	Url string `json:"url"`
}

// EnvironmentalImpactCategoryType Categories of environmental impact values
type EnvironmentalImpactCategoryType string

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

// Image defines model for image.
type Image struct {
	Category  ImageCategory `json:"category"`
	Thumbnail *string       `json:"thumbnail,omitempty"`
	Type      string        `json:"type"`
	Url       string        `json:"url"`
	Width     *float32      `json:"width,omitempty"`
	Height    *float32      `json:"height,omitempty"`
}

// ImageCategory defines model for Image.Category.
type ImageCategory string

// InterfaceRoleType Interface role endpoint implements.
type InterfaceRoleType string

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
	CountryCode        string                  `json:"country_code"`
	PartyId            string                  `json:"party_id"`
	ID                 string                  `json:"id" validate:"required"`
	Publish            bool                    `json:"publish"`
	PublishAllowedTo   []PublishTokenType      `json:"publish_allowed_to,omitempty"`
	Name               *string                 `json:"name,omitempty"`
	Address            string                  `json:"address"`
	City               string                  `json:"city"`
	PostalCode         *string                 `json:"postal_code,omitempty"`
	State              *string                 `json:"state,omitempty"`
	Country            string                  `json:"country"`
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
	GroupId      *string               `json:"group_id,omitempty"`
	Issuer       *string               `json:"issuer,omitempty"`
	Type         *PublishTokenTypeType `json:"type,omitempty"`
	Uid          *string               `json:"uid,omitempty"`
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

// ModuleIDType OCPI 2.2.1 modules
type ModuleIDType string

// Price defines model for price.
type Price struct {
	ExclVat float32  `json:"excl_vat"`
	InclVat *float32 `json:"incl_vat,omitempty"`
}

// ReservationRestrictionType defines model for reservationRestrictionType.
type ReservationRestrictionType string

// ReserveNow defines model for reserveNow.
type ReserveNowRequest struct {
	AuthorizationReference *string `json:"authorization_reference,omitempty"`
	EvseUid                *string `json:"evse_uid,omitempty"`
	ExpiryDate             string  `json:"expiry_date"`
	LocationId             string  `json:"location_id"`
	ReservationId          string  `json:"reservation_id"`
	ResponseUrl            string  `json:"response_url"`
	Token                  Token   `json:"token"`
}

// Session defines model for session.
type Session struct {
	AuthMethod             SessionAuthMethod       `json:"auth_method"`
	AuthorizationReference *string                 `json:"authorization_reference,omitempty"`
	CdrToken               CdrBodyCdrToken         `json:"cdr_token"`
	ChargingPeriods        *SessionChargingPeriods `json:"charging_periods,omitempty"`
	ConnectorId            string                  `json:"connector_id"`
	CountryCode            string                  `json:"country_code"`
	Currency               string                  `json:"currency"`
	EndDateTime            *ocpi.DateTime          `json:"end_date_time,omitempty"`
	EvseUid                string                  `json:"evse_uid"`
	Id                     string                  `json:"id"`
	Kwh                    json.Number             `json:"kwh"`
	LastUpdated            ocpi.DateTime           `json:"last_updated"`
	LocationId             string                  `json:"location_id"`
	MeterId                *string                 `json:"meter_id,omitempty"`
	PartyId                string                  `json:"party_id"`
	StartDateTime          ocpi.DateTime           `json:"start_date_time"`
	Status                 SessionStatus           `json:"status"`
	TotalCosts             *SessionTotalCosts      `json:"total_costs,omitempty"`
}

type PatchedSession struct {
	LastUpdated            ocpi.DateTime           `json:"last_updated"`
	AuthMethod             *SessionAuthMethod      `json:"auth_method,omitempty"`
	AuthorizationReference *string                 `json:"authorization_reference,omitempty"`
	CdrToken               *CdrBodyCdrToken        `json:"cdr_token,omitempty"`
	ChargingPeriods        *SessionChargingPeriods `json:"charging_periods,omitempty"`
	ConnectorId            *string                 `json:"connector_id,omitempty"`
	CountryCode            *string                 `json:"country_code,omitempty"`
	Currency               *string                 `json:"currency,omitempty"`
	EndDateTime            *ocpi.DateTime          `json:"end_date_time,omitempty"`
	EvseUid                *string                 `json:"evse_uid,omitempty"`
	Id                     *string                 `json:"id,omitempty"`
	Kwh                    *json.Number            `json:"kwh,omitempty"`
	LocationId             *string                 `json:"location_id,omitempty"`
	MeterId                *string                 `json:"meter_id,omitempty"`
	PartyId                *string                 `json:"party_id,omitempty"`
	StartDateTime          *ocpi.DateTime          `json:"start_date_time,omitempty"`
	Status                 *SessionStatus          `json:"status,omitempty"`
	TotalCosts             *SessionTotalCosts      `json:"total_costs,omitempty"`
}

// SessionAuthMethod defines model for Session.AuthMethod.
type SessionAuthMethod string

// SessionStatus defines model for Session.Status.
type SessionStatus string

// SessionResponse defines model for sessionResponse.
type SessionsResponse = ocpi.Response[[]Session]

// SessionResponse defines model for sessionResponse.
type SessionResponse = ocpi.Response[Session]

// SessionChargingPeriods defines model for session_charging_periods.
type SessionChargingPeriods struct {
	Dimensions    *SessionChargingPeriodsDimensions `json:"dimensions,omitempty"`
	StartDateTime string                            `json:"start_date_time"`
	TariffId      *string                           `json:"tariff_id,omitempty"`
}

// SessionChargingPeriodsDimensions defines model for session_charging_periods_dimensions.
type SessionChargingPeriodsDimensions struct {
	Type   SessionChargingPeriodsDimensionsType `json:"type"`
	Volume float32                              `json:"volume"`
}

// SessionChargingPeriodsDimensionsType defines model for SessionChargingPeriodsDimensions.Type.
type SessionChargingPeriodsDimensionsType string

// SessionTotalCosts defines model for session_total_costs.
type SessionTotalCosts struct {
	ExclVat float32  `json:"excl_vat"`
	InclVat *float32 `json:"incl_vat,omitempty"`
}

// SetChargingProfile defines model for setChargingProfile.
type SetChargingProfile struct {
	ChargingProfile ChargingProfile `json:"charging_profile"`
	ResponseUrl     string          `json:"response_url"`
}

// StartSessionRequest defines model for startSession.
type StartSessionRequest struct {
	AuthorizationReference *string `json:"authorization_reference,omitempty"`
	ConnectorId            *string `json:"connector_id,omitempty"`
	EvseUid                *string `json:"evse_uid,omitempty"`
	LocationId             string  `json:"location_id"`
	ResponseUrl            string  `json:"response_url"`
	Token                  Token   `json:"token"`
}

// StopSessionRequest defines model for stopSession.
type StopSessionRequest struct {
	ResponseUrl string  `json:"response_url"`
	SessionId   *string `json:"session_id,omitempty"`
}

// Tariff defines model for tariff.
type Tariff struct {
	CountryCode   string           `json:"country_code"`
	Currency      string           `json:"currency"`
	Elements      []TariffElements `json:"elements,omitempty"`
	EndDateTime   *string          `json:"end_date_time,omitempty"`
	EnergyMix     *TariffEnergyMix `json:"energy_mix,omitempty"`
	Id            string           `json:"id"`
	LastUpdated   ocpi.DateTime    `json:"last_updated"`
	MaxPrice      *Price           `json:"max_price,omitempty"`
	MinPrice      *Price           `json:"min_price,omitempty"`
	PartyId       string           `json:"party_id"`
	StartDateTime *string          `json:"start_date_time,omitempty"`
	TariffAltText *DisplayText     `json:"tariff_alt_text,omitempty"`
	TariffAltUrl  *string          `json:"tariff_alt_url,omitempty"`
	Type          *TariffType      `json:"type,omitempty"`
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

// TariffElements defines model for tariff_elements.
type TariffElements struct {
	PriceComponents []TariffElementsPriceComponents `json:"price_components,omitempty"`
	Restrictions    *TariffElementsRestrictions     `json:"restrictions,omitempty"`
}

// TariffElementsPriceComponents defines model for TariffElements.PriceComponents.
type TariffElementsPriceComponents struct {
	Type     TariffElementsPriceComponentsType `json:"type"`
	StepSize uint16                            `json:"step_size"`
	Price    json.Number                       `json:"price"`
}

// TariffElementsRestrictions defines model for tariff_elements_restrictions.
type TariffElementsRestrictions struct {
	DayOfWeek   []TariffElementsRestrictionsDayOfWeek `json:"day_of_week,omitempty"`
	EndDate     *string                               `json:"end_date,omitempty"`
	EndTime     *string                               `json:"end_time,omitempty"`
	MaxCurrent  *float32                              `json:"max_current,omitempty"`
	MaxDuration *int                                  `json:"max_duration,omitempty"`
	MaxKwh      *float32                              `json:"max_kwh,omitempty"`
	MaxPower    *float32                              `json:"max_power,omitempty"`
	MinCurrent  *float32                              `json:"min_current,omitempty"`
	MinDuration *int                                  `json:"min_duration,omitempty"`
	MinKwh      *float32                              `json:"min_kwh,omitempty"`
	MinPower    *float32                              `json:"min_power,omitempty"`
	Reservation *ReservationRestrictionType           `json:"reservation,omitempty"`
	StartDate   *string                               `json:"start_date,omitempty"`
	StartTime   *string                               `json:"start_time,omitempty"`
}

// TariffElementsRestrictionsDayOfWeek defines model for TariffElementsRestrictions.DayOfWeek.
type TariffElementsRestrictionsDayOfWeek string

// TariffEnergyMix defines model for tariff_energy_mix.
type TariffEnergyMix struct {
	EnergyProductName *string                    `json:"energy_product_name,omitempty"`
	EnergySources     []EnergySource             `json:"energy_sources,omitempty"`
	EnvironImpact     *TariffEnvironmentalImpact `json:"environ_impact,omitempty"`
	IsGreenEnergy     bool                       `json:"is_green_energy"`
	SupplierName      *string                    `json:"supplier_name,omitempty"`
}

// TariffEnvironmentalImpact defines model for tariff_energy_mix_environ_impact.
type TariffEnvironmentalImpact struct {
	Amount float32 `json:"amount"`

	// Category Categories of environmental impact values
	Category EnvironmentalImpactCategoryType `json:"category"`
}

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
	ContractId         string                   `json:"contract_id"`
	CountryCode        string                   `json:"country_code"`
	DefaultProfileType *TokenDefaultProfileType `json:"default_profile_type,omitempty"`
	EnergyContract     *TokenEnergyContract     `json:"energy_contract,omitempty"`
	GroupId            *string                  `json:"group_id,omitempty"`
	Issuer             string                   `json:"issuer"`
	Language           *string                  `json:"language,omitempty"`
	LastUpdated        ocpi.DateTime            `json:"last_updated"`
	PartyId            string                   `json:"party_id"`
	Type               TokenType                `json:"type"`
	Uid                string                   `json:"uid"`
	Valid              bool                     `json:"valid"`
	VisualNumber       *string                  `json:"visual_number,omitempty"`
	Whitelist          WhitelistType            `json:"whitelist"`
}

type PatchedToken struct {
	ContractId         string                   `json:"contract_id"`
	CountryCode        string                   `json:"country_code"`
	DefaultProfileType *TokenDefaultProfileType `json:"default_profile_type,omitempty"`
	EnergyContract     *TokenEnergyContract     `json:"energy_contract,omitempty"`
	GroupId            *string                  `json:"group_id,omitempty"`
	Issuer             string                   `json:"issuer"`
	Language           *string                  `json:"language,omitempty"`
	LastUpdated        ocpi.DateTime            `json:"last_updated"`
	PartyId            string                   `json:"party_id"`
	Type               TokenType                `json:"type"`
	Uid                string                   `json:"uid"`
	Valid              bool                     `json:"valid"`
	VisualNumber       *string                  `json:"visual_number,omitempty"`
	Whitelist          WhitelistType            `json:"whitelist"`
}

// TokenDefaultProfileType defines model for Token.DefaultProfileType.
type TokenDefaultProfileType string

// TokenType defines model for Token.Type.
type TokenType string

// WhitelistType defines model for Token.Whitelist.
type WhitelistType string

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
	ConnectorId string `json:"connector_id"`
	EvseUid     string `json:"evse_uid"`
	LocationId  string `json:"location_id"`
	ResponseUrl string `json:"response_url"`
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

// CommandType defines parameters for type of commands.
type CommandType string

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

// GetOcpiLocationsCountryCodePartyIdLocationIdParams defines parameters for GetOcpiLocationsCountryCodePartyIdLocationId.
type GetOcpiLocationsCountryCodePartyIdLocationIdParams struct {
	// EvseUid Evse.uid, required when requesting an EVSE or Connector object.
	EvseUid *string `form:"evse_uid,omitempty" json:"evse_uid,omitempty"`

	// ConnectorId Connector.id, required when requesting a Connector object.
	ConnectorId *string `form:"connector_id,omitempty" json:"connector_id,omitempty"`
}

// PatchOcpiLocationsCountryCodePartyIdLocationIdParams defines parameters for PatchOcpiLocationsCountryCodePartyIdLocationId.
type PatchOcpiLocationsCountryCodePartyIdLocationIdParams struct {
	// EvseUid Evse.uid, required when requesting an EVSE or Connector object.
	EvseUid *string `form:"evse_uid,omitempty" json:"evse_uid,omitempty"`

	// ConnectorId Connector.id, required when requesting a Connector object.
	ConnectorId *string `form:"connector_id,omitempty" json:"connector_id,omitempty"`
}

// PutOcpiLocationsCountryCodePartyIdLocationIdParams defines parameters for PutOcpiLocationsCountryCodePartyIdLocationId.
type PutOcpiLocationsCountryCodePartyIdLocationIdParams struct {
	// EvseUid Evse.uid, required when requesting an EVSE or Connector object.
	EvseUid *string `form:"evse_uid,omitempty" json:"evse_uid,omitempty"`

	// ConnectorId Connector.id, required when requesting a Connector object.
	ConnectorId *string `form:"connector_id,omitempty" json:"connector_id,omitempty"`
}

// GetOcpiLocationsLocationIdParams defines parameters for GetOcpiLocationsLocationId.
type GetOcpiLocationsLocationIdParams struct {
	// EvseUid Evse.uid, required when requesting an EVSE or Connector object.
	EvseUid *string `form:"evse_uid,omitempty" json:"evse_uid,omitempty"`

	// ConnectorId Connector.id, required when requesting a Connector object.
	ConnectorId *string `form:"connector_id,omitempty" json:"connector_id,omitempty"`
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

// GetOcpiTokensCountryCodePartyIdTokenUidParams defines parameters for GetOcpiTokensCountryCodePartyIdTokenUid.
type GetOcpiTokensCountryCodePartyIdTokenUidParams struct {
	// Type Token.type of the Token to retrieve. Default if omitted: RFID
	Type *GetOcpiTokensCountryCodePartyIdTokenUidParamsType `form:"type,omitempty" json:"type,omitempty"`
}

// GetOcpiTokensCountryCodePartyIdTokenUidParamsType defines parameters for GetOcpiTokensCountryCodePartyIdTokenUid.
type GetOcpiTokensCountryCodePartyIdTokenUidParamsType string

// PatchOcpiTokensCountryCodePartyIdTokenUidParams defines parameters for PatchOcpiTokensCountryCodePartyIdTokenUid.
type PatchOcpiTokensCountryCodePartyIdTokenUidParams struct {
	// Type Token.type of the Token to retrieve. Default if omitted: RFID
	Type *PatchOcpiTokensCountryCodePartyIdTokenUidParamsType `form:"type,omitempty" json:"type,omitempty"`
}

// PatchOcpiTokensCountryCodePartyIdTokenUidParamsType defines parameters for PatchOcpiTokensCountryCodePartyIdTokenUid.
type PatchOcpiTokensCountryCodePartyIdTokenUidParamsType string

// PutOcpiTokensCountryCodePartyIdTokenUidParams defines parameters for PutOcpiTokensCountryCodePartyIdTokenUid.
type PutOcpiTokensCountryCodePartyIdTokenUidParams struct {
	// Type Token.type of the Token to retrieve. Default if omitted: RFID
	Type *PutOcpiTokensCountryCodePartyIdTokenUidParamsType `form:"type,omitempty" json:"type,omitempty"`
}

// PutOcpiTokensCountryCodePartyIdTokenUidParamsType defines parameters for PutOcpiTokensCountryCodePartyIdTokenUid.
type PutOcpiTokensCountryCodePartyIdTokenUidParamsType string
