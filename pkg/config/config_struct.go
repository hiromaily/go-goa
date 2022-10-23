package config

import "github.com/hiromaily/go-goa/pkg/jwts"

// Root is config root
type Root struct {
	Logger *Logger `toml:"logger" validate:"required"`
	Jwt    *JWT    `toml:"jwt" validate:"required"`
	MySQL  *MySQL  `toml:"mysql" validate:"required"`
}

// Logger is zerolog setting
type Logger struct {
	Service      string `toml:"service" validate:"required"`
	Level        int    `toml:"level" validate:"required"`
	IsStackTrace bool   `toml:"is_stacktrace"`
	FileName     string `toml:"file_name"`
}

// Jwt is JWT setting
type JWT struct {
	Mode       jwts.JWTAlgo `toml:"mode" validate:"oneof=hmac rsa"`
	Audience   string       `toml:"audience" validate:"required"`
	Secret     string       `toml:"secret_code"`
	PrivateKey string       `toml:"private_key"`
	PublicKey  string       `toml:"public_key"`
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
