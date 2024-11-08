package types

type WeightUnit int16

const (
	WeightUnitLb WeightUnit = iota + 1
	WeightUnitOz
	WeightUnitG
	WeightUnitKg
)

func (c WeightUnit) String() string {
	switch c {
	case WeightUnitLb:
		return "lb"
	case WeightUnitOz:
		return "oz"
	case WeightUnitG:
		return "g"
	case WeightUnitKg:
		return "kg"
	}

	return ""
}
