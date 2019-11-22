package msg

import "github.com/somewhere/model"

type GetProductsReq struct {
	ProductID int `form:"product_id"`
}

type GetProductsResp struct {
	List []*model.TProduct `json:"list"`
	StdResp
}

type AddProductsReq struct {
	StoreID      int    `json:"store_id"`
	ProductName  string `json:"product_name" binding:"required"`
	ProductPrice int    `json:"product_price"`
}

type AddProductsResp struct {
	ProductID int `json:"product_id"`
	StdResp
}

type UpdateProductsReq struct {
	ProductID    int    `json:"product_id"`
	StoreID      int    `json:"store_id"`
	ProductName  string `json:"product_name" binding:"required"`
	ProductPrice int    `json:"product_price" binding:"required"`
}

type UpdateProductsResp struct {
	ProductID int `json:"update_sucess_num"`
	StdResp
}

type DeleteProductsReq struct {
	ProductID int `form:"product_id" binding:"required"`
}

type DeleteProductsResp struct {
	ProductID int `json:"delete_success_num"`
	StdResp
}
