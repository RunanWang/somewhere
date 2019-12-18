package msg

import "github.com/somewhere/model"

type GetUsersReq struct {
	UserID int `form:"user_id"`
}

type GetUsersResp struct {
	List []model.TUser `json:"list"`
	StdResp
}

type AddUsersReq struct {
	Name       string  `json:"name" binding:"required"`
	UserName   string  `json:"user_name" binding:"required"`
	UserAge    int     `json:"user_age"`
	Gender     int     `json:"user_gender"`
	City       string  `json:"user_city"`
	Historysum float64 `json:"user_historysum"`
}

type AddUsersResp struct {
	UserID string `json:"user_id"`
	StdResp
}

type UpdateUsersReq struct {
	UserID     string  `json:"user_id"`
	UserName   string  `json:"user_name" binding:"required"`
	UserAge    int     `json:"user_age"`
	Gender     int     `json:"user_gender"`
	City       string  `json:"user_city"`
	Historysum float64 `json:"user_historysum"`
}

type UpdateUsersResp struct {
	UserID string `json:"update_sucess_num"`
	StdResp
}

type DeleteUsersReq struct {
	UserID string `json:"user_id" binding:"required"`
	Name   string `json:"name" binding:"required"`
}

type DeleteUsersResp struct {
	UserID string `json:"delete_success_num"`
	StdResp
}
