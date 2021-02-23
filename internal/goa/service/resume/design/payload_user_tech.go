package design

import (
	. "goa.design/goa/v3/dsl"
)

//-----------------------------------------------------------------------------
// Define fields
//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// UserPayload defines the data structure used in the create user request body.
//-----------------------------------------------------------------------------
var PayloadUserTech = Type("PayloadUserTech", func() {
	Attribute("tech_name", String, "Tech name", fieldTechName)
})
