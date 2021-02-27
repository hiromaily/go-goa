// Code generated by goa v3.2.6, DO NOT EDIT.
//
// auth HTTP server encoders and decoders
//
// Command:
// $ goa gen resume/design

package server

import (
	"context"
	"io"
	"net/http"
	authviews "resume/gen/auth/views"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeLoginResponse returns an encoder for responses returned by the auth
// login endpoint.
func EncodeLoginResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*authviews.Authorized)
		enc := encoder(ctx, w)
		body := NewLoginResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeLoginRequest returns a decoder for requests sent to the auth login
// endpoint.
func DecodeLoginRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body LoginRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateLoginRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewLoginPayloadLogin(&body)

		return payload, nil
	}
}
