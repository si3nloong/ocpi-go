package ocpi230

import (
	"github.com/go-playground/validator/v10"
	"github.com/si3nloong/ocpi-go/internal/util"
	"github.com/si3nloong/ocpi-go/ocpi"
)

func init() {
	ocpi.RegisterValidation("authMethod230", func(fl validator.FieldLevel) bool {
		return util.Contains([]AuthMethod{
			AuthMethodAuthRequest,
			AuthMethodCommand,
			AuthMethodWhitelist,
		}, AuthMethod(fl.Field().String()))
	})
	ocpi.RegisterValidation("connectorFormat230", func(fl validator.FieldLevel) bool {
		return util.Contains([]ConnectorFormat{
			ConnectorFormatCable,
			ConnectorFormatSocket,
		}, ConnectorFormat(fl.Field().String()))
	})
	ocpi.RegisterValidation("connectionStatus230", func(fl validator.FieldLevel) bool {
		return util.Contains([]ConnectionStatus{
			ConnectionStatusConnected,
			ConnectionStatusOffline,
			ConnectionStatusPlanned,
			ConnectionStatusSuspended,
		}, ConnectionStatus(fl.Field().String()))
	})
	ocpi.RegisterValidation("dayOfWeek230", func(fl validator.FieldLevel) bool {
		return util.Contains([]DayOfWeek{
			DayOfWeekMonday,
			DayOfWeekTuesday,
			DayOfWeekWednesday,
			DayOfWeekThursday,
			DayOfWeekFriday,
			DayOfWeekSaturday,
			DayOfWeekSunday,
		}, DayOfWeek(fl.Field().String()))
	})
	ocpi.RegisterValidation("energySourceCategory230", func(fl validator.FieldLevel) bool {
		return util.Contains([]EnergySourceCategory{
			EnergySourceCategoryNuclear,
			EnergySourceCategoryGeneralFossil,
			EnergySourceCategoryCoal,
			EnergySourceCategoryGas,
			EnergySourceCategoryGeneralGreen,
			EnergySourceCategorySolar,
			EnergySourceCategoryWind,
			EnergySourceCategoryWater,
		}, EnergySourceCategory(fl.Field().String()))
	})
	ocpi.RegisterValidation("reservationRestrictionType230", func(fl validator.FieldLevel) bool {
		return util.Contains([]ReservationRestrictionType{
			ReservationRestrictionTypeReservation,
			ReservationRestrictionTypeReservationExpires,
		}, ReservationRestrictionType(fl.Field().String()))
	})
	ocpi.RegisterValidation("role230", func(fl validator.FieldLevel) bool {
		return util.Contains([]Role{
			RoleCPO,
			RoleEMSP,
			RoleNAP,
			RoleNSP,
			RoleOther,
			RoleSCSP,
		}, Role(fl.Field().String()))
	})
	ocpi.RegisterValidation("status230", func(fl validator.FieldLevel) bool {
		return util.Contains([]Status{
			StatusAvailable,
			StatusBlocked,
			StatusCharging,
			StatusInOperative,
			StatusOutOfOrder,
			StatusPlanned,
			StatusRemoved,
			StatusReserved,
			StatusUnknown,
		}, Status(fl.Field().String()))
	})
	ocpi.RegisterValidation("sessionStatus230", func(fl validator.FieldLevel) bool {
		return util.Contains([]SessionStatus{
			SessionStatusActive,
			SessionStatusCompleted,
			SessionStatusInvalid,
			SessionStatusPending,
			SessionStatusReservation,
		}, SessionStatus(fl.Field().String()))
	})
	ocpi.RegisterValidation("tariffDimensionType230", func(fl validator.FieldLevel) bool {
		return util.Contains([]TariffDimensionType{
			TariffDimensionTypeEnergy,
			TariffDimensionTypeFlat,
			TariffDimensionTypeParkingTime,
			TariffDimensionTypeTime,
		}, TariffDimensionType(fl.Field().String()))
	})
	ocpi.RegisterValidation("profileType230", func(fl validator.FieldLevel) bool {
		return util.Contains([]ProfileType{
			ProfileTypeCheap,
			ProfileTypeFast,
			ProfileTypeGreen,
			ProfileTypeRegular,
		}, ProfileType(fl.Field().String()))
	})
	ocpi.RegisterValidation("powerType230", func(fl validator.FieldLevel) bool {
		return util.Contains([]PowerType{
			PowerTypeAC1Phase,
			PowerTypeAC2Phase,
			PowerTypeAC2PhaseSplit,
			PowerTypeAC3Phase,
			PowerTypeDC,
		}, PowerType(fl.Field().String()))
	})
}
