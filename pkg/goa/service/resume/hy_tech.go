package resumeapi

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"goa.design/goa/v3/security"

	"github.com/hiromaily/go-goa/pkg/jwts"
	ptr "github.com/hiromaily/go-goa/pkg/pointer"
	"github.com/hiromaily/go-goa/pkg/repository"
	hytech "resume/gen/hy_tech"
)

// hy_tech service example implementation.
// The example methods log the requests and return zero values.
type hyTechsrvc struct {
	jwt      jwts.JWTer
	techRepo repository.TechRepository
}

// NewHyTech returns the hy_tech service implementation.
func NewHyTech(jwt jwts.JWTer, techRepo repository.TechRepository) hytech.Service {
	return &hyTechsrvc{jwt, techRepo}
}

// JWTAuth implements the authorization logic for service "hy_tech" for the
// "jwt" security scheme.
func (s *hyTechsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	log.Info().
		Str("token", token).
		Strs("scheme.Scopes", scheme.Scopes).
		Msg("hyTech.JWTAuth")

	if err := s.jwt.ValidateToken(token); err != nil {
		return ctx, err
	}

	return ctx, nil
}

// TechList returns all techs
func (s *hyTechsrvc) TechList(ctx context.Context, p *hytech.TechListPayload) (res hytech.TechCollection, view string, err error) {
	log.Info().Msg("hyTech.techList")

	techs, err := s.techRepo.TechList()
	if err != nil {
		return nil, "", errors.Wrap(err, "fail to call techRepo.UserList()")
	}
	if len(techs) == 0 {
		return nil, "", hytech.MakeNotFound(errors.New("tech not found"))
	}
	res = techs
	view = "default"

	return
}

// GetTech　returns tech by given tech id
func (s *hyTechsrvc) GetTech(ctx context.Context, p *hytech.GetTechPayload) (res *hytech.Tech, view string, err error) {
	log.Info().Msg("hyTech.getTech")

	tech, err := s.techRepo.GetTech(p.TechID)
	if err != nil {
		// Not Found
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, "", hytech.MakeNotFound(errors.New("tech not found"))
		}
		return nil, "", errors.Wrapf(err, "fail to call techRepo.GetTech(%d)", p.TechID)
	}
	res = tech
	view = "default"

	return
}

// CreateTech creates new tech
// FIXME: only admin user can create user
func (s *hyTechsrvc) CreateTech(ctx context.Context, p *hytech.CreateTechPayload) (res *hytech.Tech, view string, err error) {
	log.Info().Msg("hyTech.createTech")

	techID, err := s.techRepo.InsertTech(p.TechName)
	if err != nil {
		return nil, "", errors.Wrap(err, "fail to call InsertTech()")
	}
	res = &hytech.Tech{
		TechID: ptr.Int(techID),
	}
	view = "id"

	return
}

// UpdateTech updates tech data
// FIXME: only login user can update own information
func (s *hyTechsrvc) UpdateTech(ctx context.Context, p *hytech.UpdateTechPayload) (err error) {
	log.Info().Msg("hyTech.updateTech")

	// Validate
	if p.TechName == "" {
		return hytech.MakeBadRequest(errors.New("parameter is invalid"))
	}

	isTech, err := s.techRepo.IsTech(p.TechID)
	if !isTech || err != nil {
		return hytech.MakeNotFound(errors.New("tech not found"))
	}

	_, err = s.techRepo.UpdateTech(p.TechID, p.TechName)
	if err != nil {
		return errors.Wrap(err, "fail to call UpdateTech()")
	}

	return
}

// DeleteTech deletes tech data
// FIXME: only admin user can delete user
func (s *hyTechsrvc) DeleteTech(ctx context.Context, p *hytech.DeleteTechPayload) (err error) {
	log.Info().Msg("hyTech.deleteTech")

	isTech, err := s.techRepo.IsTech(p.TechID)
	if !isTech || err != nil {
		return hytech.MakeNotFound(errors.New("tech not found"))
	}

	if _, err := s.techRepo.DeleteTech(p.TechID); err != nil {
		return errors.Wrap(err, "fail to call DeleteTech()")
	}

	return
}
