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
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-playground/validator/v10 v10.4.1
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/google/uuid v1.1.2
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.9.0 // indirect
	github.com/magiconair/properties v1.8.4 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/spf13/afero v1.5.1 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.1.3 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.7.1 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/null v8.0.0+incompatible
	github.com/volatiletech/sqlboiler v3.7.1+incompatible
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	goa.design/goa/v3 v3.2.6
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	golang.org/x/mod v0.4.1 // indirect
	golang.org/x/net v0.0.0-20210222171744-9060382bd457 // indirect
	golang.org/x/sys v0.0.0-20210228012217-479acdf4ea46 // indirect
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/tools v0.1.0 // indirect
	google.golang.org/genproto v0.0.0-20210223151946-22b48be4551b // indirect
	google.golang.org/grpc v1.35.0 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	resume v0.0.0
)
