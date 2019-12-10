package model

import (
	"github.com/globalsign/mgo/bson"
	"github.com/somewhere/db"
)

type TStores struct {
	ID    bson.ObjectId `json:"store_id" bson:"_id"`
	Name  string        `json:"store_name" bson:"store_name"`
	Level float64       `json:"store_level" bson:"store_level"`
	City  string        `json:"store_city" bson:"store_city"`
	Timestamp  int64    `json:"timestamp" bson:"timestamp"`
}

func (t *TStores) AddStore() error {
	col := db.MgoDb.C("stores")
	err := col.Insert(t)
	return err
}

func (t *TStores) GetStoreByName() (stores []*TStores, err error) {
	row := db.SqlDb.QueryRow("SELECT * FROM stores where name = ?", t.Name)
	if err != nil {
		return
	}

	var aStore TStores
	err = row.Scan(&aStore.ID, &aStore.Name, &aStore.Level)
	if err != nil {
		return
	}
	stores = append(stores, &aStore)

	return
}

func (t *TStores) GetStoreByID() (stores []*TStores, err error) {
	row := db.SqlDb.QueryRow("SELECT * FROM stores where id = ?", t.ID)
	if err != nil {
		return
	}
	var aStore TStores
	err = row.Scan(&aStore.ID, &aStore.Name, &aStore.Level)
	if err != nil {
		return
	}
	stores = append(stores, &aStore)

	return
}

func GetAllStores() (stores []TStores, err error) {
	col := db.MgoDb.C("stores")
	var ret []TStores
	err = col.Find(bson.M{}).All(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (t *TStores) UpdateStore() error {
	col := db.MgoDb.C("stores")
	err := col.Update(bson.M{"_id": t.ID}, bson.M{"$set": bson.M{"store_name": t.Name, "store_level": t.Level, "store_city": t.City}})
	return err
}

func (t *TStores) DeleteStore() error {
	col := db.MgoDb.C("stores")
	err := col.Remove(bson.M{"_id": t.ID})
	return err
}
