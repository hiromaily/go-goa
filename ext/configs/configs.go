package configs

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	u "github.com/hiromaily/golibs/utils"
	"io/ioutil"
	"os"
)

/* singleton */
var (
	conf         *Config
	tomlFileName string
)

// Config is of root
type Config struct {
	Environment string
	Jwt         *JwtConfig
	MySQL       *MySQLConfig
}

// JwtConfig is for JWT configuration
type JwtConfig struct {
	Key string `toml:"key"`
}

// MySQLConfig is for MySQL Server
type MySQLConfig struct {
	Encrypted bool   `toml:"encrypted"`
	Host      string `toml:"host"`
	Port      uint16 `toml:"port"`
	DbName    string `toml:"dbname"`
	User      string `toml:"user"`
	Pass      string `toml:"pass"`
}

var checkTOMLKeys = [][]string{
	{"environment"},
	{"jwt", "key"},
	{"mysql", "encrypted"},
	{"mysql", "host"},
	{"mysql", "port"},
	{"mysql", "dbname"},
	{"mysql", "user"},
	{"mysql", "pass"},
}

func init() {
	tomlFileName = os.Getenv("GOPATH") + "/src/github.com/hiromaily/go-goa/ext/tomls/settings.toml"
}

//check validation of config
func validateConfig(conf *Config, md *toml.MetaData) error {
	//for protection when debugging on non production environment
	var errStrings []string

	format := "[%s]"
	inValid := false
	for _, keys := range checkTOMLKeys {
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
		return errors.New("Error: Check Text has wrong number of parameter")
	}
	if len(errStrings) != 0 {
		return fmt.Errorf("Error: There are lacks of keys : %#v \n", errStrings)
	}

	return nil
}

// load configfile
func loadConfig(fileName string) (*Config, error) {
	if fileName != "" {
		tomlFileName = fileName
	}

	d, err := ioutil.ReadFile(tomlFileName)
	if err != nil {
		return nil, fmt.Errorf(
			"Error reading %s: %s", tomlFileName, err)
	}

	var config Config
	md, err := toml.Decode(string(d), &config)
	if err != nil {
		return nil, fmt.Errorf(
			"Error parsing %s: %s(%v)", tomlFileName, err, md)
	}

	//check validation of config
	err = validateConfig(&config, &md)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// New is create instance
func New(fileName string) *Config {
	var err error
	if conf == nil {
		conf, err = loadConfig(fileName)
	}
	if err != nil {
		panic(err)
	}
	return conf
}

// GetConf is to get config instance
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
