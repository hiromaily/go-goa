package repository

import (
	"context"
	"database/sql"
	ptr "github.com/hiromaily/go-goa/pkg/pointer"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "github.com/hiromaily/go-goa/pkg/model/rdb"
	hycompany "resume/gen/hy_company"
)

// CompanyRepository interface
type CompanyRepository interface {
	CompanyList() ([]*hycompany.Company, error)
	GetCompany(companyID int) (*hycompany.Company, error)
	IsCompany(companyID int) (bool, error)
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

type CustomCompany struct {
	CompanyID   int
	CompanyName string
	Address     string
	CountryName string
}

func (r *companyRepository) CompanyList() ([]*hycompany.Company, error) {
	ctx := context.Background()

	// sql := "SELECT id as company_id, name FROM t_companies WHERE delete_flg=?"
	items, err := models.TCompanies(
		qm.Select("id, name, address, country_id"),
		qm.Where("is_deleted=?", 0),
	).All(ctx, r.dbConn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.TCompanies().All()")
	}

	converted := make([]*hycompany.Company, len(items))
	for i, item := range items {
		converted[i] = &hycompany.Company{
			CompanyID:   &item.ID,
			CompanyName: &item.Name,
			Address:     item.Address.Ptr(),
		}
	}
	return converted, nil
}

func (r *companyRepository) GetCompany(companyID int) (*hycompany.Company, error) {
	var res CustomCompany

	sql := `
		SELECT cp.id as company_id, cp.name as company_name, cp.address, c.name as country_name
		FROM t_company AS cp
		LEFT JOIN m_country AS c ON cp.country_id = c.id
		WHERE cp.id=?
		AND cp.is_deleted=?
		AND c.is_deleted=?
	`
	if err := queries.Raw(sql, companyID, "0", "0").Bind(context.Background(), r.dbConn, &res); err != nil {
		return nil, errors.Wrapf(err, "failed to call queries.Raw(): %s", sql)
	}

	converted := &hycompany.Company{
		CompanyID:   ptr.Int(res.CompanyID),
		CompanyName: ptr.String(res.CompanyName),
		CountryName: ptr.String(res.CountryName),
		Address:     ptr.String(res.Address),
	}
	return converted, nil
	//ctx := context.Background()
	//q := []qm.QueryMod{
	//	qm.Select("id, name, address, c.name as country_name"),
	//	qm.From(u.tableName),
	//	qm.LeftOuterJoin("m_country as c on t_company.country_id = c.id"),
	//	qm.Where("is_deleted=?", 0),
	//	qm.And("id=?", companyID),
	//}
	//item, err := models.TCompanies(q...).One(ctx, u.dbConn)
	//if err != nil {
	//	return nil, errors.Wrap(err, "failed to call models.TCompanies().One()")
	//}
	//
	//return &hycompany.Company{
	//	CompanyID:   &item.ID,
	//	CompanyName: &item.Name,
	//	CountryName: nil,
	//	Address:     &item.Address.String,
	//}, nil
}

func (r *companyRepository) IsCompany(companyID int) (bool, error) {
	ctx := context.Background()
	q := []qm.QueryMod{
		qm.Select("id"),
		qm.Where("is_deleted=?", 0),
		qm.And("id=?", companyID),
	}
	ret, err := models.TCompanies(q...).Exists(ctx, r.dbConn)
	if err != nil {
		return false, errors.Wrap(err, "failed to call models.TCompanies().One()")
	}

	return ret, nil
}

func (r *companyRepository) getCompanyIDByName(name string) (int, error) {
	ctx := context.Background()

	item, err := models.TCompanies(
		qm.Select("id"),
		qm.Where("name=?", name),
		qm.And("is_deleted=?", 0),
	).One(ctx, r.dbConn)
	if err != nil {
		return 0, errors.Wrap(err, "failed to call models.TCompanies().One()")
	}

	return item.ID, nil
}

// InsertCompany inserts company and company detail
func (r *companyRepository) InsertCompany(name, address string, countryID int16) (int, error) {
	ctx := context.Background()
	// company
	companyItem := &models.TCompany{
		Name:      name,
		CountryID: countryID,
		Address:   null.StringFrom(address),
		IsDeleted: null.StringFrom("0"),
	}
	if err := companyItem.Insert(ctx, r.dbConn, boil.Infer()); err != nil {
		return 0, errors.Wrap(err, "failed to call companyItem.Insert()")
	}
	// get LastInsertId()
	return r.getCompanyIDByName(companyItem.Name)
}

func (r *companyRepository) UpdateCompany(companyID int, name, address string, countryID int16) (int, error) {
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
	).UpdateAll(ctx, r.dbConn, updCols)
	return int(id), err
}

func (r *companyRepository) DeleteCompany(companyID int) (int, error) {
	ctx := context.Background()

	// Set updating columns
	updCols := map[string]interface{}{}
	updCols[models.TCompanyColumns.IsDeleted] = null.StringFrom("1")
	updCols[models.TCompanyColumns.UpdatedAt] = null.TimeFrom(time.Now())

	id, err := models.TCompanies(
		qm.Where("id=?", companyID),
	).UpdateAll(ctx, r.dbConn, updCols)
	return int(id), err
}
