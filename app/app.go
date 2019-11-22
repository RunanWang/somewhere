package app

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/somewhere/middleware"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/somewhere/config"
	"github.com/somewhere/db"
	"github.com/somewhere/handler/stores"
	"github.com/somewhere/handler/users"
)

type App struct {
	engine *gin.Engine
}

func NewApp() *App {
	gin.SetMode(gin.ReleaseMode)
	return &App{
		engine: gin.New(),
	}
}

func (t *App) Initialize() {

	var configPath string
	flag.StringVar(&configPath, "config", "./conf/config.toml", "config path")
	flag.Parse()

	fmt.Println("configPath", configPath)
	config.InitConfig(configPath)
	t.initLogger()

	db.InitDatabase()
	db.InitSQLDatabase()
	t.initRouter()
}

func (t *App) initLogger() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	if config.Config.ServiceConfig.LogLevel == "Debug" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

}

func (t *App) initRouter() {
	r := t.engine
	r.Use(gin.Recovery())
	r.Use(middleware.Common)

	rootGroup := r.Group("somewhere")

	storesGroup := rootGroup.Group("/stores")
	storesGroup.GET("", stores.GetStores)
	storesGroup.POST("", stores.AddStore)
	storesGroup.PUT("", stores.UpdateStore)
	storesGroup.DELETE("", stores.DeleteStore)

	userGroup := rootGroup.Group("/users")
	userGroup.GET("", users.GetUsers)
	userGroup.POST("", users.AddUser)
	userGroup.PUT("", users.UpdateUser)
	userGroup.DELETE("", users.DeleteUser)
}

func (t *App) Run() {

	err := t.engine.Run(config.Config.ServiceConfig.Address)
	if err != nil {
		panic(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	sig := <-ch
	log.Info("Received signal %s exiting\n", sig)
}
