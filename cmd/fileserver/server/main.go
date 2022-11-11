package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/pkg/errors"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF = flag.String("host", "localhost", "Server host (valid values: localhost)")
		//domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		portF   = flag.String("port", "8080", "HTTP port (overrides host HTTP port specified in service design)")
		secureF = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF    = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var logger *log.Logger
	logger = log.New(os.Stderr, "[resume] ", log.Ltime)

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
	u, err := getURL(*hostF, *portF, *secureF)
	if err != nil {
		logger.Fatal(err.Error())
	}
	handleHTTPServer(ctx, u, &wg, errc, logger, *dbgF)

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}

func getURL(host, port string, isSecure bool) (*url.URL, error) {
	addr := fmt.Sprintf("http://%s:%s", host, port)
	u, err := url.Parse(addr)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid URL %s: %#v", addr)
	}
	if isSecure {
		u.Scheme = "https"
	}
	return u, nil
}
