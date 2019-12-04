package model

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/globalsign/mgo/bson"
	"github.com/somewhere/db"
)

type TRecommend struct {
	UserID bson.ObjectId      `json:"user_id"`
	List   []TRecommendDetail `json:"list"`
}

type TRecommendDetail struct {
	ProductID bson.ObjectId `json:"item_id"`
}

func (t *TRecommend) GetRecommend() ([]TProduct, error) {
	_, err := db.RedisDb.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(db.RedisDb.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
	var list []TProduct
	return list, err
}
