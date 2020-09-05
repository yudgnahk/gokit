package templates

// MainContent ...
const MainContent = `package main

import (
    "log"
	"os"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"MODULE_NAME/adapters"
	"MODULE_NAME/configs"
	//_ "MODULE_NAME/cmd/docs"
	"MODULE_NAME/controllers"
)

// Package const define
const (
	AppExitCode = 99
)

// @title MODULE_NAME APIs
// @version 1.0
// @description This is a server of MODULE_NAME Service.
// @termsOfService http://swagger.io/terms/

// @contact.name Khang Ha
// @contact.url http://www.swagger.io/support
// @contact.email khanghld@onemount.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /BASE_PATH/v1

// @securityDefinitions.apikey BasicAuthToken
// @in header
// @name Authorization

// @securityDefinitions.apikey JWTAccessToken
// @in header
// @name Authorization
func main() {
	db := database.NewDB()
	if err := db.Open(configs.AppConfig.DB.ConnectionString()); err != nil {
		log.Fatalf("Creating connection to DB: %v", err)
	}

	var r = gin.New()

	// Init Controller
	var (
		appController = controllers.NewAppController()
	)

	v1 := r.Group("/BASE_PATH/v1")
	{
		app := v1.Group("/")
		{
			app.GET("/health", appController.Health)
		}
	}

	if configs.AppConfig.RunMode == gin.DebugMode && configs.AppConfig.Env != "PRODUCTION" {
		r.GET("/BASE_PATH/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	defer func() {
		db.Close()
	}()

	if err := r.Run(configs.AppConfig.AddressListener()); err != nil {
		log.Fatalf("Opening HTTP server: %v", err)
	}
}

func init() {
	_, err := configs.New()
	if err != nil {
		os.Exit(AppExitCode)
	}
}
`