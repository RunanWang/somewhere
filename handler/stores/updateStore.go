package stores

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"
)

func UpdateStore(c *gin.Context) {

	var (
		UpdateStoreReq  msg.UpdateStoresReq
		UpdateStoreResp msg.UpdateStoresResp
		err             error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in update store handler")
	err = c.Bind(&UpdateStoreReq)
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
		"req": UpdateStoreReq,
	})
	num, err := service.UpdateStore(c, &UpdateStoreReq)
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

	UpdateStoreResp.StoreID = num
	UpdateStoreResp.ErrorCode = 0
	UpdateStoreResp.RequestID = c.MustGet("request_id").(string)
	logger = logger.WithFields(log.Fields{
		"resp": UpdateStoreResp,
	})
	service.CommonInfoResp(c, UpdateStoreResp)
}
