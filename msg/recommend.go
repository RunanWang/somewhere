package msg

import (
	"github.com/somewhere/model"
)

type GetRecommendReq struct {
	UserID string `form:"user_id"`
}

type GetRecommendResp struct {
	List []model.TProduct `json:"list"`
	StdResp
}
