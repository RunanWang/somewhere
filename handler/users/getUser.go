package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"
)

func GetUsers(c *gin.Context) {

	var (
		getUserReq  msg.GetUsersReq
		getUserResp msg.GetUsersResp
		err         error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in get User handler")

	err = c.Bind(&getUserReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}
	logger = logger.WithFields(log.Fields{
		"req": getUserReq,
	})

	list, err := service.GetUsers(c, &getUserReq)
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

	getUserResp.ErrorCode = 0
	getUserResp.RequestID = c.MustGet("request_id").(string)
	getUserResp.List = list
	logger = logger.WithFields(log.Fields{
		"resp": getUserResp,
	})
	service.CommonInfoResp(c, getUserResp)
}

func GetUserInfo(c *gin.Context) {
	// 单独获取一个人的信息，用于在用户权限下修改自己的信息
	var (
		getUserReq  msg.GetUserInfoReq
		getUserResp msg.GetUserInfoResp
		err         error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in get User Info handler")

	err = c.Bind(&getUserReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}
	logger = logger.WithFields(log.Fields{
		"req": getUserReq,
	})

	userInfo, err := service.GetUserInfo(c, &getUserReq)
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

	getUserResp.ErrorCode = 0
	getUserResp.RequestID = c.MustGet("request_id").(string)
	getUserResp.TUser = userInfo
	logger = logger.WithFields(log.Fields{
		"resp": getUserResp,
	})
	service.CommonInfoResp(c, getUserResp)
}

func GetUsersByPage(c *gin.Context) {

	var (
		getUserReq  msg.GetUsersByPageReq
		getUserResp msg.GetUsersResp
		err         error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in get User handler")

	err = c.Bind(&getUserReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}
	logger = logger.WithFields(log.Fields{
		"req": getUserReq,
	})
	fmt.Println(getUserReq)
	list, err := service.GetUsersByPage(c, &getUserReq)
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

	getUserResp.ErrorCode = 0
	getUserResp.RequestID = c.MustGet("request_id").(string)
	getUserResp.List = list
	logger = logger.WithFields(log.Fields{
		"resp": getUserResp,
	})
	service.CommonInfoResp(c, getUserResp)
}
