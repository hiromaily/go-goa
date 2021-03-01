package config

// Root is config root
type Root struct {
	Logger *Logger `toml:"logger" validate:"required"`
	Jwt    *JWT    `toml:"jwt" validate:"required"`
	MySQL  *MySQL  `toml:"mysql" validate:"required"`
}

// Logger is zap logger property
type Logger struct {
	Service      string `toml:"service" validate:"required"`
	Env          string `toml:"env" validate:"oneof=dev prod custom"`
	Level        string `toml:"level" validate:"required"`
	IsStackTrace bool   `toml:"is_stacktrace"`
}

// Jwt is JWT property
type JWT struct {
	Secret string `toml:"secret_code" validate:"required"`
}

// MySQL is MySQL Server property
type MySQL struct {
	Encrypted bool   `toml:"encrypted"`
	Host      string `toml:"host"`
	Port      uint16 `toml:"port"`
	DBName    string `toml:"dbname"`
	User      string `toml:"user"`
	Pass      string `toml:"pass"`
}
