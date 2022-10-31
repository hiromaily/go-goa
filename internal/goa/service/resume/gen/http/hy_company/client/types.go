// Code generated by goa v3.10.2, DO NOT EDIT.
//
// hy_company HTTP client types
//
// Command:
// $ goa gen resume/design

package client

import (
	hycompany "resume/gen/hy_company"
	hycompanyviews "resume/gen/hy_company/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// CreateCompanyRequestBody is the type of the "hy_company" service
// "createCompany" endpoint HTTP request body.
type CreateCompanyRequestBody struct {
	// Country ID
	CountryID int `form:"country_id" json:"country_id" xml:"country_id"`
	// Company name
	CompanyName string `form:"company_name" json:"company_name" xml:"company_name"`
	// Company Address
	Address string `form:"address" json:"address" xml:"address"`
}

// UpdateCompanyRequestBody is the type of the "hy_company" service
// "updateCompany" endpoint HTTP request body.
type UpdateCompanyRequestBody struct {
	// Country ID
	CountryID *int `form:"country_id,omitempty" json:"country_id,omitempty" xml:"country_id,omitempty"`
	// Company name
	CompanyName *string `form:"company_name,omitempty" json:"company_name,omitempty" xml:"company_name,omitempty"`
	// Company Address
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
}

// CompanyListResponseBody is the type of the "hy_company" service
// "companyList" endpoint HTTP response body.
type CompanyListResponseBody []*CompanyResponse

// GetCompanyResponseBody is the type of the "hy_company" service "getCompany"
// endpoint HTTP response body.
type GetCompanyResponseBody struct {
	// Key ID
	CompanyID *int `form:"company_id,omitempty" json:"company_id,omitempty" xml:"company_id,omitempty"`
	// Company name
	CompanyName *string `form:"company_name,omitempty" json:"company_name,omitempty" xml:"company_name,omitempty"`
	CountryName *string `form:"country_name,omitempty" json:"country_name,omitempty" xml:"country_name,omitempty"`
	// Company Address
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Datetime
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Datetime
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// CreateCompanyResponseBody is the type of the "hy_company" service
// "createCompany" endpoint HTTP response body.
type CreateCompanyResponseBody struct {
	// Key ID
	CompanyID *int `form:"company_id,omitempty" json:"company_id,omitempty" xml:"company_id,omitempty"`
	// Company name
	CompanyName *string `form:"company_name,omitempty" json:"company_name,omitempty" xml:"company_name,omitempty"`
	CountryName *string `form:"country_name,omitempty" json:"country_name,omitempty" xml:"country_name,omitempty"`
	// Company Address
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Datetime
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Datetime
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// CompanyListNotFoundResponseBody is the type of the "hy_company" service
// "companyList" endpoint HTTP response body for the "NotFound" error.
type CompanyListNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// GetCompanyNotFoundResponseBody is the type of the "hy_company" service
// "getCompany" endpoint HTTP response body for the "NotFound" error.
type GetCompanyNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// CreateCompanyBadRequestResponseBody is the type of the "hy_company" service
// "createCompany" endpoint HTTP response body for the "BadRequest" error.
type CreateCompanyBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdateCompanyBadRequestResponseBody is the type of the "hy_company" service
// "updateCompany" endpoint HTTP response body for the "BadRequest" error.
type UpdateCompanyBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdateCompanyNotFoundResponseBody is the type of the "hy_company" service
// "updateCompany" endpoint HTTP response body for the "NotFound" error.
type UpdateCompanyNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// DeleteCompanyNotFoundResponseBody is the type of the "hy_company" service
// "deleteCompany" endpoint HTTP response body for the "NotFound" error.
type DeleteCompanyNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// CompanyResponse is used to define fields on response body types.
type CompanyResponse struct {
	// Key ID
	CompanyID *int `form:"company_id,omitempty" json:"company_id,omitempty" xml:"company_id,omitempty"`
	// Company name
	CompanyName *string `form:"company_name,omitempty" json:"company_name,omitempty" xml:"company_name,omitempty"`
	CountryName *string `form:"country_name,omitempty" json:"country_name,omitempty" xml:"country_name,omitempty"`
	// Company Address
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Datetime
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Datetime
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// NewCreateCompanyRequestBody builds the HTTP request body from the payload of
// the "createCompany" endpoint of the "hy_company" service.
func NewCreateCompanyRequestBody(p *hycompany.CreateCompanyPayload) *CreateCompanyRequestBody {
	body := &CreateCompanyRequestBody{
		CountryID:   p.CountryID,
		CompanyName: p.CompanyName,
		Address:     p.Address,
	}
	return body
}

// NewUpdateCompanyRequestBody builds the HTTP request body from the payload of
// the "updateCompany" endpoint of the "hy_company" service.
func NewUpdateCompanyRequestBody(p *hycompany.UpdateCompanyPayload) *UpdateCompanyRequestBody {
	body := &UpdateCompanyRequestBody{
		CountryID:   p.CountryID,
		CompanyName: p.CompanyName,
		Address:     p.Address,
	}
	return body
}

// NewCompanyListCompanyCollectionOK builds a "hy_company" service
// "companyList" endpoint result from a HTTP "OK" response.
func NewCompanyListCompanyCollectionOK(body CompanyListResponseBody) hycompanyviews.CompanyCollectionView {
	v := make([]*hycompanyviews.CompanyView, len(body))
	for i, val := range body {
		v[i] = unmarshalCompanyResponseToHycompanyviewsCompanyView(val)
	}

	return v
}

// NewCompanyListNotFound builds a hy_company service companyList endpoint
// NotFound error.
func NewCompanyListNotFound(body *CompanyListNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewGetCompanyCompanyOK builds a "hy_company" service "getCompany" endpoint
// result from a HTTP "OK" response.
func NewGetCompanyCompanyOK(body *GetCompanyResponseBody) *hycompanyviews.CompanyView {
	v := &hycompanyviews.CompanyView{
		CompanyID:   body.CompanyID,
		CompanyName: body.CompanyName,
		CountryName: body.CountryName,
		Address:     body.Address,
		CreatedAt:   body.CreatedAt,
		UpdatedAt:   body.UpdatedAt,
	}

	return v
}

// NewGetCompanyNotFound builds a hy_company service getCompany endpoint
// NotFound error.
func NewGetCompanyNotFound(body *GetCompanyNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewCreateCompanyCompanyCreated builds a "hy_company" service "createCompany"
// endpoint result from a HTTP "Created" response.
func NewCreateCompanyCompanyCreated(body *CreateCompanyResponseBody) *hycompanyviews.CompanyView {
	v := &hycompanyviews.CompanyView{
		CompanyID:   body.CompanyID,
		CompanyName: body.CompanyName,
		CountryName: body.CountryName,
		Address:     body.Address,
		CreatedAt:   body.CreatedAt,
		UpdatedAt:   body.UpdatedAt,
	}

	return v
}

// NewCreateCompanyBadRequest builds a hy_company service createCompany
// endpoint BadRequest error.
func NewCreateCompanyBadRequest(body *CreateCompanyBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdateCompanyBadRequest builds a hy_company service updateCompany
// endpoint BadRequest error.
func NewUpdateCompanyBadRequest(body *UpdateCompanyBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdateCompanyNotFound builds a hy_company service updateCompany endpoint
// NotFound error.
func NewUpdateCompanyNotFound(body *UpdateCompanyNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewDeleteCompanyNotFound builds a hy_company service deleteCompany endpoint
// NotFound error.
func NewDeleteCompanyNotFound(body *DeleteCompanyNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateCompanyListNotFoundResponseBody runs the validations defined on
// companyList_NotFound_response_body
func ValidateCompanyListNotFoundResponseBody(body *CompanyListNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateGetCompanyNotFoundResponseBody runs the validations defined on
// getCompany_NotFound_response_body
func ValidateGetCompanyNotFoundResponseBody(body *GetCompanyNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateCreateCompanyBadRequestResponseBody runs the validations defined on
// createCompany_BadRequest_response_body
func ValidateCreateCompanyBadRequestResponseBody(body *CreateCompanyBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdateCompanyBadRequestResponseBody runs the validations defined on
// updateCompany_BadRequest_response_body
func ValidateUpdateCompanyBadRequestResponseBody(body *UpdateCompanyBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdateCompanyNotFoundResponseBody runs the validations defined on
// updateCompany_NotFound_response_body
func ValidateUpdateCompanyNotFoundResponseBody(body *UpdateCompanyNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateDeleteCompanyNotFoundResponseBody runs the validations defined on
// deleteCompany_NotFound_response_body
func ValidateDeleteCompanyNotFoundResponseBody(body *DeleteCompanyNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateCompanyResponse runs the validations defined on CompanyResponse
func ValidateCompanyResponse(body *CompanyResponse) (err error) {
	if body.CompanyID != nil {
		if *body.CompanyID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.company_id", *body.CompanyID, 1, true))
		}
	}
	if body.CompanyName != nil {
		if utf8.RuneCountInString(*body.CompanyName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.company_name", *body.CompanyName, utf8.RuneCountInString(*body.CompanyName), 2, true))
		}
	}
	if body.CompanyName != nil {
		if utf8.RuneCountInString(*body.CompanyName) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.company_name", *body.CompanyName, utf8.RuneCountInString(*body.CompanyName), 40, false))
		}
	}
	if body.Address != nil {
		if utf8.RuneCountInString(*body.Address) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.address", *body.Address, utf8.RuneCountInString(*body.Address), 2, true))
		}
	}
	if body.Address != nil {
		if utf8.RuneCountInString(*body.Address) > 80 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.address", *body.Address, utf8.RuneCountInString(*body.Address), 80, false))
		}
	}
	if body.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.created_at", *body.CreatedAt, goa.FormatDateTime))
	}
	if body.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.updated_at", *body.UpdatedAt, goa.FormatDateTime))
	}
	return
}
