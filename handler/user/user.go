package user

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"gitlab.bj.sensetime.com/SenseGo/cali_server/msg"
	"gitlab.bj.sensetime.com/SenseGo/cali_server/service"
)

func GenerateKey(c *gin.Context) {

	var (
		generateKeyReq  msg.TGenerateKeyReq
		generateKeyResp msg.TGenerateResp
		err             error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in GenerateKey handler")

	err = c.Bind(&generateKeyReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}

	err = service.GenerateKey(c, &generateKeyReq, &generateKeyResp)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, cerror.ErrInternalError)
		return
	}
	generateKeyResp.ErrorCode = 0
	generateKeyResp.RequestID = c.MustGet("request_id").(string)
	logger = logger.WithFields(log.Fields{
		"resp": generateKeyResp,
	})
	service.CommonInfoResp(c, generateKeyResp)
}
