package model

import (
	"github.com/globalsign/mgo/bson"
	"github.com/somewhere/db"
)

type TRecord struct {
	UserID    int   `json:"user_id" bson:"user_id"`
	ProID     int   `json:"pro_id" bson:"pro_id"`
	Status    int   `json:"is_trade" bson:"is_trade"`
	Timestamp int64 `json:"timestamp" bson:"timestamp"`
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

// func GetAllRecords() (Records []*TRecord, err error) {

// 	rows, err := db.SqlDb.Query("SELECT * from records")
// 	if err != nil {
// 		return nil, err
// 	}
// 	for rows.Next() {
// 		var aRecord TRecord
// 		err = rows.Scan(&aRecord.ProID, &aRecord.UserID, &aRecord.Status)
// 		if err != nil {
// 			return
// 		}
// 		Records = append(Records, &aRecord)
// 	}
// 	return Records, nil
// }
