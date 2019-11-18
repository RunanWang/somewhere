package msg

import "gitlab.bj.sensetime.com/SenseGo/cali_server/model"

type GetLineReq struct {
	StoreID string `form:"store_id" binding:"required"`
}

type GetLineResp struct {
	List []model.TCalibration `json:"list"`
	StdResp
}

type CommitLineReq struct {
	StoreID string `json:"store_id" binding:"required"`
	// Rtsp    []string `json:"rtsp" binding:"required"`
	List []model.TCommitLine `json:"list" binding:"required"`
}

type CommitLineResp struct {
	Count int `json:"count"`
	StdResp
}
