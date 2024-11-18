package types

type LogTypes int16

const (
	LogTypeLogin                  = LogTypes(100)
	LogTypeSubscriptionInit       = LogTypes(103)
	LogTypeSubscriptionUpdate     = LogTypes(104)
	LogTypeFeedUpdate             = LogTypes(105)
	LogTypeCampaignInit           = LogTypes(106)
	LogTypeCampaignUpdate         = LogTypes(107)
	LogTypeFeedImportEnd          = LogTypes(121)
	LogTypeAccountPaymentMissing  = LogTypes(123)
	LogTypeInactiveCampaigns      = LogTypes(124)
	LogTypeLimitedCampaigns       = LogTypes(125)
	LogTypeUnlinkCustomerReseller = LogTypes(126)
	LogTypeFeedMissingUrl         = LogTypes(129)
	LogTypeFeedImportError        = LogTypes(130)
	LogTypeMerchantCenterWarning  = LogTypes(131)
	LogTypeCampaignNoConversions  = LogTypes(132)
	LogTypeCampaignLowRoi         = LogTypes(133)
	LogTypeCampaignUpdateError    = LogTypes(134)
	LogTypeEntityUpdate           = LogTypes(135)
	LogTypeFeedPixelMissing       = LogTypes(136)
	LogTypePlatformReset          = LogTypes(137)
	LogTypeProductAiSuccess       = LogTypes(138)
)
