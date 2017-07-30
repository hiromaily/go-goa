package config

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	enc "github.com/hiromaily/golibs/cipher/encryption"
	u "github.com/hiromaily/golibs/utils"
	"io/ioutil"
	"os"
)

var tomlFileName = os.Getenv("GOPATH") + "/src/github.com/hiromaily/golibs/config/settings.toml"

var conf *Config

// Config is of root
type Config struct {
	Environment int
	MySQL       *MySQLConfig
	Redis       *RedisConfig
	Mongo       *MongoConfig `toml:"mongodb"`
	Cassa       *CassaConfig `toml:"cassandra"`
	Mail        *MailConfig
	Aws         *AwsConfig
}

// MySQLConfig is for MySQL server
type MySQLConfig struct {
	Encrypted bool   `toml:"encrypted"`
	Host      string `toml:"host"`
	Port      uint16 `toml:"port"`
	DbName    string `toml:"dbname"`
	User      string `toml:"user"`
	Pass      string `toml:"pass"`
}

// RedisConfig is for Redis server
type RedisConfig struct {
	Encrypted bool   `toml:"encrypted"`
	Host      string `toml:"host"`
	Port      uint16 `toml:"port"`
	Pass      string `toml:"pass"`
}

// MongoConfig is for MongoDB server
type MongoConfig struct {
	Encrypted bool   `toml:"encrypted"`
	Host      string `toml:"host"`
	Port      uint16 `toml:"port"`
	DbName    string `toml:"dbname"`
	User      string `toml:"user"`
	Pass      string `toml:"pass"`
}

// CassaConfig is for Cassandra server
type CassaConfig struct {
	Encrypted bool   `toml:"encrypted"`
	Host      string `toml:"host"`
	Port      uint16 `toml:"port"`
	KeySpace  string `toml:"keyspace"`
}

// MailConfig is for mail
type MailConfig struct {
	Encrypted bool                `toml:"encrypted"`
	Address   string              `toml:"address"`
	Password  string              `toml:"password"`
	Timeout   string              `toml:"timeout"`
	SMTP      *SMTPConfig         `toml:"smtp"`
	Content   []MailContentConfig `toml:"content"`
}

// SMTPConfig is for SMTP server of mail
type SMTPConfig struct {
	Server string `toml:"server"`
	Port   int    `toml:"port"`
}

// MailContentConfig is for mail contents
type MailContentConfig struct {
	Subject string `toml:"subject"`
	Tplfile string `toml:"tplfile"`
}

// AwsConfig is for Aamazon Web Service
type AwsConfig struct {
	Encrypted bool       `toml:"encrypted"`
	AccessKey string     `toml:"access_key"`
	SecretKey string     `toml:"secret_key"`
	Region    string     `toml:"region"`
	Sqs       *SqsConfig `toml:"sqs"`
}

// SqsConfig is for SQS of AWS
type SqsConfig struct {
	Endpoint      string        `toml:"endpoint"`
	QueueName     string        `toml:"queue_name"`
	DeadQueueName string        `toml:"deadque_name"`
	MsgAttr       MsgAttrConfig `toml:"msgattr"`
}

// MsgAttrConfig is for part of SQS
type MsgAttrConfig struct {
	OpType      string `toml:"operation_type"`
	ContentType string `toml:"content_type"`
}

var checkTomlKeys = [][]string{
	{"environment"},
	{"mysql", "encrypted"},
	{"mysql", "host"},
	{"mysql", "port"},
	{"mysql", "dbname"},
	{"mysql", "user"},
	{"mysql", "pass"},
	{"redis", "encrypted"},
	{"redis", "host"},
	{"redis", "port"},
	{"redis", "pass"},
	{"mongodb", "encrypted"},
	{"mongodb", "host"},
	{"mongodb", "port"},
	{"mongodb", "dbname"},
	{"mongodb", "user"},
	{"mongodb", "pass"},
	{"cassandra", "encrypted"},
	{"cassandra", "host"},
	{"cassandra", "port"},
	{"cassandra", "keyspace"},
	{"mail", "encrypted"},
	{"mail", "address"},
	{"mail", "password"},
	{"mail", "timeout"},
	{"mail", "smtp", "server"},
	{"mail", "smtp", "port"},
	//{"mail", "content", "subject"},
	//{"mail", "content", "tplfile"},
	{"aws", "encrypted"},
	{"aws", "access_key"},
	{"aws", "secret_key"},
	{"aws", "region"},
	{"aws", "sqs", "endpoint"},
	{"aws", "sqs", "queue_name"},
	{"aws", "sqs", "deadque_name"},
	{"aws", "sqs", "msgattr", "operation_type"},
	{"aws", "sqs", "msgattr", "content_type"},
}

//check validation of config
func validateConfig(conf *Config, md *toml.MetaData) error {
	//for protection when debugging on non production environment
	var errStrings []string

	//Check added new items on toml
	// environment
	//if !md.IsDefined("environment") {
	//	errStrings = append(errStrings, "environment")
	//}

	format := "[%s]"
	inValid := false
	for _, keys := range checkTomlKeys {
		if !md.IsDefined(keys...) {
			switch len(keys) {
			case 1:
				format = "[%s]"
			case 2:
				format = "[%s] %s"
			case 3:
				format = "[%s.%s] %s"
			default:
				//invalid check string
				inValid = true
				break
			}
			keysIfc := u.SliceStrToInterface(keys)
			errStrings = append(errStrings, fmt.Sprintf(format, keysIfc...))
		}
	}

	// Error
	if inValid {
		return errors.New("error: Check Text has wrong number of parameter")
	}
	if len(errStrings) != 0 {
		return fmt.Errorf("error: There are lacks of keys : %#v \n", errStrings)
	}

	return nil
}

// load configfile
func loadConfig(path string) (*Config, error) {
	if path != "" {
		tomlFileName = path
	}

	d, err := ioutil.ReadFile(tomlFileName)
	if err != nil {
		return nil, fmt.Errorf(
			"error reading %s: %s", tomlFileName, err)
	}

	var config Config
	md, err := toml.Decode(string(d), &config)
	if err != nil {
		return nil, fmt.Errorf(
			"error parsing %s: %s(%v)", tomlFileName, err, md)
	}

	//check validation of config
	err = validateConfig(&config, &md)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// New is to create config instance
func New(file string, cipherFlg bool) *Config {
	var err error
	conf, err = loadConfig(file)
	if err != nil {
		panic(err)
	}

	if cipherFlg {
		Cipher()
	}

	return conf
}

// GetConf is to get config instance. singleton architecture
func GetConf() *Config {
	var err error
	if conf == nil {
		conf, err = loadConfig("")
	}
	if err != nil {
		panic(err)
	}

	return conf
}

// SetTOMLPath is to set TOML file path
func SetTOMLPath(path string) {
	tomlFileName = path
}

// ResetConf is to clear config instance
func ResetConf() {
	conf = nil
}

// Cipher is to decrypt crypted string on config
func Cipher() {
	crypt := enc.GetCrypt()

	if conf.MySQL.Encrypted {
		c := conf.MySQL
		c.Host, _ = crypt.DecryptBase64(c.Host)
		c.DbName, _ = crypt.DecryptBase64(c.DbName)
		c.User, _ = crypt.DecryptBase64(c.User)
		c.Pass, _ = crypt.DecryptBase64(c.Pass)
	}

	if conf.Redis.Encrypted {
		c := conf.Redis
		c.Host, _ = crypt.DecryptBase64(c.Host)
		c.Pass, _ = crypt.DecryptBase64(c.Pass)
	}

	if conf.Mongo.Encrypted {
		c := conf.Mongo
		c.Host, _ = crypt.DecryptBase64(c.Host)
		c.DbName, _ = crypt.DecryptBase64(c.DbName)
		c.User, _ = crypt.DecryptBase64(c.User)
		c.Pass, _ = crypt.DecryptBase64(c.Pass)
	}

	if conf.Mail.Encrypted {
		c := conf.Mail
		c.Address, _ = crypt.DecryptBase64(c.Address)
		c.Password, _ = crypt.DecryptBase64(c.Password)

		c2 := conf.Mail.SMTP
		c2.Server, _ = crypt.DecryptBase64(c2.Server)
	}

	if conf.Aws.Encrypted {
		c := conf.Aws
		c.AccessKey, _ = crypt.DecryptBase64(c.AccessKey)
		c.SecretKey, _ = crypt.DecryptBase64(c.SecretKey)
	}
}
