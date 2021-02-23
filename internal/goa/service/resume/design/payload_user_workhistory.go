package design

import (
	. "goa.design/goa/v3/dsl"
)

//-----------------------------------------------------------------------------
// Define fields
//-----------------------------------------------------------------------------
// UserWorkHistory
var fieldJobTitle = func() {
	Description("Job Title")
	MinLength(2)
	MaxLength(40)
	Example("Developer")
}

var fieldContryCode = func() {
	Description("Country code")
	MinLength(2)
	MaxLength(2)
	Example("nl")
}

var fieldTerm = func() {
	Description("worked period")
	MinLength(10)
	MaxLength(20)
	Example("Jul 2017 - Aug 2017")
}

var fieldDescription = func() {
	Description("job description")
	Example(`["Developed resume site for job seeking."]`)
}

var fieldTechs = func() {
	Description("used techs")
	Example(`["Golang with goa", "Riot.js", "Semantic UI", "MySQL", "Docker", "Travis-CI"]`)
}

//`id`          int(11) NOT NULL AUTO_INCREMENT COMMENT'ID',
//`user_id`     int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'User ID',
//`company_branch_id`  int(11) COLLATE utf8_unicode_ci NOT NULL COMMENT'Company Branch ID',
//`title`       varchar(40) COLLATE utf8_unicode_ci NOT NULL COMMENT'Title',
//`description` json NOT NULL COMMENT'Description',
//`started_at`  date DEFAULT NULL COMMENT'Started Date',
//`ended_at`    date DEFAULT NULL COMMENT'Ended Date',
//`delete_flg`  char(1) COLLATE utf8_unicode_ci DEFAULT'0' COMMENT'delete flg',
//`created_at`  datetime DEFAULT CURRENT_TIMESTAMP COMMENT'created date',
//`updated_at`  datetime DEFAULT CURRENT_TIMESTAMP COMMENT'updated date',

//-----------------------------------------------------------------------------
// UserPayload defines the data structure used in the create user request body.
//-----------------------------------------------------------------------------
var PayloadUserWorkHistory = Type("PayloadUserWorkHistory", func() {
	Attribute("title", String, "Job Title", fieldJobTitle)
	Attribute("company", String, "Company", fieldCompanyName)
	Attribute("country", String, "Country", fieldContryCode)
	Attribute("term", String, "Country", fieldTerm)
	Attribute("description", Any, "Description Json", fieldDescription)
	Attribute("techs", Any, "Used Techs as Array", fieldTechs)
})
