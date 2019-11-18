package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.bj.sensetime.com/SenseGo/cali_server/err"
)

func CommonErrorResp(c *gin.Context, aErr *cerror.Error) {
	logger := c.MustGet("logger").(*log.Entry)
	logger.WithFields(log.Fields{
		"error_code": aErr.Code,
		"error_msg":  aErr.Msg,
	}).Warn("FAIL")

	c.JSON(http.StatusOK, aErr)
}

func CommonInfoResp(c *gin.Context, info interface{}) {
	logger := c.MustGet("logger").(*log.Entry)
	startTime := c.MustGet("start_time").(time.Time)
	logger.WithFields(
		log.Fields{"run_time": time.Now().Sub(startTime) / 1000000}).Info()
	c.JSON(http.StatusOK, info)
}
