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
	Attribute("tech_name", String, "Tech Name", fieldTechName)
})

// RTTech is the tech resource result type.
var RTTech = ResultType("application/vnd.tech+json", func() {
	Description("A tech information")

	ContentType("application/json")

	Reference(PayloadTech) // name
	Reference(TypeFooter)

	Attributes(func() {
		Attribute("tech_id", Int, "Tech ID", fieldID)
		Attribute("tech_name")
		Attribute("created_at")
		Attribute("updated_at")

		//Required("tech_name")
	})

	View("default", func() {
		Attribute("tech_id")
		Attribute("tech_name")
	})

	View("id", func() {
		Description("only tech's id")
		Attribute("tech_id")
	})
})
