// Code generated by goa v3.2.6, DO NOT EDIT.
//
// hy_user HTTP server encoders and decoders
//
// Command:
// $ goa gen resume/design

package server

import (
	"context"
	"io"
	"net/http"
	hyuserviews "resume/gen/hy_user/views"
	"strconv"
	"strings"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeUserListResponse returns an encoder for responses returned by the
// hy_user userList endpoint.
func EncodeUserListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(hyuserviews.UserCollection)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewUserResponseCollection(res.Projected)
		case "id":
			body = NewUserResponseIDCollection(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUserListRequest returns a decoder for requests sent to the hy_user
// userList endpoint.
func DecodeUserListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token *string
		)
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewUserListPayload(token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeGetUserResponse returns an encoder for responses returned by the
// hy_user getUser endpoint.
func EncodeGetUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*hyuserviews.User)
		w.Header().Set("goa-view", res.View)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewGetUserResponseBody(res.Projected)
		case "id":
			body = NewGetUserResponseBodyID(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetUserRequest returns a decoder for requests sent to the hy_user
// getUser endpoint.
func DecodeGetUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			userID int
			token  *string
			err    error

			params = mux.Vars(r)
		)
		{
			userIDRaw := params["userID"]
			v, err2 := strconv.ParseInt(userIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("userID", userIDRaw, "integer"))
			}
			userID = int(v)
		}
		if userID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("userID", userID, 1, true))
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetUserPayload(userID, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeCreateUserResponse returns an encoder for responses returned by the
// hy_user createUser endpoint.
func EncodeCreateUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeCreateUserRequest returns a decoder for requests sent to the hy_user
// createUser endpoint.
func DecodeCreateUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateUserRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			token *string
		)
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewCreateUserPayload(&body, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeUpdateUserResponse returns an encoder for responses returned by the
// hy_user updateUser endpoint.
func EncodeUpdateUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeUpdateUserRequest returns a decoder for requests sent to the hy_user
// updateUser endpoint.
func DecodeUpdateUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdateUserRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateUserRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			userID int
			token  *string

			params = mux.Vars(r)
		)
		{
			userIDRaw := params["userID"]
			v, err2 := strconv.ParseInt(userIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("userID", userIDRaw, "integer"))
			}
			userID = int(v)
		}
		if userID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("userID", userID, 1, true))
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdateUserPayload(&body, userID, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeDeleteUserResponse returns an encoder for responses returned by the
// hy_user deleteUser endpoint.
func EncodeDeleteUserResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDeleteUserRequest returns a decoder for requests sent to the hy_user
// deleteUser endpoint.
func DecodeDeleteUserRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			userID int
			token  *string
			err    error

			params = mux.Vars(r)
		)
		{
			userIDRaw := params["userID"]
			v, err2 := strconv.ParseInt(userIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("userID", userIDRaw, "integer"))
			}
			userID = int(v)
		}
		if userID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("userID", userID, 1, true))
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeleteUserPayload(userID, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// marshalHyuserviewsUserViewToUserResponse builds a value of type
// *UserResponse from a value of type *hyuserviews.UserView.
func marshalHyuserviewsUserViewToUserResponse(v *hyuserviews.UserView) *UserResponse {
	res := &UserResponse{
		ID:       v.ID,
		UserName: *v.UserName,
		Email:    *v.Email,
	}

	return res
}

// marshalHyuserviewsUserViewToUserResponseID builds a value of type
// *UserResponseID from a value of type *hyuserviews.UserView.
func marshalHyuserviewsUserViewToUserResponseID(v *hyuserviews.UserView) *UserResponseID {
	res := &UserResponseID{
		ID: v.ID,
	}

	return res
}
