package service

import (
	"github.com/gin-gonic/gin"
	"github.com/somewhere/model"
	"github.com/somewhere/msg"
)

func AddStore(c *gin.Context, addStoreReq *msg.AddStoresReq) (int, error) {
	storeModel := &model.TStores{
		ID:    addStoreReq.StoreID,
		Name:  addStoreReq.StoreName,
		Level: addStoreReq.StoreLevel,
	}

	return storeModel.AddStore()
}

func GetStores(c *gin.Context, getStoresReq *msg.GetStoresReq) ([]*model.TStores, error) {
	if getStoresReq.StoreID <= 0 {
		return model.GetAllStores()
	} else {
		storesModel := &model.TStores{
			ID: getStoresReq.StoreID,
		}
		return storesModel.GetStoreByID()
	}
}

func UpdateStore(c *gin.Context, updateStoresReq *msg.UpdateStoresReq) (int, error) {

	storeModel := &model.TStores{
		ID:    updateStoresReq.StoreID,
		Name:  updateStoresReq.StoreName,
		Level: updateStoresReq.StoreLevel,
	}

	return storeModel.UpdateStore()
}

func DelStore(c *gin.Context, delStoreReq *msg.DeleteStoresReq) (int, error) {
	storeModel := &model.TStores{
		ID: delStoreReq.StoreID,
	}

	return storeModel.DeleteStore()
}
