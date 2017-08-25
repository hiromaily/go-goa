package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//-----------------------------------------------------------------------------
// Define fields
//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// UserPayload defines the data structure used in the create user request body.
//-----------------------------------------------------------------------------
var UserTechPayload = Type("UserTechPayload", func() {
	Attribute("tech_name", String, "Tech name", fieldTechName)
})
