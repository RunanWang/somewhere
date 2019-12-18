package service

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
	"github.com/somewhere/model"
	"github.com/somewhere/msg"
)

func AddUser(c *gin.Context, addUserReq *msg.AddUsersReq) (string, error) {
	fmt.Println(addUserReq)
	UserModel := &model.TUser{
		ID:         bson.NewObjectId(),
		Name:       addUserReq.Name,
		NickName:   addUserReq.UserName,
		Gender:     addUserReq.Gender,
		Age:        addUserReq.UserAge,
		City:       addUserReq.City,
		Timestamp:  time.Now().Unix(),
		Historysum: addUserReq.Historysum,
	}
	logger := c.MustGet("logger").(*log.Entry)
	err := UserModel.AddUser()
	AuthModel := &model.TAuth{
		ID:       bson.NewObjectId(),
		Username: addUserReq.Name,
		Password: "111111",
		Role:     "user",
	}
	err = AuthModel.AddAuth()
	logger = logger.WithFields(log.Fields{
		"add_item_error": err,
	})
	c.Set("logger", logger)
	return UserModel.ID.Hex(), err
}

func GetUsers(c *gin.Context, getUsersReq *msg.GetUsersReq) ([]model.TUser, error) {
	return model.GetAllUsers()
}

func UpdateUser(c *gin.Context, updateUsersReq *msg.UpdateUsersReq) (string, error) {
	UserModel := &model.TUser{
		ID:         bson.ObjectIdHex(updateUsersReq.UserID),
		NickName:   updateUsersReq.UserName,
		Gender:     updateUsersReq.Gender,
		Age:        updateUsersReq.UserAge,
		City:       updateUsersReq.City,
		Historysum: updateUsersReq.Historysum,
	}
	return UserModel.ID.Hex(), UserModel.UpdateUser()
}

func DeleteUser(c *gin.Context, delUserReq *msg.DeleteUsersReq) (string, error) {
	UserModel := &model.TUser{
		ID: bson.ObjectIdHex(delUserReq.UserID),
	}
	AuthModel := &model.TAuth{
		Username: delUserReq.Name,
	}
	err := AuthModel.DeleteAuthByName()
	err = UserModel.DeleteUser()
	return UserModel.ID.Hex(), err
}
