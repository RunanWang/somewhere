package msg

import (
	"github.com/somewhere/model"
)

type GetRecordsReq struct {
	ProductID string `form:"product_id"`
	UserID    string `form:"user_id"`
}

type GetRecordsResp struct {
	List []model.TRecord `json:"list"`
	StdResp
}

type AddRecordReq struct {
	UserID string `json:"user_id" `
	ProID  string `json:"pro_id" `
	Query  string `json:"query" `
	Status int    `json:"is_trade" `
}

type AddRecordResp struct {
	RecordID int `json:"success"`
	StdResp
}
