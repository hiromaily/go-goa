package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// CompanyPayload defines the data structure used in the create user request body.
var CompanyPayload = Type("CompanyPayload", func() {
	//`id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'Company ID',
	//`name`       varchar(20) COLLATE utf8_unicode_ci NOT NULL COMMENT'Company Name',

	//`id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'Company detail ID',
	//`company_id` int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'Company ID',
	//`hq_flg`     char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'headquarters flg',
	//`country_id` smallint COLLATE utf8_unicode_ci NOT NULL COMMENT'Country ID',
	//`address`    varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT'Address',

	//Attribute("id", Integer, "Company Detail ID", func() {
	//	Minimum(1)
	//	Maximum(999999)
	//	Example(100)
	//})
	Attribute("company_id", Integer, "Company ID", func() {
		Minimum(1)
		Maximum(999999)
		Example(10)
	})
	Attribute("name", String, "Company Name", func() {
		MinLength(2)
		MaxLength(40)
		Example("Company")
	})
	Attribute("hq_flg", String, "Headquarters flg", func() {
		MinLength(1)
		MaxLength(1)
		Example("1")
	})
	Attribute("country_id", String, "Country's ID", func() {
		MinLength(2)
		MaxLength(60)
		Example("Japan")
	})
	Attribute("address", String, "Address of company", func() {
		MinLength(2)
		MaxLength(80)
		Example("Shinagawa Tokyo")
	})
})
