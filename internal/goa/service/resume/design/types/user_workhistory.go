package types

import (
	. "goa.design/goa/v3/dsl"
)

var fieldJobTitle = func() {
	Description("Job Title")
	MinLength(2)
	MaxLength(40)
	Example("Developer")
}

var fieldCountryCode = func() {
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

// PayloadUserWorkHistory defines the data structure used in the create user-work-history request body
var PayloadUserWorkHistory = Type("PayloadUserWorkHistory", func() {
	Attribute("title", String, "Job Title", fieldJobTitle)
	Attribute("company_name", String, "Company", fieldCompanyName)
	Attribute("country_name", String, "Country", fieldCountryCode)
	Attribute("term", String, "Working term", fieldTerm)
	Attribute("description", Any, "Description Json", fieldDescription)
	Attribute("techs", Any, "Used Techs as Array", fieldTechs)
})

// RTUserWorkHistory is the user work history resource result type.
var RTUserWorkHistory = ResultType("application/vnd.userworkhistory+json", func() {
	// Response Description
	Description("A user information")

	ContentType("application/json")

	Reference(PayloadUserWorkHistory)

	Attributes(func() {
		Attribute("title")
		Attribute("company_name")
		Attribute("country_name")
		Attribute("term")
		Attribute("description")
		Attribute("techs")

		//Required("title", "company_name", "country_name")
	})

	View("default", func() {
		Attribute("title")
		Attribute("company_name")
		Attribute("country_name")
		Attribute("term")
		Attribute("description")
		Attribute("techs")
	})
})
