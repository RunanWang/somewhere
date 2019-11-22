package users

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"
)

func UpdateUser(c *gin.Context) {

	var (
		UpdateUserReq  msg.UpdateUsersReq
		UpdateUserResp msg.UpdateUsersResp
		err            error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in update User handler")
	err = c.Bind(&UpdateUserReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}
	logger = logger.WithFields(log.Fields{
		"req": UpdateUserReq,
	})
	num, err := service.UpdateUser(c, &UpdateUserReq)
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

	UpdateUserResp.UserID = num
	UpdateUserResp.ErrorCode = 0
	UpdateUserResp.RequestID = c.MustGet("request_id").(string)
	logger = logger.WithFields(log.Fields{
		"resp": UpdateUserResp,
	})
	service.CommonInfoResp(c, UpdateUserResp)
}
