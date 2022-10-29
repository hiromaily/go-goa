// Code generated by goa v3.10.1, DO NOT EDIT.
//
// hy_userWorkHistory HTTP client types
//
// Command:
// $ goa gen resume/design

package client

import (
	hyuserworkhistoryviews "resume/gen/hy_user_work_history/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// GetUserWorkHistoryResponseBody is the type of the "hy_userWorkHistory"
// service "getUserWorkHistory" endpoint HTTP response body.
type GetUserWorkHistoryResponseBody []*UserworkhistoryResponse

// GetUserWorkHistoryNotFoundResponseBody is the type of the
// "hy_userWorkHistory" service "getUserWorkHistory" endpoint HTTP response
// body for the "NotFound" error.
type GetUserWorkHistoryNotFoundResponseBody struct {
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

// UserworkhistoryResponse is used to define fields on response body types.
type UserworkhistoryResponse struct {
	// Job Title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Company name
	Company *string `form:"company,omitempty" json:"company,omitempty" xml:"company,omitempty"`
	// Country code
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// worked period
	Term *string `form:"term,omitempty" json:"term,omitempty" xml:"term,omitempty"`
	// job description
	Description interface{} `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// used techs
	Techs interface{} `form:"techs,omitempty" json:"techs,omitempty" xml:"techs,omitempty"`
}

// NewGetUserWorkHistoryUserworkhistoryCollectionOK builds a
// "hy_userWorkHistory" service "getUserWorkHistory" endpoint result from a
// HTTP "OK" response.
func NewGetUserWorkHistoryUserworkhistoryCollectionOK(body GetUserWorkHistoryResponseBody) hyuserworkhistoryviews.UserworkhistoryCollectionView {
	v := make([]*hyuserworkhistoryviews.UserworkhistoryView, len(body))
	for i, val := range body {
		v[i] = unmarshalUserworkhistoryResponseToHyuserworkhistoryviewsUserworkhistoryView(val)
	}

	return v
}

// NewGetUserWorkHistoryNotFound builds a hy_userWorkHistory service
// getUserWorkHistory endpoint NotFound error.
func NewGetUserWorkHistoryNotFound(body *GetUserWorkHistoryNotFoundResponseBody) *goa.ServiceError {
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

// ValidateGetUserWorkHistoryNotFoundResponseBody runs the validations defined
// on getUserWorkHistory_NotFound_response_body
func ValidateGetUserWorkHistoryNotFoundResponseBody(body *GetUserWorkHistoryNotFoundResponseBody) (err error) {
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

// ValidateUserworkhistoryResponse runs the validations defined on
// UserworkhistoryResponse
func ValidateUserworkhistoryResponse(body *UserworkhistoryResponse) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Company == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("company", "body"))
	}
	if body.Country == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("country", "body"))
	}
	if body.Title != nil {
		if utf8.RuneCountInString(*body.Title) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.title", *body.Title, utf8.RuneCountInString(*body.Title), 2, true))
		}
	}
	if body.Title != nil {
		if utf8.RuneCountInString(*body.Title) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.title", *body.Title, utf8.RuneCountInString(*body.Title), 40, false))
		}
	}
	if body.Company != nil {
		if utf8.RuneCountInString(*body.Company) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.company", *body.Company, utf8.RuneCountInString(*body.Company), 2, true))
		}
	}
	if body.Company != nil {
		if utf8.RuneCountInString(*body.Company) > 40 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.company", *body.Company, utf8.RuneCountInString(*body.Company), 40, false))
		}
	}
	if body.Country != nil {
		if utf8.RuneCountInString(*body.Country) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.country", *body.Country, utf8.RuneCountInString(*body.Country), 2, true))
		}
	}
	if body.Country != nil {
		if utf8.RuneCountInString(*body.Country) > 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.country", *body.Country, utf8.RuneCountInString(*body.Country), 2, false))
		}
	}
	if body.Term != nil {
		if utf8.RuneCountInString(*body.Term) < 10 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.term", *body.Term, utf8.RuneCountInString(*body.Term), 10, true))
		}
	}
	if body.Term != nil {
		if utf8.RuneCountInString(*body.Term) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.term", *body.Term, utf8.RuneCountInString(*body.Term), 20, false))
		}
	}
	return
}
