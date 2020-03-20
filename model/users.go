package model

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/somewhere/db"
)

type TUser struct {
	ID         bson.ObjectId `json:"user_id" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	NickName   string        `json:"user_name" bson:"user_name"`
	Age        int           `json:"user_age" bson:"user_age"`
	Gender     int           `json:"user_gender" bson:"user_gender"`
	City       string        `json:"user_city" bson:"user_city"`
	Timestamp  int64         `json:"user_timestamp" bson:"user_timestamp"`
	Historysum float64       `json:"user_historysum" bson:"user_historysum"`
}

func (t *TUser) AddUser() error {
	col := db.MgoDb.C("users")
	err := col.Insert(t)
	if err != nil {
		return err
	}
	err = Basic.AddUser()
	return err
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

func GetUserByID(userID string) (user TUser, err error) {
	col := db.MgoDb.C("users")
	var ret TUser
	err = col.Find(bson.M{"_id": bson.ObjectIdHex(userID)}).One(&ret)
	if err != nil {
		fmt.Println(err)
		return ret, err
	}
	return ret, nil
}

func GetUsersByPage(pageNum int, pageSize int) (users []TUser, err error) {
	c := db.MgoDb.C("users")
	pipeM := []bson.M{
		// {"$match": bson.M{"status": "true"}},
		{"$skip": (pageNum - 1) * pageSize},
		{"$limit": pageSize},
		// {"$sort": bson.M{"height": -1}},
	}
	pipe := c.Pipe(pipeM)
	err = pipe.All(&users)
	return users, err
}

func (t *TUser) UpdateUser() error {
	col := db.MgoDb.C("users")
	err := col.Update(bson.M{"_id": t.ID}, bson.M{"$set": bson.M{"user_name": t.NickName, "user_gender": t.Gender, "user_age": t.Age, "user_city": t.City, "user_historysum": t.Historysum}})
	if err != nil {
		return err
	}
	return nil
}

func (t *TUser) DeleteUser() error {
	col := db.MgoDb.C("users")
	err := col.Remove(bson.M{"_id": t.ID})
	if err != nil {
		return err
	}
	err = Basic.DeleteUser()
	return err
}

func GetUserIDByName(userName string) (userID string) {
	var user TUser
	c := db.MgoDb.C("users")
	pipeM := []bson.M{
		{"$match": bson.M{"name": userName}},
		// {"$skip": (pageNum - 1) * pageSize},
		// {"$limit": pageSize},
		// {"$sort": bson.M{"height": -1}},
	}
	pipe := c.Pipe(pipeM)
	err := pipe.One(&user)
	if err != nil {
		fmt.Println(err)
		fmt.Println(userName)
		return ""
	}
	return user.ID.Hex()
}
