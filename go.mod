module github.com/hiromaily/go-goa

go 1.15

replace resume => ./internal/goa/service/resume

// resume is module name
// path `./internal/goa/service/resume` need to include go.mod

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/ericlagergren/decimal v0.0.0-20191206042408-88212e6cfca9 // indirect
	github.com/friendsofgo/errors v0.9.2
	github.com/go-playground/validator/v10 v10.4.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/google/uuid v1.1.2
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.9.0 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/volatiletech/null/v8 v8.1.2
	github.com/volatiletech/sqlboiler/v4 v4.4.0
	github.com/volatiletech/strmangle v0.0.1
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	goa.design/goa/v3 v3.2.6
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	golang.org/x/sys v0.0.0-20210228012217-479acdf4ea46 // indirect
	resume v0.0.0
)
