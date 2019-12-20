package products

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"
)

func GetProducts(c *gin.Context) {

	var (
		getProductReq  msg.GetProductsReq
		getProductResp msg.GetProductsResp
		err            error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in get Product handler")

	err = c.Bind(&getProductReq)
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
		"req": getProductReq,
	})

	list, err := service.GetProducts(c, &getProductReq)
	if err != nil {
		logger = c.MustGet("logger").(*log.Entry)
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		c.Set("logger", logger)

		if _, isMysql := err.(*mysql.MySQLError); isMysql {
			service.CommonErrorResp(c, cerror.ErrInternalError)
		} else {
			service.CommonErrorResp(c, cerror.ErrInvalidParam)
		}

		return
	}

	getProductResp.ErrorCode = 0
	getProductResp.RequestID = c.MustGet("request_id").(string)
	getProductResp.List = list
	logger = logger.WithFields(log.Fields{
		"resp": getProductResp,
	})
	service.CommonInfoResp(c, getProductResp)
}

func GetProductsByPage(c *gin.Context) {

	var (
		getProductReq  msg.GetProductsByPageReq
		getProductResp msg.GetProductsResp
		err            error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in get Product handler")

	err = c.Bind(&getProductReq)
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
		"req": getProductReq,
	})

	list, err := service.GetProductsByPage(c, &getProductReq)
	if err != nil {
		logger = c.MustGet("logger").(*log.Entry)
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		c.Set("logger", logger)

		if _, isMysql := err.(*mysql.MySQLError); isMysql {
			service.CommonErrorResp(c, cerror.ErrInternalError)
		} else {
			service.CommonErrorResp(c, cerror.ErrInvalidParam)
		}

		return
	}

	getProductResp.ErrorCode = 0
	getProductResp.RequestID = c.MustGet("request_id").(string)
	getProductResp.List = list
	logger = logger.WithFields(log.Fields{
		"resp": getProductResp,
	})
	service.CommonInfoResp(c, getProductResp)
}
