package db

import (
	"github.com/globalsign/mgo"
	"github.com/somewhere/config"
)

var Db *mgo.Database

func InitDatabase() error {

	sesstion, err := mgo.Dial(config.Config.DbConfig.URI)
	if err != nil {
		return err
	}

	sesstion.SetMode(mgo.Eventual, true)

	Db = sesstion.DB(config.Config.DbConfig.DB)

	return createIndex()
}

func createIndex() error {
	col := Db.C("cali")
	index := mgo.Index{
		Key:    []string{"store_id", "rtsp"},
		Unique: true,
	}

	err := col.EnsureIndex(index)

	return err
}
