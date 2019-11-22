package service

import (
	"github.com/gin-gonic/gin"
	"github.com/somewhere/model"
	"github.com/somewhere/msg"
)

func AddProduct(c *gin.Context, addProductReq *msg.AddProductsReq) (int, error) {
	ProductModel := &model.TProduct{
		Name:    addProductReq.ProductName,
		Price:   addProductReq.ProductPrice,
		StoreID: addProductReq.StoreID,
	}

	return ProductModel.AddProduct()
}

func GetProducts(c *gin.Context, getProductsReq *msg.GetProductsReq) ([]*model.TProduct, error) {
	if getProductsReq.ProductID <= 0 {
		return model.GetAllProducts()
	} else {
		ProductsModel := &model.TProduct{
			ID: getProductsReq.ProductID,
		}
		return ProductsModel.GetProductByID()
	}
}

func UpdateProduct(c *gin.Context, updateProductsReq *msg.UpdateProductsReq) (int, error) {

	ProductModel := &model.TProduct{
		ID:      updateProductsReq.ProductID,
		Name:    updateProductsReq.ProductName,
		StoreID: updateProductsReq.StoreID,
		Price:   updateProductsReq.ProductPrice,
	}

	return ProductModel.UpdateProduct()
}

func DeleteProduct(c *gin.Context, delProductReq *msg.DeleteProductsReq) (int, error) {
	ProductModel := &model.TProduct{
		ID: delProductReq.ProductID,
	}

	return ProductModel.DeleteProduct()
}
