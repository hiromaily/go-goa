// Code generated by goa v3.10.1, DO NOT EDIT.
//
// hy_user HTTP client encoders and decoders
//
// Command:
// $ goa gen resume/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	hyuser "resume/gen/hy_user"
	hyuserviews "resume/gen/hy_user/views"
	"strings"

	goahttp "goa.design/goa/v3/http"
)

// BuildUserListRequest instantiates a HTTP request object with method and path
// set to call the "hy_user" service "userList" endpoint
func (c *Client) BuildUserListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UserListHyUserPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("hy_user", "userList", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUserListRequest returns an encoder for requests sent to the hy_user
// userList server.
func EncodeUserListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*hyuser.UserListPayload)
		if !ok {
			return goahttp.ErrInvalidType("hy_user", "userList", "*hyuser.UserListPayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeUserListResponse returns a decoder for responses returned by the
// hy_user userList endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeUserListResponse may return the following errors:
//   - "NotFound" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeUserListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UserListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("hy_user", "userList", err)
			}
			p := NewUserListUserCollectionOK(body)
			view := resp.Header.Get("goa-view")
			vres := hyuserviews.UserCollection{Projected: p, View: view}
			if err = hyuserviews.ValidateUserCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("hy_user", "userList", err)
			}
			res := hyuser.NewUserCollection(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body UserListNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("hy_user", "userList", err)
			}
			err = ValidateUserListNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("hy_user", "userList", err)
			}
			return nil, NewUserListNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("hy_user", "userList", resp.StatusCode, string(body))
		}
	}
}

// BuildGetUserRequest instantiates a HTTP request object with method and path
// set to call the "hy_user" service "getUser" endpoint
func (c *Client) BuildGetUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		userID int
	)
	{
		p, ok := v.(*hyuser.GetUserPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("hy_user", "getUser", "*hyuser.GetUserPayload", v)
		}
		userID = p.UserID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetUserHyUserPath(userID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("hy_user", "getUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetUserRequest returns an encoder for requests sent to the hy_user
// getUser server.
func EncodeGetUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*hyuser.GetUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("hy_user", "getUser", "*hyuser.GetUserPayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeGetUserResponse returns a decoder for responses returned by the
// hy_user getUser endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeGetUserResponse may return the following errors:
//   - "NotFound" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeGetUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("hy_user", "getUser", err)
			}
			p := NewGetUserUserOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &hyuserviews.User{Projected: p, View: view}
			if err = hyuserviews.ValidateUser(vres); err != nil {
				return nil, goahttp.ErrValidationError("hy_user", "getUser", err)
			}
			res := hyuser.NewUser(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body GetUserNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("hy_user", "getUser", err)
			}
			err = ValidateGetUserNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("hy_user", "getUser", err)
			}
			return nil, NewGetUserNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("hy_user", "getUser", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateUserRequest instantiates a HTTP request object with method and
// path set to call the "hy_user" service "createUser" endpoint
func (c *Client) BuildCreateUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateUserHyUserPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("hy_user", "createUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateUserRequest returns an encoder for requests sent to the hy_user
// createUser server.
func EncodeCreateUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*hyuser.CreateUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("hy_user", "createUser", "*hyuser.CreateUserPayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewCreateUserRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("hy_user", "createUser", err)
		}
		return nil
	}
}

// DecodeCreateUserResponse returns a decoder for responses returned by the
// hy_user createUser endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeCreateUserResponse may return the following errors:
//   - "BadRequest" (type *goa.ServiceError): http.StatusBadRequest
//   - error: internal error
func DecodeCreateUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body CreateUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("hy_user", "createUser", err)
			}
			p := NewCreateUserUserCreated(&body)
			view := resp.Header.Get("goa-view")
			vres := &hyuserviews.User{Projected: p, View: view}
			if err = hyuserviews.ValidateUser(vres); err != nil {
				return nil, goahttp.ErrValidationError("hy_user", "createUser", err)
			}
			res := hyuser.NewUser(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body CreateUserBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("hy_user", "createUser", err)
			}
			err = ValidateCreateUserBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("hy_user", "createUser", err)
			}
			return nil, NewCreateUserBadRequest(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("hy_user", "createUser", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateUserRequest instantiates a HTTP request object with method and
// path set to call the "hy_user" service "updateUser" endpoint
func (c *Client) BuildUpdateUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		userID int
	)
	{
		p, ok := v.(*hyuser.UpdateUserPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("hy_user", "updateUser", "*hyuser.UpdateUserPayload", v)
		}
		userID = p.UserID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateUserHyUserPath(userID)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("hy_user", "updateUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateUserRequest returns an encoder for requests sent to the hy_user
// updateUser server.
func EncodeUpdateUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*hyuser.UpdateUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("hy_user", "updateUser", "*hyuser.UpdateUserPayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		body := NewUpdateUserRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("hy_user", "updateUser", err)
		}
		return nil
	}
}

// DecodeUpdateUserResponse returns a decoder for responses returned by the
// hy_user updateUser endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeUpdateUserResponse may return the following errors:
//   - "BadRequest" (type *goa.ServiceError): http.StatusBadRequest
//   - "NotFound" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeUpdateUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		case http.StatusBadRequest:
			var (
				body UpdateUserBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("hy_user", "updateUser", err)
			}
			err = ValidateUpdateUserBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("hy_user", "updateUser", err)
			}
			return nil, NewUpdateUserBadRequest(&body)
		case http.StatusNotFound:
			var (
				body UpdateUserNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("hy_user", "updateUser", err)
			}
			err = ValidateUpdateUserNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("hy_user", "updateUser", err)
			}
			return nil, NewUpdateUserNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("hy_user", "updateUser", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteUserRequest instantiates a HTTP request object with method and
// path set to call the "hy_user" service "deleteUser" endpoint
func (c *Client) BuildDeleteUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		userID int
	)
	{
		p, ok := v.(*hyuser.DeleteUserPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("hy_user", "deleteUser", "*hyuser.DeleteUserPayload", v)
		}
		userID = p.UserID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteUserHyUserPath(userID)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("hy_user", "deleteUser", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteUserRequest returns an encoder for requests sent to the hy_user
// deleteUser server.
func EncodeDeleteUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*hyuser.DeleteUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("hy_user", "deleteUser", "*hyuser.DeleteUserPayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeDeleteUserResponse returns a decoder for responses returned by the
// hy_user deleteUser endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeDeleteUserResponse may return the following errors:
//   - "NotFound" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeDeleteUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		case http.StatusNotFound:
			var (
				body DeleteUserNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("hy_user", "deleteUser", err)
			}
			err = ValidateDeleteUserNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("hy_user", "deleteUser", err)
			}
			return nil, NewDeleteUserNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("hy_user", "deleteUser", resp.StatusCode, string(body))
		}
	}
}

// unmarshalUserResponseToHyuserviewsUserView builds a value of type
// *hyuserviews.UserView from a value of type *UserResponse.
func unmarshalUserResponseToHyuserviewsUserView(v *UserResponse) *hyuserviews.UserView {
	res := &hyuserviews.UserView{
		UserID:    v.UserID,
		UserName:  v.UserName,
		Email:     v.Email,
		Password:  v.Password,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}

	return res
}
