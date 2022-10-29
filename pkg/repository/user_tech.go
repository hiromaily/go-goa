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
	GetUserLikeTechs(userID int) ([]UserTech, error)
	GetUserDisLikeTechs(userID int) ([]UserTech, error)
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

type UserTech struct {
	ID   int    `boil:"tech_id" json:"tech_id"`
	Name string `boil:"tech_name" json:"tech_name"`
}

func (r *userTechRepository) getUserTechs(userID int, tableName string) ([]UserTech, error) {
	var (
		res []UserTech
		sql = fmt.Sprintf(`
		SELECT
		  lt.tech_id, t.name as tech_name
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

func (r *userTechRepository) GetUserLikeTechs(userID int) ([]UserTech, error) {
	return r.getUserTechs(userID, r.tableNameLike)
}

func (r *userTechRepository) GetUserDisLikeTechs(userID int) ([]UserTech, error) {
	return r.getUserTechs(userID, r.tableNameDisLike)
}
