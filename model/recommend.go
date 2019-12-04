package model

import (
	"github.com/globalsign/mgo/bson"
)

type TRecommend struct {
	UserID bson.ObjectId      `json:"user_id"`
	List   []TRecommendDetail `json:"list"`
}

type TRecommendDetail struct {
	ProductID bson.ObjectId `json:"item_id"`
}

func (t *TRecommend) GetRecommend() ([]TProduct, error) {
	var list []TProduct
	return list, nil
}
