package controllers

import (
	"fmt"
	"github.com/goadesign/goa"
	c "github.com/hiromaily/go-goa/ext/context"
	m "github.com/hiromaily/go-goa/ext/models"
	"github.com/hiromaily/go-goa/goa/app"
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

	//type Userworkhistory struct {
	//	// Job Title
	//	Title string `form:"title" json:"title" xml:"title"`
	//	// Company name
	//	Company string `form:"company" json:"company" xml:"company"`
	//	// Country code
	//	Country string `form:"country" json:"country" xml:"country"`
	//	// worked period
	//	Term *string `form:"term,omitempty" json:"term,omitempty" xml:"term,omitempty"`
	//	// job description
	//	Description *interface{} `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	//	// used techs
	//	Techs *interface{} `form:"techs,omitempty" json:"techs,omitempty" xml:"techs,omitempty"`
	//}
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
