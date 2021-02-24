package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/hiromaily/go-goa/internal/goa/service/resume/gen/auth"
	"github.com/hiromaily/go-goa/internal/goa/service/resume/gen/health"
	hycompany "github.com/hiromaily/go-goa/internal/goa/service/resume/gen/hy_company"
	hycompanybranch "github.com/hiromaily/go-goa/internal/goa/service/resume/gen/hy_companybranch"
	hytech "github.com/hiromaily/go-goa/internal/goa/service/resume/gen/hy_tech"
	hyuser "github.com/hiromaily/go-goa/internal/goa/service/resume/gen/hy_user"
	hyuserworkhistory "github.com/hiromaily/go-goa/internal/goa/service/resume/gen/hy_user_work_history"
	hyusertech "github.com/hiromaily/go-goa/internal/goa/service/resume/gen/hy_usertech"
	"github.com/hiromaily/go-goa/pkg/goa/service"
)

// replace
// resume/example to
// github.com/hiromaily/go-goa/internal/goa/service/resume

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[resumeapi] ", log.Ltime)
	}

	// Initialize the services.
	var (
		authSvc              auth.Service
		hyCompanySvc         hycompany.Service
		hyCompanybranchSvc   hycompanybranch.Service
		healthSvc            health.Service
		hyTechSvc            hytech.Service
		hyUserSvc            hyuser.Service
		hyUsertechSvc        hyusertech.Service
		hyUserWorkHistorySvc hyuserworkhistory.Service
	)
	{
		authSvc = service.NewAuth(logger)
		hyCompanySvc = service.NewHyCompany(logger)
		hyCompanybranchSvc = service.NewHyCompanybranch(logger)
		healthSvc = service.NewHealth(logger)
		hyTechSvc = service.NewHyTech(logger)
		hyUserSvc = service.NewHyUser(logger)
		hyUsertechSvc = service.NewHyUsertech(logger)
		hyUserWorkHistorySvc = service.NewHyUserWorkHistory(logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		authEndpoints              *auth.Endpoints
		hyCompanyEndpoints         *hycompany.Endpoints
		hyCompanybranchEndpoints   *hycompanybranch.Endpoints
		healthEndpoints            *health.Endpoints
		hyTechEndpoints            *hytech.Endpoints
		hyUserEndpoints            *hyuser.Endpoints
		hyUsertechEndpoints        *hyusertech.Endpoints
		hyUserWorkHistoryEndpoints *hyuserworkhistory.Endpoints
	)
	{
		authEndpoints = auth.NewEndpoints(authSvc)
		hyCompanyEndpoints = hycompany.NewEndpoints(hyCompanySvc)
		hyCompanybranchEndpoints = hycompanybranch.NewEndpoints(hyCompanybranchSvc)
		healthEndpoints = health.NewEndpoints(healthSvc)
		hyTechEndpoints = hytech.NewEndpoints(hyTechSvc)
		hyUserEndpoints = hyuser.NewEndpoints(hyUserSvc)
		hyUsertechEndpoints = hyusertech.NewEndpoints(hyUsertechSvc)
		hyUserWorkHistoryEndpoints = hyuserworkhistory.NewEndpoints(hyUserWorkHistorySvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://localhost:8080/api"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", u.Host, err)
					os.Exit(1)
				}
				u.Host = net.JoinHostPort(h, *httpPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, ":80")
			}
			handleHTTPServer(ctx, u, authEndpoints, hyCompanyEndpoints, hyCompanybranchEndpoints, healthEndpoints, hyTechEndpoints, hyUserEndpoints, hyUsertechEndpoints, hyUserWorkHistoryEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: localhost)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
