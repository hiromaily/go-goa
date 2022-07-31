package types

import (
	. "goa.design/goa/v3/dsl"
)

var fieldTechName = func() {
	Description("Tech name")
	MinLength(1)
	MaxLength(40)
	Example("Golang")
}

// PayloadTech defines the data structure used in the create tech request body
var PayloadTech = Type("PayloadTech", func() {
	Attribute("name", String, "Tech Name", fieldTechName)
})

// RTTech is the tech resource result type.
var RTTech = ResultType("application/vnd.tech+json", func() {
	Description("A tech information")

	ContentType("application/json")

	Reference(PayloadTech) // name
	Reference(TypeFooter)

	Attributes(func() {
		Attribute("id", Int, "Tech ID", fieldID)
		Attribute("name")
		Attribute("created_at")
		Attribute("updated_at")

		Required("name")
	})

	View("default", func() {
		Attribute("id")
		Attribute("name")
	})

	View("id", func() {
		Description("id is the view used for C U D")
		Attribute("id")
	})
})
