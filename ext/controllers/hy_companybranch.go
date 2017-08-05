package controllers

import (
	"fmt"
	"github.com/goadesign/goa"
	c "github.com/hiromaily/go-goa/ext/context"
	m "github.com/hiromaily/go-goa/ext/models"
	"github.com/hiromaily/go-goa/goa/app"
)

// HyCompanybranchController implements the hy_companybranch resource.
type HyCompanybranchController struct {
	*goa.Controller
	ctx *c.Ctx
}

// NewHyCompanybranchController creates a hy_companybranch controller.
func NewHyCompanybranchController(service *goa.Service, ctx *c.Ctx) *HyCompanybranchController {
	return &HyCompanybranchController{
		Controller: service.NewController("HyCompanybranchController"),
		ctx:        ctx,
	}
}

// GetCompanyBranch runs the GetCompanyBranch action.
func (c *HyCompanybranchController) GetCompanyBranch(ctx *app.GetCompanyBranchHyCompanybranchContext) error {
	fmt.Println("[hy_companybranch][GetCompanyBranch]")

	//var companies []*app.Company
	company := &app.Company{}

	svc := &m.Company{Db: c.ctx.Db}
	err := svc.GetCompanyBranch(ctx.ID, company)
	if err != nil {
		return err
	}

	if company.ID == nil {
		return ctx.NotFound()
	}

	//res := &app.Company{}
	return ctx.OK(company)
}

// CreateCompanyBranch runs the CreateCompanyBranch action.
func (c *HyCompanybranchController) CreateCompanyBranch(ctx *app.CreateCompanyBranchHyCompanybranchContext) error {
	fmt.Println("[hy_companybranch][CreateCompanyBranch]")

	svc := &m.Company{Db: c.ctx.Db}
	ID, err := svc.InsertCompanyBranch(ctx.ID, ctx.Payload)
	if err != nil {
		return err
	}

	res := &app.CompanyDetailid{ID: &ID}
	return ctx.OKDetailid(res)
}

// UpdateCompanyBranch runs the UpdateCompanyBranch action.
func (c *HyCompanybranchController) UpdateCompanyBranch(ctx *app.UpdateCompanyBranchHyCompanybranchContext) error {
	// HyCompanybranchController_UpdateCompanyBranch: start_implement

	// Put your logic here

	// HyCompanybranchController_UpdateCompanyBranch: end_implement
	res := &app.Company{}
	return ctx.OK(res)
}

// DeleteCompanyBranch runs the DeleteCompanyBranch action.
func (c *HyCompanybranchController) DeleteCompanyBranch(ctx *app.DeleteCompanyBranchHyCompanybranchContext) error {
	// HyCompanybranchController_DeleteCompanyBranch: start_implement

	// Put your logic here

	// HyCompanybranchController_DeleteCompanyBranch: end_implement
	res := &app.Company{}
	return ctx.OK(res)
}
