package types

type Gender int16

const (
	GenderMale Gender = iota + 1
	GenderFemale
	GenderUnisex
)

func (c Gender) String() string {
	switch c {
	case GenderMale:
		return "male"
	case GenderFemale:
		return "female"
	case GenderUnisex:
		return "unisex"
	}

	return ""
}
