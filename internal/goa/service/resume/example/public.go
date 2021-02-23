package resumeapi

import (
	"log"
	public "resume/example/gen/public"
)

// public service example implementation.
// The example methods log the requests and return zero values.
type publicsrvc struct {
	logger *log.Logger
}

// NewPublic returns the public service implementation.
func NewPublic(logger *log.Logger) public.Service {
	return &publicsrvc{logger}
}
