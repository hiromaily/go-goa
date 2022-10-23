package resumeapi

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"goa.design/goa/v3/security"

	"github.com/hiromaily/go-goa/pkg/repository"
	hyuserworkhistory "resume/gen/hy_user_work_history"
)

// hy_userWorkHistory service example implementation.
// The example methods log the requests and return zero values.
type hyUserWorkHistorysrvc struct {
	userWorkHistoryRepo repository.UserWorkHistoryRepository
}

// NewHyUserWorkHistory returns the hy_userWorkHistory service implementation.
func NewHyUserWorkHistory(userWorkHistoryRepo repository.UserWorkHistoryRepository) hyuserworkhistory.Service {
	return &hyUserWorkHistorysrvc{userWorkHistoryRepo}
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
	//	//TODO:check user ID and this part should be set in middleware.
	//	//fmt.Println("[user id]:", ctx.Value("user.jwt"))
	//	if ctx.Value("user.jwt") == nil || u.Itoi(ctx.Value("user.jwt")) == 0 ||
	//		u.Itoi(ctx.Value("user.jwt")) != ctx.UserID {
	//		return ctx.Unauthorized()
	//	}
	//
	//	var userWorks []*app.Userworkhistory
	//
	//	svc := &m.UserWorkHistory{Db: c.ctx.Db}
	//	err := svc.GetUserWorks(ctx.UserID, &userWorks)
	//	if err != nil {
	//		return err
	//	}
	//
	//	if len(userWorks) == 0 {
	//		return ctx.NotFound()
	//	}
	//
	//	res := app.UserworkhistoryCollection(userWorks)
	//	return ctx.OK(res)
	log.Info().Msg("hyUserWorkHistory.getUserWorkHistory")
	return
}
