package api

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	ctx "github.com/hiromaily/go-goa/ext/context"

	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func init() {
	getVersion()
}

func getVersion() {
	wd, err := os.Getwd()
	if strings.HasSuffix(wd, "test") {
		wd = ".."
	}
	dat, err := ioutil.ReadFile(wd + "/VERSION")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(dat)
	//Config.Version = string(dat)
}

func NewApi(ctx *ctx.GoaApi) *goa.Service {
	service := goa.New("api")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	return service
}
