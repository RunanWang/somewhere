package stores

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"
)

func GetStores(c *gin.Context) {

	var (
		getStoreReq  msg.GetStoresReq
		getStoreResp msg.GetStoresResp
		err          error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in get store handler")

	err = c.Bind(&getStoreReq)
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
		"req": getStoreReq,
	})

	list, err := service.GetStores(c, &getStoreReq)
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

	getStoreResp.ErrorCode = 0
	getStoreResp.RequestID = c.MustGet("request_id").(string)
	getStoreResp.List = list
	logger = logger.WithFields(log.Fields{
		"resp": getStoreResp,
	})
	service.CommonInfoResp(c, getStoreResp)
}
