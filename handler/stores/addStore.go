package stores

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	cerror "github.com/somewhere/err"
	"github.com/somewhere/msg"
	"github.com/somewhere/service"
)

func AddStore(c *gin.Context) {

	var (
		addStoreReq  msg.AddStoresReq
		addStoreResp msg.AddStoresResp
		err          error
	)

	logger := c.MustGet("logger").(*log.Entry)
	logger.Tracef("in add store handler")

	err = c.Bind(&addStoreReq)
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
		"req": addStoreReq,
	})

	id, err := service.AddStore(c, &addStoreReq)
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

	addStoreResp.StoreID = id
	addStoreResp.ErrorCode = 0
	addStoreResp.RequestID = c.MustGet("request_id").(string)
	logger = logger.WithFields(log.Fields{
		"resp": addStoreResp,
	})
	service.CommonInfoResp(c, addStoreResp)
}
