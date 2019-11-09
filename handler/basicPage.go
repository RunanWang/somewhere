package handler

import(
	"github.com/gin-gonic/gin";
	"net/http"
)

func WelcomePage(c *gin.Context){
	c.String(http.StatusOK,"welcome!")
}

func IndexHandler(c *gin.Context){
	c.HTML(http.StatusOK,"index.tmpl",gin.H{
		"title":"Hello",
	})
}