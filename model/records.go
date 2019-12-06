package model

import (
	"github.com/globalsign/mgo/bson"
	"github.com/somewhere/db"
)

type TRecord struct {
	RecordID  bson.ObjectId `json:"record_id" bson:"_id"`
	UserID    bson.ObjectId `json:"user_id" bson:"user_id"`
	ItemID    bson.ObjectId `json:"item_id" bson:"item_id"`
	Status    int           `json:"is_trade" bson:"is_trade"`
	Timestamp int64         `json:"timestamp" bson:"timestamp"`
}

func (t *TRecord) AddRecord() error {
	col := db.MgoDb.C("records")
	return col.Insert(t)
}

func GetAllRecords() ([]TRecord, error) {
	col := db.MgoDb.C("records")
	var ret []TRecord
	err := col.Find(bson.M{}).All(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (t *TRecord) GetRecordsByItemID() ([]TRecord, error) {
	col := db.MgoDb.C("records")
	var ret []TRecord
	err := col.Find(bson.M{"item_id": t.ItemID}).All(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (t *TRecord) GetRecordsByUserID() ([]TRecord, error) {
	col := db.MgoDb.C("records")
	var ret []TRecord
	err := col.Find(bson.M{"user_id": t.UserID}).All(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
