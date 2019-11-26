package records

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"
)

func GetRecords(c *gin.Context) {

	var (
		getRecordReq  msg.GetRecordsReq
		getRecordResp msg.GetRecordsResp
		err           error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in get Record handler")

	err = c.Bind(&getRecordReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}
	logger = logger.WithFields(log.Fields{
		"req": getRecordReq,
	})

	list, err := service.GetRecords(c, &getRecordReq)
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

	getRecordResp.ErrorCode = 0
	getRecordResp.RequestID = c.MustGet("request_id").(string)
	getRecordResp.List = list
	logger = logger.WithFields(log.Fields{
		"resp": getRecordResp,
	})
	service.CommonInfoResp(c, getRecordResp)
}
