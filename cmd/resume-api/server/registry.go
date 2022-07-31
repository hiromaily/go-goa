package main

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"

	"github.com/hiromaily/go-goa/pkg/config"
	"github.com/hiromaily/go-goa/pkg/encryption"
	"github.com/hiromaily/go-goa/pkg/jwts"
	"github.com/hiromaily/go-goa/pkg/logger"
	"github.com/hiromaily/go-goa/pkg/repository"
	"github.com/hiromaily/go-goa/pkg/storage/mysql"
)

type registry struct {
	conf     *config.Root
	logger   *zap.Logger
	jwter    jwts.JWTer
	db       *sql.DB
	hasher   encryption.Hasher
	userRepo repository.UserRepository
}

func createRegistry(confPath string) *registry {
	reg := &registry{}
	var err error

	// config
	reg.conf, err = config.NewConfig(confPath, false)
	if err != nil {
		panic(err)
	}

	// logger
	reg.newLogger()
	reg.newJWT()

	// database
	reg.newMySQLClient()

	// create repository
	// TODO: create proper repository for each services
	reg.newUserRepo()
	reg.newTechRepo()

}

func (r *registry) newLogger() *zap.Logger {
	if r.logger == nil {
		r.logger = logger.NewZapLogger(r.conf.Logger)
	}
	return r.logger
}

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
		// TODO: hash salt
		r.hasher = encryption.NewScrypt("salt")
	}
	return r.hasher
}

func (r *registry) newUserRepo() repository.UserRepository {
	if r.userRepo == nil {
		r.userRepo = repository.NewUserRepository(
			r.newMySQLClient(),
			r.newLogger(),
			r.newHasher(),
		)
	}
	return r.userRepo
}

func (r *registry) newTechRepo() repository.TechRepository {
	if r.userRepo == nil {
		r.userRepo = repository.NewUserRepository(
			r.newMySQLClient(),
			r.newLogger(),
		)
	}
	return r.userRepo
}
