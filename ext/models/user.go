package models

import (
	"github.com/hiromaily/golibs/db/gorm"
	//c "github.com/hiromaily/go-goa/ext/context"
	//"github.com/jinzhu/gorm"
	//"golang.org/x/crypto/bcrypt"
	//"gopkg.in/guregu/null.v3"
)

// User is user object in Database
type User struct {
	//Ctx *c.Ctx
	Db *gorm.GR
}

// Count is to count
func (m *User) Count() (cnt int) {
	//m.Ctx.Db
	//m.Db.DB
	return 0
}

// Login is for login
func (m *User) Login(username, password string) error {
	//m.Ctx.Db
	//m.Db.DB
	return nil
}
