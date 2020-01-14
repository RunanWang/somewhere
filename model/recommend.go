package model

import (
	"encoding/json"
	"errors"

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

func (t *TRecommend) AddRecommend() error {
	userList, err := GetAllProducts()
	if err != nil {
		return err
	}
	userID := t.UserID
	for _, item := range userList {
		ans, err := json.Marshal(item)
		if err != nil {
			return err
		}
		_, err = db.RedisDb.Do("lpush", userID, ans)
		if err != nil {
			return err
		}
	}
	n, _ := db.RedisDb.Do("EXPIRE", userID, 10*60)
	if n != int64(1) {
		return errors.New("error in expire")
	}
	return nil
}
