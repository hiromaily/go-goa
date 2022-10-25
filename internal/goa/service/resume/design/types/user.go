package types

import (
	. "goa.design/goa/v3/dsl"
)

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

// PayloadUser defines the data structure used in the create user request body
var PayloadUser = Type("PayloadUser", func() {
	Attribute("user_name", String, "User name", fieldUserName)
	Attribute("email", String, "User E-mail", fieldEmail)
	Attribute("password", String, "User password", fieldPassword)
})

// RTUser is the user type for response
var RTUser = ResultType("application/vnd.user+json", func() {
	Description("User information")

	ContentType("application/json")

	// Reference sets a type or result type reference. The value itself can be a
	// type or a result type. The reference type attributes define the default
	// properties for attributes with the same name in the type using the reference.
	Reference(PayloadUser)
	Reference(TypeFooter)

	Attributes(func() {
		Attribute("id", Int, "User ID", fieldID)
		Attribute("user_name")
		Attribute("email")
		Attribute("password")
		Attribute("created_at")
		Attribute("updated_at")

		// Required adds a "required" validation to the attribute.
		//Required("user_name", "email", "password")
	})

	// View defines a rendering of the result type
	View("default", func() {
		Attribute("id")
		Attribute("user_name")
		Attribute("email")
	})

	View("id", func() {
		Attribute("id")
		Required("id")
	})
})

// RTUserCompany is the user company type for response
// TODO: unimplemented
var RTUserCompany = ResultType("application/vnd.usercomany+json", func() {
	Description("A user who belongs to which companies")

	Attributes(func() {
		Attribute("id", Int, "ID of user company", func() {
			Example(1)
		})
		Attribute("href", String, "API href of bottle", func() {
			Example("/user/1/company")
		})

		Attribute("user_id", Int, "ID of user id", func() {
			Example(1)
		})

		Attribute("company_id", Int, "ID of user id", func() {
			Example(1)
		})

		Reference(TypeFooter)

		Required("id", "href", "user_id", "company_id")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("user_id")
		Attribute("company_id")
	})
})
