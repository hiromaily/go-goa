// Code generated by goa v3.10.1, DO NOT EDIT.
//
// hy_company HTTP server types
//
// Command:
// $ goa gen resume/design

package server

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
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Country ID
	CountryID *int `form:"country_id,omitempty" json:"country_id,omitempty" xml:"country_id,omitempty"`
	// Company Address
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
}

// UpdateCompanyRequestBody is the type of the "hy_company" service
// "updateCompany" endpoint HTTP request body.
type UpdateCompanyRequestBody struct {
	// Company name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Country ID
	CountryID *int `form:"country_id,omitempty" json:"country_id,omitempty" xml:"country_id,omitempty"`
	// Company Address
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
}

// CompanyResponseCollection is the type of the "hy_company" service
// "companyList" endpoint HTTP response body.
type CompanyResponseCollection []*CompanyResponse

// CompanyResponseDetailidCollection is the type of the "hy_company" service
// "companyList" endpoint HTTP response body.
type CompanyResponseDetailidCollection []*CompanyResponseDetailid

// CompanyResponseIDCollection is the type of the "hy_company" service
// "companyList" endpoint HTTP response body.
type CompanyResponseIDCollection []*CompanyResponseID

// CompanyResponseIdnameCollection is the type of the "hy_company" service
// "companyList" endpoint HTTP response body.
type CompanyResponseIdnameCollection []*CompanyResponseIdname

// CompanyResponse is used to define fields on response body types.
type CompanyResponse struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// ID
	CompanyID *int `form:"company_id,omitempty" json:"company_id,omitempty" xml:"company_id,omitempty"`
	// Company name
	Name        string  `form:"name" json:"name" xml:"name"`
	IsHq        *string `form:"is_hq,omitempty" json:"is_hq,omitempty" xml:"is_hq,omitempty"`
	CountryName *string `form:"country_name,omitempty" json:"country_name,omitempty" xml:"country_name,omitempty"`
	// Company Address
	Address string `form:"address" json:"address" xml:"address"`
}

// CompanyResponseDetailid is used to define fields on response body types.
type CompanyResponseDetailid struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// CompanyResponseID is used to define fields on response body types.
type CompanyResponseID struct {
	// ID
	CompanyID *int `form:"company_id,omitempty" json:"company_id,omitempty" xml:"company_id,omitempty"`
}

// CompanyResponseIdname is used to define fields on response body types.
type CompanyResponseIdname struct {
	// ID
	CompanyID *int `form:"company_id,omitempty" json:"company_id,omitempty" xml:"company_id,omitempty"`
	// Company name
	Name string `form:"name" json:"name" xml:"name"`
}

// NewCompanyResponseCollection builds the HTTP response body from the result
// of the "companyList" endpoint of the "hy_company" service.
func NewCompanyResponseCollection(res hycompanyviews.CompanyCollectionView) CompanyResponseCollection {
	body := make([]*CompanyResponse, len(res))
	for i, val := range res {
		body[i] = marshalHycompanyviewsCompanyViewToCompanyResponse(val)
	}
	return body
}

// NewCompanyResponseDetailidCollection builds the HTTP response body from the
// result of the "companyList" endpoint of the "hy_company" service.
func NewCompanyResponseDetailidCollection(res hycompanyviews.CompanyCollectionView) CompanyResponseDetailidCollection {
	body := make([]*CompanyResponseDetailid, len(res))
	for i, val := range res {
		body[i] = marshalHycompanyviewsCompanyViewToCompanyResponseDetailid(val)
	}
	return body
}

// NewCompanyResponseIDCollection builds the HTTP response body from the result
// of the "companyList" endpoint of the "hy_company" service.
func NewCompanyResponseIDCollection(res hycompanyviews.CompanyCollectionView) CompanyResponseIDCollection {
	body := make([]*CompanyResponseID, len(res))
	for i, val := range res {
		body[i] = marshalHycompanyviewsCompanyViewToCompanyResponseID(val)
	}
	return body
}

// NewCompanyResponseIdnameCollection builds the HTTP response body from the
// result of the "companyList" endpoint of the "hy_company" service.
func NewCompanyResponseIdnameCollection(res hycompanyviews.CompanyCollectionView) CompanyResponseIdnameCollection {
	body := make([]*CompanyResponseIdname, len(res))
	for i, val := range res {
		body[i] = marshalHycompanyviewsCompanyViewToCompanyResponseIdname(val)
	}
	return body
}

// NewCompanyListPayload builds a hy_company service companyList endpoint
// payload.
func NewCompanyListPayload(token *string) *hycompany.CompanyListPayload {
	v := &hycompany.CompanyListPayload{}
	v.Token = token

	return v
}

// NewGetCompanyGroupPayload builds a hy_company service getCompanyGroup
// endpoint payload.
func NewGetCompanyGroupPayload(body *GetCompanyGroupRequestBody, companyID int, token *string) *hycompany.GetCompanyGroupPayload {
	v := &hycompany.GetCompanyGroupPayload{
		IsHq: body.IsHq,
	}
	v.CompanyID = &companyID
	v.Token = token

	return v
}

// NewCreateCompanyPayload builds a hy_company service createCompany endpoint
// payload.
func NewCreateCompanyPayload(body *CreateCompanyRequestBody, token *string) *hycompany.CreateCompanyPayload {
	v := &hycompany.CreateCompanyPayload{
		Name:      *body.Name,
		CountryID: *body.CountryID,
		Address:   *body.Address,
	}
	v.Token = token

	return v
}

// NewUpdateCompanyPayload builds a hy_company service updateCompany endpoint
// payload.
func NewUpdateCompanyPayload(body *UpdateCompanyRequestBody, companyID int, token *string) *hycompany.UpdateCompanyPayload {
	v := &hycompany.UpdateCompanyPayload{
		Name:      *body.Name,
		CountryID: *body.CountryID,
		Address:   *body.Address,
	}
	v.CompanyID = &companyID
	v.Token = token

	return v
}

// NewDeleteCompanyPayload builds a hy_company service deleteCompany endpoint
// payload.
func NewDeleteCompanyPayload(companyID int, token *string) *hycompany.DeleteCompanyPayload {
	v := &hycompany.DeleteCompanyPayload{}
	v.CompanyID = &companyID
	v.Token = token

	return v
}

// ValidateGetCompanyGroupRequestBody runs the validations defined on
// GetCompanyGroupRequestBody
func ValidateGetCompanyGroupRequestBody(body *GetCompanyGroupRequestBody) (err error) {
	if body.IsHq != nil {
		if !(*body.IsHq == "1" || *body.IsHq == "0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.is_hq", *body.IsHq, []interface{}{"1", "0"}))
		}
	}
	return
}

// ValidateCreateCompanyRequestBody runs the validations defined on
// CreateCompanyRequestBody
func ValidateCreateCompanyRequestBody(body *CreateCompanyRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.CountryID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("country_id", "body"))
	}
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
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
	if body.CountryID != nil {
		if *body.CountryID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.country_id", *body.CountryID, 1, true))
		}
	}
	if body.CountryID != nil {
		if *body.CountryID > 999 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.country_id", *body.CountryID, 999, false))
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
	return
}

// ValidateUpdateCompanyRequestBody runs the validations defined on
// UpdateCompanyRequestBody
func ValidateUpdateCompanyRequestBody(body *UpdateCompanyRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.CountryID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("country_id", "body"))
	}
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
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
	if body.CountryID != nil {
		if *body.CountryID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.country_id", *body.CountryID, 1, true))
		}
	}
	if body.CountryID != nil {
		if *body.CountryID > 999 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.country_id", *body.CountryID, 999, false))
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
	return
}
