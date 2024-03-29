// Code generated by goa v3.10.2, DO NOT EDIT.
//
// hy_user HTTP server
//
// Command:
// $ goa gen resume/design

package server

import (
	"context"
	"net/http"
	"regexp"
	hyuser "resume/gen/hy_user"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the hy_user service endpoint HTTP handlers.
type Server struct {
	Mounts     []*MountPoint
	UserList   http.Handler
	GetUser    http.Handler
	CreateUser http.Handler
	UpdateUser http.Handler
	DeleteUser http.Handler
	CORS       http.Handler
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

// New instantiates HTTP handlers for all the hy_user service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *hyuser.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"UserList", "GET", "/api/user"},
			{"GetUser", "GET", "/api/user/{user_id}"},
			{"CreateUser", "POST", "/api/user"},
			{"UpdateUser", "PUT", "/api/user/{user_id}"},
			{"DeleteUser", "DELETE", "/api/user/{user_id}"},
			{"CORS", "OPTIONS", "/api/user"},
			{"CORS", "OPTIONS", "/api/user/{user_id}"},
		},
		UserList:   NewUserListHandler(e.UserList, mux, decoder, encoder, errhandler, formatter),
		GetUser:    NewGetUserHandler(e.GetUser, mux, decoder, encoder, errhandler, formatter),
		CreateUser: NewCreateUserHandler(e.CreateUser, mux, decoder, encoder, errhandler, formatter),
		UpdateUser: NewUpdateUserHandler(e.UpdateUser, mux, decoder, encoder, errhandler, formatter),
		DeleteUser: NewDeleteUserHandler(e.DeleteUser, mux, decoder, encoder, errhandler, formatter),
		CORS:       NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "hy_user" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.UserList = m(s.UserList)
	s.GetUser = m(s.GetUser)
	s.CreateUser = m(s.CreateUser)
	s.UpdateUser = m(s.UpdateUser)
	s.DeleteUser = m(s.DeleteUser)
	s.CORS = m(s.CORS)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return hyuser.MethodNames[:] }

// Mount configures the mux to serve the hy_user endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountUserListHandler(mux, h.UserList)
	MountGetUserHandler(mux, h.GetUser)
	MountCreateUserHandler(mux, h.CreateUser)
	MountUpdateUserHandler(mux, h.UpdateUser)
	MountDeleteUserHandler(mux, h.DeleteUser)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the hy_user endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountUserListHandler configures the mux to serve the "hy_user" service
// "userList" endpoint.
func MountUserListHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleHyUserOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/user", f)
}

// NewUserListHandler creates a HTTP handler which loads the HTTP request and
// calls the "hy_user" service "userList" endpoint.
func NewUserListHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUserListRequest(mux, decoder)
		encodeResponse = EncodeUserListResponse(encoder)
		encodeError    = EncodeUserListError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "userList")
		ctx = context.WithValue(ctx, goa.ServiceKey, "hy_user")
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

// MountGetUserHandler configures the mux to serve the "hy_user" service
// "getUser" endpoint.
func MountGetUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleHyUserOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/user/{user_id}", f)
}

// NewGetUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "hy_user" service "getUser" endpoint.
func NewGetUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetUserRequest(mux, decoder)
		encodeResponse = EncodeGetUserResponse(encoder)
		encodeError    = EncodeGetUserError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "getUser")
		ctx = context.WithValue(ctx, goa.ServiceKey, "hy_user")
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

// MountCreateUserHandler configures the mux to serve the "hy_user" service
// "createUser" endpoint.
func MountCreateUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleHyUserOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/api/user", f)
}

// NewCreateUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "hy_user" service "createUser" endpoint.
func NewCreateUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateUserRequest(mux, decoder)
		encodeResponse = EncodeCreateUserResponse(encoder)
		encodeError    = EncodeCreateUserError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "createUser")
		ctx = context.WithValue(ctx, goa.ServiceKey, "hy_user")
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

// MountUpdateUserHandler configures the mux to serve the "hy_user" service
// "updateUser" endpoint.
func MountUpdateUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleHyUserOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/api/user/{user_id}", f)
}

// NewUpdateUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "hy_user" service "updateUser" endpoint.
func NewUpdateUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateUserRequest(mux, decoder)
		encodeResponse = EncodeUpdateUserResponse(encoder)
		encodeError    = EncodeUpdateUserError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "updateUser")
		ctx = context.WithValue(ctx, goa.ServiceKey, "hy_user")
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

// MountDeleteUserHandler configures the mux to serve the "hy_user" service
// "deleteUser" endpoint.
func MountDeleteUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleHyUserOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/api/user/{user_id}", f)
}

// NewDeleteUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "hy_user" service "deleteUser" endpoint.
func NewDeleteUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteUserRequest(mux, decoder)
		encodeResponse = EncodeDeleteUserResponse(encoder)
		encodeError    = EncodeDeleteUserError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "deleteUser")
		ctx = context.WithValue(ctx, goa.ServiceKey, "hy_user")
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

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service hy_user.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleHyUserOrigin(h)
	mux.Handle("OPTIONS", "/api/user", h.ServeHTTP)
	mux.Handle("OPTIONS", "/api/user/{user_id}", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleHyUserOrigin applies the CORS response headers corresponding to the
// origin for the service hy_user.
func HandleHyUserOrigin(h http.Handler) http.Handler {
	spec0 := regexp.MustCompile("localhost")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Max-Age", "600")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "*")
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}
