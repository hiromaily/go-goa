// Code generated by goa v3.7.13, DO NOT EDIT.
//
// hy_user HTTP client types
//
// Command:
// $ goa gen resume/design

package client

import (
	hyuser "resume/gen/hy_user"
	hyuserviews "resume/gen/hy_user/views"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// CreateUserRequestBody is the type of the "hy_user" service "createUser"
// endpoint HTTP request body.
type CreateUserRequestBody struct {
	// User ID
	UserID *int `form:"userID,omitempty" json:"userID,omitempty" xml:"userID,omitempty"`
	// User name
	UserName string `form:"user_name" json:"user_name" xml:"user_name"`
	// E-mail of user
	Email string `form:"email" json:"email" xml:"email"`
	// Password
	Password string `form:"password" json:"password" xml:"password"`
}

// UpdateUserRequestBody is the type of the "hy_user" service "updateUser"
// endpoint HTTP request body.
type UpdateUserRequestBody struct {
	// User name
	UserName string `form:"user_name" json:"user_name" xml:"user_name"`
	// E-mail of user
	Email string `form:"email" json:"email" xml:"email"`
	// Password
	Password string `form:"password" json:"password" xml:"password"`
}

// UserListResponseBody is the type of the "hy_user" service "userList"
// endpoint HTTP response body.
type UserListResponseBody []*UserResponse

// GetUserResponseBody is the type of the "hy_user" service "getUser" endpoint
// HTTP response body.
type GetUserResponseBody struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// User name
	UserName *string `form:"user_name,omitempty" json:"user_name,omitempty" xml:"user_name,omitempty"`
	// E-mail of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// Datetime
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Datetime
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// UserResponse is used to define fields on response body types.
type UserResponse struct {
	// ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// User name
	UserName *string `form:"user_name,omitempty" json:"user_name,omitempty" xml:"user_name,omitempty"`
	// E-mail of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// Datetime
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Datetime
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// NewCreateUserRequestBody builds the HTTP request body from the payload of
// the "createUser" endpoint of the "hy_user" service.
func NewCreateUserRequestBody(p *hyuser.CreateUserPayload) *CreateUserRequestBody {
	body := &CreateUserRequestBody{
		UserID:   p.UserID,
		UserName: p.UserName,
		Email:    p.Email,
		Password: p.Password,
	}
	return body
}

// NewUpdateUserRequestBody builds the HTTP request body from the payload of
// the "updateUser" endpoint of the "hy_user" service.
func NewUpdateUserRequestBody(p *hyuser.UpdateUserPayload) *UpdateUserRequestBody {
	body := &UpdateUserRequestBody{
		UserName: p.UserName,
		Email:    p.Email,
		Password: p.Password,
	}
	return body
}

// NewUserListUserCollectionOK builds a "hy_user" service "userList" endpoint
// result from a HTTP "OK" response.
func NewUserListUserCollectionOK(body UserListResponseBody) hyuserviews.UserCollectionView {
	v := make([]*hyuserviews.UserView, len(body))
	for i, val := range body {
		v[i] = unmarshalUserResponseToHyuserviewsUserView(val)
	}

	return v
}

// NewGetUserUserOK builds a "hy_user" service "getUser" endpoint result from a
// HTTP "OK" response.
func NewGetUserUserOK(body *GetUserResponseBody) *hyuserviews.UserView {
	v := &hyuserviews.UserView{
		ID:        body.ID,
		UserName:  body.UserName,
		Email:     body.Email,
		Password:  body.Password,
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
	}

	return v
}

// ValidateUserResponse runs the validations defined on UserResponse
func ValidateUserResponse(body *UserResponse) (err error) {
	if body.UserName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_name", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	if body.ID != nil {
		if *body.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.id", *body.ID, 1, true))
		}
	}
	if body.UserName != nil {
		if utf8.RuneCountInString(*body.UserName) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.user_name", *body.UserName, utf8.RuneCountInString(*body.UserName), 2, true))
		}
	}
	if body.UserName != nil {
		if utf8.RuneCountInString(*body.UserName) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.user_name", *body.UserName, utf8.RuneCountInString(*body.UserName), 20, false))
		}
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	if body.Password != nil {
		if utf8.RuneCountInString(*body.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", *body.Password, utf8.RuneCountInString(*body.Password), 8, true))
		}
	}
	if body.Password != nil {
		if utf8.RuneCountInString(*body.Password) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.password", *body.Password, utf8.RuneCountInString(*body.Password), 20, false))
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
