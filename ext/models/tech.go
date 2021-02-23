package models

//
//import (
//	"fmt"
//	"github.com/hiromaily/go-goa/goa/app"
//	"github.com/hiromaily/golibs/db/gorm"
//)
//
//// User is user object in Database
//type Tech struct {
//	//Ctx *c.Ctx
//	Db *gorm.GR
//}
//
//type ParamTech struct {
//	Id   int    `gorm:"column:id"`
//	Name string `gorm:"column:name"`
//}
//
//const TableTech = "t_techs"
//
//func (m *Tech) TechList(techs *[]*app.Tech) error {
//	if err := m.Db.DB.Raw("SELECT id, name FROM t_techs WHERE delete_flg=?", "0").Scan(techs).Error; err != nil {
//		fmt.Println("[error]", err)
//		return err
//	}
//
//	return nil
//}
//
//func (m *Tech) GetTech(techID int, tech *app.Tech) error {
//	techs := []*app.Tech{}
//	if err := m.Db.DB.Raw("SELECT id, name FROM t_techs WHERE delete_flg=? AND id=? limit 1", "0", techID).Scan(&techs).Error; err != nil {
//		fmt.Println("[error]", err)
//		return err
//	}
//
//	if len(techs) == 1 {
//		*tech = *techs[0]
//	}
//	return nil
//}
//
//func (m *Tech) InsertTech(tech *app.CreateTechHyTechPayload) (int, error) {
//	insTech := ParamTech{Name: tech.Name}
//	if err := m.Db.DB.Table(TableTech).Save(&insTech).Error; err != nil {
//		return 0, err
//	}
//
//	return insTech.Id, nil
//}
//
//func (m *Tech) UpdateTech(techID int, tech *app.UpdateTechHyTechPayload) error {
//	updTech := ParamTech{Name: tech.Name}
//	if err := m.Db.DB.Table(TableTech).Where("id = ?", techID).Update(&updTech).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (m *Tech) DeleteTech(techID int) error {
//	if err := m.Db.DB.Table(TableTech).Where("id = ?", techID).Delete(&ParamTech{}).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
