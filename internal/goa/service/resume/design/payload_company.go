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
	Attribute("name", String, "Company Name", fieldCompanyName)
	Attribute("country_id", Int, "Country's ID", fieldCountryID)
	Attribute("address", String, "Address of company", fieldAddress)
})
