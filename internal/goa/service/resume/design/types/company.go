package types

import (
	. "goa.design/goa/v3/dsl"
)

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

// PayloadCompany defines the data structure used in the create company request body
var PayloadCompany = Type("PayloadCompany", func() {
	Attribute("name", String, "Company Name", fieldCompanyName)
	Attribute("country_id", Int, "Country's ID", fieldCountryID)
	Attribute("address", String, "Address of company", fieldAddress)
})

// PayloadCompanyPartial defines the data structure used in the create company request body
var PayloadCompanyPartial = Type("PayloadCompanyPartial", func() {
	Attribute("country_id", Int, "Country's ID", fieldCountryID)
	Attribute("address", String, "Address of company", fieldAddress)
})

// RTCompany is the company resource result type.
var RTCompany = ResultType("application/vnd.company+json", func() {
	Description("A company information")

	ContentType("application/json")

	Reference(PayloadCompany) // name, country_id, address
	Reference(TypeFooter)     // created_at, updated_at

	Attributes(func() {
		Attribute("id", Int, "Company Detail ID", fieldID)
		Attribute("company_id", Int, "Company ID", fieldID)
		Attribute("name")
		Attribute("is_hq")
		Attribute("country_name")
		Attribute("address")
		Attribute("created_at")
		Attribute("updated_at")

		Required("name", "address")
	})

	View("default", func() {
		Attribute("id")
		Attribute("company_id")
		Attribute("name")
		Attribute("is_hq")
		Attribute("country_name")
		Attribute("address")
		//Attribute("created_at")
		//Attribute("created_by")
	})

	View("detailid", func() {
		Description("only company's detail id")
		Attribute("id")
	})

	View("id", func() {
		Description("only company's id")
		Attribute("company_id")
	})

	View("idname", func() {
		Description("only company's id and name")
		Attribute("company_id")
		Attribute("name")
	})
})
