package stores

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"
)

func AddStore(c *gin.Context) {

	var (
		addStoreReq  msg.TAddCameraReq
		addStoreResp msg.TAddCameraResp
		err          error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in add camera handler")

	err = c.Bind(&addCameraReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, logger, cerror.ErrInvalidParam)
		return
	}
	logger = logger.WithFields(log.Fields{
		"req": addCameraReq,
	})

	id, err := service.AddCamera(c, &addCameraReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})

		if _, isMysql := err.(*mysql.MySQLError); isMysql {
			service.CommonErrorResp(c, logger, cerror.ErrInternalError)
		} else {
			service.CommonErrorResp(c, logger, cerror.ErrInvalidParam)
		}

		return
	}

	addCameraResp.ID = id
	addCameraResp.ErrorCode = 0
	addCameraResp.RequestID = c.MustGet("request_id").(string)
	logger = logger.WithFields(log.Fields{
		"resp": addCameraResp,
	})
	service.CommonInfoResp(c, logger, addCameraResp)
}
