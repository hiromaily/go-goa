// Code generated by goa v3.2.6, DO NOT EDIT.
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

// GetCompanyGroupRequestBody is the type of the "hy_company" service
// "getCompanyGroup" endpoint HTTP request body.
type GetCompanyGroupRequestBody struct {
	// Head Quarters flag
	IsHq *string `form:"is_hq,omitempty" json:"is_hq,omitempty" xml:"is_hq,omitempty"`
}

// CreateCompanyRequestBody is the type of the "hy_company" service
// "createCompany" endpoint HTTP request body.
type CreateCompanyRequestBody struct {
	// Company name
	Name string `form:"name" json:"name" xml:"name"`
	// Country ID
	CountryID int `form:"country_id" json:"country_id" xml:"country_id"`
	// Company Address
	Address string `form:"address" json:"address" xml:"address"`
}

// UpdateCompanyRequestBody is the type of the "hy_company" service
// "updateCompany" endpoint HTTP request body.
type UpdateCompanyRequestBody struct {
	// Company name
	Name string `form:"name" json:"name" xml:"name"`
	// Country ID
	CountryID int `form:"country_id" json:"country_id" xml:"country_id"`
	// Company Address
	Address string `form:"address" json:"address" xml:"address"`
}

// CompanyListResponseBody is the type of the "hy_company" service
// "companyList" endpoint HTTP response body.
type CompanyListResponseBody []*CompanyResponse

// GetCompanyGroupResponseBody is the type of the "hy_company" service
// "getCompanyGroup" endpoint HTTP response body.
type GetCompanyGroupResponseBody []*CompanyResponse

// CompanyResponse is used to define fields on response body types.
type CompanyResponse struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// ID
	CompanyID *int `form:"company_id,omitempty" json:"company_id,omitempty" xml:"company_id,omitempty"`
	// Company name
	Name        *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	IsHq        *string `form:"is_hq,omitempty" json:"is_hq,omitempty" xml:"is_hq,omitempty"`
	CountryName *string `form:"country_name,omitempty" json:"country_name,omitempty" xml:"country_name,omitempty"`
	// Company Address
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// Datetime
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Datetime
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// NewGetCompanyGroupRequestBody builds the HTTP request body from the payload
// of the "getCompanyGroup" endpoint of the "hy_company" service.
func NewGetCompanyGroupRequestBody(p *hycompany.GetCompanyGroupPayload) *GetCompanyGroupRequestBody {
	body := &GetCompanyGroupRequestBody{
		IsHq: p.IsHq,
	}
	return body
}

// NewCreateCompanyRequestBody builds the HTTP request body from the payload of
// the "createCompany" endpoint of the "hy_company" service.
func NewCreateCompanyRequestBody(p *hycompany.CreateCompanyPayload) *CreateCompanyRequestBody {
	body := &CreateCompanyRequestBody{
		Name:      p.Name,
		CountryID: p.CountryID,
		Address:   p.Address,
	}
	return body
}

// NewUpdateCompanyRequestBody builds the HTTP request body from the payload of
// the "updateCompany" endpoint of the "hy_company" service.
func NewUpdateCompanyRequestBody(p *hycompany.UpdateCompanyPayload) *UpdateCompanyRequestBody {
	body := &UpdateCompanyRequestBody{
		Name:      p.Name,
		CountryID: p.CountryID,
		Address:   p.Address,
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

// NewGetCompanyGroupCompanyCollectionOK builds a "hy_company" service
// "getCompanyGroup" endpoint result from a HTTP "OK" response.
func NewGetCompanyGroupCompanyCollectionOK(body GetCompanyGroupResponseBody) hycompanyviews.CompanyCollectionView {
	v := make([]*hycompanyviews.CompanyView, len(body))
	for i, val := range body {
		v[i] = unmarshalCompanyResponseToHycompanyviewsCompanyView(val)
	}
	return v
}

// ValidateCompanyResponse runs the validations defined on CompanyResponse
func ValidateCompanyResponse(body *CompanyResponse) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
	}
	if body.ID != nil {
		if *body.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.id", *body.ID, 1, true))
		}
	}
	if body.CompanyID != nil {
		if *body.CompanyID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.company_id", *body.CompanyID, 1, true))
		}
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 2, true))
		}
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 40, false))
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
