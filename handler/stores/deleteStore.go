package stores

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"
)

func DeleteStore(c *gin.Context) {

	var (
		DeleteStoreReq  msg.DeleteStoresReq
		DeleteStoreResp msg.DeleteStoresResp
		err             error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in Delete store handler")
	err = c.Bind(&DeleteStoreReq)
	if err != nil {
		logger = logger.WithFields(log.Fields{
			"error": err.Error(),
		})
		service.CommonErrorResp(c, cerror.ErrInvalidParam)
		return
	}
	logger = logger.WithFields(log.Fields{
		"req": DeleteStoreReq,
	})
	num, err := service.DeleteStore(c, &DeleteStoreReq)
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

	DeleteStoreResp.StoreID = num
	DeleteStoreResp.ErrorCode = 0
	DeleteStoreResp.RequestID = c.MustGet("request_id").(string)
	logger = logger.WithFields(log.Fields{
		"resp": DeleteStoreResp,
	})
	service.CommonInfoResp(c, DeleteStoreResp)
}
