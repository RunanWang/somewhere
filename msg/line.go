package msg

import "github.com/somewhere/model"

type GetLineReq struct {
	StoreID string `form:"store_id" binding:"required"`
}

type GetLineResp struct {
	List []model.TCalibration `json:"list"`
	StdResp
}

type CommitLineResp struct {
	Count int `json:"count"`
	StdResp
}
