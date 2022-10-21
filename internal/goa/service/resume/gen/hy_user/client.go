// Code generated by goa v3.10.1, DO NOT EDIT.
//
// hy_user client
//
// Command:
// $ goa gen resume/design

package hyuser

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "hy_user" service client.
type Client struct {
	UserListEndpoint   goa.Endpoint
	GetUserEndpoint    goa.Endpoint
	CreateUserEndpoint goa.Endpoint
	UpdateUserEndpoint goa.Endpoint
	DeleteUserEndpoint goa.Endpoint
}

// NewClient initializes a "hy_user" service client given the endpoints.
func NewClient(userList, getUser, createUser, updateUser, deleteUser goa.Endpoint) *Client {
	return &Client{
		UserListEndpoint:   userList,
		GetUserEndpoint:    getUser,
		CreateUserEndpoint: createUser,
		UpdateUserEndpoint: updateUser,
		DeleteUserEndpoint: deleteUser,
	}
}

// UserList calls the "userList" endpoint of the "hy_user" service.
// UserList may return the following errors:
//   - "NoContent" (type *goa.ServiceError)
//   - "BadRequest" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) UserList(ctx context.Context, p *UserListPayload) (res UserCollection, err error) {
	var ires interface{}
	ires, err = c.UserListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(UserCollection), nil
}

// GetUser calls the "getUser" endpoint of the "hy_user" service.
// GetUser may return the following errors:
//   - "BadRequest" (type *goa.ServiceError)
//   - "NotFound" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) GetUser(ctx context.Context, p *GetUserPayload) (res *User, err error) {
	var ires interface{}
	ires, err = c.GetUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*User), nil
}

// CreateUser calls the "createUser" endpoint of the "hy_user" service.
// CreateUser may return the following errors:
//   - "BadRequest" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) CreateUser(ctx context.Context, p *CreateUserPayload) (err error) {
	_, err = c.CreateUserEndpoint(ctx, p)
	return
}

// UpdateUser calls the "updateUser" endpoint of the "hy_user" service.
// UpdateUser may return the following errors:
//   - "BadRequest" (type *goa.ServiceError)
//   - "NotFound" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) UpdateUser(ctx context.Context, p *UpdateUserPayload) (err error) {
	_, err = c.UpdateUserEndpoint(ctx, p)
	return
}

// DeleteUser calls the "deleteUser" endpoint of the "hy_user" service.
// DeleteUser may return the following errors:
//   - "BadRequest" (type *goa.ServiceError)
//   - "NotFound" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) DeleteUser(ctx context.Context, p *DeleteUserPayload) (err error) {
	_, err = c.DeleteUserEndpoint(ctx, p)
	return
}
