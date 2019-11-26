package db

import (
	"github.com/globalsign/mgo"
	"github.com/somewhere/config"
)

var MgoDb *mgo.Database

func InitDatabase() error {

	sesstion, err := mgo.Dial(config.Config.DbConfig.URI)
	if err != nil {
		return err
	}

	sesstion.SetMode(mgo.Eventual, true)

	MgoDb = sesstion.DB(config.Config.DbConfig.DB)

	return createIndex()
}

func createIndex() error {
	col := MgoDb.C("records")
	index := mgo.Index{
		Key:    []string{"user_id, timestamp"},
		Unique: false,
	}
	err := col.EnsureIndex(index)
	if err != nil {
		return err
	}

	col = MgoDb.C("users")
	index = mgo.Index{
		Key:    []string{"user_name"},
		Unique: true,
	}
	err = col.EnsureIndex(index)
	if err != nil {
		return err
	}

	col = MgoDb.C("stores")
	index = mgo.Index{
		Key:    []string{"store_name"},
		Unique: true,
	}
	err = col.EnsureIndex(index)
	if err != nil {
		return err
	}

	col = MgoDb.C("items")
	index = mgo.Index{
		Key:    []string{"store_id", "item_id"},
		Unique: true,
	}
	err = col.EnsureIndex(index)
	return err
}
