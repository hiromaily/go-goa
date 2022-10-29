// Code generated by goa v3.10.1, DO NOT EDIT.
//
// hy_userWorkHistory HTTP server
//
// Command:
// $ goa gen resume/design

package server

import (
	"context"
	"net/http"
	hyuserworkhistory "resume/gen/hy_user_work_history"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the hy_userWorkHistory service endpoint HTTP handlers.
type Server struct {
	Mounts             []*MountPoint
	GetUserWorkHistory http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the hy_userWorkHistory service
// endpoints using the provided encoder and decoder. The handlers are mounted
// on the given mux using the HTTP verb and path defined in the design.
// errhandler is called whenever a response fails to be encoded. formatter is
// used to format errors returned by the service methods prior to encoding.
// Both errhandler and formatter are optional and can be nil.
func New(
	e *hyuserworkhistory.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"GetUserWorkHistory", "GET", "/user/{user_id}/workhistory"},
		},
		GetUserWorkHistory: NewGetUserWorkHistoryHandler(e.GetUserWorkHistory, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "hy_userWorkHistory" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.GetUserWorkHistory = m(s.GetUserWorkHistory)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return hyuserworkhistory.MethodNames[:] }

// Mount configures the mux to serve the hy_userWorkHistory endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountGetUserWorkHistoryHandler(mux, h.GetUserWorkHistory)
}

// Mount configures the mux to serve the hy_userWorkHistory endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountGetUserWorkHistoryHandler configures the mux to serve the
// "hy_userWorkHistory" service "getUserWorkHistory" endpoint.
func MountGetUserWorkHistoryHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/user/{user_id}/workhistory", f)
}

// NewGetUserWorkHistoryHandler creates a HTTP handler which loads the HTTP
// request and calls the "hy_userWorkHistory" service "getUserWorkHistory"
// endpoint.
func NewGetUserWorkHistoryHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetUserWorkHistoryRequest(mux, decoder)
		encodeResponse = EncodeGetUserWorkHistoryResponse(encoder)
		encodeError    = EncodeGetUserWorkHistoryError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "getUserWorkHistory")
		ctx = context.WithValue(ctx, goa.ServiceKey, "hy_userWorkHistory")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
