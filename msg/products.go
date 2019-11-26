package msg

import (
	"github.com/somewhere/model"
)

type GetProductsReq struct {
	ProductID int `form:"product_id"`
}

type GetProductsResp struct {
	List []model.TProduct `json:"list"`
	StdResp
}

type AddProductsReq struct {
	StoreID   string  `json:"store_id" bson:"store_id"`
	Name      string  `json:"item_name" bson:"item_name"`
	Price     float64 `json:"item_price" bson:"item_price"`
	Score     float64 `json:"item_score" bson:"item_score"`
	SaleCount int     `json:"item_salecount" bson:"item_salecount"`
	Brand     string  `json:"item_brand" bson:"item_brand"`
}

type AddProductsResp struct {
	ProductID string `json:"item_id"`
	StdResp
}

type UpdateProductsReq struct {
	ID        string  `json:"item_id" bson:"item_id"`
	StoreID   string  `json:"store_id" bson:"store_id"`
	Name      string  `json:"item_name" bson:"item_name"`
	Price     float64 `json:"item_price" bson:"item_price"`
	Score     float64 `json:"item_score" bson:"item_score"`
	SaleCount int     `json:"item_salecount" bson:"item_salecount"`
	Brand     string  `json:"item_brand" bson:"item_brand"`
}

type UpdateProductsResp struct {
	ProductID int `json:"update_sucess_num"`
	StdResp
}

type DeleteProductsReq struct {
	ID string `json:"item_id" bson:"item_id"`
}

type DeleteProductsResp struct {
	ProductID int `json:"delete_success_num"`
	StdResp
}
