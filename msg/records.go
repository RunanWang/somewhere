package msg

import "github.com/somewhere/model"

type GetRecordsReq struct {
	RecordID int `form:"Record_id"`
}

type GetRecordsResp struct {
	List []*model.TRecord `json:"list"`
	StdResp
}

type AddRecordReq struct {
	UserID int `json:"user_id" `
	ProID  int `json:"pro_id" `
	Status int `json:"is_trde" `
}

type AddRecordResp struct {
	RecordID int `json:"Record_id"`
	StdResp
}
