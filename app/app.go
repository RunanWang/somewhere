package app

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/somewhere/handler"
	"github.com/somewhere/middleware"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/somewhere/config"
	"github.com/somewhere/db"
	"github.com/somewhere/handler/algo"
	"github.com/somewhere/handler/products"
	"github.com/somewhere/handler/recommend"
	"github.com/somewhere/handler/records"
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
	db.InitRedisDatabase()
	t.initRouter()
	err := handler.BasicInit()
	if err != nil {
		fmt.Println("No basic init!")
	}
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
	r.Use(middleware.CorsHandler())
	r.Use(gin.Recovery())
	r.Use(middleware.Common)

	///////////////////////////////////////////////////////////////////////////////////////
	// 鉴权API接口
	var authMiddleware = middleware.GinJWTMiddlewareInit(middleware.AllUserAuthorizator)
	r.POST("/login", authMiddleware.LoginHandler)
	r.NoRoute(authMiddleware.MiddlewareFunc(), middleware.NoRouteHandler)
	auth := r.Group("/auth")
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	api := r.Group("/user")
	api.Use(authMiddleware.MiddlewareFunc())
	{
		api.GET("/info", handler.GetUserInfo)
		api.POST("/logout", handler.Logout)
	}
	//////////////////////////////////////////////////////////////////////////////////////
	// 普通API接口
	r.POST("/reg", handler.RegisterHandler)
	r.GET("/basic", handler.GetBasic)
	rootGroup := r.Group("somewhere")
	storesGroup := rootGroup.Group("/stores")
	storesGroup.GET("", stores.GetStores)
	storesGroup.POST("", stores.AddStore)
	storesGroup.PUT("", stores.UpdateStore)
	storesGroup.DELETE("", stores.DeleteStore)

	userStoresGroup := rootGroup.Group("/storespage")
	userStoresGroup.POST("", stores.GetStoresByPage)

	storeInfoGroup := rootGroup.Group("/storeinfo")
	storeInfoGroup.POST("", stores.GetStoreInfo)

	userGroup := rootGroup.Group("/users")
	userGroup.GET("", users.GetUsers)
	userGroup.POST("", users.AddUser)
	userGroup.PUT("", users.UpdateUser)
	userGroup.DELETE("", users.DeleteUser)

	userInfoGroup := rootGroup.Group("/userinfo")
	userInfoGroup.POST("", users.GetUserInfo)

	userPageGroup := rootGroup.Group("/userspage")
	userPageGroup.POST("", users.GetUsersByPage)

	proGroup := rootGroup.Group("/products")
	proGroup.GET("", products.GetProducts)
	proGroup.POST("", products.AddProduct)
	proGroup.PUT("", products.UpdateProduct)
	proGroup.DELETE("", products.DeleteProduct)

	proPageGroup := rootGroup.Group("/productspage")
	proPageGroup.POST("", products.GetProductsByPage)

	recGroup := rootGroup.Group("/records")
	recGroup.GET("", records.GetRecords)
	recGroup.POST("", records.AddRecord)

	recoGroup := rootGroup.Group("/recommend")
	recoGroup.POST("", recommend.GetRecommend)

	algoGroup := rootGroup.Group("/algo")
	algoGroup.POST("/train", algo.TrainModel)
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
