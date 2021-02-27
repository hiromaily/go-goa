package resumeapi

import (
	"context"
	"fmt"
	"log"
	hyuserworkhistory "resume/gen/hy_user_work_history"

	"goa.design/goa/v3/security"
)

// hy_userWorkHistory service example implementation.
// The example methods log the requests and return zero values.
type hyUserWorkHistorysrvc struct {
	logger *log.Logger
}

// NewHyUserWorkHistory returns the hy_userWorkHistory service implementation.
func NewHyUserWorkHistory(logger *log.Logger) hyuserworkhistory.Service {
	return &hyUserWorkHistorysrvc{logger}
}

// JWTAuth implements the authorization logic for service "hy_userWorkHistory"
// for the "jwt" security scheme.
func (s *hyUserWorkHistorysrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	//
	// TBD: add authorization logic.
	//
	// In case of authorization failure this function should return
	// one of the generated error structs, e.g.:
	//
	//    return ctx, myservice.MakeUnauthorizedError("invalid token")
	//
	// Alternatively this function may return an instance of
	// goa.ServiceError with a Name field value that matches one of
	// the design error names, e.g:
	//
	//    return ctx, goa.PermanentError("unauthorized", "invalid token")
	//
	return ctx, fmt.Errorf("not implemented")
}

// get user's work history
func (s *hyUserWorkHistorysrvc) GetUserWorkHistory(ctx context.Context, p *hyuserworkhistory.GetUserWorkHistoryPayload) (res hyuserworkhistory.UserworkhistoryCollection, err error) {
	s.logger.Print("hyUserWorkHistory.getUserWorkHistory")
	return
}
