package design

import (
	. "goa.design/goa/v3/dsl"
)

//-----------------------------------------------------------------------------
// Define fields
//-----------------------------------------------------------------------------
// Company
var fieldCompanyName = func() {
	Description("Company name")
	MinLength(2)
	MaxLength(40)
	Example("Company")
}

var fieldCountryID = func() {
	Description("Country ID")
	Minimum(1)
	Maximum(999)
	Example(110)
}

var fieldAddress = func() {
	Description("Company Address")
	MinLength(2)
	MaxLength(80)
	Example("Shinagawa Tokyo")
}

//-----------------------------------------------------------------------------
// CompanyPayload defines the data structure used in the create user request body.
//-----------------------------------------------------------------------------
var PayloadCompany = Type("PayloadCompany", func() {
	//`id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'Company ID',
	//`name`       varchar(20) COLLATE utf8_unicode_ci NOT NULL COMMENT'Company Name',

	//`id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'Company detail ID',
	//`company_id` int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'Company ID',
	//`hq_flg`     char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'headquarters flg',
	//`country_id` smallint COLLATE utf8_unicode_ci NOT NULL COMMENT'Country ID',
	//`address`    varchar(80) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT'Address',

	//Attribute("id", Integer, "Company Detail ID", fieldID)
	//Attribute("company_id", Integer, "Company ID", fieldID)
	Attribute("name", String, "Company Name", fieldCompanyName)
	//Attribute("hq_flg", String, "Headquarters flg", fieldFlg)
	Attribute("country_id", Int, "Country's ID", fieldCountryID)
	Attribute("address", String, "Address of company", fieldAddress)
})

//var PayloadCompanyTiny = Type("PayloadCompanyTiny", func() {
//	Attribute("country_id", Int, "Country's ID", fieldCountryID)
//	Attribute("address", String, "Address of company", fieldAddress)
//})
