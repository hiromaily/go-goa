package resumeapi

import (
	"context"
	"fmt"
	"log"
	hycompany "resume/gen/hy_company"

	"goa.design/goa/v3/security"
)

// hy_company service example implementation.
// The example methods log the requests and return zero values.
type hyCompanysrvc struct {
	logger *log.Logger
}

// NewHyCompany returns the hy_company service implementation.
func NewHyCompany(logger *log.Logger) hycompany.Service {
	return &hyCompanysrvc{logger}
}

// JWTAuth implements the authorization logic for service "hy_company" for the
// "jwt" security scheme.
func (s *hyCompanysrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
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

// List all companies
func (s *hyCompanysrvc) CompanyList(ctx context.Context, p *hycompany.CompanyListPayload) (res hycompany.CompanyCollection, view string, err error) {
	//	var companies []*app.CompanyIdname
	//
	//	svc := &m.Company{Db: c.ctx.Db}
	//	err := svc.CompanyList(&companies)
	//	if err != nil {
	//		return err
	//	}
	//
	//	if len(companies) == 0 {
	//		return ctx.NoContent()
	//	}
	//
	//	res := app.CompanyIdnameCollection(companies)
	//	return ctx.OKIdname(res)

	view = "idname"
	s.logger.Print("hyCompany.companyList")
	return
}

// Retrieve company with given company_id
func (s *hyCompanysrvc) GetCompanyGroup(ctx context.Context, p *hycompany.GetCompanyGroupPayload) (res hycompany.CompanyCollection, view string, err error) {
	//	var companies []*app.CompanyIdname
	//
	//	svc := &m.Company{Db: c.ctx.Db}
	//	err := svc.CompanyList(&companies)
	//	if err != nil {
	//		return err
	//	}
	//
	//	if len(companies) == 0 {
	//		return ctx.NoContent()
	//	}
	//
	//	res := app.CompanyIdnameCollection(companies)
	//	return ctx.OKIdname(res)

	view = "default"
	s.logger.Print("hyCompany.getCompanyGroup")
	return
}

// Create new company
func (s *hyCompanysrvc) CreateCompany(ctx context.Context, p *hycompany.CreateCompanyPayload) (err error) {
	//	svc := &m.Company{Db: c.ctx.Db}
	//	companyID, err := svc.InsertCompany(ctx.Payload)
	//	if err != nil {
	//		return err
	//	}
	//
	//	res := &app.CompanyID{CompanyID: &companyID}
	//	return ctx.OKId(res)

	s.logger.Print("hyCompany.createCompany")
	return
}

// Change company properties
func (s *hyCompanysrvc) UpdateCompany(ctx context.Context, p *hycompany.UpdateCompanyPayload) (err error) {
	//	svc := &m.Company{Db: c.ctx.Db}
	//	err := svc.UpdateCompany(ctx.CompanyID, ctx.Payload)
	//	if err != nil {
	//		return err
	//	}
	//
	//	res := &app.CompanyID{CompanyID: &ctx.CompanyID}
	//	return ctx.OKId(res)

	s.logger.Print("hyCompany.updateCompany")
	return
}

// Delete company
func (s *hyCompanysrvc) DeleteCompany(ctx context.Context, p *hycompany.DeleteCompanyPayload) (err error) {
	//	svc := &m.Company{Db: c.ctx.Db}
	//	err := svc.DeleteCompany(ctx.CompanyID)
	//	if err != nil {
	//		return err
	//	}
	//
	//	res := &app.CompanyID{CompanyID: &ctx.CompanyID}
	//	return ctx.OKId(res)

	s.logger.Print("hyCompany.deleteCompany")
	return
}
