package main

import (
	"github.com/gin-gonic/gin";
	"net/http"
)

func main(){
	router := gin.Default()
	router.GET("",welcomePage)
	router.Run()
}

func welcomePage(c *gin.Context){
	c.String(http.StatusOK,"welcome!")
}