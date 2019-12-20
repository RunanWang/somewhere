package service

import (
	"time"

	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/somewhere/model"
	"github.com/somewhere/msg"
)

func AddProduct(c *gin.Context, addProductReq *msg.AddProductsReq) (string, error) {
	ProductModel := &model.TProduct{
		ID:        bson.NewObjectId(),
		Name:      addProductReq.Name,
		Price:     addProductReq.Price,
		StoreID:   bson.ObjectIdHex(addProductReq.StoreID),
		Score:     addProductReq.Score,
		SaleCount: addProductReq.SaleCount,
		Brand:     addProductReq.Brand,
		Timestamp: time.Now().Unix(),
	}
	logger := c.MustGet("logger").(*log.Entry)
	err := ProductModel.AddProduct()
	logger = logger.WithFields(log.Fields{
		"add_item_error": err,
	})
	c.Set("logger", logger)
	return ProductModel.ID.Hex(), err
}

func GetProducts(c *gin.Context, getProductsReq *msg.GetProductsReq) ([]model.TProduct, error) {
	if getProductsReq.StoreID == "" {
		return model.GetAllProducts()
	}
	var temp model.TProduct
	temp.StoreID = bson.ObjectIdHex(getProductsReq.StoreID)
	return temp.GetProductsByStoreID()
}

func GetProductsByPage(c *gin.Context, getProductsReq *msg.GetProductsByPageReq) ([]model.TProduct, error) {
	pageNum := getProductsReq.PageNum
	pageSize := getProductsReq.PageSize
	StoreID := getProductsReq.StoreID
	if StoreID == "" {
		return model.GetAllProductsByPage(pageNum, pageSize)
	}
	tempStoreID := bson.ObjectIdHex(StoreID)
	return model.GetProductsByPage(pageNum, pageSize, tempStoreID)
}

func UpdateProduct(c *gin.Context, updateProductsReq *msg.UpdateProductsReq) error {
	ProductModel := &model.TProduct{
		ID:        bson.ObjectIdHex(updateProductsReq.ID),
		Name:      updateProductsReq.Name,
		Price:     updateProductsReq.Price,
		StoreID:   bson.ObjectIdHex(updateProductsReq.StoreID),
		Score:     updateProductsReq.Score,
		SaleCount: updateProductsReq.SaleCount,
		Brand:     updateProductsReq.Brand,
		Timestamp: time.Now().Unix(),
	}
	return ProductModel.UpdateProduct()
}

func DeleteProduct(c *gin.Context, delProductReq *msg.DeleteProductsReq) (int, error) {
	ProductModel := &model.TProduct{
		ID: bson.ObjectIdHex(delProductReq.ID),
	}
	return 0, ProductModel.DeleteProduct()
}
