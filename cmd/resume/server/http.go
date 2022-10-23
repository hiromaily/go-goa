package main

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"

	"github.com/hiromaily/go-goa/pkg/logger"
	auth "resume/gen/auth"
	health "resume/gen/health"
	authsvr "resume/gen/http/auth/server"
	healthsvr "resume/gen/http/health/server"
	hycompanysvr "resume/gen/http/hy_company/server"
	hytechsvr "resume/gen/http/hy_tech/server"
	hyusersvr "resume/gen/http/hy_user/server"
	hyuserworkhistorysvr "resume/gen/http/hy_user_work_history/server"
	hyusertechsvr "resume/gen/http/hy_usertech/server"
	hycompany "resume/gen/hy_company"
	hytech "resume/gen/hy_tech"
	hyuser "resume/gen/hy_user"
	hyuserworkhistory "resume/gen/hy_user_work_history"
	hyusertech "resume/gen/hy_usertech"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, authEndpoints *auth.Endpoints, hyCompanyEndpoints *hycompany.Endpoints, healthEndpoints *health.Endpoints, hyTechEndpoints *hytech.Endpoints, hyUserEndpoints *hyuser.Endpoints, hyUsertechEndpoints *hyusertech.Endpoints, hyUserWorkHistoryEndpoints *hyuserworkhistory.Endpoints, wg *sync.WaitGroup, errc chan error, debug bool, logger logger.Logger) {
	goaLogger := log.New(os.Stderr, "[resume] ", log.Ltime)

	// Setup goa log adapter.
	adapter := middleware.NewLogger(goaLogger)

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	mux := goahttp.NewMuxer()

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		authServer              *authsvr.Server
		hyCompanyServer         *hycompanysvr.Server
		healthServer            *healthsvr.Server
		hyTechServer            *hytechsvr.Server
		hyUserServer            *hyusersvr.Server
		hyUsertechServer        *hyusertechsvr.Server
		hyUserWorkHistoryServer *hyuserworkhistorysvr.Server
	)
	{
		eh := errorHandler(logger)
		authServer = authsvr.New(authEndpoints, mux, dec, enc, eh, nil)
		hyCompanyServer = hycompanysvr.New(hyCompanyEndpoints, mux, dec, enc, eh, nil)
		healthServer = healthsvr.New(healthEndpoints, mux, dec, enc, eh, nil)
		hyTechServer = hytechsvr.New(hyTechEndpoints, mux, dec, enc, eh, nil)
		hyUserServer = hyusersvr.New(hyUserEndpoints, mux, dec, enc, eh, nil)
		hyUsertechServer = hyusertechsvr.New(hyUsertechEndpoints, mux, dec, enc, eh, nil)
		hyUserWorkHistoryServer = hyuserworkhistorysvr.New(hyUserWorkHistoryEndpoints, mux, dec, enc, eh, nil)
		if debug {
			servers := goahttp.Servers{
				authServer,
				hyCompanyServer,
				healthServer,
				hyTechServer,
				hyUserServer,
				hyUsertechServer,
				hyUserWorkHistoryServer,
			}
			servers.Use(httpmdlwr.Debug(mux, os.Stdout))
		}
	}
	// Configure the mux.
	authsvr.Mount(mux, authServer)
	hycompanysvr.Mount(mux, hyCompanyServer)
	healthsvr.Mount(mux, healthServer)
	hytechsvr.Mount(mux, hyTechServer)
	hyusersvr.Mount(mux, hyUserServer)
	hyusertechsvr.Mount(mux, hyUsertechServer)
	hyuserworkhistorysvr.Mount(mux, hyUserWorkHistoryServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler, ReadHeaderTimeout: time.Second * 60}
	for _, m := range authServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range hyCompanyServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range healthServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range hyTechServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range hyUserServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range hyUsertechServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range hyUserWorkHistoryServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Printf("HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Printf("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			logger.Printf("failed to shutdown: %v", err)
		}
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger logger.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
