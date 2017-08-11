package context

import (
	conf "github.com/hiromaily/go-goa/ext/configs"
	"github.com/hiromaily/golibs/db/gorm"
)

// Ctx is context object
type Ctx struct {
	Conf *conf.Config
	Db   *gorm.GR
}

func (c *Ctx) initDB() error {
	dbInfo := conf.GetConf().MySQL
	err := gorm.New(dbInfo.Host, dbInfo.DbName, dbInfo.User, dbInfo.Pass, dbInfo.Port)
	if err != nil {
		return err
	}

	c.Db = gorm.GetDB()
	return nil
}

// SetupContext is to setup context
func SetupContext(c *conf.Config) (*Ctx, error) {
	ctx := &Ctx{Conf: c}
	err := ctx.initDB()

	return ctx, err
}
