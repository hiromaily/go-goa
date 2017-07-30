package gorm

import (
	"errors"
	"fmt"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// GR is for database information and gorm instance
type GR struct {
	DB         *gorm.DB
	ServerInfo //embeded
}

// ServerInfo is for database information
type ServerInfo struct {
	host   string
	port   uint16
	dbname string
	user   string
	pass   string
}

var dbInfo GR

// New is for create instance
func New(host, dbname, user, pass string, port uint16) {

	var err error
	if dbInfo.DB == nil {
		dbInfo.host = host
		dbInfo.port = port
		dbInfo.dbname = dbname
		dbInfo.user = user
		dbInfo.pass = pass

		dbInfo.DB, err = dbInfo.Connection()
	}
	fmt.Printf("dbInfo.db %#v\n", *dbInfo.DB)
	if err != nil {
		panic(err.Error())
	}

	return
}

// GetDB is to get instance. singleton architecture
func GetDB() *GR {
	if dbInfo.DB == nil {
		panic(errors.New("DB instance is nil"))
	}
	return &dbInfo
}

func (gr *GR) getDsn() string {
	//If use nil on Date column, set *time.Time
	//Be careful when parsing is required on Date type
	// e.g. db, err := sql.Open("mysql", "root:@/?parseTime=true")
	param := "?charset=utf8&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
		gr.user, gr.pass, gr.host, gr.port, gr.dbname, param)
}

// Connection is to connect to MySQL server
// Be careful, sql.Open() doesn't return err. Use db.Ping() to check DB condition.
func (gr *GR) Connection() (*gorm.DB, error) {
	//return sql.Open("mysql", getDsn())
	return gorm.Open("mysql", gr.getDsn())
}

// Close is to close connection
func (gr *GR) Close() {
	gr.DB.Close()
}
