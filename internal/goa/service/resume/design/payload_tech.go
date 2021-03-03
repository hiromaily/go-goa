package design

import (
	. "goa.design/goa/v3/dsl"
)

//-----------------------------------------------------------------------------
// Define fields
//-----------------------------------------------------------------------------
// Company
var fieldTechName = func() {
	Description("Tech name")
	MinLength(1)
	MaxLength(40)
	Example("Golang")
}

//-----------------------------------------------------------------------------
// TechPayload defines the data structure used in the create user request body.
//-----------------------------------------------------------------------------
var PayloadTech = Type("PayloadTech", func() {
	Attribute("name", String, "Tech Name", fieldTechName)
})
