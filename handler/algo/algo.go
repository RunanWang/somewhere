package algo

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"
	"github.com/somewhere/utils"
)

func TrainModel(c *gin.Context) {

	var (
		trainModelResp msg.StdResp
		err            error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in train model handler")

	err = utils.TrainModel()
	if err != nil {
		logger = c.MustGet("logger").(*log.Entry)
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		c.Set("logger", logger)
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}

	trainModelResp.ErrorCode = 0
	trainModelResp.ErrorMsg = ""
	logger = logger.WithFields(log.Fields{
		"resp": trainModelResp,
	})
	service.CommonInfoResp(c, trainModelResp)
}
