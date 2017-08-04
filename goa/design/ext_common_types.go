package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//-----------------------------------------------------------------------------
// Define fields
//-----------------------------------------------------------------------------
// Common
var fieldID = func() {
	Description("ID")
	Minimum(1)
	Example(10)
}

var fieldFlg = func() {
	Description("Flg")
	MinLength(1)
	MaxLength(1)
	Example("1")
}

var fieldDatetime = func() {
	Description("Datetime")
	Format("date-time")
	Example("2017-03-10T15:00:00Z")
}

//-----------------------------------------------------------------------------
// Define Type
//-----------------------------------------------------------------------------

//TODO: CommonResponse defines common data of response
var CommonResponse = Type("CommonResponse", func() {
	Attribute("created_at", String, "Date of creation", fieldDatetime)
	Attribute("updated_at", String, "Date of update", fieldDatetime)
})

//var BottlePayload = Type("BottlePayload", func() {
//	Attribute("vintage", Integer, func() {
//		Minimum(1900)
//		Maximum(2020)
//		Example(2012)
//	})
//	Attribute("color", func() {
//		Enum("red", "white", "rose", "yellow", "sparkling")
//	})
//})
