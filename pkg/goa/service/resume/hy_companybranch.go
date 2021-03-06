package resumeapi

import (
	"context"
	"fmt"
	"log"
	hycompanybranch "resume/gen/hy_companybranch"

	"goa.design/goa/v3/security"
)

// hy_companybranch service example implementation.
// The example methods log the requests and return zero values.
type hyCompanybranchsrvc struct {
	logger *log.Logger
}

// NewHyCompanybranch returns the hy_companybranch service implementation.
func NewHyCompanybranch(logger *log.Logger) hycompanybranch.Service {
	return &hyCompanybranchsrvc{logger}
}

// JWTAuth implements the authorization logic for service "hy_companybranch"
// for the "jwt" security scheme.
func (s *hyCompanybranchsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
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

// Get company branch with given id
func (s *hyCompanybranchsrvc) GetCompanyBranch(ctx context.Context, p *hycompanybranch.GetCompanyBranchPayload) (err error) {
	//	//var companies []*app.Company
	//	company := &app.Company{}
	//
	//	svc := &m.Company{Db: c.ctx.Db}
	//	err := svc.GetCompanyBranch(ctx.ID, company)
	//	if err != nil {
	//		return err
	//	}
	//
	//	if company.ID == nil {
	//		return ctx.NotFound()
	//	}
	//
	//	//res := &app.Company{}
	//	return ctx.OK(company)

	s.logger.Print("hyCompanybranch.getCompanyBranch")
	return
}

// Create new company branch
func (s *hyCompanybranchsrvc) CreateCompanyBranch(ctx context.Context, p *hycompanybranch.CreateCompanyBranchPayload) (err error) {
	//	svc := &m.Company{Db: c.ctx.Db}
	//	ID, err := svc.InsertCompanyBranch(ctx.ID, ctx.Payload)
	//	if err != nil {
	//		return err
	//	}
	//
	//	res := &app.CompanyDetailid{ID: &ID}
	//	return ctx.OKDetailid(res)

	s.logger.Print("hyCompanybranch.createCompanyBranch")
	return
}

// Change company branch properties
func (s *hyCompanybranchsrvc) UpdateCompanyBranch(ctx context.Context, p *hycompanybranch.UpdateCompanyBranchPayload) (err error) {
	//	svc := &m.Company{Db: c.ctx.Db}
	//	err := svc.UpdateCompanyBranch(ctx.ID, ctx.Payload)
	//	if err != nil {
	//		return err
	//	}
	//
	//	res := &app.CompanyDetailid{ID: &ctx.ID}
	//	return ctx.OKDetailid(res)

	s.logger.Print("hyCompanybranch.updateCompanyBranch")
	return
}

// Delete company branch
func (s *hyCompanybranchsrvc) DeleteCompanyBranch(ctx context.Context, p *hycompanybranch.DeleteCompanyBranchPayload) (err error) {
	//	svc := &m.Company{Db: c.ctx.Db}
	//	err := svc.DeleteCompanyBranch(ctx.ID)
	//	if err != nil {
	//		return err
	//	}
	//
	//	res := &app.CompanyDetailid{ID: &ctx.ID}
	//	return ctx.OKDetailid(res)

	s.logger.Print("hyCompanybranch.deleteCompanyBranch")
	return
}
