package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

// UserTechRepository interface
type UserTechRepository interface {
	GetUserLikeTechs(userID int) ([]string, error)
	GetUserDisLikeTechs(userID int) ([]string, error)
}

type userTechRepository struct {
	dbConn           *sql.DB
	tableNameLike    string
	tableNameDisLike string
}

// NewUserTechRepository returns UserTechRepository
func NewUserTechRepository(dbConn *sql.DB) UserTechRepository {
	return &userTechRepository{
		dbConn:           dbConn,
		tableNameLike:    "t_user_like_tech",
		tableNameDisLike: "t_user_dislike_tech",
	}
}

func (r *userTechRepository) getUserTechs(userID int, tableName string) ([]string, error) {
	var (
		res []string
		sql = fmt.Sprintf(`
		SELECT
		  t.name as tech_name
		FROM %s as lt
		  LEFT JOIN t_tech AS t ON lt.tech_id = t.id
		WHERE lt.is_deleted="0"
		AND t.is_deleted="0"
		AND lt.user_id=%d`, tableName, userID)
	)

	if err := queries.Raw(sql).Bind(context.Background(), r.dbConn, &res); err != nil {
		return nil, errors.Wrapf(err, "failed to call queries.Raw(): %s", sql)
	}
	return res, nil
}

func (r *userTechRepository) GetUserLikeTechs(userID int) ([]string, error) {
	return r.getUserTechs(userID, r.tableNameLike)
}

func (r *userTechRepository) GetUserDisLikeTechs(userID int) ([]string, error) {
	return r.getUserTechs(userID, r.tableNameDisLike)
}
