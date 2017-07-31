package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Define Response data
// It contains the media type data structures used by resource actions to build the responses
// Add variable and set MediaType object

// reference
// https://goa.design/reference/goa/design/apidsl/#func-mediatype-a-name-apidsl-mediatype-a

//-----------------------------------------------------------------------------
// Authorized is the authority resource media type.
//-----------------------------------------------------------------------------
var Authorized = MediaType("application/vnd.authorized+json", func() {
	// Response Description
	Description("An authorized response")

	Attributes(func() {
		Attribute("token", String, "JWT token", func() {
			Example("token.string")
		})
		Required("token")
	})

	//default MediaType
	View("default", func() {
		Attribute("token")
	})
})

//-----------------------------------------------------------------------------
// User is the user resource media type.
//-----------------------------------------------------------------------------
var User = MediaType("application/vnd.user+json", func() {
	// Response Description
	Description("A user information")

	// Reference can be used in: MediaType, Type
	// Though set Reference, Attribute is required
	Reference(UserPayload)
	Reference(CommonResponse)

	//`user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'User ID',
	//`first_name` varchar(20) COLLATE utf8_unicode_ci NOT NULL COMMENT 'First Name',
	//`last_name` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'Last Name',
	//`email` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'E-Mail Address',
	//`password` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'Password',
	//`oauth2_flg` char(1) COLLATE utf8_unicode_ci DEFAULT '0' COMMENT 'oauth_flg flg',
	//`delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT '0' COMMENT 'delete flg',
	//`create_datetime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'created date',
	//`update_datetime` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'updated date',

	Attributes(func() {
		Attribute("user_id", Integer, "User ID", func() {
			Minimum(1)
			Example(1)
		})
		Attribute("first_name")
		Attribute("last_name")
		Attribute("email")
		Attribute("password")
		Attribute("oauth2_flg")
		Attribute("create_datetime")
		Attribute("update_datetime")

		//when field is zero or empty, data is not return unless it's not set in Required
		Required("first_name", "last_name", "email", "password")
	})

	//View defines a rendering of the media type
	//Media types may have multiple views　(it can change response pattern)
	//View default is for UserList, GetUser,
	View("default", func() {
		Attribute("user_id")
		Attribute("first_name")
		Attribute("last_name")
		Attribute("email")
	})

	//View id is for UserList, GetUser,
	View("id", func() {
		Description("tiny is the view used to id list")
		Attribute("user_id")
	})
})

//-----------------------------------------------------------------------------
// Company is the company resource media type.
//-----------------------------------------------------------------------------
//TODO: unimplemented
var Company = MediaType("application/vnd.company+json", func() {
	Description("A company information")

	Reference(CompanyPayload)
	Reference(CommonResponse)

	Attributes(func() {
		Attribute("id", Integer, "ID of company", func() {
			Example(1)
		})
		Attribute("href", String, "API href of company", func() {
			Example("/company/1")
		})

		Attribute("name")
		Attribute("country")
		Attribute("address")
		Attribute("created_at")
		Attribute("updated_at")

		//when field is zero or empty, data is not return unless it's not set in Required
		Required("id", "href", "name", "country", "address")
	})

	//View defines a rendering of the media type
	//Media types may have multiple views　(it can change response pattern)
	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("country")
		Attribute("address")
		//Attribute("created_at")
		//Attribute("created_by")
	})

	View("tiny", func() {
		Description("tiny is the view used to id list")
		Attribute("id")
		Attribute("href")
		Attribute("name")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})
})

//-----------------------------------------------------------------------------
// UserCompany is the company resource media type.
//-----------------------------------------------------------------------------
//TODO: unimplemented
var UserCompany = MediaType("application/vnd.usercomany+json", func() {
	Description("A user who belongs to which companies")

	// So define types and take advantage of the Reference keyword
	// when the same attributes are shared between types and media types
	//Reference(BottlePayload)

	Attributes(func() {
		Attribute("id", Integer, "ID of user company", func() {
			Example(1)
		})
		Attribute("href", String, "API href of bottle", func() {
			Example("/user/1/company")
		})

		Attribute("user_id", Integer, "ID of user id", func() {
			Example(1)
		})

		Attribute("company_id", Integer, "ID of user id", func() {
			Example(1)
		})

		//Attribute("rating", Integer, "Rating of bottle between 1 and 5", func() {
		//	Minimum(1)
		//	Maximum(5)
		//})

		//From User object
		//Attribute("user", User, "User who belong to company")
		//From Company object
		//Attribute("company", Company, "user who work(ed) for company")

		Reference(CommonResponse)

		Required("id", "href", "user_id", "company_id")
		//Required("created_at")
	})

	//Links are rendered using the special "link" view.
	// Media types that are linked to must define that view. Here is an example
	// showing all the possible media type sub-definitions:

	//When including other media type from outside, this link would be required.

	//Links(func() {
	//	Link("user")
	//})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("user_id")
		Attribute("company_id")

		//Attribute("account", func() {
		//	View("tiny")
		//})
		//Attribute("links")
	})
})
