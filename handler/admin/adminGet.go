package admin

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"
)

func AdminGet(c *gin.Context) {

	var (
		getAdminReq  msg.GetLineReq
		getAdminResp msg.GetLineResp
		err          error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in get line handler")

	err = c.Bind(&getAdminReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}

	err = service.GetAdmin(c, &getLineReq, &getLineResp)
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

func AdminPost(c *gin.Context) {

	var (
		getAdminReq  msg.GetLineReq
		getAdminResp msg.GetLineResp
		err          error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in get line handler")

	err = c.Bind(&getAdminReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}

	err = service.GetAdmin(c, &getLineReq, &getLineResp)
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
