// Code generated by goa v3.10.2, DO NOT EDIT.
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
	UserName *string `form:"user_name,omitempty" json:"user_name,omitempty" xml:"user_name,omitempty"`
	// E-mail of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// Password
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// UserListResponseBody is the type of the "hy_user" service "userList"
// endpoint HTTP response body.
type UserListResponseBody []*UserResponse

// GetUserResponseBody is the type of the "hy_user" service "getUser" endpoint
// HTTP response body.
type GetUserResponseBody struct {
	// Key ID
	UserID *int `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
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

// CreateUserResponseBody is the type of the "hy_user" service "createUser"
// endpoint HTTP response body.
type CreateUserResponseBody struct {
	// Key ID
	UserID *int `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
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

// UserListNotFoundResponseBody is the type of the "hy_user" service "userList"
// endpoint HTTP response body for the "NotFound" error.
type UserListNotFoundResponseBody struct {
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

// UserListUnauthorizedResponseBody is the type of the "hy_user" service
// "userList" endpoint HTTP response body for the "Unauthorized" error.
type UserListUnauthorizedResponseBody struct {
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

// GetUserNotFoundResponseBody is the type of the "hy_user" service "getUser"
// endpoint HTTP response body for the "NotFound" error.
type GetUserNotFoundResponseBody struct {
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

// GetUserUnauthorizedResponseBody is the type of the "hy_user" service
// "getUser" endpoint HTTP response body for the "Unauthorized" error.
type GetUserUnauthorizedResponseBody struct {
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

// CreateUserBadRequestResponseBody is the type of the "hy_user" service
// "createUser" endpoint HTTP response body for the "BadRequest" error.
type CreateUserBadRequestResponseBody struct {
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

// CreateUserUnauthorizedResponseBody is the type of the "hy_user" service
// "createUser" endpoint HTTP response body for the "Unauthorized" error.
type CreateUserUnauthorizedResponseBody struct {
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

// UpdateUserBadRequestResponseBody is the type of the "hy_user" service
// "updateUser" endpoint HTTP response body for the "BadRequest" error.
type UpdateUserBadRequestResponseBody struct {
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

// UpdateUserNotFoundResponseBody is the type of the "hy_user" service
// "updateUser" endpoint HTTP response body for the "NotFound" error.
type UpdateUserNotFoundResponseBody struct {
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

// UpdateUserUnauthorizedResponseBody is the type of the "hy_user" service
// "updateUser" endpoint HTTP response body for the "Unauthorized" error.
type UpdateUserUnauthorizedResponseBody struct {
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

// DeleteUserNotFoundResponseBody is the type of the "hy_user" service
// "deleteUser" endpoint HTTP response body for the "NotFound" error.
type DeleteUserNotFoundResponseBody struct {
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

// DeleteUserUnauthorizedResponseBody is the type of the "hy_user" service
// "deleteUser" endpoint HTTP response body for the "Unauthorized" error.
type DeleteUserUnauthorizedResponseBody struct {
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

// UserResponse is used to define fields on response body types.
type UserResponse struct {
	// Key ID
	UserID *int `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
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

// NewUserListNotFound builds a hy_user service userList endpoint NotFound
// error.
func NewUserListNotFound(body *UserListNotFoundResponseBody) *goa.ServiceError {
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

// NewUserListUnauthorized builds a hy_user service userList endpoint
// Unauthorized error.
func NewUserListUnauthorized(body *UserListUnauthorizedResponseBody) *goa.ServiceError {
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

// NewGetUserUserOK builds a "hy_user" service "getUser" endpoint result from a
// HTTP "OK" response.
func NewGetUserUserOK(body *GetUserResponseBody) *hyuserviews.UserView {
	v := &hyuserviews.UserView{
		UserID:    body.UserID,
		UserName:  body.UserName,
		Email:     body.Email,
		Password:  body.Password,
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
	}

	return v
}

// NewGetUserNotFound builds a hy_user service getUser endpoint NotFound error.
func NewGetUserNotFound(body *GetUserNotFoundResponseBody) *goa.ServiceError {
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

// NewGetUserUnauthorized builds a hy_user service getUser endpoint
// Unauthorized error.
func NewGetUserUnauthorized(body *GetUserUnauthorizedResponseBody) *goa.ServiceError {
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

// NewCreateUserUserCreated builds a "hy_user" service "createUser" endpoint
// result from a HTTP "Created" response.
func NewCreateUserUserCreated(body *CreateUserResponseBody) *hyuserviews.UserView {
	v := &hyuserviews.UserView{
		UserID:    body.UserID,
		UserName:  body.UserName,
		Email:     body.Email,
		Password:  body.Password,
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
	}

	return v
}

// NewCreateUserBadRequest builds a hy_user service createUser endpoint
// BadRequest error.
func NewCreateUserBadRequest(body *CreateUserBadRequestResponseBody) *goa.ServiceError {
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

// NewCreateUserUnauthorized builds a hy_user service createUser endpoint
// Unauthorized error.
func NewCreateUserUnauthorized(body *CreateUserUnauthorizedResponseBody) *goa.ServiceError {
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

// NewUpdateUserBadRequest builds a hy_user service updateUser endpoint
// BadRequest error.
func NewUpdateUserBadRequest(body *UpdateUserBadRequestResponseBody) *goa.ServiceError {
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

// NewUpdateUserNotFound builds a hy_user service updateUser endpoint NotFound
// error.
func NewUpdateUserNotFound(body *UpdateUserNotFoundResponseBody) *goa.ServiceError {
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

// NewUpdateUserUnauthorized builds a hy_user service updateUser endpoint
// Unauthorized error.
func NewUpdateUserUnauthorized(body *UpdateUserUnauthorizedResponseBody) *goa.ServiceError {
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

// NewDeleteUserNotFound builds a hy_user service deleteUser endpoint NotFound
// error.
func NewDeleteUserNotFound(body *DeleteUserNotFoundResponseBody) *goa.ServiceError {
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

// NewDeleteUserUnauthorized builds a hy_user service deleteUser endpoint
// Unauthorized error.
func NewDeleteUserUnauthorized(body *DeleteUserUnauthorizedResponseBody) *goa.ServiceError {
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

// ValidateUserListNotFoundResponseBody runs the validations defined on
// userList_NotFound_response_body
func ValidateUserListNotFoundResponseBody(body *UserListNotFoundResponseBody) (err error) {
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

// ValidateUserListUnauthorizedResponseBody runs the validations defined on
// userList_Unauthorized_response_body
func ValidateUserListUnauthorizedResponseBody(body *UserListUnauthorizedResponseBody) (err error) {
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

// ValidateGetUserNotFoundResponseBody runs the validations defined on
// getUser_NotFound_response_body
func ValidateGetUserNotFoundResponseBody(body *GetUserNotFoundResponseBody) (err error) {
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

// ValidateGetUserUnauthorizedResponseBody runs the validations defined on
// getUser_Unauthorized_response_body
func ValidateGetUserUnauthorizedResponseBody(body *GetUserUnauthorizedResponseBody) (err error) {
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

// ValidateCreateUserBadRequestResponseBody runs the validations defined on
// createUser_BadRequest_response_body
func ValidateCreateUserBadRequestResponseBody(body *CreateUserBadRequestResponseBody) (err error) {
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

// ValidateCreateUserUnauthorizedResponseBody runs the validations defined on
// createUser_Unauthorized_response_body
func ValidateCreateUserUnauthorizedResponseBody(body *CreateUserUnauthorizedResponseBody) (err error) {
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

// ValidateUpdateUserBadRequestResponseBody runs the validations defined on
// updateUser_BadRequest_response_body
func ValidateUpdateUserBadRequestResponseBody(body *UpdateUserBadRequestResponseBody) (err error) {
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

// ValidateUpdateUserNotFoundResponseBody runs the validations defined on
// updateUser_NotFound_response_body
func ValidateUpdateUserNotFoundResponseBody(body *UpdateUserNotFoundResponseBody) (err error) {
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

// ValidateUpdateUserUnauthorizedResponseBody runs the validations defined on
// updateUser_Unauthorized_response_body
func ValidateUpdateUserUnauthorizedResponseBody(body *UpdateUserUnauthorizedResponseBody) (err error) {
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

// ValidateDeleteUserNotFoundResponseBody runs the validations defined on
// deleteUser_NotFound_response_body
func ValidateDeleteUserNotFoundResponseBody(body *DeleteUserNotFoundResponseBody) (err error) {
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

// ValidateDeleteUserUnauthorizedResponseBody runs the validations defined on
// deleteUser_Unauthorized_response_body
func ValidateDeleteUserUnauthorizedResponseBody(body *DeleteUserUnauthorizedResponseBody) (err error) {
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

// ValidateUserResponse runs the validations defined on UserResponse
func ValidateUserResponse(body *UserResponse) (err error) {
	if body.UserID != nil {
		if *body.UserID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.user_id", *body.UserID, 1, true))
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
