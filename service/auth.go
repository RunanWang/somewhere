package service

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
	"github.com/somewhere/model"
	"github.com/somewhere/msg"
)

func AddAuth(c *gin.Context, addAuthReq *msg.AddAuthReq) (string, error) {
	AuthModel := &model.TAuth{
		ID:       bson.NewObjectId(),
		Username: addAuthReq.Name,
		Password: addAuthReq.Password,
		Role:     addAuthReq.Role,
	}
	logger := c.MustGet("logger").(*log.Entry)
	err := AuthModel.AddAuth()
	logger = logger.WithFields(log.Fields{
		"add_item_error": err,
	})
	c.Set("logger", logger)
	return AuthModel.ID.Hex(), err
}
