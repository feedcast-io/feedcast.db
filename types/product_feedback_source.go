package types

type ProductFeedbackSource int16

const (
	ProductFeedbackSourceGoogle ProductFeedbackSource = iota + 1
	ProductFeedbackSourceMeta
	ProductFeedbackSourceBing
	ProductFeedbackSourceFreeListing
)
