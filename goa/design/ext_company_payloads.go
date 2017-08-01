package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// CompanyPayload defines the data structure used in the create user request body.
var CompanyPayload = Type("CompanyPayload", func() {
	Attribute("name", String, "Name of company", func() {
		MinLength(2)
		MaxLength(40)
		Example("Sony")
	})
	Attribute("country", String, "Country of HQ", func() {
		MinLength(2)
		MaxLength(40)
		Example("Japan")
	})
	Attribute("address", String, "Address of company", func() {
		MinLength(2)
		MaxLength(40)
		Example("Tokyo")
	})
})
