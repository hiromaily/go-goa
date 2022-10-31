// Code generated by goa v3.10.2, DO NOT EDIT.
//
// hy_tech HTTP server encoders and decoders
//
// Command:
// $ goa gen resume/design

package server

import (
	"context"
	"errors"
	"io"
	"net/http"
	hytechviews "resume/gen/hy_tech/views"
	"strconv"
	"strings"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeTechListResponse returns an encoder for responses returned by the
// hy_tech techList endpoint.
func EncodeTechListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(hytechviews.TechCollection)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewTechResponseCollection(res.Projected)
		case "id":
			body = NewTechResponseIDCollection(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeTechListRequest returns a decoder for requests sent to the hy_tech
// techList endpoint.
func DecodeTechListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			token *string
		)
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		payload := NewTechListPayload(token)
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

// EncodeTechListError returns an encoder for errors returned by the techList
// hy_tech endpoint.
func EncodeTechListError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "NotFound":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewTechListNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetTechResponse returns an encoder for responses returned by the
// hy_tech getTech endpoint.
func EncodeGetTechResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*hytechviews.Tech)
		w.Header().Set("goa-view", res.View)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewGetTechResponseBody(res.Projected)
		case "id":
			body = NewGetTechResponseBodyID(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetTechRequest returns a decoder for requests sent to the hy_tech
// getTech endpoint.
func DecodeGetTechRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			techID int
			token  *string
			err    error

			params = mux.Vars(r)
		)
		{
			techIDRaw := params["tech_id"]
			v, err2 := strconv.ParseInt(techIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("techID", techIDRaw, "integer"))
			}
			techID = int(v)
		}
		if techID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("techID", techID, 1, true))
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetTechPayload(techID, token)
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

// EncodeGetTechError returns an encoder for errors returned by the getTech
// hy_tech endpoint.
func EncodeGetTechError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "NotFound":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewGetTechNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateTechResponse returns an encoder for responses returned by the
// hy_tech createTech endpoint.
func EncodeCreateTechResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*hytechviews.Tech)
		w.Header().Set("goa-view", res.View)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/json")
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewCreateTechOKResponseBody(res.Projected)
		case "id":
			body = NewCreateTechOKResponseBodyID(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeCreateTechRequest returns a decoder for requests sent to the hy_tech
// createTech endpoint.
func DecodeCreateTechRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateTechRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateTechRequestBody(&body)
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
		payload := NewCreateTechPayload(&body, token)
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

// EncodeCreateTechError returns an encoder for errors returned by the
// createTech hy_tech endpoint.
func EncodeCreateTechError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "BadRequest":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewCreateTechBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdateTechResponse returns an encoder for responses returned by the
// hy_tech updateTech endpoint.
func EncodeUpdateTechResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeUpdateTechRequest returns a decoder for requests sent to the hy_tech
// updateTech endpoint.
func DecodeUpdateTechRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdateTechRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateTechRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			techID int
			token  *string

			params = mux.Vars(r)
		)
		{
			techIDRaw := params["tech_id"]
			v, err2 := strconv.ParseInt(techIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("techID", techIDRaw, "integer"))
			}
			techID = int(v)
		}
		if techID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("techID", techID, 1, true))
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdateTechPayload(&body, techID, token)
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

// EncodeUpdateTechError returns an encoder for errors returned by the
// updateTech hy_tech endpoint.
func EncodeUpdateTechError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "BadRequest":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUpdateTechBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "NotFound":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUpdateTechNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteTechResponse returns an encoder for responses returned by the
// hy_tech deleteTech endpoint.
func EncodeDeleteTechResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeDeleteTechRequest returns a decoder for requests sent to the hy_tech
// deleteTech endpoint.
func DecodeDeleteTechRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			techID int
			token  *string
			err    error

			params = mux.Vars(r)
		)
		{
			techIDRaw := params["tech_id"]
			v, err2 := strconv.ParseInt(techIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("techID", techIDRaw, "integer"))
			}
			techID = int(v)
		}
		if techID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("techID", techID, 1, true))
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeleteTechPayload(techID, token)
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

// EncodeDeleteTechError returns an encoder for errors returned by the
// deleteTech hy_tech endpoint.
func EncodeDeleteTechError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "NotFound":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewDeleteTechNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalHytechviewsTechViewToTechResponse builds a value of type
// *TechResponse from a value of type *hytechviews.TechView.
func marshalHytechviewsTechViewToTechResponse(v *hytechviews.TechView) *TechResponse {
	res := &TechResponse{
		TechID:   v.TechID,
		TechName: v.TechName,
	}

	return res
}

// marshalHytechviewsTechViewToTechResponseID builds a value of type
// *TechResponseID from a value of type *hytechviews.TechView.
func marshalHytechviewsTechViewToTechResponseID(v *hytechviews.TechView) *TechResponseID {
	res := &TechResponseID{
		TechID: v.TechID,
	}

	return res
}
