// Code generated by goa v3.10.1, DO NOT EDIT.
//
// hy_tech HTTP server types
//
// Command:
// $ goa gen resume/design

package server

import (
	hytech "resume/gen/hy_tech"
	hytechviews "resume/gen/hy_tech/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// CreateTechRequestBody is the type of the "hy_tech" service "createTech"
// endpoint HTTP request body.
type CreateTechRequestBody struct {
	// Tech name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// UpdateTechRequestBody is the type of the "hy_tech" service "updateTech"
// endpoint HTTP request body.
type UpdateTechRequestBody struct {
	// Tech name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// TechResponseCollection is the type of the "hy_tech" service "techList"
// endpoint HTTP response body.
type TechResponseCollection []*TechResponse

// TechResponseIDCollection is the type of the "hy_tech" service "techList"
// endpoint HTTP response body.
type TechResponseIDCollection []*TechResponseID

// GetTechResponseBody is the type of the "hy_tech" service "getTech" endpoint
// HTTP response body.
type GetTechResponseBody struct {
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

// GetTechResponseBodyDetailid is the type of the "hy_tech" service "getTech"
// endpoint HTTP response body.
type GetTechResponseBodyDetailid struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// GetTechResponseBodyID is the type of the "hy_tech" service "getTech"
// endpoint HTTP response body.
type GetTechResponseBodyID struct {
	// ID
	CompanyID *int `form:"company_id,omitempty" json:"company_id,omitempty" xml:"company_id,omitempty"`
}

// GetTechResponseBodyIdname is the type of the "hy_tech" service "getTech"
// endpoint HTTP response body.
type GetTechResponseBodyIdname struct {
	// ID
	CompanyID *int `form:"company_id,omitempty" json:"company_id,omitempty" xml:"company_id,omitempty"`
	// Company name
	Name string `form:"name" json:"name" xml:"name"`
}

// TechResponse is used to define fields on response body types.
type TechResponse struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Tech name
	Name string `form:"name" json:"name" xml:"name"`
}

// TechResponseID is used to define fields on response body types.
type TechResponseID struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// NewTechResponseCollection builds the HTTP response body from the result of
// the "techList" endpoint of the "hy_tech" service.
func NewTechResponseCollection(res hytechviews.TechCollectionView) TechResponseCollection {
	body := make([]*TechResponse, len(res))
	for i, val := range res {
		body[i] = marshalHytechviewsTechViewToTechResponse(val)
	}
	return body
}

// NewTechResponseIDCollection builds the HTTP response body from the result of
// the "techList" endpoint of the "hy_tech" service.
func NewTechResponseIDCollection(res hytechviews.TechCollectionView) TechResponseIDCollection {
	body := make([]*TechResponseID, len(res))
	for i, val := range res {
		body[i] = marshalHytechviewsTechViewToTechResponseID(val)
	}
	return body
}

// NewGetTechResponseBody builds the HTTP response body from the result of the
// "getTech" endpoint of the "hy_tech" service.
func NewGetTechResponseBody(res *hytechviews.CompanyView) *GetTechResponseBody {
	body := &GetTechResponseBody{
		ID:          res.ID,
		CompanyID:   res.CompanyID,
		Name:        *res.Name,
		IsHq:        res.IsHq,
		CountryName: res.CountryName,
		Address:     *res.Address,
	}
	return body
}

// NewGetTechResponseBodyDetailid builds the HTTP response body from the result
// of the "getTech" endpoint of the "hy_tech" service.
func NewGetTechResponseBodyDetailid(res *hytechviews.CompanyView) *GetTechResponseBodyDetailid {
	body := &GetTechResponseBodyDetailid{
		ID: res.ID,
	}
	return body
}

// NewGetTechResponseBodyID builds the HTTP response body from the result of
// the "getTech" endpoint of the "hy_tech" service.
func NewGetTechResponseBodyID(res *hytechviews.CompanyView) *GetTechResponseBodyID {
	body := &GetTechResponseBodyID{
		CompanyID: res.CompanyID,
	}
	return body
}

// NewGetTechResponseBodyIdname builds the HTTP response body from the result
// of the "getTech" endpoint of the "hy_tech" service.
func NewGetTechResponseBodyIdname(res *hytechviews.CompanyView) *GetTechResponseBodyIdname {
	body := &GetTechResponseBodyIdname{
		CompanyID: res.CompanyID,
		Name:      *res.Name,
	}
	return body
}

// NewTechListPayload builds a hy_tech service techList endpoint payload.
func NewTechListPayload(token *string) *hytech.TechListPayload {
	v := &hytech.TechListPayload{}
	v.Token = token

	return v
}

// NewGetTechPayload builds a hy_tech service getTech endpoint payload.
func NewGetTechPayload(techID int, token *string) *hytech.GetTechPayload {
	v := &hytech.GetTechPayload{}
	v.TechID = &techID
	v.Token = token

	return v
}

// NewCreateTechPayload builds a hy_tech service createTech endpoint payload.
func NewCreateTechPayload(body *CreateTechRequestBody, token *string) *hytech.CreateTechPayload {
	v := &hytech.CreateTechPayload{
		Name: *body.Name,
	}
	v.Token = token

	return v
}

// NewUpdateTechPayload builds a hy_tech service updateTech endpoint payload.
func NewUpdateTechPayload(body *UpdateTechRequestBody, techID int, token *string) *hytech.UpdateTechPayload {
	v := &hytech.UpdateTechPayload{
		Name: *body.Name,
	}
	v.TechID = &techID
	v.Token = token

	return v
}

// NewDeleteTechPayload builds a hy_tech service deleteTech endpoint payload.
func NewDeleteTechPayload(techID int, token *string) *hytech.DeleteTechPayload {
	v := &hytech.DeleteTechPayload{}
	v.TechID = &techID
	v.Token = token

	return v
}

// ValidateCreateTechRequestBody runs the validations defined on
// CreateTechRequestBody
func ValidateCreateTechRequestBody(body *CreateTechRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 1, true))
		}
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 40, false))
		}
	}
	return
}

// ValidateUpdateTechRequestBody runs the validations defined on
// UpdateTechRequestBody
func ValidateUpdateTechRequestBody(body *UpdateTechRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 1, true))
		}
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 40, false))
		}
	}
	return
}
