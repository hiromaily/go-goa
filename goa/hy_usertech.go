package goa

import (
	"github.com/goadesign/goa"
	"github.com/hiromaily/go-goa/goa/app"
)

// HyUsertechController implements the hy_usertech resource.
type HyUsertechController struct {
	*goa.Controller
}

// NewHyUsertechController creates a hy_usertech controller.
func NewHyUsertechController(service *goa.Service) *HyUsertechController {
	return &HyUsertechController{Controller: service.NewController("HyUsertechController")}
}

// GetUserDislikeTech runs the GetUserDislikeTech action.
func (c *HyUsertechController) GetUserDislikeTech(ctx *app.GetUserDislikeTechHyUsertechContext) error {
	// HyUsertechController_GetUserDislikeTech: start_implement

	// Put your logic here

	res := app.UsertechCollection{}
	return ctx.OK(res)
	// HyUsertechController_GetUserDislikeTech: end_implement
}

// GetUserLikeTech runs the GetUserLikeTech action.
func (c *HyUsertechController) GetUserLikeTech(ctx *app.GetUserLikeTechHyUsertechContext) error {
	// HyUsertechController_GetUserLikeTech: start_implement

	// Put your logic here

	res := app.UsertechCollection{}
	return ctx.OK(res)
	// HyUsertechController_GetUserLikeTech: end_implement
}
