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
	"time"
)

var (
	tomlPath   = flag.String("f", "", "Toml file path")
	portNum    = flag.String("P", "8080", "Port of server")
	retryCount = flag.Int("rc", 5, "retry count before starting")
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

	//wait until mysql run
	if cnf.Environment == "heroku" {
		//timer
		if os.Getenv("PORT") != "" {
			*portNum = os.Getenv("PORT")
			lg.Debug("exported Port is %s", *portNum)
		}
	}

	// Create service
	var ctx *c.Ctx
	var err error
	for i := 0; i < *retryCount; i++ {
		lg.Info("connecting to db server ...")
		ctx, err = c.SetupContext(cnf)
		if err != nil {
			lg.Errorf("db connection failed. %v", err)
			time.Sleep(3 * time.Second)
			continue
		}
		break
	}
	if err != nil {
		panic("database can not be connected.")
	} else {
		lg.Info("connected!!")
	}

	service := newAPI(ctx)

	// Start service
	if err := service.ListenAndServe(fmt.Sprintf(":%s", *portNum)); err != nil {
		service.LogError("startup", "err", err)
	}
}

// log settings
func initLog() {
	lg.InitializeLog(lg.DebugStatus, lg.TimeShortFile,
		"[go-goa]", "", "hiromaily")
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
func newAPI(ctx *c.Ctx) *goa.Service {
	service := goa.New("api")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	//
	validationHandler, _ := goa.NewMiddleware(md.AuthMiddlewareHandler)

	// Mount security middlewares
	key := []byte(ctx.Conf.Jwt.Key)
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

	//HyUserTech
	hyUserTechController := ctl.NewHyUsertechController(service, ctx)
	service.Context = context.WithValue(service.Context, "HyUserTechController", hyUserTechController)
	app.MountHyUsertechController(service, hyUserTechController)

	//HyUserWorkHistory
	hyUserWorkHistoryController := ctl.NewHyUserWorkHistoryController(service, ctx)
	service.Context = context.WithValue(service.Context, "HyUserWorkHistoryController", hyUserWorkHistoryController)
	app.MountHyUserWorkHistoryController(service, hyUserWorkHistoryController)

	//HyTech
	hyTechController := ctl.NewHyTechController(service, ctx)
	service.Context = context.WithValue(service.Context, "HyTechController", hyTechController)
	app.MountHyTechController(service, hyTechController)

	//HyCompany
	hyCompanyController := ctl.NewHyCompanyController(service, ctx)
	service.Context = context.WithValue(service.Context, "HyCompanyController", hyCompanyController)
	app.MountHyCompanyController(service, hyCompanyController)

	//HyCompanyBranch
	hyCompanyBranchController := ctl.NewHyCompanybranchController(service, ctx)
	service.Context = context.WithValue(service.Context, "HyCompanyBranchController", hyCompanyBranchController)
	app.MountHyCompanybranchController(service, hyCompanyBranchController)

	return service
}
