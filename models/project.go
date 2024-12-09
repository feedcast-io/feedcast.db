package models

type Project struct {
	ID         int32
	Name       string
	Merchant   Merchant
	MerchantId int32
}
