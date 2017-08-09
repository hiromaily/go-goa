package goa

import (
	"github.com/goadesign/goa"
	"github.com/hiromaily/go-goa/goa/app"
)

// HyUserWorkHistoryController implements the hy_userWorkHistory resource.
type HyUserWorkHistoryController struct {
	*goa.Controller
}

// NewHyUserWorkHistoryController creates a hy_userWorkHistory controller.
func NewHyUserWorkHistoryController(service *goa.Service) *HyUserWorkHistoryController {
	return &HyUserWorkHistoryController{Controller: service.NewController("HyUserWorkHistoryController")}
}

// GetUserWorkHistory runs the GetUserWorkHistory action.
func (c *HyUserWorkHistoryController) GetUserWorkHistory(ctx *app.GetUserWorkHistoryHyUserWorkHistoryContext) error {
	// HyUserWorkHistoryController_GetUserWorkHistory: start_implement

	// Put your logic here

	// HyUserWorkHistoryController_GetUserWorkHistory: end_implement
	res := app.UserworkhistoryCollection{}
	return ctx.OK(res)
}
