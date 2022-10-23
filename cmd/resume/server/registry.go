package main

import (
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"resume/gen/auth"
	"resume/gen/health"
	hycompany "resume/gen/hy_company"
	hytech "resume/gen/hy_tech"
	hyuser "resume/gen/hy_user"
	hyusertech "resume/gen/hy_usertech"

	"github.com/hiromaily/go-goa/pkg/config"
	"github.com/hiromaily/go-goa/pkg/encryption"
	resume "github.com/hiromaily/go-goa/pkg/goa/service/resume"
	"github.com/hiromaily/go-goa/pkg/jwts"
	"github.com/hiromaily/go-goa/pkg/mysql"
	"github.com/hiromaily/go-goa/pkg/repository"
)

type registry struct {
	conf         *config.Root
	jwter        jwts.JWTer
	db           *sql.DB
	hasher       encryption.Hasher
	userRepo     repository.UserRepository
	techRepo     repository.TechRepository
	companyRepo  repository.CompanyRepository
	userTechRepo repository.UserTechRepository
}

func NewRegistry(confPath string) *registry {
	reg := &registry{}

	// JWT
	reg.newJWT()

	// Database
	reg.newMySQLClient()

	// Create repository
	// TODO: create proper repository for each services
	reg.newUserRepo()
	reg.newTechRepo()
	reg.newCompanyRepo()

	return reg
}

func (r *registry) NewAuth() auth.Service {
	return resume.NewAuth(r.userRepo)
}

func (r *registry) NewHyUser() hyuser.Service {
	return resume.NewHyUser(r.userRepo)
}

func (r *registry) NewHyCompany() hycompany.Service {
	return resume.NewHyCompany(r.companyRepo)
}

func (r *registry) NewHyTech() hytech.Service {
	return resume.NewHyTech(r.techRepo)
}

func (r *registry) NewHyUserTech() hyusertech.Service {
	return resume.NewHyUsertech(r.userTechRepo)
}

func (r *registry) NewHealth() health.Service {
	return resume.NewHealth()
}

// TODO
// hyUserWorkHistorySvc = resume.NewHyUserWorkHistory(logger, db)

//func (r *registry) newLogger() *zap.Logger {
//	if r.logger == nil {
//		r.logger = logger.NewZapLogger(r.conf.Logger)
//	}
//	return r.logger
//}

func (r *registry) newJWT() jwts.JWTer {
	if r.jwter == nil {
		auth := r.conf.Jwt
		var signAlgo jwts.SigAlgoer
		if auth.Mode == jwts.AlgoHMAC && auth.Secret != "" {
			signAlgo = jwts.NewHMAC(auth.Secret)
		} else if auth.Mode == jwts.AlgoRSA && auth.PrivateKey != "" && auth.PublicKey != "" {
			var err error
			signAlgo, err = jwts.NewRSA(auth.PrivateKey, auth.PublicKey)
			if err != nil {
				panic(err)
			}
		} else {
			panic(errors.New("invalid jwt config"))
		}
		r.jwter = jwts.NewJWT(auth.Audience, signAlgo)
	}
	return r.jwter
}

func (r *registry) newMySQLClient() *sql.DB {
	if r.db == nil {
		dbConn, err := mysql.NewMySQL(r.conf.MySQL)
		if err != nil {
			panic(err)
		}
		r.db = dbConn
	}
	if r.conf.MySQL.Debug {
		boil.DebugMode = true
	}
	return r.db
}

func (r *registry) newHasher() encryption.Hasher {
	if r.hasher == nil {
		// TODO: define hash salt in config
		r.hasher = encryption.NewScrypt("salt")
	}
	return r.hasher
}

func (r *registry) newUserRepo() repository.UserRepository {
	if r.userRepo == nil {
		r.userRepo = repository.NewUserRepository(
			r.newMySQLClient(),
			r.newHasher(),
		)
	}
	return r.userRepo
}

func (r *registry) newTechRepo() repository.TechRepository {
	if r.techRepo == nil {
		r.techRepo = repository.NewTechRepository(
			r.newMySQLClient(),
		)
	}
	return r.techRepo
}

func (r *registry) newCompanyRepo() repository.CompanyRepository {
	if r.companyRepo == nil {
		r.companyRepo = repository.NewCompanyRepository(
			r.newMySQLClient(),
		)
	}
	return r.companyRepo
}

func (r *registry) newUesrTechRepo() repository.UserTechRepository {
	if r.userTechRepo == nil {
		r.userTechRepo = repository.NewUserTechRepository(
			r.newMySQLClient(),
		)
	}
	return r.userTechRepo
}
