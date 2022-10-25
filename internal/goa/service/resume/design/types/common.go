package types

import (
	. "goa.design/goa/v3/dsl"
)

// Common
var fieldID = func() {
	Description("Key ID")
	Minimum(1)
	Example(10)
}

//var fieldFlg = func() {
//	Description("1/0 flag ")
//	MinLength(1)
//	MaxLength(1)
//	Example("1")
//}

var fieldDatetime = func() {
	Description("Datetime")
	Format("date-time")
	Example("2017-03-10T15:00:00Z")
}

var TypeFooter = Type("TypeFooter", func() {
	Attribute("created_at", String, "Date of creation", fieldDatetime)
	Attribute("updated_at", String, "Date of update", fieldDatetime)
})
