package service

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/somewhere/model"
	"github.com/somewhere/msg"
)

func GetRecommend(c *gin.Context, getRecommendReq *msg.GetRecommendReq) ([]model.TProduct, error) {
	var rec model.TRecommend
	rec.UserID = bson.ObjectIdHex(getRecommendReq.UserID)
	return rec.GetRecommend()
}
