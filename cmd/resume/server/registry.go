package main

import (
	"database/sql"
	"errors"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/hiromaily/go-goa/pkg/config"
	"github.com/hiromaily/go-goa/pkg/encryption"
	resume "github.com/hiromaily/go-goa/pkg/goa/service/resume"
	"github.com/hiromaily/go-goa/pkg/jwts"
	"github.com/hiromaily/go-goa/pkg/mysql"
	"github.com/hiromaily/go-goa/pkg/repository"
	"resume/gen/auth"
	"resume/gen/health"
	hycompany "resume/gen/hy_company"
	hytech "resume/gen/hy_tech"
	hyuser "resume/gen/hy_user"
	hyuserworkhistory "resume/gen/hy_user_work_history"
	hyusertech "resume/gen/hy_usertech"
)

type registry struct {
	conf                *config.Root
	jwter               jwts.JWTer
	db                  *sql.DB
	hasher              encryption.Hasher
	userRepo            repository.UserRepository
	techRepo            repository.TechRepository
	companyRepo         repository.CompanyRepository
	userTechRepo        repository.UserTechRepository
	userWorkHistoryRepo repository.UserWorkHistoryRepository
}

func NewRegistry(conf *config.Root) *registry {
	if conf == nil {
		panic(errors.New("conf is invalid"))
	}

	reg := &registry{conf: conf}

	// JWT
	reg.newJWT()

	// Database
	reg.newMySQLClient()

	// Create repository
	// reg.newUserRepo()
	// reg.newCompanyRepo()
	// reg.newTechRepo()
	// reg.newUserTechRepo()
	// reg.newUserWorkHistoryRepo()

	return reg
}

func (r *registry) NewAuth() auth.Service {
	return resume.NewAuth(r.newJWT(), r.newUserRepo())
}

func (r *registry) NewHyUser() hyuser.Service {
	return resume.NewHyUser(r.newJWT(), r.newUserRepo())
}

func (r *registry) NewHyCompany() hycompany.Service {
	return resume.NewHyCompany(r.newJWT(), r.newCompanyRepo())
}

func (r *registry) NewHyTech() hytech.Service {
	return resume.NewHyTech(r.newJWT(), r.newTechRepo())
}

func (r *registry) NewHyUserTech() hyusertech.Service {
	return resume.NewHyUsertech(r.newJWT(), r.newUserTechRepo())
}

func (r *registry) NewHyUserWorkHistory() hyuserworkhistory.Service {
	return resume.NewHyUserWorkHistory(r.newJWT(), r.newUserWorkHistoryRepo())
}

func (r *registry) NewHealth() health.Service {
	return resume.NewHealth()
}

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
		if r.conf.Hash.Salt == "" {
			panic(errors.New("salt key must be set in config"))
		}
		r.hasher = encryption.NewScrypt(r.conf.Hash.Salt)
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

func (r *registry) newUserTechRepo() repository.UserTechRepository {
	if r.userTechRepo == nil {
		r.userTechRepo = repository.NewUserTechRepository(
			r.newMySQLClient(),
		)
	}
	return r.userTechRepo
}

func (r *registry) newUserWorkHistoryRepo() repository.UserWorkHistoryRepository {
	if r.userWorkHistoryRepo == nil {
		r.userWorkHistoryRepo = repository.NewUserWorkHistoryRepository(
			r.newMySQLClient(),
		)
	}
	return r.userWorkHistoryRepo
}
