package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//Response data
//it contains the media type data structures used by resource actions to build the responses

// User is the user resource media type.
var User = MediaType("application/vnd.user+json", func() {
	Description("A user information")

	//Though set Reference, Attribute is required
	Reference(UserPayload)
	Reference(ResponseCommon)

	Attributes(func() {
		Attribute("id", Integer, "ID of user", func() {
			Example(1)
		})
		Attribute("href", String, "API href of user", func() {
			Example("/user/1")
		})

		Attribute("name")
		Attribute("email")
		Attribute("created_at")
		Attribute("updated_at")

		//when field is zero or empty, data is not return unless it's not set in Required
		Required("id", "href", "name", "email")
	})

	//View defines a rendering of the media type
	//Media types may have multiple views　(it can change response pattern)
	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("email")
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

// Company is the company resource media type.
var Company = MediaType("application/vnd.company+json", func() {
	Description("A company information")

	Reference(CompanyPayload)
	Reference(ResponseCommon)

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

//TODO: unimplemented
// UserCompany is the company resource media type.
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

		Reference(ResponseCommon)

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
