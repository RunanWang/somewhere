package line

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	cerror "gitlab.bj.sensetime.com/SenseGo/cali_server/err"
	"gitlab.bj.sensetime.com/SenseGo/cali_server/msg"
	"gitlab.bj.sensetime.com/SenseGo/cali_server/service"
)

func GetLine(c *gin.Context) {

	var (
		getLineReq  msg.GetLineReq
		getLineResp msg.GetLineResp
		err         error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in get line handler")

	err = c.Bind(&getLineReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}

	err = service.GetLine(c, &getLineReq, &getLineResp)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, cerror.ErrInternalError)
		return
	}
	getLineResp.ErrorCode = 0
	getLineResp.RequestID = c.MustGet("request_id").(string)
	logger = logger.WithFields(log.Fields{
		"resp": getLineResp,
	})
	service.CommonInfoResp(c, getLineResp)
}
