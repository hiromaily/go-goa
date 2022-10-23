package repository

import (
	"context"
	"database/sql"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	hycompany "resume/gen/hy_company"
	"time"

	models "github.com/hiromaily/go-goa/pkg/model/rdb"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// CompanyRepository interface
type CompanyRepository interface {
	CompanyList() ([]*hycompany.Company, error)
	InsertCompany(name, address string, countryID int16) (int, error)
	UpdateCompany(companyID int, name, address string, countryID int16) (int, error)
	DeleteCompany(companyID int) (int, error)
}

type companyRepository struct {
	dbConn    *sql.DB
	tableName string
}

// NewCompanyRepository returns CompanyRepository
func NewCompanyRepository(dbConn *sql.DB) CompanyRepository {
	return &companyRepository{
		dbConn:    dbConn,
		tableName: "t_company",
	}
}

func (c *companyRepository) CompanyList() ([]*hycompany.Company, error) {
	ctx := context.Background()

	// sql := "SELECT id as company_id, name FROM t_companies WHERE delete_flg=?"
	items, err := models.TCompanies(
		qm.Select("id, name"),
		qm.Where("is_deleted=?", 0),
	).All(ctx, c.dbConn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.TCompanies().All()")
	}

	// convert []*models.TCompany to []*hycompany.Company
	converted := make([]*hycompany.Company, len(items))
	for i, item := range items {
		// return only companyID, name
		converted[i] = &hycompany.Company{
			CompanyID: &item.ID,
			Name:      item.Name,
		}
	}
	return converted, nil
}

func (c *companyRepository) getCompanyIDByName(name string) (int, error) {
	ctx := context.Background()

	item, err := models.TCompanies(
		qm.Select("id"),
		qm.Where("name=?", name),
		qm.And("is_deleted=?", 0),
	).One(ctx, c.dbConn)
	if err != nil {
		return 0, errors.Wrap(err, "failed to call models.TCompanies().One()")
	}

	return item.ID, nil
}

// InsertCompany inserts company and company detail
func (c *companyRepository) InsertCompany(name, address string, countryID int16) (int, error) {

	ctx := context.Background()
	// company
	companyItem := &models.TCompany{
		Name:      name,
		CountryID: countryID,
		Address:   null.StringFrom(address),
		IsDeleted: null.StringFrom("0"),
	}
	if err := companyItem.Insert(ctx, c.dbConn, boil.Infer()); err != nil {
		return 0, errors.Wrap(err, "failed to call companyItem.Insert()")
	}
	// get LastInsertId()
	return c.getCompanyIDByName(companyItem.Name)
}

func (u *companyRepository) UpdateCompany(companyID int, name, address string, countryID int16) (int, error) {
	if companyID == 0 {
		return 0, errors.New("companyID is invalid")
	}

	ctx := context.Background()

	// Set updating columns
	updCols := map[string]interface{}{}
	if name != "" {
		updCols[models.TCompanyColumns.Name] = name
	}
	if address != "" {
		updCols[models.TCompanyColumns.Address] = address
	}
	if countryID != 0 {
		updCols[models.TCompanyColumns.CountryID] = countryID
	}
	updCols[models.TCompanyColumns.UpdatedAt] = null.TimeFrom(time.Now().UTC())

	id, err := models.TCompanies(
		qm.Where("id=?", companyID),
	).UpdateAll(ctx, u.dbConn, updCols)
	return int(id), err
}

func (u *companyRepository) DeleteCompany(companyID int) (int, error) {
	ctx := context.Background()

	// Set updating columns
	updCols := map[string]interface{}{}
	updCols[models.TCompanyColumns.IsDeleted] = null.StringFrom("1")
	updCols[models.TCompanyColumns.UpdatedAt] = null.TimeFrom(time.Now())

	id, err := models.TCompanies(
		qm.Where("id=?", companyID),
	).UpdateAll(ctx, u.dbConn, updCols)
	return int(id), err
}
