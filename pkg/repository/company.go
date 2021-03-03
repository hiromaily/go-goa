package repository

import (
	"context"
	"database/sql"
	hycompany "resume/gen/hy_company"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/queries/qm"
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
		qm.Where("delete_flg=?", 0),
	).All(ctx, c.dbConn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.TCompanies().All()")
	}

	// TODO: convert []*models.TCompany to []*hycompany.Company
	converted := make([]*hycompany.Company, len(items))
	for i, item := range items {
		converted[i] = &hycompany.Company{
			ID: &item.ID,
			//CompanyID: &item.Name *int
			//Name:        item.Name
			//HqFlg       *string
			//CountryName *string
			//Address string
			//CreatedAt *string
			//UpdatedAt *string
		}
	}
	return converted, nil
}

//func (m *Company) CompanyList(companies *[]*app.CompanyIdname) error {
//	if err := m.Db.DB.Raw("SELECT id as company_id, name FROM t_companies WHERE delete_flg=?", "0").Scan(companies).Error; err != nil {
//		return err
//	}
//	return nil
//}

func (c *companyRepository) GetCompanyGroup(companyID int, IsHQ *string) ([]*models.TCompany, error) {
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

	//items, err := models.TCompanies(
	//	qm.Select("id, name"),
	//	qm.Where("delete_flg=?", 0),
	//).All(ctx, c.dbConn)
	//if err != nil {
	//	return nil, errors.Wrap(err, "failed to call models.TCompanies().All()")
	//}
	return nil, nil
}

//func (m *Company) GetCompanyGroup(companyID int, hqFlg *string, companies *[]*app.Company) error {
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
//	if hqFlg != nil {
//		sql = fmt.Sprintf(sql, fmt.Sprintf("AND cd.hq_flg='%s' ", *hqFlg))
//	} else {
//		sql = fmt.Sprintf(sql, "")
//	}
//
//	if err := m.Db.DB.Raw(sql, "0", "0", "0", companyID).Scan(companies).Error; err != nil {
//		fmt.Println("[error]", err)
//		return err
//	}
//
//	return nil
//}

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
//
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
//
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
