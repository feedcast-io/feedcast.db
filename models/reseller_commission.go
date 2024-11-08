package models

type ResellerCommission struct {
	ID                      int32
	Reseller                Reseller
	ResellerId              int32
	CommissionRate          int32
	MinInvoiceCall          int32
	WelcomeCommissionAmount int32
}
