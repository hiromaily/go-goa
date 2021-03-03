package design

import (
	. "goa.design/goa/v3/dsl"
)

//-----------------------------------------------------------------------------
// Define fields
//-----------------------------------------------------------------------------
// User
var fieldUserName = func() {
	Description("User name")
	MinLength(2)
	MaxLength(20)
	Example("Hiroki")
}

var fieldEmail = func() {
	Description("E-mail of user")
	Format("email")
	Example("hy@gmail.com")
}

var fieldPassword = func() {
	Description("Password")
	MinLength(8)
	MaxLength(20)
	Example("xxxxxxxx")
}

//-----------------------------------------------------------------------------
// UserPayload defines the data structure used in the create user request body.
//-----------------------------------------------------------------------------
var PayloadUser = Type("PayloadUser", func() {
	Attribute("user_name", String, "User name", fieldUserName)
	Attribute("email", String, "E-mail of user", fieldEmail)
	Attribute("password", String, "Password", fieldPassword)
})
