package controllers

//
//import (
//	"goa.design/goa/v3"
//	c "github.com/hiromaily/go-goa/ext/context"
//	m "github.com/hiromaily/go-goa/ext/models"
//	"github.com/hiromaily/go-goa/goa/app"
//)
//
//// HyTechController implements the hy_tech resource.
//type HyTechController struct {
//	*goa.Controller
//	ctx *c.Ctx //Added
//}
//
//// NewHyTechController creates a hy_tech controller.
//func NewHyTechController(service *goa.Service, ctx *c.Ctx) *HyTechController {
//	return &HyTechController{
//		Controller: service.NewController("HyTechController"),
//		ctx:        ctx, //Added
//	}
//}
//
//// GetTech runs the GetTech action.
//func (c *HyTechController) GetTech(ctx *app.GetTechHyTechContext) error {
//	tech := &app.Tech{}
//
//	svc := &m.Tech{Db: c.ctx.Db}
//	err := svc.GetTech(ctx.TechID, tech)
//	if err != nil {
//		return err
//	}
//
//	if tech.ID == nil {
//		//404
//		return ctx.NotFound()
//	}
//
//	return ctx.OK(tech)
//}
//
//// TechList runs the TechList action.
//func (c *HyTechController) TechList(ctx *app.TechListHyTechContext) error {
//	var techs []*app.Tech
//
//	svc := &m.Tech{Db: c.ctx.Db}
//	err := svc.TechList(&techs)
//	if err != nil {
//		return err
//	}
//
//	if len(techs) == 0 {
//		return ctx.NoContent()
//	}
//
//	res := app.TechCollection(techs)
//	return ctx.OK(res)
//}
//
//// CreateTech runs the CreateTech action.
//func (c *HyTechController) CreateTech(ctx *app.CreateTechHyTechContext) error {
//	svc := &m.Tech{Db: c.ctx.Db}
//	techID, err := svc.InsertTech(ctx.Payload)
//	if err != nil {
//		return err
//	}
//
//	res := &app.TechID{ID: &techID}
//	return ctx.OKId(res)
//}
//
//// UpdateTech runs the UpdateTech action.
//func (c *HyTechController) UpdateTech(ctx *app.UpdateTechHyTechContext) error {
//	svc := &m.Tech{Db: c.ctx.Db}
//	err := svc.UpdateTech(ctx.TechID, ctx.Payload)
//	if err != nil {
//		return err
//	}
//
//	res := &app.TechID{ID: &ctx.TechID}
//	return ctx.OKId(res)
//}
//
//// DeleteTech runs the DeleteTech action.
//func (c *HyTechController) DeleteTech(ctx *app.DeleteTechHyTechContext) error {
//	svc := &m.Tech{Db: c.ctx.Db}
//	err := svc.DeleteTech(ctx.TechID)
//	if err != nil {
//		return err
//	}
//
//	res := &app.TechID{ID: &ctx.TechID}
//	return ctx.OKId(res)
//}
