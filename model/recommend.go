package model

import (
	"encoding/json"
	"errors"

	"github.com/garyburd/redigo/redis"
	"github.com/somewhere/db"
)

type TRecommend struct {
	UserID   string `json:"user_id"`
	Query    string `json:"query"`
	PageSize int    `json:"page_size"`
	PageNum  int    `json:"page_num"`
}

func (t *TRecommend) GetRecommend() ([]TProduct, error) {
	var ansRec []TProduct
	// key := fmt.Sprint(t.UserID, "_", t.Query)
	db.InitRedisDatabase()
	is_key_exit, err := redis.Bool(db.RedisDb.Do("EXISTS", t.UserID))
	if err != nil {
		return ansRec, err
	}
	if !is_key_exit {
		err = t.AddRecommend()
		if err != nil {
			return ansRec, err
		}
	}
	startNum := (t.PageNum - 1) * t.PageSize
	endNum := (t.PageNum) * t.PageSize
	userList, err := redis.Values(db.RedisDb.Do("lrange", t.UserID, startNum, endNum))
	if err != nil {
		return ansRec, err
	}
	for _, byteRec := range userList {
		ans := &TProduct{}
		v, ok := byteRec.([]byte)
		if ok {
			err = json.Unmarshal([]byte(v), ans)
			if err != nil {
				return ansRec, err
			}
			ansRec = append(ansRec, *ans)
		} else {
			return ansRec, err
		}
	}
	return ansRec, nil
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
