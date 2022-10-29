package resumeapi

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"goa.design/goa/v3/security"

	"github.com/hiromaily/go-goa/pkg/jwts"
	"github.com/hiromaily/go-goa/pkg/repository"
	hyuserworkhistory "resume/gen/hy_user_work_history"
)

// hy_userWorkHistory service example implementation.
// The example methods log the requests and return zero values.
type hyUserWorkHistorysrvc struct {
	jwt                 jwts.JWTer
	userWorkHistoryRepo repository.UserWorkHistoryRepository
}

// NewHyUserWorkHistory returns the hy_userWorkHistory service implementation.
func NewHyUserWorkHistory(jwt jwts.JWTer, userWorkHistoryRepo repository.UserWorkHistoryRepository) hyuserworkhistory.Service {
	return &hyUserWorkHistorysrvc{jwt, userWorkHistoryRepo}
}

// JWTAuth implements the authorization logic for service "hy_userWorkHistory"
// for the "jwt" security scheme.
func (s *hyUserWorkHistorysrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	log.Info().
		Str("token", token).
		Strs("scheme.Scopes", scheme.Scopes).
		Msg("hyUserWorkHistory.JWTAuth")

	if err := s.jwt.ValidateToken(token); err != nil {
		return ctx, err
	}

	return ctx, nil
}

// GetUserWorkHistory returns user's working history
func (s *hyUserWorkHistorysrvc) GetUserWorkHistory(ctx context.Context, p *hyuserworkhistory.GetUserWorkHistoryPayload) (res hyuserworkhistory.UserworkhistoryCollection, err error) {
	log.Info().Msg("hyUserWorkHistory.getUserWorkHistory")

	userHistory, err := s.userWorkHistoryRepo.GetUserWorks(p.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "fail to call userTechs.GetUserWorks()")
	}

	if len(userHistory) == 0 {
		return nil, hyuserworkhistory.MakeNotFound(errors.New("user work history not found"))
	}
	res = userHistory

	return
}
