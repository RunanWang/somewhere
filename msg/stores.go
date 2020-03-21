package msg

import "github.com/somewhere/model"

type GetStoresReq struct {
	StoreID int `form:"store_id"`
}

type GetStoresByPageReq struct {
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
}

type GetStoresResp struct {
	List []model.TStores `json:"data"`
	StdResp
}

type AddStoresReq struct {
	Name       string  `json:"name" binding:"required"`
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
	StoreID string `json:"store_id" binding:"required"`
	Name    string `json:"name" binding:"required"`
}

type DeleteStoresResp struct {
	StoreID int `json:"delete_success_num"`
	StdResp
}

type GetStoreInfoReq struct {
	StoreID string `json:"store_id"`
}

type GetStoreInfoResp struct {
	model.TStores
	StdResp
}
