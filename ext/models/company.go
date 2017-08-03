package models

import (
	//"errors"
	"github.com/hiromaily/go-goa/goa/app"
	"github.com/hiromaily/golibs/db/gorm"
	//lg "github.com/hiromaily/golibs/log"
	"fmt"
)

// User is user object in Database
type Company struct {
	//Ctx *c.Ctx
	Db *gorm.GR
}

func (m *Company) CompanyList(companies *[]*app.CompanyIdname) {
	m.Db.DB.Raw("SELECT id, name FROM t_companies WHERE delete_flg=?", "0").Scan(companies)
}

func (m *Company) GetCompanyGroup(companyID int, hqFlg *string, companies *[]*app.Company) error {
	sql := `
SELECT cd.id as id, cd.company_id as company_id, c.name as name, cd.hq_flg as hq_flg,
  country.name as country_name, cd.address as address
 FROM t_companies AS c
 LEFT JOIN t_company_detail AS cd ON c.id = cd.company_id
 LEFT JOIN m_countries AS country ON cd.country_id = country.id
 WHERE c.delete_flg=?
 AND cd.delete_flg=?
 AND country.delete_flg=? %s
 AND c.id=?
`
	if hqFlg != nil {
		sql = fmt.Sprintf(sql, fmt.Sprintf("AND cd.hq_flg='%s' ", *hqFlg))
	} else {
		sql = fmt.Sprintf(sql, "")
	}

	if err := m.Db.DB.Raw(sql, "0", "0", "0", companyID).Scan(companies).Error; err != nil {
		fmt.Println("[error]", err)
		return err
	}

	return nil
}
