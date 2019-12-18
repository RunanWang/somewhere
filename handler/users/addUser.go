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

func AddUser(c *gin.Context) {

	var (
		addUserReq  msg.AddUsersReq
		addUserResp msg.AddUsersResp
		err         error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in add User handler")
	fmt.Println(addUserReq)
	err = c.Bind(&addUserReq)
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
		"req": addUserReq,
	})

	id, err := service.AddUser(c, &addUserReq)
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

	addUserResp.UserID = id
	addUserResp.ErrorCode = 0
	addUserResp.RequestID = c.MustGet("request_id").(string)
	logger = logger.WithFields(log.Fields{
		"resp": addUserResp,
	})
	service.CommonInfoResp(c, addUserResp)
}
