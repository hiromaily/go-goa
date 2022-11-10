// Code generated by goa v3.10.2, DO NOT EDIT.
//
// hy_usertech HTTP server types
//
// Command:
// $ goa gen resume/design

package server

import (
	hyusertech "resume/gen/hy_usertech"
	hyusertechviews "resume/gen/hy_usertech/views"

	goa "goa.design/goa/v3/pkg"
)

// UsertechResponseCollection is the type of the "hy_usertech" service
// "getUserLikeTech" endpoint HTTP response body.
type UsertechResponseCollection []*UsertechResponse

// UsertechResponseTechNameCollection is the type of the "hy_usertech" service
// "getUserLikeTech" endpoint HTTP response body.
type UsertechResponseTechNameCollection []*UsertechResponseTechName

// GetUserLikeTechNotFoundResponseBody is the type of the "hy_usertech" service
// "getUserLikeTech" endpoint HTTP response body for the "NotFound" error.
type GetUserLikeTechNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetUserLikeTechUnauthorizedResponseBody is the type of the "hy_usertech"
// service "getUserLikeTech" endpoint HTTP response body for the "Unauthorized"
// error.
type GetUserLikeTechUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetUserDisLikeTechNotFoundResponseBody is the type of the "hy_usertech"
// service "getUserDisLikeTech" endpoint HTTP response body for the "NotFound"
// error.
type GetUserDisLikeTechNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetUserDisLikeTechUnauthorizedResponseBody is the type of the "hy_usertech"
// service "getUserDisLikeTech" endpoint HTTP response body for the
// "Unauthorized" error.
type GetUserDisLikeTechUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UsertechResponse is used to define fields on response body types.
type UsertechResponse struct {
	// Key ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Tech name
	TechName *string `form:"tech_name,omitempty" json:"tech_name,omitempty" xml:"tech_name,omitempty"`
}

// UsertechResponseTechName is used to define fields on response body types.
type UsertechResponseTechName struct {
	// Tech name
	TechName *string `form:"tech_name,omitempty" json:"tech_name,omitempty" xml:"tech_name,omitempty"`
}

// NewUsertechResponseCollection builds the HTTP response body from the result
// of the "getUserLikeTech" endpoint of the "hy_usertech" service.
func NewUsertechResponseCollection(res hyusertechviews.UsertechCollectionView) UsertechResponseCollection {
	body := make([]*UsertechResponse, len(res))
	for i, val := range res {
		body[i] = marshalHyusertechviewsUsertechViewToUsertechResponse(val)
	}
	return body
}

// NewUsertechResponseTechNameCollection builds the HTTP response body from the
// result of the "getUserLikeTech" endpoint of the "hy_usertech" service.
func NewUsertechResponseTechNameCollection(res hyusertechviews.UsertechCollectionView) UsertechResponseTechNameCollection {
	body := make([]*UsertechResponseTechName, len(res))
	for i, val := range res {
		body[i] = marshalHyusertechviewsUsertechViewToUsertechResponseTechName(val)
	}
	return body
}

// NewGetUserLikeTechNotFoundResponseBody builds the HTTP response body from
// the result of the "getUserLikeTech" endpoint of the "hy_usertech" service.
func NewGetUserLikeTechNotFoundResponseBody(res *goa.ServiceError) *GetUserLikeTechNotFoundResponseBody {
	body := &GetUserLikeTechNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetUserLikeTechUnauthorizedResponseBody builds the HTTP response body
// from the result of the "getUserLikeTech" endpoint of the "hy_usertech"
// service.
func NewGetUserLikeTechUnauthorizedResponseBody(res *goa.ServiceError) *GetUserLikeTechUnauthorizedResponseBody {
	body := &GetUserLikeTechUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetUserDisLikeTechNotFoundResponseBody builds the HTTP response body from
// the result of the "getUserDisLikeTech" endpoint of the "hy_usertech" service.
func NewGetUserDisLikeTechNotFoundResponseBody(res *goa.ServiceError) *GetUserDisLikeTechNotFoundResponseBody {
	body := &GetUserDisLikeTechNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetUserDisLikeTechUnauthorizedResponseBody builds the HTTP response body
// from the result of the "getUserDisLikeTech" endpoint of the "hy_usertech"
// service.
func NewGetUserDisLikeTechUnauthorizedResponseBody(res *goa.ServiceError) *GetUserDisLikeTechUnauthorizedResponseBody {
	body := &GetUserDisLikeTechUnauthorizedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetUserLikeTechPayload builds a hy_usertech service getUserLikeTech
// endpoint payload.
func NewGetUserLikeTechPayload(userID int, token *string) *hyusertech.GetUserLikeTechPayload {
	v := &hyusertech.GetUserLikeTechPayload{}
	v.UserID = userID
	v.Token = token

	return v
}

// NewGetUserDisLikeTechPayload builds a hy_usertech service getUserDisLikeTech
// endpoint payload.
func NewGetUserDisLikeTechPayload(userID int, token *string) *hyusertech.GetUserDisLikeTechPayload {
	v := &hyusertech.GetUserDisLikeTechPayload{}
	v.UserID = userID
	v.Token = token

	return v
}
