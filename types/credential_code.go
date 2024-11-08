package types

type CredentialCode string

const (
	CredentialCodeAdsFeedcast      = CredentialCode("adwords.feedcast")
	CredentialCodeAdsMcc           = CredentialCode("adwords.mcc")
	CredentialCodeShoppingCss      = CredentialCode("shopping.feedcast")
	CredentialCodeShoppingNoCss    = CredentialCode("shopping.no-feedcast")
	CredentialCodeShoppingMcc      = CredentialCode("shopping.mcc")
	CredentialCodeFacebookFeedcast = CredentialCode("facebook.feedcast")
	CredentialCodeSheetsFeedcast   = CredentialCode("sheets.feedcast")
	CredentialCodeBingAds          = CredentialCode("bing.ads")
	CredentialCodeBingContentApi   = CredentialCode("bing.contentapi")
	CredentialCodeWooCommerce      = CredentialCode("woocommerce")
	CredentialCodePrestaShop       = CredentialCode("prestashop")
)
