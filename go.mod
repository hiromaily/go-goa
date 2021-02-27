module github.com/hiromaily/go-goa

go 1.15

replace resume => ./internal/goa/service/resume
// resume is module name
// path `./internal/goa/service/resume` need to include go.mod

require (
	goa.design/goa/v3 v3.2.6
	golang.org/x/mod v0.4.1 // indirect
	golang.org/x/net v0.0.0-20210222171744-9060382bd457 // indirect
	golang.org/x/sys v0.0.0-20210223212115-eede4237b368 // indirect
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/tools v0.1.0 // indirect
	google.golang.org/genproto v0.0.0-20210223151946-22b48be4551b // indirect
	google.golang.org/grpc v1.35.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	resume v0.0.0
)
