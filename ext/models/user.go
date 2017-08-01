package models

import (
	"errors"
	"github.com/hiromaily/go-goa/goa/app"
	hs "github.com/hiromaily/golibs/cipher/hash"
	"github.com/hiromaily/golibs/db/gorm"
	//lg "github.com/hiromaily/golibs/log"
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
	id       int
	userName string
}

type ParamUser struct {
	Id       int    `gorm:"column:id"`
	UserName string `gorm:"column:user_name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

// Count is to count
func (m *User) Count() (cnt int) {
	//m.Ctx.Db
	//m.Db.DB
	return 0
}

// Login is for login
func (m *User) Login(email, password string) error {
	var users []LoginUser
	m.Db.DB.Raw("SELECT id, user_name FROM t_users WHERE delete_flg=? AND email=? AND password=?", "0", email, hs.GetMD5Plus(password, "")).Scan(&users)

	if len(users) == 0 {
		return errors.New("invalid input.")
	} else if len(users) > 1 {
		return errors.New("data in database would be broken.")
	}

	//lg.Debugf("users[0].userName: %v", users[0].userName)
	return nil
}

func (m *User) UserList(users *[]*app.User) {
	m.Db.DB.Raw("SELECT id, user_name, email FROM t_users WHERE delete_flg=?", "0").Scan(users)
}

func (m *User) GetUser(userID int, user *app.User) {
	users := []*app.User{}
	m.Db.DB.Raw("SELECT id, user_name, email FROM t_users WHERE delete_flg=? AND id=? limit 1", "0", userID).Scan(&users)

	if len(users) == 1 {
		*user = *users[0]
	}
	return
}

func (m *User) InsertUser(user *app.CreateUserHyUserPayload) (int, error) {
	//m.Db.DB.Exec("INSERT INTO t_users (user_name, email, password) VALUES (?, ?, ?)", user.UserName, user.Email, user.Password)
	insUser := ParamUser{UserName: user.UserName, Email: user.Email, Password: user.Password}
	m.Db.DB.Table("t_users").Save(&insUser)
	if m.Db.DB.Error != nil {
		return 0, m.Db.DB.Error
	}

	//lg.Debugf("Id: %d", insUser.Id)
	return insUser.Id, nil
}

func (m *User) UpdateUser(userID int, user *app.UpdateUserHyUserPayload) error {
	//m.Db.DB.Exec("INSERT INTO t_users (user_name, email, password) VALUES (?, ?, ?)", user.UserName, user.Email, user.Password)
	updUser := ParamUser{UserName: user.UserName, Email: user.Email, Password: user.Password}
	m.Db.DB.Table("t_users").Where("id = ?", userID).Update(&updUser)
	if m.Db.DB.Error != nil {
		return m.Db.DB.Error
	}

	return nil
}

func (m *User) DeleteUser(userID int) error {
	m.Db.DB.Table("t_users").Where("id = ?", userID).Delete(&ParamUser{})
	if m.Db.DB.Error != nil {
		return m.Db.DB.Error
	}

	return nil

}
