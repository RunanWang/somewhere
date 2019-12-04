package model

import (
	"github.com/globalsign/mgo/bson"
	"github.com/somewhere/db"
)

type TProduct struct {
	ID        bson.ObjectId `json:"item_id" bson:"_id"`
	StoreID   bson.ObjectId `json:"store_id" bson:"store_id"`
	Name      string        `json:"item_name" bson:"item_name"`
	Price     float64       `json:"item_price" bson:"item_price"`
	Score     float64       `json:"item_score" bson:"item_score"`
	SaleCount int           `json:"item_salecount" bson:"item_salecount"`
	Brand     string        `json:"item_brand" bson:"item_brand"`
	Timestamp int64         `json:"item_timestamp" bson:"item_timestamp"`
}

func (t *TProduct) AddProduct() error {
	col := db.MgoDb.C("items")
	err := col.Insert(t)
	return err
}

func (t *TProduct) GetProductByID(id bson.ObjectId) (Product TProduct, err error) {
	col := db.MgoDb.C("items")
	var ret TProduct
	err = col.Find(bson.M{"_id": id}).One(&ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func GetAllProducts() (Products []TProduct, err error) {
	col := db.MgoDb.C("items")
	var ret []TProduct
	err = col.Find(bson.M{}).All(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (t *TProduct) UpdateProduct() error {
	col := db.MgoDb.C("items")
	err := col.Update(bson.M{"_id": t.ID}, bson.M{"$set": bson.M{"store_id": t.StoreID, "item_name": t.Name, "item_price": t.Price, "item_score": t.Score, "item_salecount": t.SaleCount, "item_brand": t.Brand, "item_timestamp": t.Timestamp}})
	if err != nil {
		return err
	}
	return nil
}

func (t *TProduct) DeleteProduct() error {
	col := db.MgoDb.C("items")
	err := col.Remove(bson.M{"_id": t.ID})
	return err
}
