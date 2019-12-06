package msg

import (
	"github.com/globalsign/mgo/bson"
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
	UserID bson.ObjectId `json:"user_id" `
	ProID  bson.ObjectId `json:"pro_id" `
	Status int           `json:"is_trade" `
}

type AddRecordResp struct {
	RecordID int `json:"success"`
	StdResp
}
