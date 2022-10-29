package resumeapi

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"goa.design/goa/v3/security"

	"github.com/hiromaily/go-goa/pkg/jwts"
	"github.com/hiromaily/go-goa/pkg/repository"
	hyusertech "resume/gen/hy_usertech"
)

// hy_usertech service example implementation.
// The example methods log the requests and return zero values.
type hyUsertechsrvc struct {
	jwt          jwts.JWTer
	userTechRepo repository.UserTechRepository
}

// NewHyUsertech returns the hy_usertech service implementation.
func NewHyUsertech(jwt jwts.JWTer, userTechRepo repository.UserTechRepository) hyusertech.Service {
	return &hyUsertechsrvc{jwt, userTechRepo}
}

// JWTAuth implements the authorization logic for service "hy_usertech" for the
// "jwt" security scheme.
func (s *hyUsertechsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	log.Info().
		Str("token", token).
		Strs("scheme.Scopes", scheme.Scopes).
		Msg("hyUsertech.JWTAuth")

	if err := s.jwt.ValidateToken(token); err != nil {
		return ctx, err
	}

	return ctx, nil
}

// GetUserLikeTech returns user's favorite techs
func (s *hyUsertechsrvc) GetUserLikeTech(ctx context.Context, p *hyusertech.GetUserLikeTechPayload) (res hyusertech.UsertechCollection, view string, err error) {
	log.Info().Msg("hyUsertech.getUserLikeTech")

	userTechs, err := s.userTechRepo.GetUserLikeTechs(p.UserID)
	if err != nil {
		return nil, "", errors.Wrap(err, "fail to call userTechs.GetUserLikeTechs()")
	}
	return s.getUserTech(userTechs)
}

// GetUserDisLikeTech returns user's unfavorite techs
func (s *hyUsertechsrvc) GetUserDisLikeTech(ctx context.Context, p *hyusertech.GetUserDisLikeTechPayload) (res hyusertech.UsertechCollection, view string, err error) {
	log.Info().Msg("hyUsertech.getUserDisLikeTech")

	userTechs, err := s.userTechRepo.GetUserDisLikeTechs(p.UserID)
	if err != nil {
		return nil, "", errors.Wrap(err, "fail to call userTechs.GetUserDisLikeTechs()")
	}
	return s.getUserTech(userTechs)
}

func (s *hyUsertechsrvc) getUserTech(userTechs []repository.UserTech) (res hyusertech.UsertechCollection, view string, err error) {
	log.Info().Msg("hyUsertech.getUserTech")

	if len(userTechs) == 0 {
		return nil, "", hyusertech.MakeNotFound(errors.New("user tech not found"))
	}
	// convert
	techs := make([]*hyusertech.Usertech, len(userTechs))
	for i, _ := range userTechs {
		techs[i] = &hyusertech.Usertech{
			ID:       &userTechs[i].ID,
			TechName: &userTechs[i].Name,
		}
	}
	res = techs
	view = "techName"

	return
}
