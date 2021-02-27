// Code generated by goa v3.2.6, DO NOT EDIT.
//
// auth HTTP client encoders and decoders
//
// Command:
// $ goa gen resume/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	auth "resume/gen/auth"
	authviews "resume/gen/auth/views"

	goahttp "goa.design/goa/v3/http"
)

// BuildLoginRequest instantiates a HTTP request object with method and path
// set to call the "auth" service "login" endpoint
func (c *Client) BuildLoginRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: LoginAuthPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("auth", "login", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeLoginRequest returns an encoder for requests sent to the auth login
// server.
func EncodeLoginRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*auth.LoginPayload)
		if !ok {
			return goahttp.ErrInvalidType("auth", "login", "*auth.LoginPayload", v)
		}
		body := NewLoginRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("auth", "login", err)
		}
		return nil
	}
}

// DecodeLoginResponse returns a decoder for responses returned by the auth
// login endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeLoginResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body LoginResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("auth", "login", err)
			}
			p := NewLoginAuthorizedOK(&body)
			view := "default"
			vres := &authviews.Authorized{Projected: p, View: view}
			if err = authviews.ValidateAuthorized(vres); err != nil {
				return nil, goahttp.ErrValidationError("auth", "login", err)
			}
			res := auth.NewAuthorized(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("auth", "login", resp.StatusCode, string(body))
		}
	}
}
