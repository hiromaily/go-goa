package controllers

import (
	"github.com/goadesign/goa"
	"github.com/hiromaily/go-goa/goa/app"
	c "github.com/hiromaily/go-goa/ext/context"
	//m "github.com/hiromaily/go-goa/ext/models"
)

// HyTechController implements the hy_tech resource.
type HyTechController struct {
	*goa.Controller
	ctx *c.Ctx //Added
}

// NewHyTechController creates a hy_tech controller.
func NewHyTechController(service *goa.Service, ctx *c.Ctx) *HyTechController {
	return &HyTechController{
		Controller: service.NewController("HyTechController"),
		ctx:        ctx, //Added
	}
}

// CreateTech runs the CreateTech action.
func (c *HyTechController) CreateTech(ctx *app.CreateTechHyTechContext) error {
	// HyTechController_CreateTech: start_implement

	// Put your logic here

	// HyTechController_CreateTech: end_implement
	res := &app.Tech{}
	return ctx.OK(res)
}

// DeleteTech runs the DeleteTech action.
func (c *HyTechController) DeleteTech(ctx *app.DeleteTechHyTechContext) error {
	// HyTechController_DeleteTech: start_implement

	// Put your logic here

	// HyTechController_DeleteTech: end_implement
	res := &app.Tech{}
	return ctx.OK(res)
}

// GetTech runs the GetTech action.
func (c *HyTechController) GetTech(ctx *app.GetTechHyTechContext) error {
	// HyTechController_GetTech: start_implement

	// Put your logic here

	// HyTechController_GetTech: end_implement
	res := &app.Tech{}
	return ctx.OK(res)
}

// TechList runs the TechList action.
func (c *HyTechController) TechList(ctx *app.TechListHyTechContext) error {
	// HyTechController_TechList: start_implement

	// Put your logic here

	// HyTechController_TechList: end_implement
	res := app.TechCollection{}
	return ctx.OK(res)
}

// UpdateTech runs the UpdateTech action.
func (c *HyTechController) UpdateTech(ctx *app.UpdateTechHyTechContext) error {
	// HyTechController_UpdateTech: start_implement

	// Put your logic here

	// HyTechController_UpdateTech: end_implement
	res := &app.Tech{}
	return ctx.OK(res)
}
