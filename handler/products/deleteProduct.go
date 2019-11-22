package products

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"
)

func DeleteProduct(c *gin.Context) {

	var (
		DeleteProductReq  msg.DeleteProductsReq
		DeleteProductResp msg.DeleteProductsResp
		err               error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in Delete Product handler")
	err = c.Bind(&DeleteProductReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}
	logger = logger.WithFields(log.Fields{
		"req": DeleteProductReq,
	})
	num, err := service.DeleteProduct(c, &DeleteProductReq)
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

	DeleteProductResp.ProductID = num
	DeleteProductResp.ErrorCode = 0
	DeleteProductResp.RequestID = c.MustGet("request_id").(string)
	logger = logger.WithFields(log.Fields{
		"resp": DeleteProductResp,
	})
	service.CommonInfoResp(c, DeleteProductResp)
}
