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
