package ocpi230

import (
	"github.com/go-playground/validator/v10"
	"github.com/si3nloong/ocpi-go/internal/util"
	"github.com/si3nloong/ocpi-go/ocpi"
)

func init() {
	ocpi.RegisterValidation("connectorCapability230", func(fl validator.FieldLevel) bool {
		return util.Contains([]ConnectorCapability{
			ConnectorCapabilityISO151182PlugAndCharge,
			ConnectorCapabilityISO1511820PlugAndCharge,
		}, ConnectorCapability(fl.Field().String()))
	})
	ocpi.RegisterValidation("connectorFormat230", func(fl validator.FieldLevel) bool {
		return util.Contains([]ConnectorFormat{
			ConnectorFormatCable,
			ConnectorFormatSocket,
		}, ConnectorFormat(fl.Field().String()))
	})
	ocpi.RegisterValidation("profileType230", func(fl validator.FieldLevel) bool {
		return util.Contains([]ProfileType{
			ProfileTypeCheap,
			ProfileTypeFast,
			ProfileTypeGreen,
			ProfileTypeRegular,
		}, ProfileType(fl.Field().String()))
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
	ocpi.RegisterValidation("powerType230", func(fl validator.FieldLevel) bool {
		return util.Contains([]PowerType{
			PowerTypeAC1Phase,
			PowerTypeAC2Phase,
			PowerTypeAC2PhaseSplit,
			PowerTypeAC3Phase,
			PowerTypeDC,
		}, PowerType(fl.Field().String()))
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
}
