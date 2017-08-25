package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

//-----------------------------------------------------------------------------
// Define fields
//-----------------------------------------------------------------------------
// Company
var fieldTechName = func() {
	Description("Tech name")
	MinLength(1)
	MaxLength(40)
	Example("Golang")
}

//-----------------------------------------------------------------------------
// TechPayload defines the data structure used in the create user request body.
//-----------------------------------------------------------------------------
var TechPayload = Type("TechPayload", func() {
	//`id`         int(11) NOT NULL AUTO_INCREMENT COMMENT'Tech ID',
	//`name`       varchar(40) COLLATE utf8_unicode_ci NOT NULL COMMENT'Tech Name',
	//`delete_flg` char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flg',
	//`created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
	//`updated_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',

	Attribute("name", String, "Tech Name", fieldTechName)
})
