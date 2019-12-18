package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"
)

func BasicInit() error {
	return service.BasicInit()
}

func GetBasic(c *gin.Context) {

	var (
		getBasicResp msg.GetBasicResp
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in get Basic handler")

	basic := service.GetBasic()

	getBasicResp.ErrorCode = 0
	getBasicResp.RequestID = c.MustGet("request_id").(string)
	getBasicResp.TBasic = basic
	logger = logger.WithFields(log.Fields{
		"resp": getBasicResp,
	})
	service.CommonInfoResp(c, getBasicResp)
}
