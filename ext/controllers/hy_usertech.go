package controllers

import (
	"fmt"
	"github.com/goadesign/goa"
	c "github.com/hiromaily/go-goa/ext/context"
	m "github.com/hiromaily/go-goa/ext/models"
	"github.com/hiromaily/go-goa/goa/app"
)

// HyUsertechController implements the hy_usertech resource.
type HyUsertechController struct {
	*goa.Controller
	ctx *c.Ctx //Added
}

// NewHyUsertechController creates a hy_usertech controller.
func NewHyUsertechController(service *goa.Service, ctx *c.Ctx) *HyUsertechController {
	return &HyUsertechController{
		Controller: service.NewController("HyUsertechController"),
		ctx:        ctx, //Added
	}
}

// GetUserLikeTech runs the GetUserLikeTech action.
func (c *HyUsertechController) GetUserLikeTech(ctx *app.GetUserLikeTechHyUsertechContext) error {
	fmt.Println("[hy_usertech][GetUserLikeTech]")

	const TableTechLike = "t_user_like_techs"

	//type UsertechTech struct {
	//	// Tech name
	//	TechName string `form:"tech_name" json:"tech_name" xml:"tech_name"`
	//}
	var userTechs []*app.UsertechTech

	svc := &m.UserTech{Db: c.ctx.Db}
	err := svc.GetUserTechs(ctx.UserID, &userTechs, TableTechLike)
	if err != nil {
		return err
	}

	if len(userTechs) == 0 {
		return ctx.NotFound()
	}

	//type UserCollection []*User
	res := app.UsertechTechCollection(userTechs)
	return ctx.OKTech(res)
}

// GetUserDislikeTech runs the GetUserDislikeTech action.
func (c *HyUsertechController) GetUserDislikeTech(ctx *app.GetUserDislikeTechHyUsertechContext) error {
	fmt.Println("[hy_usertech][GetUserDislikeTech]")

	const TableTechDislike = "t_user_dislike_techs"

	var userTechs []*app.UsertechTech

	svc := &m.UserTech{Db: c.ctx.Db}
	err := svc.GetUserTechs(ctx.UserID, &userTechs, TableTechDislike)
	if err != nil {
		return err
	}

	if len(userTechs) == 0 {
		return ctx.NotFound()
	}

	//type UserCollection []*User
	res := app.UsertechTechCollection(userTechs)
	return ctx.OKTech(res)
}
