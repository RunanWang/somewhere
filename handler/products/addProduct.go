package products

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"
)

func AddProduct(c *gin.Context) {

	var (
		addProductReq  msg.AddProductsReq
		addProductResp msg.AddProductsResp
		err            error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in add Product handler")

	err = c.Bind(&addProductReq)
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
		"req": addProductReq,
	})

	id, err := service.AddProduct(c, &addProductReq)
	if err != nil {
		logger = c.MustGet("logger").(*log.Entry)
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		c.Set("logger", logger)
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}

	addProductResp.ProductID = id
	addProductResp.ErrorCode = 0
	addProductResp.RequestID = c.MustGet("request_id").(string)
	logger = logger.WithFields(log.Fields{
		"resp": addProductResp,
	})
	service.CommonInfoResp(c, addProductResp)
}
