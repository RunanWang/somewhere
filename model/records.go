package model

import (
	"github.com/globalsign/mgo/bson"
	"github.com/somewhere/db"
)

type TRecord struct {
	UserID  int `json:"user_id" bson:"user_id"`
	StoreID int `json:"store_id" bson:"store_id"`
	ProID   int `json:"pro_id" bson:"pro_id"`
	Status  int `json:"is_trde" bson:"is_trade"`
}

func (t *TRecord) AddRecord() error {
	col := db.MgoDb.C("records")
	return col.Insert(t)
}

func GetRecords(StoreID string) ([]TRecord, error) {

	col := db.MgoDb.C("records")

	var ret []TRecord
	err := col.Find(bson.M{"store_id": StoreID}).All(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
