package msg

import "github.com/somewhere/model"

type GetStoresReq struct {
	StoreID int `form:"store_id" binding:"required"`
}

type GetStoresResp struct {
	List []*model.TStores `json:"list"`
	StdResp
}

type AddStoresReq struct {
	StoreID    int    `json:"store_id" binding:"required"`
	StoreName  string `json:"store_name" binding:"required"`
	StoreLevel int    `json:"store_level" binding:"required"`
}

type AddStoresResp struct {
	StoreID int `json:"store_id"`
	StdResp
}

type UpdateStoresReq struct {
	StoreID    int    `json:"store_id" binding:"required"`
	StoreName  string `json:"store_name" binding:"required"`
	StoreLevel int    `json:"store_level" binding:"required"`
}

type UpdateStoresResp struct {
	StoreID int `json:"store_id"`
	StdResp
}

type DeleteStoresReq struct {
	StoreID int `form:"store_id" binding:"required"`
}

type DeleteStoresResp struct {
	StoreID int `json:"store_id"`
	StdResp
}
