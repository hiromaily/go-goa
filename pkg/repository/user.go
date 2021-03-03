package repository

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"go.uber.org/zap"

	"github.com/hiromaily/go-goa/pkg/encryption"
	models "github.com/hiromaily/go-goa/pkg/model/rdb"
)

//
//import (
//	"errors"
//	"github.com/hiromaily/go-goa/goa/app"
//	hs "github.com/hiromaily/golibs/cipher/hash"
//	"github.com/hiromaily/golibs/db/gorm"
//	//lg "github.com/hiromaily/golibs/log"
//	//c "github.com/hiromaily/go-goa/ext/context"
//	//"github.com/jinzhu/gorm"
//	//"golang.org/x/crypto/bcrypt"
//	//"gopkg.in/guregu/null.v3"
//	"fmt"
//)

// UserRepository interface
type UserRepository interface {
	Login(email, password string) (int, error)
}

type userRepository struct {
	dbConn    *sql.DB
	tableName string
	logger    *zap.Logger
	hash      encryption.Hasher
}

// NewUserRepository returns CompanyRepository
func NewUserRepository(dbConn *sql.DB, logger *zap.Logger) CompanyRepository {
	return &companyRepository{
		dbConn:    dbConn,
		tableName: "t_company",
		logger:    logger,
	}
}

//type LoginUser struct {
//	ID       int
//	UserName string
//}
//
//type ParamUser struct {
//	Id       int    `gorm:"column:id"`
//	UserName string `gorm:"column:user_name"`
//	Email    string `gorm:"column:email"`
//	Password string `gorm:"column:password"`
//}
//

func (u *userRepository) Login(email, password string) (int, error) {
	type LoginUser struct {
		ID       int    `boil:"id"`
		Email    string `boil:"email"`
		Password string `boil:"password"`
	}

	ctx := context.Background()

	var user LoginUser
	// sql := "SELECT id, user_name FROM t_users WHERE delete_flg=? AND email=? AND password=?"
	err := models.TUsers(
		qm.Select("id, email, password"),
		qm.Where("email=?", email),
		qm.And("is_deleted=?", 0),
	).Bind(ctx, u.dbConn, &user)
	if err != nil {
		return 0, errors.Wrap(err, "failed to call models.TUsers().Bind()")
	}

	// check
	if user.Password != u.hash.Hash(password) {
		return 0, errors.Errorf("password is invalid")
	}
	return user.ID, nil

}

//// Login is for login
//func (m *User) Login(email, password string) (int, error) {
//	var users []LoginUser
//
//	if err := m.Db.DB.Raw("SELECT id, user_name FROM t_users WHERE delete_flg=? AND email=? AND password=?", "0", email, hs.GetMD5Plus(password, "")).Scan(&users).Error; err != nil {
//		return 0, err
//	}
//
//	if len(users) == 0 || users[0].ID == 0 {
//		return 0, nil
//	} else if len(users) > 1 {
//		return 0, errors.New("data in database would be broken.")
//	}
//
//	return users[0].ID, nil
//}
//
//func (m *User) UserList(users *[]*app.User) error {
//	if err := m.Db.DB.Raw("SELECT id, user_name, email FROM t_users WHERE delete_flg=?", "0").Scan(users).Error; err != nil {
//		fmt.Println("[error]", err)
//		return err
//	}
//
//	return nil
//}
//
//func (m *User) GetUser(userID int, user *app.User) error {
//	users := []*app.User{}
//	if err := m.Db.DB.Raw("SELECT id, user_name, email FROM t_users WHERE delete_flg=? AND id=? limit 1", "0", userID).Scan(&users).Error; err != nil {
//		fmt.Println("[error]", err)
//		return err
//	}
//
//	if len(users) == 1 {
//		*user = *users[0]
//	}
//	return nil
//}
//
//func (m *User) InsertUser(user *app.CreateUserHyUserPayload) (int, error) {
//	//m.Db.DB.Exec("INSERT INTO t_users (user_name, email, password) VALUES (?, ?, ?)", user.UserName, user.Email, user.Password)
//	insUser := ParamUser{UserName: user.UserName, Email: user.Email, Password: user.Password}
//	if err := m.Db.DB.Table(TableUser).Save(&insUser).Error; err != nil {
//		return 0, err
//	}
//
//	return insUser.Id, nil
//}
//
//func (m *User) UpdateUser(userID int, user *app.UpdateUserHyUserPayload) error {
//	//updUser := ParamUser{UserName: user.UserName, Email: user.Email, Password: user.Password}
//	updUser := ParamUser{UserName: user.UserName, Email: user.Email}
//	if user.Password != "**********" {
//		updUser.Password = user.Password
//	}
//	if err := m.Db.DB.Table(TableUser).Where("id = ?", userID).Update(&updUser).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (m *User) DeleteUser(userID int) error {
//	if err := m.Db.DB.Table(TableUser).Where("id = ?", userID).Delete(&ParamUser{}).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
