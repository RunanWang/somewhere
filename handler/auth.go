package handler

import (
	"net/http"

	"github.com/somewhere/model"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	var auth model.TAuth
	claims := jwt.ExtractClaims(c)
	userName := claims["userName"].(string)
	auth.Username = userName
	avatar := model.GetUserID(userName)

	code := 200
	userRoles := auth.GetRoles()
	data := model.UserMsg{Roles: userRoles, Introduction: "", Avatar: avatar, Name: userName}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  "ok",
		"data": data,
	})
}

func Logout(c *gin.Context) {
	code := 200
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  "ok",
		"data": "success",
	})
}
