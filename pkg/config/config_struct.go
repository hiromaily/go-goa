package config

import "github.com/hiromaily/go-goa/pkg/jwts"

// Root is config root
type Root struct {
	Logger *Logger `toml:"logger" validate:"required"`
	Jwt    *JWT    `toml:"jwt" validate:"required"`
	Hash   *Hash   `toml:"hash" validate:"required"`
	MySQL  *MySQL  `toml:"mysql" validate:"required"`
}

// Logger is zerolog setting
type Logger struct {
	Service      string `toml:"service" validate:"required"`
	Level        int    `toml:"level"` // 0 is possible value
	IsStackTrace bool   `toml:"is_stacktrace"`
	FileName     string `toml:"file_name"`
}

// JWT is JWT setting
type JWT struct {
	Mode       jwts.JWTAlgo `toml:"mode" validate:"oneof=hmac rsa"`
	Audience   string       `toml:"audience" validate:"required"`
	Secret     string       `toml:"secret_code"`
	PrivateKey string       `toml:"private_key"`
	PublicKey  string       `toml:"public_key"`
}

// Hash is hash setting
type Hash struct {
	Salt string `toml:"salt"`
}

// MySQL is MySQL Server setting
type MySQL struct {
	Encrypted bool   `toml:"encrypted"`
	Host      string `toml:"host"`
	Port      uint16 `toml:"port"`
	DBName    string `toml:"dbname"`
	User      string `toml:"user"`
	Pass      string `toml:"pass"`
	Debug     bool   `toml:"debug"`
}
