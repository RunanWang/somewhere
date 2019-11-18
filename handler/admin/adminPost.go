package admin

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	cerror "gitlab.bj.sensetime.com/SenseGo/cali_server/err"
	"gitlab.bj.sensetime.com/SenseGo/cali_server/msg"
	"gitlab.bj.sensetime.com/SenseGo/cali_server/service"
)

func CommitLine(c *gin.Context) {
	var (
		commitLineReq  msg.CommitLineReq
		commitLineResp msg.CommitLineResp
		err            error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in commitLine handler")

	err = c.Bind(&commitLineReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}

	err = service.CommitLine(c, &commitLineReq, &commitLineResp)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		c.Set("logger", logger)
		service.CommonErrorResp(c, cerror.ErrInternalError)
		return
	}

	commitLineResp.ErrorCode = 0
	commitLineResp.RequestID = c.MustGet("request_id").(string)
	logger = logger.WithFields(log.Fields{
		"resp": commitLineResp,
	})
	service.CommonInfoResp(c, commitLineResp)
}
