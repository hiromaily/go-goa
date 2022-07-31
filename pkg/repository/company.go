package repository

import (
	"context"
	"database/sql"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	hycompany "resume/gen/hy_company"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"

	models "github.com/hiromaily/go-goa/pkg/model/rdb"
)

// CompanyRepository interface
type CompanyRepository interface {
	CompanyList() ([]*hycompany.Company, error)
}

type companyRepository struct {
	dbConn    *sql.DB
	tableName string
	logger    *zap.Logger
}

// NewCompanyRepository returns CompanyRepository
func NewCompanyRepository(dbConn *sql.DB, logger *zap.Logger) CompanyRepository {
	return &companyRepository{
		dbConn:    dbConn,
		tableName: "t_company",
		logger:    logger,
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

func (c *companyRepository) GetCompanyGroup(companyID int, isHQ *string) ([]*hycompany.Company, error) {
	ctx := context.Background()

	//	sql := `
	//SELECT cd.id as id, cd.company_id as company_id, c.name as name, cd.hq_flg as hq_flg,
	//  country.name as country_name, cd.address as address
	// FROM t_companies AS c
	// LEFT JOIN t_company_detail AS cd ON c.id = cd.company_id
	// LEFT JOIN m_countries AS country ON cd.country_id = country.id
	// WHERE c.delete_flg=?
	// AND cd.delete_flg=?
	// AND country.delete_flg=? %s
	// AND c.id=?
	//`
	q := []qm.QueryMod{
		qm.Select("cd.id as id, cd.company_id as company_id, c.name as name, cd.is_hq as is_hq, country.name as country_name, cd.address as address"),
		qm.From("t_companies AS c"),
		qm.LeftOuterJoin("group_members on group_members.user_id = users.id"),
		qm.LeftOuterJoin("m_countries AS country ON cd.country_id = country.id"),
		qm.Where("c.is_deleted=?", 0),
		qm.And("cd.is_deleted=?", 0),
		qm.And("country.is_deleted=?", 0),
		qm.And("c.id=?", companyID),
	}
	if isHQ != nil {
		q = append(q, qm.And("cd.is_hq=?", *isHQ))
	}

	var companies []*hycompany.Company
	err := models.TCompanies(q...).BindG(ctx, &companies)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.TCompanies().BindG()")
	}
	return companies, nil
}

//
//type ParamCompany struct {
//	ID   int    `gorm:"column:id"`
//	Name string `gorm:"column:name"`
//}
//
//type ParamCompanyDetail struct {
//	ID        int    `gorm:"column:id"`
//	CompanyID int    `gorm:"column:company_id"`
//	HqFlg     string `gorm:"column:hq_flg"`
//	CountryID int    `gorm:"column:country_id"`
//	Address   string `gorm:"column:address"`
//}
//
//type ParamCompanyDetailTiny struct {
//	ID        int    `gorm:"column:id"`
//	CountryID int    `gorm:"column:country_id"`
//	Address   string `gorm:"column:address"`
//}
//
//
//const TableCompany = "t_companies"
//const TableCompanyDetail = "t_company_detail"
//

// InsertCompany inserts company and company detail
// TODO: transaction
func (c *companyRepository) InsertCompany(name string, branchItem *models.TCompanyBranch) (int, error) {

	ctx := context.Background()
	// company
	companyItem := &models.TCompany{
		Name:      name,
		IsDeleted: null.StringFrom("0"),
	}
	if err := companyItem.Insert(ctx, c.dbConn, boil.Infer()); err != nil {
		return 0, errors.Wrap(err, "failed to call companyItem.Insert()")
	}
	// get LastInsertId()
	id, err := c.getCompanyIDByName(companyItem.Name)
	if err != nil {
		return 0, err
	}

	// company branch
	//		CompanyID: insCompany.ID,
	//		HqFlg:     "1",
	//		CountryID: company.CountryID,
	//		Address:   company.Address,
	if err := branchItem.Insert(ctx, c.dbConn, boil.Infer()); err != nil {
		return 0, errors.Wrap(err, "failed to call branchItem.Insert()")
	}

	return id, nil
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

//func (m *Company) InsertCompany(company *app.CreateCompanyHyCompanyPayload) (int, error) {
//	//TODO:transaction is required
//	tx := m.Db.DB.Begin()
//
//	insCompany := ParamCompany{Name: company.Name}
//	//if err := m.Db.DB.Table(TableCompany).Save(&insCompany).Error; err != nil {
//	if err := tx.Table(TableCompany).Save(&insCompany).Error; err != nil {
//		tx.Rollback()
//		return 0, err
//	}
//	ins2Company := ParamCompanyDetail{
//		CompanyID: insCompany.ID,
//		HqFlg:     "1",
//		CountryID: company.CountryID,
//		Address:   company.Address,
//	}
//	//if err := m.Db.DB.Table(TableCompanyDetail).Save(&ins2Company).Error; err != nil {
//	if err := tx.Table(TableCompanyDetail).Save(&ins2Company).Error; err != nil {
//		tx.Rollback()
//		return 0, err
//	}
//	tx.Commit()
//	return insCompany.ID, nil
//}

//func (m *Company) UpdateCompany(companyID int, company *app.UpdateCompanyHyCompanyPayload) error {
//	tx := m.Db.DB.Begin()
//
//	updCompany := ParamCompany{Name: company.Name}
//	if err := tx.Table(TableCompany).Where("id = ?", companyID).Update(&updCompany).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//	upd2Company := ParamCompanyDetail{CountryID: company.CountryID, Address: company.Address, HqFlg: "1", CompanyID: companyID}
//	if err := tx.Table(TableCompanyDetail).Where("company_id = ? AND hq_flg = ?", companyID, "1").Update(&upd2Company).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	tx.Commit()
//	return nil
//}
//
//func (m *Company) DeleteCompany(companyID int) error {
//	tx := m.Db.DB.Begin()
//	if err := tx.Table(TableCompanyDetail).Where("company_id = ?", companyID).Delete(&ParamCompanyDetail{}).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//	if err := tx.Table(TableCompany).Where("id = ?", companyID).Delete(&ParamCompany{}).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	tx.Commit()
//	return nil
//}
//
//func (m *Company) GetCompanyBranch(branchID int, company *app.Company) error {
//	companies := []*app.Company{}
//	sql := `
//SELECT cd.id as id, cd.company_id as company_id, c.name as name, cd.hq_flg as hq_flg,
//  country.name as country_name, cd.address as address
// FROM t_company_detail AS cd
// LEFT JOIN t_companies AS c ON c.id = cd.company_id
// LEFT JOIN m_countries AS country ON cd.country_id = country.id
// WHERE c.delete_flg=?
// AND cd.delete_flg=?
// AND country.delete_flg=?
// AND cd.id=?
//`
//
//	if err := m.Db.DB.Raw(sql, "0", "0", "0", branchID).Scan(&companies).Error; err != nil {
//		return err
//	}
//
//	if len(companies) == 1 {
//		*company = *companies[0]
//	}
//
//	return nil
//}
//
//func (m *Company) InsertCompanyBranch(companyID int, company *app.CreateCompanyBranchHyCompanybranchPayload) (int, error) {
//	tx := m.Db.DB.Begin()
//
//	insCompany := ParamCompanyDetail{
//		HqFlg:     "0",
//		CompanyID: companyID,
//		CountryID: company.CountryID,
//		Address:   company.Address,
//	}
//	if err := tx.Table(TableCompanyDetail).Save(&insCompany).Error; err != nil {
//		tx.Rollback()
//		return 0, err
//	}
//	tx.Commit()
//	return insCompany.ID, nil
//}
//
//func (m *Company) UpdateCompanyBranch(ID int, company *app.UpdateCompanyBranchHyCompanybranchPayload) error {
//	tx := m.Db.DB.Begin()
//
//	updCompany := ParamCompanyDetailTiny{CountryID: company.CountryID, Address: company.Address}
//	fmt.Println("[updCompany]", updCompany)
//	if err := tx.Table(TableCompanyDetail).Where("id=? AND delete_flg=?", ID, "0").Update(&updCompany).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	tx.Commit()
//	return nil
//}
//
//func (m *Company) DeleteCompanyBranch(ID int) error {
//	tx := m.Db.DB.Begin()
//	if err := tx.Table(TableCompanyDetail).Where("id = ? AND hq_flg<>?", ID, "1").Delete(&ParamCompanyDetail{}).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	tx.Commit()
//	return nil
//}
