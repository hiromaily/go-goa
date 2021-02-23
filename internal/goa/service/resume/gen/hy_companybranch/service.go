// Code generated by goa v3.2.6, DO NOT EDIT.
//
// hy_companybranch service
//
// Command:
// $ goa gen resume/design

package hycompanybranch

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// The company branch service returns company branch data
type Service interface {
	// Get company branch with given id
	GetCompanyBranch(context.Context, *GetCompanyBranchPayload) (err error)
	// Create new company branch
	CreateCompanyBranch(context.Context, *CreateCompanyBranchPayload) (err error)
	// Change company branch properties
	UpdateCompanyBranch(context.Context, *UpdateCompanyBranchPayload) (err error)
	// Delete company branch
	DeleteCompanyBranch(context.Context, *DeleteCompanyBranchPayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "hy_companybranch"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"getCompanyBranch", "createCompanyBranch", "updateCompanyBranch", "deleteCompanyBranch"}

// GetCompanyBranchPayload is the payload type of the hy_companybranch service
// getCompanyBranch method.
type GetCompanyBranchPayload struct {
	// JWT token used to perform authorization
	Token *string
	// Company detail ID
	CompanyDetailID *int
}

// CreateCompanyBranchPayload is the payload type of the hy_companybranch
// service createCompanyBranch method.
type CreateCompanyBranchPayload struct {
	// JWT token used to perform authorization
	Token *string
	// Company ID
	CompanyID *int
	// Country ID
	CountryID int
	// Company Address
	Address string
}

// UpdateCompanyBranchPayload is the payload type of the hy_companybranch
// service updateCompanyBranch method.
type UpdateCompanyBranchPayload struct {
	// JWT token used to perform authorization
	Token *string
	// Company detail ID
	CompanyDetailID *int
	// Country ID
	CountryID int
	// Company Address
	Address string
}

// DeleteCompanyBranchPayload is the payload type of the hy_companybranch
// service deleteCompanyBranch method.
type DeleteCompanyBranchPayload struct {
	// JWT token used to perform authorization
	Token *string
	// Company detail ID
	CompanyDetailID *int
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "NotFound",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "BadRequest",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
