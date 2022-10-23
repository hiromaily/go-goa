package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	hytech "resume/gen/hy_tech"

	models "github.com/hiromaily/go-goa/pkg/model/rdb"
)

// TechRepository interface
type TechRepository interface {
	TechList() ([]*hytech.Tech, error)
	GetTech(techID int) (*hytech.Tech, error)
	InsertTech(name string) (int, error)
	UpdateTech(techID int, name string) (int, error)
	DeleteTech(techID int) (int, error)
}

type techRepository struct {
	tableName string
	dbConn    *sql.DB
}

// NewTechRepository returns TechRepository
func NewTechRepository(dbConn *sql.DB) TechRepository {
	return &techRepository{
		dbConn:    dbConn,
		tableName: "t_tech",
	}
}

func (u *techRepository) TechList() ([]*hytech.Tech, error) {
	ctx := context.Background()

	items, err := models.TTeches(
		qm.Select("id, name"),
		qm.Where("is_deleted=?", 0),
	).All(ctx, u.dbConn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.TTechs().All()")
	}

	converted := make([]*hytech.Tech, len(items))
	for i, item := range items {
		converted[i] = &hytech.Tech{
			ID:   &item.ID,
			Name: item.Name,
		}
	}
	return converted, nil
}

func (u *techRepository) GetTech(techID int) (*hytech.Tech, error) {
	ctx := context.Background()
	// sql := "SELECT id, name FROM t_techs WHERE delete_flg=? AND id=?"
	q := []qm.QueryMod{
		qm.Select("id, name"),
		qm.Where("is_deleted=?", 0),
		qm.And("id=?", techID),
	}
	item, err := models.TTeches(q...).One(ctx, u.dbConn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.TTechs().One()")
	}
	return &hytech.Tech{
		ID:   &item.ID,
		Name: item.Name,
	}, nil
}

func (u *techRepository) getTechByName(name string) (*models.TTech, error) {
	ctx := context.Background()

	q := []qm.QueryMod{
		qm.Where("is_deleted=?", 0),
		qm.And("name=?", name),
	}
	item, err := models.TTeches(q...).One(ctx, u.dbConn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.TTechs().One()")
	}
	return item, nil
}

func (u *techRepository) InsertTech(name string) (int, error) {
	item := &models.TTech{
		Name: name,
	}

	ctx := context.Background()
	if err := item.Insert(ctx, u.dbConn, boil.Infer()); err != nil {
		return 0, errors.Wrap(err, "failed to call item.Insert()")
	}
	tech, err := u.getTechByName(name)
	return 0, errors.Wrap(err, "failed to call getTechByName()")

	return tech.ID, nil
}

func (u *techRepository) UpdateTech(techID int, name string) (int, error) {
	if techID == 0 {
		return 0, errors.New("techID is invalid")
	}

	ctx := context.Background()

	// Set updating columns
	updCols := map[string]interface{}{}
	if name != "" {
		updCols[models.TTechColumns.Name] = name
	}
	updCols[models.TTechColumns.UpdatedAt] = null.TimeFrom(time.Now().UTC())

	id, err := models.TTeches(
		qm.Where("id=?", techID),
	).UpdateAll(ctx, u.dbConn, updCols)
	return int(id), err
}

func (u *techRepository) DeleteTech(techID int) (int, error) {
	ctx := context.Background()

	// Set updating columns
	updCols := map[string]interface{}{}
	updCols[models.TTechColumns.IsDeleted] = null.StringFrom("1")
	updCols[models.TTechColumns.UpdatedAt] = null.TimeFrom(time.Now())

	id, err := models.TTeches(
		qm.Where("id=?", techID),
	).UpdateAll(ctx, u.dbConn, updCols)
	return int(id), err
}
