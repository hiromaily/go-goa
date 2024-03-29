package design

import (
	"fmt"
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
		Path(fmt.Sprintf("%s/company", baseAPIDir))
	})

	Error("NotFound")
	Error("BadRequest")
	Error("Unauthorized")

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
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("getCompany", func() {
		Description("returns company by given company_id")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("company_id", Int, "Company ID", func() {
				Minimum(1)
			})
			Required("company_id")
		})
		Result(types.RTCompany)

		HTTP(func() {
			GET("/{company_id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("createCompany", func() {
		Description("Create new company")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Extend(types.PayloadCompany)
			Required("company_name", "country_id", "address")
		})
		Result(types.RTCompany)

		HTTP(func() {
			POST("")
			Response(StatusCreated)
			Response("BadRequest", StatusBadRequest)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("updateCompany", func() {
		Description("Update company data")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("company_id", Int, "Company ID", func() {
				Minimum(1)
			})
			Extend(types.PayloadCompany)
			Required("company_id")
		})

		HTTP(func() {
			PUT("/{company_id}")
			Response(StatusOK)
			Response("BadRequest", StatusBadRequest)
			Response("NotFound", StatusNotFound)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("deleteCompany", func() {
		Description("Delete company")

		Payload(func() {
			Token("token", String, "JWT token used to perform authorization")
			Attribute("company_id", Int, "Company ID", func() {
				Minimum(1)
			})
			Required("company_id")
		})

		HTTP(func() {
			DELETE("/{company_id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
			Response("Unauthorized", StatusUnauthorized)
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
