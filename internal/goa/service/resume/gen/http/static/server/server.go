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
)

// Server lists the static service endpoint HTTP handlers.
type Server struct {
	Mounts                                      []*MountPoint
	Docs                                        http.Handler
	InternalGoaServiceResumeGenHTTPOpenapiJSON  http.Handler
	InternalGoaServiceResumeGenHTTPOpenapi3JSON http.Handler
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
	fileSystemDocs http.FileSystem,
	fileSystemInternalGoaServiceResumeGenHTTPOpenapiJSON http.FileSystem,
	fileSystemInternalGoaServiceResumeGenHTTPOpenapi3JSON http.FileSystem,
) *Server {
	if fileSystemDocs == nil {
		fileSystemDocs = http.Dir(".")
	}
	if fileSystemInternalGoaServiceResumeGenHTTPOpenapiJSON == nil {
		fileSystemInternalGoaServiceResumeGenHTTPOpenapiJSON = http.Dir(".")
	}
	if fileSystemInternalGoaServiceResumeGenHTTPOpenapi3JSON == nil {
		fileSystemInternalGoaServiceResumeGenHTTPOpenapi3JSON = http.Dir(".")
	}
	return &Server{
		Mounts: []*MountPoint{
			{"./docs/", "GET", "/"},
			{"./internal/goa/service/resume/gen/http/openapi.json", "GET", "/openapi.json"},
			{"./internal/goa/service/resume/gen/http/openapi3.json", "GET", "/openapi3.json"},
		},
		Docs: http.FileServer(fileSystemDocs),
		InternalGoaServiceResumeGenHTTPOpenapiJSON:  http.FileServer(fileSystemInternalGoaServiceResumeGenHTTPOpenapiJSON),
		InternalGoaServiceResumeGenHTTPOpenapi3JSON: http.FileServer(fileSystemInternalGoaServiceResumeGenHTTPOpenapi3JSON),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "static" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return static.MethodNames[:] }

// Mount configures the mux to serve the static endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountDocs(mux, goahttp.Replace("/", "/./docs/", h.Docs))
	MountInternalGoaServiceResumeGenHTTPOpenapiJSON(mux, goahttp.Replace("", "/./internal/goa/service/resume/gen/http/openapi.json", h.InternalGoaServiceResumeGenHTTPOpenapiJSON))
	MountInternalGoaServiceResumeGenHTTPOpenapi3JSON(mux, goahttp.Replace("", "/./internal/goa/service/resume/gen/http/openapi3.json", h.InternalGoaServiceResumeGenHTTPOpenapi3JSON))
}

// Mount configures the mux to serve the static endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountDocs configures the mux to serve GET request made to "/".
func MountDocs(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/", h.ServeHTTP)
	mux.Handle("GET", "/*filepath", h.ServeHTTP)
}

// MountInternalGoaServiceResumeGenHTTPOpenapiJSON configures the mux to serve
// GET request made to "/openapi.json".
func MountInternalGoaServiceResumeGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi.json", h.ServeHTTP)
}

// MountInternalGoaServiceResumeGenHTTPOpenapi3JSON configures the mux to serve
// GET request made to "/openapi3.json".
func MountInternalGoaServiceResumeGenHTTPOpenapi3JSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi3.json", h.ServeHTTP)
}
