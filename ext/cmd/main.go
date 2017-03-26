package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/security/jwt"
	conf "github.com/hiromaily/go-goa/ext/configs"
	c "github.com/hiromaily/go-goa/ext/context"
	ctl "github.com/hiromaily/go-goa/ext/controllers"
	md "github.com/hiromaily/go-goa/ext/middlewares"
	m "github.com/hiromaily/go-goa/ext/models"
	g "github.com/hiromaily/go-goa/goa"
	"github.com/hiromaily/go-goa/goa/app"
	lg "github.com/hiromaily/golibs/log"
	"io/ioutil"
	"os"
	"strings"
)

var (
	tomlPath = flag.String("f", "", "Toml file path")
	portNum  = flag.Int("P", 0, "Port of server")
)

func init() {
	//command-line
	flag.Parse()

	//log settings
	initLog()

	//version
	//getVersion()

}

func main() {
	//config
	cnf := conf.New(*tomlPath)

	// Create service
	ctx := c.SetupContext(cnf)
	service := newApi(ctx)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}

// log settings
func initLog() {
	lg.InitializeLog(lg.DebugStatus, lg.LogOff, 99,
		"[go-goa]", "/var/log/go/go-goa.log")
}

// version
func getVersion() {
	//get current directory
	wd, _ := os.Getwd()
	if strings.HasSuffix(wd, "test") {
		wd = ".."
	}
	dat, err := ioutil.ReadFile(wd + "/VERSION")
	if err != nil {
		lg.Fatal(err)
	}
	fmt.Print(dat)
	//Config.Version = string(dat)
}

// Create service object for Goa
func newApi(ctx *c.Ctx) *goa.Service {
	service := goa.New("api")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	//
	validationHandler, _ := goa.NewMiddleware(md.AuthMiddlewareHandler)

	// Mount security middlewares
	key := []byte("keys")
	keyResolver := jwt.NewSimpleResolver([]jwt.Key{key})
	app.UseJWTMiddleware(service, jwt.New(keyResolver, validationHandler, app.NewJWTSecurity()))

	// Extend service.Context with model objects
	userModel := &m.User{Db: ctx.Db}
	service.Context = context.WithValue(service.Context, "user", userModel)

	// Mount controller
	app.MountHealthController(service, g.NewHealthController(service))
	app.MountPublicController(service, g.NewPublicController(service))

	// Mount controller with extended
	//app.MountHyCompanyController(service, g.NewHyCompanyController(service))
	//app.MountHyUserController(service, g.NewHyUserController(service))

	//Auth
	authController := ctl.NewAuthController(service, ctx)
	service.Context = context.WithValue(service.Context, "AuthController", authController)
	app.MountAuthController(service, authController)

	//HyUser
	hyUserController := ctl.NewHyUserController(service, ctx)
	service.Context = context.WithValue(service.Context, "HyUserController", hyUserController)
	app.MountHyUserController(service, hyUserController)

	//HyCompany
	hyCompanyController := ctl.NewHyCompanyController(service, ctx)
	service.Context = context.WithValue(service.Context, "HyCompanyController", hyCompanyController)
	app.MountHyCompanyController(service, hyCompanyController)

	return service
}
