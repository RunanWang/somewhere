package model

import (
	"github.com/globalsign/mgo/bson"
	"github.com/somewhere/db"
)

type TStores struct {
	ID        bson.ObjectId `json:"store_id" bson:"_id"`
	Name      string        `json:"name" bson:"name"`
	NickName  string        `json:"store_name" bson:"store_name"`
	Level     float64       `json:"store_level" bson:"store_level"`
	City      string        `json:"store_city" bson:"store_city"`
	Timestamp int64         `json:"timestamp" bson:"timestamp"`
}

func (t *TStores) AddStore() error {
	col := db.MgoDb.C("stores")
	err := col.Insert(t)
	if err != nil {
		return err
	}
	err = Basic.AddShop()
	return err
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

func GetStoresByPage(pageNum int, pageSize int) (stores []TStores, err error) {
	c := db.MgoDb.C("stores")
	pipeM := []bson.M{
		// {"$match": bson.M{"status": "true"}},
		{"$skip": (pageNum - 1) * pageSize},
		{"$limit": pageSize},
		// {"$sort": bson.M{"height": -1}},
	}
	pipe := c.Pipe(pipeM)
	err = pipe.All(&stores)
	return stores, err
}

func (t *TStores) UpdateStore() error {
	col := db.MgoDb.C("stores")
	err := col.Update(bson.M{"_id": t.ID}, bson.M{"$set": bson.M{"store_name": t.Name, "store_level": t.Level, "store_city": t.City}})
	return err
}

func (t *TStores) DeleteStore() error {
	col := db.MgoDb.C("stores")
	err := col.Remove(bson.M{"_id": t.ID})
	if err != nil {
		return err
	}
	err = Basic.DeleteShop()
	return err
}

func (t *TStores) GetStoreByStoreID() error {
	col := db.MgoDb.C("stores")
	err := col.Find(bson.M{"_id": t.ID}).One(&t)
	return err
}
