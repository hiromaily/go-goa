package main

import (
	"database/sql"
	"github.com/hiromaily/go-goa/pkg/mysql"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"resume/gen/auth"
	hycompany "resume/gen/hy_company"

	"github.com/hiromaily/go-goa/pkg/config"
	"github.com/hiromaily/go-goa/pkg/encryption"
	resume "github.com/hiromaily/go-goa/pkg/goa/service/resume"
	"github.com/hiromaily/go-goa/pkg/jwts"
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

// TODO: WIP
func (r *registry) NewAuth() auth.Service {
	return resume.NewAuth(r.userRepo)
}

// TODO: WIP
func (r *registry) NewHyCompany() hycompany.Service {
	return resume.NewHyCompany(r.companyRepo)
}

//hyCompanybranchSvc = resume.NewHyCompanybranch(logger, db)
//healthSvc = resume.NewHealth(logger, db)
//hyTechSvc = resume.NewHyTech(logger, db)
//hyUserSvc = resume.NewHyUser(logger, db)
//hyUsertechSvc = resume.NewHyUsertech(logger, db)
//hyUserWorkHistorySvc = resume.NewHyUserWorkHistory(logger, db)

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
