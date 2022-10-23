package repository

import (
	"context"
	"database/sql"
	hyuser "resume/gen/hy_user"
	"time"

	"github.com/hiromaily/go-goa/pkg/encryption"
	models "github.com/hiromaily/go-goa/pkg/model/rdb"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// UserRepository interface
type UserRepository interface {
	Login(email, password string) (int, error)
	UserList() ([]*hyuser.User, error)
	GetUser(userID int) (*hyuser.User, error)
	InsertUser(name, email, password string) (int, error)
	UpdateUser(userID int, name, email, password string) (int, error)
	DeleteUser(userID int) (int, error)
}

type userRepository struct {
	tableName string
	dbConn    *sql.DB
	hash      encryption.Hasher
}

// NewUserRepository returns UserRepository
func NewUserRepository(dbConn *sql.DB, hash encryption.Hasher) UserRepository {
	return &userRepository{
		dbConn:    dbConn,
		tableName: "t_user",
		hash:      hash,
	}
}

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

func (u *userRepository) UserList() ([]*hyuser.User, error) {
	ctx := context.Background()

	// sql = "SELECT id, user_name, email FROM t_users WHERE delete_flg=?"
	items, err := models.TUsers(
		qm.Select("id, user_name, email"),
		qm.Where("is_deleted=?", 0),
	).All(ctx, u.dbConn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.TUsers().All()")
	}

	converted := make([]*hyuser.User, len(items))
	for i, item := range items {
		converted[i] = &hyuser.User{
			ID:       &item.ID,
			UserName: item.UserName,
			Email:    item.Email.String,
		}
	}
	return converted, nil
}

func (u *userRepository) GetUser(userID int) (*hyuser.User, error) {
	ctx := context.Background()
	// sql := "SELECT id, user_name, email FROM t_users WHERE delete_flg=?"
	q := []qm.QueryMod{
		qm.Select("id, user_name, email"),
		qm.Where("is_deleted=?", 0),
		qm.And("id=?", userID),
	}
	item, err := models.TUsers(q...).One(ctx, u.dbConn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.TUsers().One()")
	}
	return &hyuser.User{
		ID:       &item.ID,
		UserName: item.UserName,
		Email:    item.Email.String,
	}, nil
}

func (u *userRepository) getUserByEmail(email string) (*models.TUser, error) {
	ctx := context.Background()

	q := []qm.QueryMod{
		qm.Where("is_deleted=?", 0),
		qm.And("email=?", email),
	}
	item, err := models.TUsers(q...).One(ctx, u.dbConn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.TUsers().One()")
	}
	return item, nil
}

func (u *userRepository) InsertUser(name, email, password string) (int, error) {
	item := &models.TUser{
		UserName: name,
		Email:    null.StringFrom(email),
		Password: null.StringFrom(u.hash.Hash(password)),
	}

	ctx := context.Background()
	// sql := "INSERT INTO t_users (user_name, email, password) VALUES (?, ?, ?)"
	if err := item.Insert(ctx, u.dbConn, boil.Infer()); err != nil {
		return 0, errors.Wrap(err, "failed to call item.Insert()")
	}
	user, err := u.getUserByEmail(email)
	if err != nil {
		return 0, errors.Wrap(err, "failed to call getUserByEmail()")
	}

	return user.ID, nil
}

func (u *userRepository) UpdateUser(userID int, name, email, password string) (int, error) {
	if userID == 0 {
		return 0, errors.New("userID is invalid")
	}

	ctx := context.Background()

	// Set updating columns
	updCols := map[string]interface{}{}
	if name != "" {
		updCols[models.TUserColumns.UserName] = name
	}
	if email != "" {
		updCols[models.TUserColumns.Email] = email
	}
	if password != "" {
		updCols[models.TUserColumns.Password] = u.hash.Hash(password)
	}
	updCols[models.TUserColumns.UpdatedAt] = null.TimeFrom(time.Now().UTC())

	id, err := models.TUsers(
		qm.Where("id=?", userID),
	).UpdateAll(ctx, u.dbConn, updCols)
	return int(id), err
}

// DeleteUser deletes user
func (u *userRepository) DeleteUser(userID int) (int, error) {
	ctx := context.Background()

	// Set updating columns
	updCols := map[string]interface{}{}
	updCols[models.TUserColumns.IsDeleted] = null.StringFrom("1")
	updCols[models.TUserColumns.UpdatedAt] = null.TimeFrom(time.Now())

	id, err := models.TUsers(
		qm.Where("id=?", userID),
	).UpdateAll(ctx, u.dbConn, updCols)
	return int(id), err
}
