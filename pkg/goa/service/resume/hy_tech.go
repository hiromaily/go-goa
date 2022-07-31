package resumeapi

import (
	"context"
	"fmt"
	"log"
	hytech "resume/gen/hy_tech"

	"goa.design/goa/v3/security"
)

// hy_tech service example implementation.
// The example methods log the requests and return zero values.
type hyTechsrvc struct {
	logger *log.Logger
}

// NewHyTech returns the hy_tech service implementation.
func NewHyTech(logger *log.Logger) hytech.Service {
	return &hyTechsrvc{logger}
}

// JWTAuth implements the authorization logic for service "hy_tech" for the
// "jwt" security scheme.
func (s *hyTechsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}

// List all techs
func (s *hyTechsrvc) TechList(ctx context.Context, p *hytech.TechListPayload) (res hytech.TechCollection, view string, err error) {
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
	view = "default"
	s.logger.Print("hyTech.techList")
	return
}

// get tech with given tech id
func (s *hyTechsrvc) GetTech(ctx context.Context, p *hytech.GetTechPayload) (res *hytech.Company, view string, err error) {
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

	res = &hytech.Company{}
	view = "default"
	s.logger.Print("hyTech.getTech")
	return
}

// Create new tech
func (s *hyTechsrvc) CreateTech(ctx context.Context, p *hytech.CreateTechPayload) (err error) {
	//	svc := &m.Tech{Db: c.ctx.Db}
	//	techID, err := svc.InsertTech(ctx.Payload)
	//	if err != nil {
	//		return err
	//	}
	//
	//	res := &app.TechID{ID: &techID}
	//	return ctx.OKId(res)

	s.logger.Print("hyTech.createTech")
	return
}

// Change tech properties
func (s *hyTechsrvc) UpdateTech(ctx context.Context, p *hytech.UpdateTechPayload) (err error) {
	//	svc := &m.Tech{Db: c.ctx.Db}
	//	err := svc.UpdateTech(ctx.TechID, ctx.Payload)
	//	if err != nil {
	//		return err
	//	}
	//
	//	res := &app.TechID{ID: &ctx.TechID}
	//	return ctx.OKId(res)

	s.logger.Print("hyTech.updateTech")
	return
}

// Delete tech
func (s *hyTechsrvc) DeleteTech(ctx context.Context, p *hytech.DeleteTechPayload) (err error) {
	//	svc := &m.Tech{Db: c.ctx.Db}
	//	err := svc.DeleteTech(ctx.TechID)
	//	if err != nil {
	//		return err
	//	}
	//
	//	res := &app.TechID{ID: &ctx.TechID}
	//	return ctx.OKId(res)

	s.logger.Print("hyTech.deleteTech")
	return
}
