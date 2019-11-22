package service

import (
	"github.com/gin-gonic/gin"
	"github.com/somewhere/model"
	"github.com/somewhere/msg"
)

func AddUser(c *gin.Context, addUserReq *msg.AddUsersReq) (int, error) {
	UserModel := &model.TUser{
		Name: addUserReq.UserName,
		Age:  addUserReq.UserAge,
	}

	return UserModel.AddUser()
}

func GetUsers(c *gin.Context, getUsersReq *msg.GetUsersReq) ([]*model.TUser, error) {
	if getUsersReq.UserID <= 0 {
		return model.GetAllUsers()
	} else {
		UsersModel := &model.TUser{
			ID: getUsersReq.UserID,
		}
		return UsersModel.GetUserByID()
	}
}

func UpdateUser(c *gin.Context, updateUsersReq *msg.UpdateUsersReq) (int, error) {

	UserModel := &model.TUser{
		ID:   updateUsersReq.UserID,
		Name: updateUsersReq.UserName,
		Age:  updateUsersReq.UserAge,
	}

	return UserModel.UpdateUser()
}

func DeleteUser(c *gin.Context, delUserReq *msg.DeleteUsersReq) (int, error) {
	UserModel := &model.TUser{
		ID: delUserReq.UserID,
	}

	return UserModel.DeleteUser()
}
