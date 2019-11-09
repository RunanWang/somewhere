package main

import (
	"github.com/gin-gonic/gin";
	"github.com/somewhere/handler"
)

func main(){
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("",handler.WelcomePage)
	router.GET("/index",handler.IndexHandler)
	router.Run()
}

