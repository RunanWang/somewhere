package model

import (
	"github.com/globalsign/mgo/bson"
	"github.com/somewhere/db"
)

type TAuth struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Role     string        `json:"role" bson:"role"`
	Username string        `json:"username" bson:"name"`
	Password string        `json:"password" bson:"password"`
}

type User struct {
	UserName   string
	UserClaims []Claims
}

type UserMsg struct {
	Roles        []string
	Introduction string
	Avatar       string
	Name         string
	ID           string
}

type Claims struct {
	ID     int    `json:"claim_id"`
	AuthID int    `json:"auth_id"`
	Type   string `json:"type"`
	Value  string `json:"value"`
}

func (t *TAuth) CheckAuth() bool {
	col := db.MgoDb.C("auth")
	var ret TAuth
	err := col.Find(bson.M{"name": t.Username}).One(&ret)
	if err != nil {
		return false
	}
	if t.Password != ret.Password {
		return false
	}
	return true
}

func GetUserID(username string) string {
	return "https://zbj-bucket1.oss-cn-shenzhen.aliyuncs.com/avatar.JPG"
}

func (t *TAuth) GetRoles() string {
	var roles string
	col := db.MgoDb.C("auth")
	var ret TAuth
	err := col.Find(bson.M{"name": t.Username}).One(&ret)
	if err != nil {
		return roles
	}
	return ret.Role
}

func (t *TAuth) GetAuth() string {
	col := db.MgoDb.C("auth")
	var ret TAuth
	col.Find(bson.M{"name": t.Username}).One(&ret)
	return ret.ID.Hex()
}

func GetUserClaims(userName string) (claims []Claims) {
	var auth TAuth
	col := db.MgoDb.C("auth")
	err := col.Find(bson.M{"name": userName}).One(&auth)
	if err != nil {
		return
	}
	var ret Claims
	ret.Value = auth.Role
	ret.Type = "role"
	ret.AuthID = 1
	ret.ID = 1
	claims = append(claims, ret)
	return
}

func (t *TAuth) AddAuth() error {
	col := db.MgoDb.C("auth")
	err := col.Insert(t)
	return err
}

func (t *TAuth) DeleteAuthByName() error {
	col := db.MgoDb.C("auth")
	err := col.Remove(bson.M{"name": t.Username})
	return err
}
