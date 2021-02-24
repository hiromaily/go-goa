package service

import (
	"context"
	"log"

	"github.com/hiromaily/go-goa/internal/goa/service/resume/gen/health"
)

// health service example implementation.
// The example methods log the requests and return zero values.
type healthsrvc struct {
	logger *log.Logger
}

// NewHealth returns the health service implementation.
func NewHealth(logger *log.Logger) health.Service {
	return &healthsrvc{logger}
}

// Health implements health.
func (s *healthsrvc) Health(ctx context.Context) (err error) {
	s.logger.Print("health.health")
	return
}
