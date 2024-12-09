package types

type MerchantHearAboutUsType int16

const (
	MerchantHearAboutUsUnknown    = MerchantHearAboutUsType(0)
	MerchantHearAboutUsGoogle     = MerchantHearAboutUsType(1)
	MerchantHearAboutUsFacebook   = MerchantHearAboutUsType(2)
	MerchantHearAboutUsLinkedIn   = MerchantHearAboutUsType(3)
	MerchantHearAboutUsBlog       = MerchantHearAboutUsType(4)
	MerchantHearAboutUsReputation = MerchantHearAboutUsType(5)
	MerchantHeadAboutUsAutogen    = MerchantHearAboutUsType(6)
)
