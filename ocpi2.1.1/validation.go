package ocpi211

import (
	"github.com/go-playground/validator/v10"
	"github.com/si3nloong/ocpi-go/internal/util"
	"github.com/si3nloong/ocpi-go/ocpi"
)

func init() {
	ocpi.RegisterValidation("authMethod211", func(fl validator.FieldLevel) bool {
		return util.Contains([]AuthMethod{
			AuthMethodAuthRequest,
			AuthMethodWhitelist,
		}, AuthMethod(fl.Field().String()))
	})
	ocpi.RegisterValidation("connectorFormat211", func(fl validator.FieldLevel) bool {
		return util.Contains([]ConnectorFormat{
			ConnectorFormatCable,
			ConnectorFormatSocket,
		}, ConnectorFormat(fl.Field().String()))
	})
	ocpi.RegisterValidation("dayOfWeek211", func(fl validator.FieldLevel) bool {
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
	ocpi.RegisterValidation("energySourceCategory211", func(fl validator.FieldLevel) bool {
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
	ocpi.RegisterValidation("locationType211", func(fl validator.FieldLevel) bool {
		return util.Contains([]LocationType{
			LocationTypeOnStreet,
			LocationTypeParkingGarage,
			LocationTypeUndergroundGarage,
			LocationTypeParkingLot,
			LocationTypeOther,
			LocationTypeUnknown,
		}, LocationType(fl.Field().String()))
	})
	ocpi.RegisterValidation("status211", func(fl validator.FieldLevel) bool {
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
	ocpi.RegisterValidation("sessionStatus211", func(fl validator.FieldLevel) bool {
		return util.Contains([]SessionStatus{
			SessionStatusActive,
			SessionStatusCompleted,
			SessionStatusInvalid,
			SessionStatusPending,
		}, SessionStatus(fl.Field().String()))
	})
	ocpi.RegisterValidation("tariffDimensionType211", func(fl validator.FieldLevel) bool {
		return util.Contains([]TariffDimensionType{
			TariffDimensionTypeEnergy,
			TariffDimensionTypeFlat,
			TariffDimensionTypeParkingTime,
			TariffDimensionTypeTime,
		}, TariffDimensionType(fl.Field().String()))
	})
	ocpi.RegisterValidation("powerType211", func(fl validator.FieldLevel) bool {
		return util.Contains([]PowerType{
			PowerTypeAC1Phase,
			PowerTypeAC3Phase,
			PowerTypeDC,
		}, PowerType(fl.Field().String()))
	})
}
