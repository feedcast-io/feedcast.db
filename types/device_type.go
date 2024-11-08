package types

type DeviceType int8

const (
	DeviceTypeUnknown = DeviceType(0)
	DeviceTypeDesktop = DeviceType(1)
	DeviceTypeMobile  = DeviceType(2)
	DeviceTypeTablet  = DeviceType(3)
)
