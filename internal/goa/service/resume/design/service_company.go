package design

import (
	. "goa.design/goa/v3/dsl"
	"resume/design/types"
)

// Service company
var _ = Service(resourcePrefix+"company", func() {
	Description("The company service returns company data")

	Security(JWT)
	// BasePath
	HTTP(func() {
		Path("/company")
	})

	Method("companyList", func() {
		Description("List all companies")
		Error("NoContent")
		Error("BadRequest")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
		})
		Result(CollectionOf(types.RTCompany))
		HTTP(func() {
			GET("")
			Response(StatusOK)
		})
	})

	Method("getCompany", func() {
		Description("Retrieve company with given company_id")
		Error("NoContent")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("company_id", Int, "Company ID")
		})
		//Result(CollectionOf(types.RTCompany))
		Result(types.RTCompany)
		HTTP(func() {
			GET("/{company_id}")
			Response(StatusOK)
		})
	})

	Method("createCompany", func() {
		Description("Create new company")
		Error("BadRequest")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Extend(types.PayloadCompany)
			Required("name", "country_id", "address")
		})
		HTTP(func() {
			POST("")
			Response(StatusOK)
			Response(StatusCreated)
		})
	})

	Method("updateCompany", func() {
		Description("Change company properties")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("company_id", Int, "Company ID")
			Extend(types.PayloadCompany)
			Required("name", "country_id", "address")
		})
		HTTP(func() {
			PUT("/{company_id}")
			Response(StatusOK)
		})
	})

	Method("deleteCompany", func() {
		Description("Delete company")
		Error("BadRequest")
		Error("NotFound")
		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("company_id", Int, "Company ID")
		})
		HTTP(func() {
			DELETE("/{company_id}")
			Response(StatusOK)
		})
	})
})

// Service company branch
//var _ = Service(resourcePrefix+"companybranch", func() {
//	Description("The company branch service returns company branch data")
//
//	Security(JWT)
//	// BasePath
//	HTTP(func() {
//		Path("/company/branch")
//	})
//
//	Method("getCompanyBranch", func() {
//		Description("Get company branch with given id")
//		Error("NotFound")
//		Error("BadRequest")
//		Payload(func() {
//			Token("token", String, "JWT token used to perform authorization")
//			Attribute("company_branch_id", Int, "Company branch ID")
//		})
//		HTTP(func() {
//			GET("/{company_branch_id}")
//			Response(StatusOK)
//		})
//	})
//
//	Method("createCompanyBranch", func() {
//		Description("Create new company branch")
//		Error("BadRequest")
//		Payload(func() {
//			Token("token", String, "JWT token used to perform authorization")
//			Attribute("company_id", Int, "Company ID")
//			Extend(types.PayloadCompanyPartial)
//			Required("country_id", "address")
//		})
//		HTTP(func() {
//			POST("")
//			Response(StatusOK)
//			Response(StatusCreated)
//		})
//	})
//
//	Method("updateCompanyBranch", func() {
//		Description("Change company branch properties")
//		Error("BadRequest")
//		Error("NotFound")
//		Payload(func() {
//			Token("token", String, "JWT token used to perform authorization")
//			Attribute("company_branch_id", Int, "Company branch ID")
//			Extend(types.PayloadCompanyPartial)
//			Required("country_id", "address")
//		})
//		HTTP(func() {
//			PUT("/{company_branch_id}")
//			Response(StatusOK)
//		})
//	})
//
//	Method("deleteCompanyBranch", func() {
//		Description("Delete company branch")
//		Error("BadRequest")
//		Error("NotFound")
//		Payload(func() {
//			Token("token", String, "JWT token used to perform authorization")
//			Attribute("company_branch_id", Int, "Company branch ID")
//		})
//		HTTP(func() {
//			DELETE("/{company_branch_id}")
//			Response(StatusOK)
//		})
//	})
//})
