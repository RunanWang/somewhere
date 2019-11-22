package msg

import "github.com/somewhere/model"

type GetUsersReq struct {
	UserID int `form:"user_id"`
}

type GetUsersResp struct {
	List []*model.TUser `json:"list"`
	StdResp
}

type AddUsersReq struct {
	UserName string `json:"user_name" binding:"required"`
	UserAge  int    `json:"user_age"`
}

type AddUsersResp struct {
	UserID int `json:"user_id"`
	StdResp
}

type UpdateUsersReq struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name" binding:"required"`
	UserAge  int    `json:"user_age" binding:"required"`
}

type UpdateUsersResp struct {
	UserID int `json:"update_sucess_num"`
	StdResp
}

type DeleteUsersReq struct {
	UserID int `form:"user_id" binding:"required"`
}

type DeleteUsersResp struct {
	UserID int `json:"delete_success_num"`
	StdResp
}
