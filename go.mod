module github.com/hiromaily/go-goa

go 1.17

replace resume => ./internal/goa/service/resume

// resume is module name
// path `./internal/goa/service/resume` need to include go.mod

require (
	github.com/BurntSushi/toml v1.2.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/friendsofgo/errors v0.9.2
	github.com/go-playground/validator/v10 v10.11.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.3.0
	github.com/pkg/errors v0.9.1
	github.com/volatiletech/null/v8 v8.1.2
	github.com/volatiletech/sqlboiler/v4 v4.12.0
	github.com/volatiletech/strmangle v0.0.4
	go.uber.org/zap v1.21.0
	goa.design/goa/v3 v3.7.13
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa
	resume v0.0.0
)

require (
	github.com/dimfeld/httppath v0.0.0-20170720192232-ee938bf73598 // indirect
	github.com/dimfeld/httptreemux/v5 v5.4.0 // indirect
	github.com/ericlagergren/decimal v0.0.0-20211103172832-aca2edc11f73 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.10.6 // indirect
	github.com/manveru/faker v0.0.0-20171103152722-9fbc68a78c4d // indirect
	github.com/mitchellh/mapstructure v1.4.2 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/randomize v0.0.1 // indirect
	github.com/zach-klippenstein/goregen v0.0.0-20160303162051-795b5e3961ea // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	goa.design/plugins/v3 v3.7.13 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.12 // indirect
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
