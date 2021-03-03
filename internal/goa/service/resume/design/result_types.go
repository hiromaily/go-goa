package design

import (
	. "goa.design/goa/v3/dsl"
)

// ResultType defines a result type used to describe a method response.
// MediaType is replaced by ResultType

//-----------------------------------------------------------------------------
// Authorized is the authority resource result type.
//-----------------------------------------------------------------------------
var RTAuthorized = ResultType("application/vnd.authorized+json", func() {
	Description("An authorized response")

	Attributes(func() {
		Attribute("token", String, "JWT token", func() {
			Example("token.string")
		})
		Attribute("id", Int, "User ID", fieldID)
		Required("token", "id")
	})

	View("default", func() {
		Attribute("token")
		Attribute("id")
	})
})

//-----------------------------------------------------------------------------
// User is the user resource result type.
//-----------------------------------------------------------------------------
var RTUser = ResultType("application/vnd.user+json", func() {
	Description("A user information")

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
		Required("user_name", "email", "password")
	})

	// View defines a rendering of the result type
	// it may have multiple views　(change response pattern)
	View("default", func() {
		Attribute("id")
		Attribute("user_name")
		Attribute("email")
	})

	View("id", func() {
		Description("id is the view used for C U D")
		Attribute("id")
	})
})

//-----------------------------------------------------------------------------
// UserTech is the user tech resource result type.
//-----------------------------------------------------------------------------
var RTUserTech = ResultType("application/vnd.usertech+json", func() {
	Description("A user information")

	ContentType("application/json")

	Reference(PayloadUserTech)

	//`id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'ID',
	//`user_id`    int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'User ID',
	//`tech_id`    int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'Tech ID',
	//`name`       varchar(40) COLLATE utf8_unicode_ci NOT NULL COMMENT'Tech Name',

	Attributes(func() {
		Attribute("id", Int, "Tech ID", fieldID)
		Attribute("tech_name")

		// Required adds a "required" validation to the attribute.
		Required("tech_name")
	})

	View("default", func() {
		Attribute("id")
		Attribute("tech_name")
	})

	View("tech", func() {
		Description("id is the view used for C U D")
		Attribute("tech_name")
	})
})

//-----------------------------------------------------------------------------
// UserTech is the user tech resource result type.
//-----------------------------------------------------------------------------
var RTUserWorkHistory = ResultType("application/vnd.userworkhistory+json", func() {
	// Response Description
	Description("A user information")

	ContentType("application/json")

	Reference(PayloadUserWorkHistory)

	Attributes(func() {
		Attribute("title")
		Attribute("company")
		Attribute("country")
		Attribute("term")
		Attribute("description")
		Attribute("techs")

		Required("title", "company", "country")
	})

	View("default", func() {
		Attribute("title")
		Attribute("company")
		Attribute("country")
		Attribute("term")
		Attribute("description")
		Attribute("techs")
	})
})

//-----------------------------------------------------------------------------
// Tech is the tech resource result type.
//-----------------------------------------------------------------------------
var RTTech = ResultType("application/vnd.tech+json", func() {
	Description("A tech information")

	ContentType("application/json")

	Reference(PayloadTech)
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

//-----------------------------------------------------------------------------
// Company is the company resource result type.
//-----------------------------------------------------------------------------
var RTCompany = ResultType("application/vnd.company+json", func() {
	Description("A company information")

	ContentType("application/json")

	Reference(PayloadCompany)
	Reference(TypeFooter)

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

//-----------------------------------------------------------------------------
// UserCompany is the company resource result type.
// TODO: unimplemented
//-----------------------------------------------------------------------------
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
