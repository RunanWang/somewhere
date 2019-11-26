package model

import (
	"github.com/globalsign/mgo/bson"
	"github.com/somewhere/db"
)

type TUser struct {
	ID         bson.ObjectId `json:"user_id" bson:"_id"`
	Name       string        `json:"user_name" bson:"user_name"`
	Age        int           `json:"user_age" bson:"user_age"`
	Gender     int           `json:"user_gender" bson:"user_gender"`
	City       string        `json:"user_city" bson:"user_city"`
	Timestamp  int64         `json:"user_timestamp" bson:"user_timestamp"`
	Historysum float64       `json:"user_historysum" bson:"user_historysum"`
}

func (t *TUser) AddUser() error {
	col := db.MgoDb.C("users")
	err := col.Insert(t)
	return err
}

func (t *TUser) GetUserByName() (users []*TUser, err error) {
	row := db.SqlDb.QueryRow("SELECT * FROM users where name = ?", t.Name)
	if err != nil {
		return
	}

	var aUser TUser
	err = row.Scan(&aUser.ID, &aUser.Name, &aUser.Age)
	if err != nil {
		return
	}
	users = append(users, &aUser)

	return
}

func (t *TUser) GetUserByID() (users []*TUser, err error) {
	row := db.SqlDb.QueryRow("SELECT * FROM users where id = ?", t.ID)
	if err != nil {
		return
	}
	var aUser TUser
	err = row.Scan(&aUser.ID, &aUser.Name, &aUser.Age)
	if err != nil {
		return
	}
	users = append(users, &aUser)

	return
}

func GetAllUsers() (users []TUser, err error) {
	col := db.MgoDb.C("users")
	var ret []TUser
	err = col.Find(bson.M{}).All(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (t *TUser) UpdateUser() error {
	col := db.MgoDb.C("users")
	err := col.Update(bson.M{"_id": t.ID}, bson.M{"$set": bson.M{"user_name": t.Name, "user_gender": t.Gender, "user_age": t.Age, "user_city": t.City, "user_historysum": t.Historysum}})
	if err != nil {
		return err
	}
	return nil
}

func (t *TUser) DeleteUser() error {
	col := db.MgoDb.C("users")
	err := col.Remove(bson.M{"_id": t.ID})
	return err
}
