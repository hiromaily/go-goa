package models

import (
	"fmt"
	"github.com/hiromaily/go-goa/goa/app"
	"github.com/hiromaily/golibs/db/gorm"
)

// User is user object in Database
type UserTech struct {
	//Ctx *c.Ctx
	Db *gorm.GR
}

const TableTechLike = "t_user_like_techs"
const TableTechDislike = "t_user_dislike_techs"

func (m *UserTech) GetUserTechs(userID int, userTechs *[]*app.UsertechTech, tableName string) error {
	sql := `
SELECT t.name as tech_name
 FROM %s AS ult
 LEFT JOIN t_techs AS t ON ult.tech_id = t.id
 WHERE ult.delete_flg=?
 AND t.delete_flg=?
 AND ult.user_id=?
`
	sql = fmt.Sprintf(sql, tableName)

	if err := m.Db.DB.Raw(sql, "0", "0", userID).Scan(userTechs).Error; err != nil {
		fmt.Println("[error]", err)
		return err
	}

	return nil
}
