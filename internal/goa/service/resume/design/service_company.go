package design

import (
	"resume/design/types"

	. "goa.design/goa/v3/dsl"
)

// ----------------------------------------------------------------------------
// Service company
// ----------------------------------------------------------------------------

var _ = Service(resourcePrefix+"company", func() {
	Description("The company service returns company data")

	Security(JWT)
	HTTP(func() {
		// BasePath
		Path("/company")
	})

	Error("NotFound")
	Error("BadRequest")

	Method("companyList", func() {
		Description("List all companies")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
		})
		Result(CollectionOf(types.RTCompany))

		HTTP(func() {
			GET("")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})

	Method("getCompany", func() {
		Description("Retrieve company with given company_id")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("companyID", Int, "Company ID", func() {
				Minimum(1)
			})
			Required("companyID")
		})
		Result(types.RTCompany)

		HTTP(func() {
			GET("/{companyID}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})

	Method("createCompany", func() {
		Description("Create new company")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Extend(types.PayloadCompany)
			Required("name", "country_id", "address")
		})
		Result(types.RTCompany)

		HTTP(func() {
			POST("")
			Response(StatusCreated)
			Response("BadRequest", StatusBadRequest)
		})
	})

	Method("updateCompany", func() {
		Description("Update company data")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("companyID", Int, "Company ID", func() {
				Minimum(1)
			})
			Extend(types.PayloadCompany)
			Required("companyID")
		})

		HTTP(func() {
			PUT("/{companyID}")
			Response(StatusOK)
			Response("BadRequest", StatusBadRequest)
			Response("NotFound", StatusNotFound)
		})
	})

	Method("deleteCompany", func() {
		Description("Delete company")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("companyID", Int, "Company ID", func() {
				Minimum(1)
			})
			Required("companyID")
		})

		HTTP(func() {
			DELETE("/{companyID}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
		})
	})
})

// ----------------------------------------------------------------------------
// Service company branch
// ----------------------------------------------------------------------------

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
