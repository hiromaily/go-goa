package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// For only footer
// ResponseCommon defines common data of response
var ResponseCommon = Type("BottlePayload", func() {
	Attribute("created_at", DateTime, "Date of creation", func() {
		//Format("date-time")
		Example("2017-03-10T15:00:00Z")
	})
	Attribute("updated_at", DateTime, "Date of update", func() {
		//Format("date-time")
		Example("2017-03-10T15:00:00Z")
	})
	Attribute("created_by", Integer, "user id", func() {
		Example(1)
	})
	Attribute("updated_by", Integer, "user id", func() {
		Example(1)
	})

})

// UserPayload defines the data structure used in the create user request body.
var UserPayload = Type("UserPayload", func() {
	Attribute("name", String, "Name of user", func() {
		Example("Hiroki")
	})
	Attribute("email", String, "E-mail of user", func() {
		Format("email")
		Example("hy@gmail.com")
	})
})

// CompanyPayload defines the data structure used in the create user request body.
var CompanyPayload = Type("CompanyPayload", func() {
	Attribute("name", String, "Name of company", func() {
		Example("Sony")
	})
	Attribute("country", String, "Country of HQ", func() {
		Example("Tokyo, Japan")
	})
	Attribute("address", String, "Address of company", func() {
		Example("Tokyo")
	})
})

// BottlePayload defines the data structure used in the create bottle request body.
// It is also the base type for the bottle media type used to render bottles.
//var BottlePayload = Type("BottlePayload", func() {
//	Attribute("name", func() {
//		MinLength(2)
//		Example("Number 8")
//	})
//	Attribute("vineyard", func() {
//		MinLength(2)
//		Example("Asti")
//	})
//	Attribute("varietal", func() {
//		MinLength(4)
//		Example("Merlot")
//	})
//	Attribute("vintage", Integer, func() {
//		Minimum(1900)
//		Maximum(2020)
//		Example(2012)
//	})
//	Attribute("color", func() {
//		Enum("red", "white", "rose", "yellow", "sparkling")
//	})
//	Attribute("sweetness", Integer, func() {
//		Minimum(1)
//		Maximum(5)
//	})
//	Attribute("country", func() {
//		MinLength(2)
//		Example("USA")
//	})
//	Attribute("region", func() {
//		Example("Napa Valley")
//	})
//	Attribute("review", func() {
//		MinLength(3)
//		MaxLength(300)
//		Example("Great and inexpensive")
//	})
//})
