package products

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"
)

func UpdateProduct(c *gin.Context) {

	var (
		UpdateProductReq  msg.UpdateProductsReq
		UpdateProductResp msg.UpdateProductsResp
		err               error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in update Product handler")
	err = c.Bind(&UpdateProductReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}
	logger = logger.WithFields(log.Fields{
		"req": UpdateProductReq,
	})
	err = service.UpdateProduct(c, &UpdateProductReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})

		if _, isMysql := err.(*mysql.MySQLError); isMysql {
			service.CommonErrorResp(c, cerror.ErrInternalError)
		} else {
			service.CommonErrorResp(c, cerror.ErrInvalidParam)
		}

		return
	}

	UpdateProductResp.ProductID = 0
	UpdateProductResp.ErrorCode = 0
	UpdateProductResp.RequestID = c.MustGet("request_id").(string)
	logger = logger.WithFields(log.Fields{
		"resp": UpdateProductResp,
	})
	service.CommonInfoResp(c, UpdateProductResp)
}
