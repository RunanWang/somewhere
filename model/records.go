package model

import (
	"github.com/somewhere/db"
)

type TRecord struct {
	RecID  int `json:"rec_id" bson:"rec_id"`
	UserID int `json:"user_id" bson:"user_id"`
	ProID  int `json:"pro_id" bson:"pro_id"`
	Status int `json:"is_trde" bson:"is_trade"`
}

func (t *TRecord) AddRecord() (int, error) {
	// col := db.MgoDb.C("records")
	// return col.Insert(t)
	stmtIns, err := db.SqlDb.Prepare("INSERT INTO records (product_id,user_id,is_trade) VALUES( ?,?,? )") // ? = placeholder
	if err != nil {
		return -1, err
	}
	defer stmtIns.Close()

	rs, err := stmtIns.Exec(t.ProID, t.UserID, t.Status)
	if err != nil {
		return -1, err
	}

	id, err := rs.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

// func GetRecords(StoreID string) ([]TRecord, error) {

// 	// col := db.MgoDb.C("records")

// 	// var ret []TRecord
// 	// err := col.Find(bson.M{"store_id": StoreID}).All(&ret)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// return ret, nil

// }

func GetAllRecords() (Records []*TRecord, err error) {

	rows, err := db.SqlDb.Query("SELECT * from records")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var aRecord TRecord
		err = rows.Scan(&aRecord.RecID, &aRecord.ProID, &aRecord.UserID, &aRecord.Status)
		if err != nil {
			return
		}
		Records = append(Records, &aRecord)
	}
	return Records, nil
}
