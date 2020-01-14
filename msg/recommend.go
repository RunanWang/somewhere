package msg

import (
	"github.com/somewhere/model"
)

type GetRecommendReq struct {
	UserID   string `json:"user_id"`
	Query    string `json:"query"`
	PageSize int    `json:"page_size"`
	PageNum  int    `json:"page_num"`
}

type GetRecommendResp struct {
	List []model.TProduct `json:"list"`
	StdResp
}
