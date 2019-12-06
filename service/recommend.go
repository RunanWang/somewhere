package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/somewhere/model"
	"github.com/somewhere/msg"
)

func GetRecommend(c *gin.Context, getRecommendReq *msg.GetRecommendReq) ([]model.TProduct, error) {
	var rec model.TRecommend
	rec.UserID = getRecommendReq.UserID
	var list []model.TProduct
	ans, err := rec.GetRecommend()

	for _, v := range ans.List {
		id := bson.ObjectIdHex(v.ProductID)
		var tempItem model.TProduct
		tempItem.ID = id
		tempItem, err = tempItem.GetProductByID()
		if err != nil {
			fmt.Println("error in gettinf item:", id, err)
			continue
		}
		list = append(list, tempItem)
	}
	return list, nil
}
