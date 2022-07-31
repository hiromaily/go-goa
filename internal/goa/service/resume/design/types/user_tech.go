package types

import (
	. "goa.design/goa/v3/dsl"
)

// PayloadUserTech defines the data structure used in the create user-tech request body
var PayloadUserTech = Type("PayloadUserTech", func() {
	Attribute("tech_name", String, "Tech name", fieldTechName)
})

// RTUserTech is the user tech resource result type.
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
