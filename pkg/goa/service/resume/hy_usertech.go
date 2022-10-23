package resumeapi

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"goa.design/goa/v3/security"

	"github.com/hiromaily/go-goa/pkg/repository"
	hyusertech "resume/gen/hy_usertech"
)

// hy_usertech service example implementation.
// The example methods log the requests and return zero values.
type hyUsertechsrvc struct {
	userTechRepo repository.UserTechRepository
}

// NewHyUsertech returns the hy_usertech service implementation.
func NewHyUsertech(userTechRepo repository.UserTechRepository) hyusertech.Service {
	return &hyUsertechsrvc{userTechRepo}
}

// JWTAuth implements the authorization logic for service "hy_usertech" for the
// "jwt" security scheme.
func (s *hyUsertechsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
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

// get user's favorite techs
func (s *hyUsertechsrvc) GetUserLikeTech(ctx context.Context, p *hyusertech.GetUserLikeTechPayload) (res hyusertech.UsertechCollection, view string, err error) {
	//	//TODO:check user ID and this part should be set in middleware.
	//	//fmt.Println("[user id]:", ctx.Value("user.jwt"))
	//	if ctx.Value("user.jwt") == nil || u.Itoi(ctx.Value("user.jwt")) == 0 ||
	//		u.Itoi(ctx.Value("user.jwt")) != ctx.UserID {
	//		return ctx.Unauthorized()
	//	}
	//
	//	const TableTechLike = "t_user_like_techs"
	//
	//	//type UsertechTech struct {
	//	//	// Tech name
	//	//	TechName string `form:"tech_name" json:"tech_name" xml:"tech_name"`
	//	//}
	//	var userTechs []*app.UsertechTech
	//
	//	svc := &m.UserTech{Db: c.ctx.Db}
	//	err := svc.GetUserTechs(ctx.UserID, &userTechs, TableTechLike)
	//	if err != nil {
	//		return err
	//	}
	//
	//	if len(userTechs) == 0 {
	//		return ctx.NotFound()
	//	}
	//
	//	//type UserCollection []*User
	//	res := app.UsertechTechCollection(userTechs)
	//	return ctx.OKTech(res)
	view = "default"
	log.Info().Msg("hyUsertech.getUserLikeTech")
	return
}

// get user's dislike techs
func (s *hyUsertechsrvc) GetUserDisLikeTech(ctx context.Context, p *hyusertech.GetUserDisLikeTechPayload) (res hyusertech.UsertechCollection, view string, err error) {
	//	//TODO:check user ID and this part should be set in middleware.
	//	//fmt.Println("[user id]:", ctx.Value("user.jwt"))
	//	if ctx.Value("user.jwt") == nil || u.Itoi(ctx.Value("user.jwt")) == 0 ||
	//		u.Itoi(ctx.Value("user.jwt")) != ctx.UserID {
	//		return ctx.Unauthorized()
	//	}
	//
	//	const TableTechDislike = "t_user_dislike_techs"
	//
	//	var userTechs []*app.UsertechTech
	//
	//	svc := &m.UserTech{Db: c.ctx.Db}
	//	err := svc.GetUserTechs(ctx.UserID, &userTechs, TableTechDislike)
	//	if err != nil {
	//		return err
	//	}
	//
	//	if len(userTechs) == 0 {
	//		return ctx.NotFound()
	//	}
	//
	//	//type UserCollection []*User
	//	res := app.UsertechTechCollection(userTechs)
	//	return ctx.OKTech(res)
	view = "default"
	log.Info().Msg("hyUsertech.getUserDisLikeTech")
	return
}
