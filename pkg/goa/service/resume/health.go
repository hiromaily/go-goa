package resumeapi

import (
	"context"

	"github.com/rs/zerolog/log"

	health "resume/gen/health"
)

// health service example implementation.
// The example methods log the requests and return zero values.
type healthsrvc struct{}

// NewHealth returns the health service implementation.
func NewHealth() health.Service {
	return &healthsrvc{}
}

// Health implements health.
func (s *healthsrvc) Health(ctx context.Context) (err error) {
	log.Info().Msg("health.health")
	return
}
