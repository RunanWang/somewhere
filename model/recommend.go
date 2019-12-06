package model

import (
	"encoding/json"

	"github.com/garyburd/redigo/redis"
	"github.com/somewhere/db"
)

type TRecommend struct {
	UserID string             `json:"user_id"`
	List   []TRecommendDetail `json:"list"`
}

type TRecommendDetail struct {
	ProductID string `json:"item_id"`
}

func (t *TRecommend) GetRecommend() (TRecommend, error) {
	var userReco TRecommend
	userReco.UserID = t.UserID
	userList, err := redis.String(db.RedisDb.Do("GET", t.UserID))
	if err != nil {
		return userReco, err
	}
	ans := &TRecommend{}
	err = json.Unmarshal([]byte(userList), ans)
	userReco.List = ans.List
	return userReco, nil
}
