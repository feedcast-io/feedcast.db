package types

type Availability int16

const (
	AvailabilityInStock Availability = iota + 1
	AvailabilityOutOfStock
	AvailabilityBackOrder
)

func (c Availability) String() string {
	switch c {
	case AvailabilityInStock:
		return "in_stock"
	case AvailabilityOutOfStock:
		return "out_of_stock"
	case AvailabilityBackOrder:
		return "backorder"
	}

	return ""
}
