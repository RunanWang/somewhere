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

func DeleteUser(c *gin.Context) {

	var (
		DeleteUserReq  msg.DeleteUsersReq
		DeleteUserResp msg.DeleteUsersResp
		err            error
	)
	fmt.Println(DeleteUserReq)
	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in Delete User handler")
	err = c.Bind(&DeleteUserReq)
	if err != nil {
		logger = c.MustGet("logger").(*log.Entry)
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		c.Set("logger", logger)
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}
	logger = logger.WithFields(log.Fields{
		"req": DeleteUserReq,
	})
	num, err := service.DeleteUser(c, &DeleteUserReq)
	if err != nil {
		logger = c.MustGet("logger").(*log.Entry)
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		c.Set("logger", logger)
		if _, isMysql := err.(*mysql.MySQLError); isMysql {
			service.CommonErrorResp(c, cerror.ErrInternalError)
		} else {
			service.CommonErrorResp(c, cerror.ErrInvalidParam)
		}

		return
	}

	DeleteUserResp.UserID = num
	DeleteUserResp.ErrorCode = 0
	DeleteUserResp.RequestID = c.MustGet("request_id").(string)
	logger = logger.WithFields(log.Fields{
		"resp": DeleteUserResp,
	})
	service.CommonInfoResp(c, DeleteUserResp)
}
