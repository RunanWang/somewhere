package service

import (
	"github.com/gin-gonic/gin"
	"github.com/somewhere/model"
	"github.com/somewhere/msg"
)

func GetRecommend(c *gin.Context, getRecommendReq *msg.GetRecommendReq) ([]model.TProduct, error) {
	var rec model.TRecommend
	rec.UserID = getRecommendReq.UserID
	rec.PageSize = getRecommendReq.PageSize
	rec.PageNum = getRecommendReq.PageNum
	rec.Query = getRecommendReq.Query
	ans, err := rec.GetRecommend()
	return ans, err
}
