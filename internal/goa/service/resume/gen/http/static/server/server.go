// Code generated by goa v3.10.1, DO NOT EDIT.
//
// static HTTP server
//
// Command:
// $ goa gen resume/design

package server

import (
	"context"
	"net/http"
	static "resume/gen/static"

	goahttp "goa.design/goa/v3/http"
	"goa.design/plugins/v3/cors"
)

// Server lists the static service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	CORS   http.Handler
	Assets http.Handler
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

// New instantiates HTTP handlers for all the static service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *static.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
	fileSystemAssets http.FileSystem,
) *Server {
	if fileSystemAssets == nil {
		fileSystemAssets = http.Dir(".")
	}
	return &Server{
		Mounts: []*MountPoint{
			{"CORS", "OPTIONS", "/assets/{*filepath}"},
			{"assets/", "GET", "/assets"},
		},
		CORS:   NewCORSHandler(),
		Assets: http.FileServer(fileSystemAssets),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "static" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.CORS = m(s.CORS)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return static.MethodNames[:] }

// Mount configures the mux to serve the static endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountCORSHandler(mux, h.CORS)
	MountAssets(mux, goahttp.Replace("/assets", "/assets/", h.Assets))
}

// Mount configures the mux to serve the static endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountAssets configures the mux to serve GET request made to "/assets".
func MountAssets(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/assets/", HandleStaticOrigin(h).ServeHTTP)
	mux.Handle("GET", "/assets/*filepath", HandleStaticOrigin(h).ServeHTTP)
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service static.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleStaticOrigin(h)
	mux.Handle("OPTIONS", "/assets/{*filepath}", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleStaticOrigin applies the CORS response headers corresponding to the
// origin for the service static.
func HandleStaticOrigin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}
