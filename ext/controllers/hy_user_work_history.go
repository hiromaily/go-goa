package controllers

import (
	"fmt"
	"github.com/goadesign/goa"
	c "github.com/hiromaily/go-goa/ext/context"
	m "github.com/hiromaily/go-goa/ext/models"
	"github.com/hiromaily/go-goa/goa/app"
	u "github.com/hiromaily/golibs/utils"
)

// HyUserWorkHistoryController implements the hy_userWorkHistory resource.
type HyUserWorkHistoryController struct {
	*goa.Controller
	ctx *c.Ctx //Added
}

// NewHyUserWorkHistoryController creates a hy_userWorkHistory controller.
func NewHyUserWorkHistoryController(service *goa.Service, ctx *c.Ctx) *HyUserWorkHistoryController {
	return &HyUserWorkHistoryController{
		Controller: service.NewController("HyUserWorkHistoryController"),
		ctx:        ctx, //Added
	}
}

// GetUserWorkHistory runs the GetUserWorkHistory action.
func (c *HyUserWorkHistoryController) GetUserWorkHistory(ctx *app.GetUserWorkHistoryHyUserWorkHistoryContext) error {
	fmt.Println("[hy_user_work_history][GetUserWorkHistory]")

	//TODO:check user ID and this part should be set in middleware.
	//fmt.Println("[user id]:", ctx.Value("user.jwt"))
	if ctx.Value("user.jwt") == nil || u.Itoi(ctx.Value("user.jwt")) == 0 ||
		u.Itoi(ctx.Value("user.jwt")) != ctx.UserID {
		return ctx.Unauthorized()
	}

	var userWorks []*app.Userworkhistory

	svc := &m.UserWorkHistory{Db: c.ctx.Db}
	err := svc.GetUserWorks(ctx.UserID, &userWorks)
	if err != nil {
		return err
	}

	if len(userWorks) == 0 {
		return ctx.NotFound()
	}

	res := app.UserworkhistoryCollection(userWorks)
	return ctx.OK(res)
}
