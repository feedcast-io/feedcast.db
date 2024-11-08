package types

type ConditionType int16

const (
	ConditionTypeNew ConditionType = iota + 1
	ConditionTypeRefurb
	ConditionTypeUsed
)

func (c ConditionType) String() string {
	switch c {
	case ConditionTypeNew:
		return "new"
	case ConditionTypeRefurb:
		return "refurbished"
	case ConditionTypeUsed:
		return "used"
	}

	return ""
}
