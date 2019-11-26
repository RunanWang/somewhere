package msg

import "github.com/somewhere/model"

type GetStoresReq struct {
	StoreID int `form:"store_id"`
}

type GetStoresResp struct {
	List []model.TStores `json:"list"`
	StdResp
}

type AddStoresReq struct {
	StoreName  string  `json:"store_name" binding:"required"`
	StoreLevel float64 `json:"store_level"`
	StoreCity  string  `json:"store_city"`
}

type AddStoresResp struct {
	StoreID string `json:"store_id"`
	StdResp
}

type UpdateStoresReq struct {
	StoreID    string  `json:"store_id"`
	StoreName  string  `json:"store_name" `
	StoreLevel float64 `json:"store_level" `
	StoreCity  string  `json:"store_city"`
}

type UpdateStoresResp struct {
	StoreID int `json:"update_sucess_num"`
	StdResp
}

type DeleteStoresReq struct {
	StoreID string `form:"store_id" binding:"required"`
}

type DeleteStoresResp struct {
	StoreID int `json:"delete_success_num"`
	StdResp
}
