package recommend

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"
)

func GetRecommend(c *gin.Context) {

	var (
		getRecommendReq  msg.GetRecommendReq
		getRecommendResp msg.GetRecommendResp
		err              error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in get recommend handler")

	err = c.Bind(&getRecommendReq)
	if err != nil {
		logger = c.MustGet("logger").(*log.Entry)
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		c.Set("logger", logger)
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}
	logger = logger.WithFields(log.Fields{
		"req": getRecommendReq,
	})

	list, err := service.GetRecommend(c, &getRecommendReq)
	if err != nil {
		logger = c.MustGet("logger").(*log.Entry)
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		c.Set("logger", logger)
		service.CommonErrorResp(c, cerror.ErrInvalidParam)

		return
	}

	getRecommendResp.List = list
	getRecommendResp.ErrorCode = 0
	getRecommendResp.RequestID = c.MustGet("request_id").(string)
	logger = logger.WithFields(log.Fields{
		"resp": getRecommendResp,
	})
	service.CommonInfoResp(c, getRecommendResp)
}

func AddRecommend(c *gin.Context) {

	var (
		getRecommendReq  msg.GetRecommendReq
		getRecommendResp msg.GetRecommendResp
		err              error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in get recommend handler")

	err = c.Bind(&getRecommendReq)
	if err != nil {
		logger = c.MustGet("logger").(*log.Entry)
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		c.Set("logger", logger)
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}
	logger = logger.WithFields(log.Fields{
		"req": getRecommendReq,
	})

	err = service.AddRecommend(c, &getRecommendReq)
	if err != nil {
		logger = c.MustGet("logger").(*log.Entry)
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		c.Set("logger", logger)
		service.CommonErrorResp(c, cerror.ErrInvalidParam)

		return
	}

	getRecommendResp.ErrorCode = 0
	getRecommendResp.RequestID = c.MustGet("request_id").(string)
	logger = logger.WithFields(log.Fields{
		"resp": getRecommendResp,
	})
	service.CommonInfoResp(c, getRecommendResp)
}
