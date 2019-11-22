package records

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"
)

func AddRecord(c *gin.Context) {

	var (
		addRecordReq  msg.AddRecordReq
		addRecordResp msg.AddRecordResp
		err           error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in add Record handler")

	err = c.Bind(&addRecordReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}
	logger = logger.WithFields(log.Fields{
		"req": addRecordReq,
	})

	id, err := service.AddRecord(c, &addRecordReq)
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

	addRecordResp.RecordID = id
	addRecordResp.ErrorCode = 0
	addRecordResp.RequestID = c.MustGet("request_id").(string)
	logger = logger.WithFields(log.Fields{
		"resp": addRecordResp,
	})
	service.CommonInfoResp(c, addRecordResp)
}
