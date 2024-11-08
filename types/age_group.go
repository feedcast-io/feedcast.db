package types

type AgeGroup int16

const (
	AgeGroupNewBorn AgeGroup = iota + 1
	AgeGroupInfant
	AgeGroupToddler
	AgeGroupKids
	AgeGroupAdult
)

func (c AgeGroup) String() string {
	switch c {
	case AgeGroupNewBorn:
		return "newborn"
	case AgeGroupInfant:
		return "infant"
	case AgeGroupToddler:
		return "toddler"
	case AgeGroupKids:
		return "kids"
	case AgeGroupAdult:
		return "adult"
	}

	return ""
}
