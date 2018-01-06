package goa

import (
	"github.com/goadesign/goa"
	"github.com/hiromaily/go-goa/goa/app"
)

// HyUserController implements the hy_user resource.
type HyUserController struct {
	*goa.Controller
}

// NewHyUserController creates a hy_user controller.
func NewHyUserController(service *goa.Service) *HyUserController {
	return &HyUserController{Controller: service.NewController("HyUserController")}
}

// CreateUser runs the CreateUser action.
func (c *HyUserController) CreateUser(ctx *app.CreateUserHyUserContext) error {
	// HyUserController_CreateUser: start_implement

	// Put your logic here

	res := &app.User{}
	return ctx.OK(res)
	// HyUserController_CreateUser: end_implement
}

// DeleteUser runs the DeleteUser action.
func (c *HyUserController) DeleteUser(ctx *app.DeleteUserHyUserContext) error {
	// HyUserController_DeleteUser: start_implement

	// Put your logic here

	res := &app.User{}
	return ctx.OK(res)
	// HyUserController_DeleteUser: end_implement
}

// GetUser runs the GetUser action.
func (c *HyUserController) GetUser(ctx *app.GetUserHyUserContext) error {
	// HyUserController_GetUser: start_implement

	// Put your logic here

	res := &app.User{}
	return ctx.OK(res)
	// HyUserController_GetUser: end_implement
}

// UpdateUser runs the UpdateUser action.
func (c *HyUserController) UpdateUser(ctx *app.UpdateUserHyUserContext) error {
	// HyUserController_UpdateUser: start_implement

	// Put your logic here

	res := &app.User{}
	return ctx.OK(res)
	// HyUserController_UpdateUser: end_implement
}

// UserList runs the UserList action.
func (c *HyUserController) UserList(ctx *app.UserListHyUserContext) error {
	// HyUserController_UserList: start_implement

	// Put your logic here

	res := app.UserCollection{}
	return ctx.OK(res)
	// HyUserController_UserList: end_implement
}
