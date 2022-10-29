package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"

	"github.com/hiromaily/go-goa/pkg/config"
	"github.com/hiromaily/go-goa/pkg/logger"
	auth "resume/gen/auth"
	health "resume/gen/health"
	hycompany "resume/gen/hy_company"
	hytech "resume/gen/hy_tech"
	hyuser "resume/gen/hy_user"
	hyuserworkhistory "resume/gen/hy_user_work_history"
	hyusertech "resume/gen/hy_usertech"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
		confPath  = flag.String("conf", "", "config path")
	)
	flag.Parse()

	// Config
	conf, err := config.NewConfig(*confPath, false)
	if err != nil {
		panic(err)
	}
	// Logger for zerolog
	myLogger := logger.NewZeroLog(conf.Logger)
	// logger.NewZeroLog(conf.Logger)

	// Registry
	reg := NewRegistry(conf)

	// Initialize the services.
	authSvc := reg.NewAuth()
	hyUserSvc := reg.NewHyUser()
	hyCompanySvc := reg.NewHyCompany()
	hyTechSvc := reg.NewHyTech()
	hyUsertechSvc := reg.NewHyUserTech()
	healthSvc := reg.NewHealth()
	hyUserWorkHistorySvc := reg.NewHyUserWorkHistory()

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	authEndpoints := auth.NewEndpoints(authSvc)
	hyCompanyEndpoints := hycompany.NewEndpoints(hyCompanySvc)
	healthEndpoints := health.NewEndpoints(healthSvc)
	hyTechEndpoints := hytech.NewEndpoints(hyTechSvc)
	hyUserEndpoints := hyuser.NewEndpoints(hyUserSvc)
	hyUsertechEndpoints := hyusertech.NewEndpoints(hyUsertechSvc)
	hyUserWorkHistoryEndpoints := hyuserworkhistory.NewEndpoints(hyUserWorkHistorySvc)

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
			addr := "http://localhost:8090"
			u, err := url.Parse(addr)
			if err != nil {
				log.Fatal().Msg(fmt.Sprintf("invalid URL %s: %#v", addr, err))
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
					log.Fatal().Msg(fmt.Sprintf("invalid URL %s: %#v", u.Host, err))
				}
				u.Host = net.JoinHostPort(h, *httpPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "80")
			}
			handleHTTPServer(ctx, u, authEndpoints, hyCompanyEndpoints, healthEndpoints, hyTechEndpoints, hyUserEndpoints, hyUsertechEndpoints, hyUserWorkHistoryEndpoints, &wg, errc, *dbgF, myLogger)
		}

	default:
		log.Fatal().Msg(fmt.Sprintf("invalid host argument: %q (valid hosts: localhost)", *hostF))
	}

	// Wait for signal.
	log.Info().Msg(fmt.Sprintf("exiting (%v)", <-errc))

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	log.Info().Msg("exited")
}
