package context

import (
	conf "github.com/hiromaily/go-goa/ext/configs"
	"github.com/hiromaily/golibs/db/mysql"
)

type GoaApi struct {
	Conf *conf.Config
	Db   *mysql.MS
}

func (api *GoaApi) initDB() {
	dbInfo := conf.GetConf().MySQL
	mysql.New(dbInfo.Host, dbInfo.DbName, dbInfo.User, dbInfo.Pass, dbInfo.Port)
	mysql.GetDB().SetMaxIdleConns(50)

	api.Db = mysql.GetDB()
}

func SetupContext(c *conf.Config) *GoaApi {
	context := &GoaApi{Conf: c}
	context.initDB()

	return context
}
