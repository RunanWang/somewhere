package handler

import (
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"github.com/somewhere/model"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"

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

func RegisterHandler(c *gin.Context) {
	var (
		addAuthReq  msg.AddAuthReq
		addAuthResp msg.AddAuthResp
		err         error
	)
	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in register handler")
	err = c.Bind(&addAuthReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}
	logger = logger.WithFields(log.Fields{
		"req": addAuthReq,
	})

	_, err = service.AddAuth(c, &addAuthReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		if _, isMysql := err.(*mysql.MySQLError); isMysql {
			service.CommonErrorResp(c, cerror.ErrInternalError)
		} else {
			service.CommonErrorResp(c, cerror.ErrInvalidParam)
		}
		return
	}

	if addAuthReq.Role == "store" {

	} else if addAuthReq.Role == "user" {
		var newUser model.TUser
		newUser.Name = addAuthReq.Name
		newUser.NickName = addAuthReq.Name
		newUser.ID = bson.NewObjectId()
		err = newUser.AddUser()
	}

	addAuthResp.ErrorCode = 0
	addAuthResp.RequestID = c.MustGet("request_id").(string)
	logger = logger.WithFields(log.Fields{
		"resp": addAuthResp,
	})
	service.CommonInfoResp(c, addAuthResp)

}
