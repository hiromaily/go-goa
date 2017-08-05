package goa

import (
	"github.com/goadesign/goa"
	"github.com/hiromaily/go-goa/goa/app"
)

// HyCompanyController implements the hy_company resource.
type HyCompanyController struct {
	*goa.Controller
}

// NewHyCompanyController creates a hy_company controller.
func NewHyCompanyController(service *goa.Service) *HyCompanyController {
	return &HyCompanyController{Controller: service.NewController("HyCompanyController")}
}

// CompanyList runs the CompanyList action.
func (c *HyCompanyController) CompanyList(ctx *app.CompanyListHyCompanyContext) error {
	// HyCompanyController_CompanyList: start_implement

	// Put your logic here

	// HyCompanyController_CompanyList: end_implement
	res := app.CompanyCollection{}
	return ctx.OK(res)
}

// CreateCompany runs the CreateCompany action.
func (c *HyCompanyController) CreateCompany(ctx *app.CreateCompanyHyCompanyContext) error {
	// HyCompanyController_CreateCompany: start_implement

	// Put your logic here

	// HyCompanyController_CreateCompany: end_implement
	res := &app.Company{}
	return ctx.OK(res)
}

// CreateCompanyBranch runs the CreateCompanyBranch action.
func (c *HyCompanyController) CreateCompanyBranch(ctx *app.CreateCompanyBranchHyCompanyContext) error {
	// HyCompanyController_CreateCompanyBranch: start_implement

	// Put your logic here

	// HyCompanyController_CreateCompanyBranch: end_implement
	res := &app.Company{}
	return ctx.OK(res)
}

// DeleteCompany runs the DeleteCompany action.
func (c *HyCompanyController) DeleteCompany(ctx *app.DeleteCompanyHyCompanyContext) error {
	// HyCompanyController_DeleteCompany: start_implement

	// Put your logic here

	// HyCompanyController_DeleteCompany: end_implement
	res := &app.Company{}
	return ctx.OK(res)
}

// GetCompanyBranch runs the GetCompanyBranch action.
func (c *HyCompanyController) GetCompanyBranch(ctx *app.GetCompanyBranchHyCompanyContext) error {
	// HyCompanyController_GetCompanyBranch: start_implement

	// Put your logic here

	// HyCompanyController_GetCompanyBranch: end_implement
	res := &app.Company{}
	return ctx.OK(res)
}

// GetCompanyGroup runs the GetCompanyGroup action.
func (c *HyCompanyController) GetCompanyGroup(ctx *app.GetCompanyGroupHyCompanyContext) error {
	// HyCompanyController_GetCompanyGroup: start_implement

	// Put your logic here

	// HyCompanyController_GetCompanyGroup: end_implement
	res := app.CompanyCollection{}
	return ctx.OK(res)
}

// UpdateCompany runs the UpdateCompany action.
func (c *HyCompanyController) UpdateCompany(ctx *app.UpdateCompanyHyCompanyContext) error {
	// HyCompanyController_UpdateCompany: start_implement

	// Put your logic here

	// HyCompanyController_UpdateCompany: end_implement
	res := &app.Company{}
	return ctx.OK(res)
}
