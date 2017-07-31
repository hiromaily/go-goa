package models

import (
	"errors"
	hs "github.com/hiromaily/golibs/cipher/hash"
	"github.com/hiromaily/golibs/db/gorm"
	lg "github.com/hiromaily/golibs/log"
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

type LoginUser struct {
	id        int
	firstName string
	lastName  string
}

// Count is to count
func (m *User) Count() (cnt int) {
	//m.Ctx.Db
	//m.Db.DB
	return 0
}

// TODO:Login is for login
func (m *User) Login(email, password string) error {
	//TODO:hash password
	var users []LoginUser
	m.Db.DB.Raw("SELECT user_id, first_name, last_name FROM t_users WHERE delete_flg=? AND email=? AND password=?", "0", email, hs.GetMD5Plus(password, "")).Scan(&users)

	lg.Debugf("len(users): %v", len(users))
	if len(users) == 0 {
		return errors.New("invalid input.")
	} else if len(users) > 1 {
		return errors.New("data in database would be broken.")
	}

	lg.Debugf("users[0].firstName: %v", users[0].firstName)
	return nil
}
