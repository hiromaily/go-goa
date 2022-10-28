package resumeapi

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"goa.design/goa/v3/security"

	"github.com/hiromaily/go-goa/pkg/jwts"
	ptr "github.com/hiromaily/go-goa/pkg/pointer"
	"github.com/hiromaily/go-goa/pkg/repository"
	hycompany "resume/gen/hy_company"
)

// hy_company service example implementation.
// The example methods log the requests and return zero values.
type hyCompanysrvc struct {
	jwt         jwts.JWTer
	companyRepo repository.CompanyRepository
}

// NewHyCompany returns the hy_company service implementation.
func NewHyCompany(jwt jwts.JWTer, companyRepo repository.CompanyRepository) hycompany.Service {
	return &hyCompanysrvc{jwt, companyRepo}
}

// JWTAuth implements the authorization logic for service "hy_company" for the
// "jwt" security scheme.
func (s *hyCompanysrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	log.Info().
		Str("token", token).
		Strs("scheme.Scopes", scheme.Scopes).
		Msg("hycompany.JWTAuth")

	if err := s.jwt.ValidateToken(token); err != nil {
		return ctx, err
	}

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
	return ctx, nil
}

// CompanyList returns all companies
func (s *hyCompanysrvc) CompanyList(ctx context.Context, p *hycompany.CompanyListPayload) (res hycompany.CompanyCollection, view string, err error) {
	log.Info().Msg("hyCompany.companyList")

	companies, err := s.companyRepo.CompanyList()

	if err != nil {
		return nil, "", errors.Wrap(err, "fail to call companyRepo.UserList()")
	}
	if len(companies) == 0 {
		return nil, "", hycompany.MakeNotFound(errors.New("company not found"))
	}
	res = companies
	view = "default"

	return
}

// GetCompany returns company by given company_id
func (s *hyCompanysrvc) GetCompany(ctx context.Context, p *hycompany.GetCompanyPayload) (res *hycompany.Company, view string, err error) {
	log.Info().Msg("hyCompany.getCompany")

	company, err := s.companyRepo.GetCompany(p.CompanyID)
	if err != nil {
		// Not Found
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, "", hycompany.MakeNotFound(errors.New("company not found"))
		}
		return nil, "", errors.Wrapf(err, "fail to call companyRepo.GetCompany(%d)", p.CompanyID)
	}
	res = company
	view = "default"

	return
}

// CreateCompany creates new company
// FIXME: only admin user can create user
func (s *hyCompanysrvc) CreateCompany(ctx context.Context, p *hycompany.CreateCompanyPayload) (res *hycompany.Company, view string, err error) {
	log.Info().Msg("hyCompany.createCompany")

	companyID, err := s.companyRepo.InsertCompany(p.CompanyName, p.Address, int16(p.CountryID))
	if err != nil {
		return nil, "", errors.Wrap(err, "fail to call InsertCompany()")
	}
	res = &hycompany.Company{
		CompanyID: ptr.Int(companyID),
	}
	view = "id"

	return
}

// UpdateCompany updates company
// FIXME: only login user can update own information
func (s *hyCompanysrvc) UpdateCompany(ctx context.Context, p *hycompany.UpdateCompanyPayload) (err error) {
	log.Info().Str("CompanyName", ptr.StringVal(p.CompanyName)).Str("Address", ptr.StringVal(p.Address)).Int("CountryID", ptr.IntVal(p.CountryID)).Msg("hyCompany.updateCompany")

	// Validate
	if p.CompanyName == nil && p.Address == nil && p.CountryID == nil {
		return hycompany.MakeBadRequest(errors.New("parameter is invalid"))
	}

	isCompany, err := s.companyRepo.IsCompany(p.CompanyID)
	if !isCompany || err != nil {
		return hycompany.MakeNotFound(errors.New("company not found"))
	}

	_, err = s.companyRepo.UpdateCompany(p.CompanyID, ptr.StringVal(p.CompanyName), ptr.StringVal(p.Address), int16(ptr.IntVal(p.CountryID)))
	if err != nil {
		return errors.Wrap(err, "fail to call UpdateCompany()")
	}

	return
}

// DeleteCompany deletes company
// FIXME: only admin user can delete user
func (s *hyCompanysrvc) DeleteCompany(ctx context.Context, p *hycompany.DeleteCompanyPayload) (err error) {
	log.Info().Msg("hyCompany.deleteCompany")

	isCompany, err := s.companyRepo.IsCompany(p.CompanyID)
	if !isCompany || err != nil {
		return hycompany.MakeNotFound(errors.New("company not found"))
	}

	if _, err := s.companyRepo.DeleteCompany(p.CompanyID); err != nil {
		return errors.Wrap(err, "fail to call DeleteCompany()")
	}

	return
}
