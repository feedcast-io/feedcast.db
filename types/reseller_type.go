package types

type ResellerType int16

const (
	ResellerTypePayer ResellerType = iota + 1
	ResellerTypeNonPayer
	ResellerTypeNonPayerWithCommission
	ResellerTypeCommercial
)
