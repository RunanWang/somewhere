package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"

	"github.com/garyburd/redigo/redis"
	"github.com/somewhere/db"
	"github.com/somewhere/utils"
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
			err = t.AddRecommendByOrder()
			if err != nil {
				return ansRec, err
			}
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
	var userList []TProduct
	originList, err := GetAllProducts()
	if err != nil {
		return err
	}
	score, err := utils.GetItemScoreFromUserID(t.UserID)
	if err != nil {
		return err
	}

	sort.Slice(score.List, func(i, j int) bool {
		return score.List[i].Score < score.List[j].Score
	})
	for _, item := range score.List {
		for _, detail := range originList {
			if item.ItemID == detail.ID.Hex() {
				userList = append(userList, detail)
				break
			}
		}
	}
	fmt.Println(userList)

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

func (t *TRecommend) AddRecommendByOrder() error {
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
