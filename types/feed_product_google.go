package types

import (
	"encoding/xml"
	"log"
	"reflect"
)

type GoogleShipping struct {
	Country         string `xml:"g:country"`
	Price           string `xml:"g:price"`
	MinHandlingTime string `xml:"g:min_handling_time,omitempty"`
	MaxHandlingTime string `xml:"g:max_handling_time,omitempty"`
	MinTransitTime  string `xml:"g:min_transit_time,omitempty"`
	MaxTransitTime  string `xml:"g:max_transit_time,omitempty"`
}

type GoogleProduct struct {
	XMLName              xml.Name               `xml:"item"`
	Reference            string                 `xml:"g:id,omitempty"`
	Title                string                 `xml:"g:title,omitempty" customerData:"title"`
	Description          string                 `xml:"g:description,omitempty" customerData:"description"`
	Link                 string                 `xml:"g:link,omitempty" customerData:"link"`
	Image                string                 `xml:"g:image_link,omitempty" customerData:"image_link"`
	Condition            string                 `xml:"g:condition,omitempty" customerData:"condition_k"`
	Available            string                 `xml:"g:availability,omitempty" customerData:"availability_k"`
	AvailableDate        string                 `xml:"g:availability_date,omitempty" customerData:"availability_date_k"`
	Gtin                 string                 `xml:"g:gtin,omitempty" customerData:"gtin"`
	Brand                string                 `xml:"g:brand,omitempty" customerData:"brand_k"`
	Mpn                  string                 `xml:"g:mpn,omitempty" customerData:"mpn"`
	Gender               string                 `xml:"g:gender,omitempty" customerData:"gender_k"`
	Color                string                 `xml:"g:color,omitempty" customerData:"color_k"`
	AgeGroup             string                 `xml:"g:age_group,omitempty" customerData:"age_group_k"`
	Bundle               string                 `xml:"g:is_bundle,omitempty" customerData:"is_bundle"`
	Label0               string                 `xml:"g:custom_label_0,omitempty" customerData:"custom_label_0_k"`
	Label1               string                 `xml:"g:custom_label_1,omitempty" customerData:"custom_label_1_k"`
	Label2               string                 `xml:"g:custom_label_2,omitempty" customerData:"custom_label_2_k"`
	Label3               string                 `xml:"g:custom_label_3,omitempty" customerData:"custom_label_3_k"`
	Label4               string                 `xml:"g:custom_label_4,omitempty" customerData:"custom_label_4_k"`
	ProductType          string                 `xml:"g:product_type,omitempty" customerData:"product_type_k"`
	AdsRedirect          string                 `xml:"g:ads_redirect,omitempty" customerData:"ads_redirect"`
	Size                 string                 `xml:"g:size,omitempty" customerData:"size_k"`
	MobileLink           string                 `xml:"g:mobile_link,omitempty" customerData:"mobile_link"`
	ItemGroupId          string                 `xml:"g:item_group_id,omitempty" customerData:"item_group_id_k"`
	Material             string                 `xml:"g:material,omitempty" customerData:"material_k"`
	Category             string                 `xml:"g:google_product_category,omitempty" customerData:"google_product_category_k"`
	ShippingWeight       string                 `xml:"g:shipping_weight,omitempty" customerData:"shipping_weight"`
	ProductWeight        string                 `xml:"g:product_weight,omitempty" customerData:"product_weight"`
	Shipping             *GoogleShipping        `xml:"g:shipping,omitempty"`
	ExcludedDestinations []string               `xml:"g:excluded_destination,omitempty"`
	Price                string                 `xml:"g:price,omitempty" customerData:"price"`
	SalePrice            string                 `xml:"g:sale_price,omitempty" customerData:"sale_price"`
	IdExists             string                 `xml:"g:identifier_exists,omitempty" customerData:"identifier_exists"`
	CustomData           map[string]interface{} `xml:"-"`
	CustomDataAi         map[string]interface{} `xml:"-"`
	Options              map[string]string      `xml:"-"`
	Language             string                 `xml:"-"`
	Country              string                 `xml:"-"`
}

// Override product attributes from CustomData map when key match product field `customerData`tag.
func (p *GoogleProduct) AppendCustomData() {
	datas := []map[string]interface{}{
		p.CustomDataAi,
		p.CustomData,
	}

	for _, d := range datas {
		if len(d) > 0 {
			t := reflect.TypeOf(*p)
			for i, maxId := 0, t.NumField(); i < maxId; i++ {
				field := t.Field(i)
				tag := field.Tag.Get("customerData")
				if len(tag) > 0 {
					customValue, ok := d[tag].(string)
					if ok && len(customValue) > 0 {
						reflect.ValueOf(p).Elem().FieldByName(field.Name).SetString(customValue)
						log.Printf("Replace field %s with value %s", field.Name, customValue)
					}
				}
			}
		}
	}
}
