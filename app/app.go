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

	userGroup := rootGroup.Group("/user")
	userGroup.GET("", user.userGet)

	adminGroup := rootGroup.Group("/admin")
	adminGroup.GET("", admin.adminGet)
	adminGroup.POST("", admin.adminPost)
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
