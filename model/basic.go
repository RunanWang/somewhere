package model

import (
	"github.com/globalsign/mgo/bson"
	"github.com/somewhere/db"
)

type TBasic struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	UserNum   int           `json:"user_num" bson:"user_num"`
	ShopNum   int           `json:"shop_num" bson:"shop_num"`
	ItemNum   int           `json:"item_num" bson:"item_num"`
	RecordNum int           `json:"record_num" bson:"record_num"`
}

var Basic TBasic

func BasicInit() error {
	col := db.MgoDb.C("basic")
	err := col.Find(bson.M{}).One(&Basic)
	return err
}

func (t *TBasic) AddUser() error {
	col := db.MgoDb.C("basic")
	err := col.Update(bson.M{"_id": Basic.ID}, bson.M{"$set": bson.M{"user_num": Basic.UserNum + 1}})
	if err != nil {
		return err
	}
	err = col.Find(bson.M{}).One(&Basic)
	return nil
}

func (t *TBasic) DeleteUser() error {
	col := db.MgoDb.C("basic")
	err := col.Update(bson.M{"_id": Basic.ID}, bson.M{"$set": bson.M{"user_num": Basic.UserNum - 1}})
	if err != nil {
		return err
	}
	err = col.Find(bson.M{}).One(&Basic)
	return nil
}

func (t *TBasic) AddShop() error {
	col := db.MgoDb.C("basic")
	err := col.Update(bson.M{"_id": Basic.ID}, bson.M{"$set": bson.M{"shop_num": Basic.ShopNum + 1}})
	if err != nil {
		return err
	}
	err = col.Find(bson.M{}).One(&Basic)
	return nil
}

func (t *TBasic) DeleteShop() error {
	col := db.MgoDb.C("basic")
	err := col.Update(bson.M{"_id": Basic.ID}, bson.M{"$set": bson.M{"shop_num": Basic.ShopNum - 1}})
	if err != nil {
		return err
	}
	err = col.Find(bson.M{}).One(&Basic)
	return nil
}

func (t *TBasic) AddItem() error {
	col := db.MgoDb.C("basic")
	err := col.Update(bson.M{"_id": Basic.ID}, bson.M{"$set": bson.M{"item_num": Basic.ItemNum + 1}})
	if err != nil {
		return err
	}
	err = col.Find(bson.M{}).One(&Basic)
	return nil
}

func (t *TBasic) DeleteItem() error {
	col := db.MgoDb.C("basic")
	err := col.Update(bson.M{"_id": Basic.ID}, bson.M{"$set": bson.M{"item_num": Basic.ItemNum - 1}})
	if err != nil {
		return err
	}
	err = col.Find(bson.M{}).One(&Basic)
	return nil
}
