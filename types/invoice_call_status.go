package types

type InvoiceCallStatus int16

const (
	InvoiceCallStatusesPending = InvoiceCallStatus(1)
	InvoiceCallStatusesPaid    = InvoiceCallStatus(2)
)
