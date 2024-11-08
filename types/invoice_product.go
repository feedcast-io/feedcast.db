package types

type InvoiceProduct string

const (
	InvoiceProductPackStarter      InvoiceProduct = "starter"
	InvoiceProductPackPremium      InvoiceProduct = "premium"
	InvoiceProductPackPro          InvoiceProduct = "pro"
	InvoiceProductTechnoFeedcast   InvoiceProduct = "techno_feedcast"
	InvoiceProductProductLimit1    InvoiceProduct = "opt_product_level_1"
	InvoiceProductProductLimit2    InvoiceProduct = "opt_product_level_2"
	InvoiceProductProductLimit3    InvoiceProduct = "opt_product_level_3"
	InvoiceProductProductLimit4    InvoiceProduct = "opt_product_level_4"
	InvoiceProductGoogleListingSeo InvoiceProduct = "opt_google_listing_seo"
	InvoiceProductInstagram        InvoiceProduct = "opt_instagram_shopping"
	InvoiceProductAddPlatforms     InvoiceProduct = "opt_additional_platforms"
	InvoiceProductAI               InvoiceProduct = "opt_ai"
)
