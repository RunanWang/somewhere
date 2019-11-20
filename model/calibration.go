package model

import (
	"github.com/globalsign/mgo/bson"
	"github.com/somewhere/db"
)

type TCalibration struct {
	StoreID     string      `json:"store_id" bson:"store_id"`
	Rtsp        string      `json:"rtsp" bson:"rtsp"`
	Calibration [][]float64 `json:"calibration" bson:"calibration"`
	Status      int         `json:"status" bson:"status"`
}

func (t *TCalibration) AddCalibration() error {
	var ret TCalibration
	col := db.Db.C("cali")
	err := col.Find(bson.M{"store_id": t.StoreID, "rtsp": t.Rtsp}).One(&ret)
	if err != nil {
		return col.Insert(t)
	}
	if ret.Status == 1 {
		return col.Update(bson.M{"store_id": t.StoreID, "rtsp": t.Rtsp}, bson.M{"$set": bson.M{"status": 2}})
	}
	return nil
}

func GetStoreCalibration(StoreID string) ([]TCalibration, error) {

	col := db.Db.C("cali")

	var ret []TCalibration
	err := col.Find(bson.M{"store_id": StoreID}).All(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
